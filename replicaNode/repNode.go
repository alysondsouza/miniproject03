package main

import (
	"auction"
	"client"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"server"
	"service"
	"strings"
	"sync"
	"time"
	"utils"
)

var done = make(chan bool)
var heart = sync.WaitGroup{}

type ReplicaNode struct {
	id              string
	ipAddress       *net.TCPAddr
	logger          *utils.Logger
	lamport         *utils.Lamport
	replicaPeers    map[string]service.ServiceClient
	frontEndClients []string
	auction         *auction.Auction
	primaryReplica  service.ServiceClient
	service.UnimplementedServiceServer
}

// go run . -name s1 -port 9000 -ips 9001:9002
// go run . -name s2 -port 9001 -ips 9000:9002
// go run . -name s3 -port 9002 -ips 9000:9001

// go run . -name node1 -address localhost -port 5000

func main() {
	// Flags: -name nodeX - address localhost -port 9000 -ips 9001:9002
	var name = flag.String("name", "nodeX", "The name of the replicaNode.")
	var address = flag.String("address", "localhost", "The address of the replicaNode.")
	var port = flag.String("port", "9000", "The number of the port.")
	var ips = flag.String("ips", "", "The other ReplicaNode's ip addresses")
	flag.Parse()

	n := NewReplicaNode(*name, *address, *port)
	n.init(*ips)

	<-done
}

// init initializes the ReplicaNode by setting up its server and connecting to the other ReplicaNodes.
// It also starts the ReplicaNode.auction.
func (n *ReplicaNode) init(ips string) {
	server.InitServer(n.ipAddress, n, n.logger)
	n.connectToReplicas(ips)

	// Connect to primary replica
	n.primaryReplica = client.CreateClient(":9000", n.logger, grpc.WithInsecure(), grpc.WithBlock())
	n.heartbeat()
	n.auction.Start()
}

// connectToReplicas connects to all the other specified ReplicaNode ip addresses and creates a client.
// These client connections are stored in the ReplicaNode.replicaPeers map.
func (n *ReplicaNode) connectToReplicas(ips string) {
	var addresses = strings.Split(ips, ":")

	for _, address := range addresses {
		replicaClient := client.CreateClient(":"+address, n.logger, grpc.WithInsecure(), grpc.WithBlock())
		response, err := replicaClient.GetId(context.Background(), &service.InfoRequest{
			RequestName: n.id,
		})

		if err != nil {
			n.logger.ErrorLogger.Fatalf("Could not find the name from IP-address %v :: %v", address, err)
		}

		n.replicaPeers[response.ResponseName] = replicaClient
	}
}

// GetId replies the ReplicaNode with the id of this node.
func (n *ReplicaNode) GetId(_ context.Context, reqName *service.InfoRequest) (*service.InfoResponse, error) {
	n.logger.InfoPrintf("%v is requesting the id of %v.", reqName.RequestName, n.id)
	return &service.InfoResponse{ResponseName: n.id}, nil
}

// RegisterClientPortOnPrimaryServer registers a FrontEndNode's ip address on this ReplicaNode (which must be the primary server).
func (n *ReplicaNode) RegisterClientPortOnPrimaryServer(_ context.Context, rc *service.RequestPort) (*service.ReturnPort, error) {
	n.frontEndClients = append(n.frontEndClients, rc.Port)
	n.broadcast()
	return &service.ReturnPort{Address: fmt.Sprintf("%v registered with port %v in primary server", rc.Name, rc.Port)}, nil
}

// Bid receives a bid from a FrontEndNode and returns a service.Ack acknowledgement back to the client.
func (n *ReplicaNode) Bid(_ context.Context, rb *service.RequestBid) (*service.Ack, error) {
	n.logger.InfoPrintf("Received a bid of %v from %v.", rb.MyBid, rb.NodeId)
	valid, err := n.auction.CheckBid(rb.GetNodeId(), int(rb.MyBid))
	var a service.Ack

	a.Status = service.Ack_SUCCESS

	if !valid {
		a.Status = service.Ack_FAIL

		if err != nil {
			a.Status = service.Ack_EXCEPTION
		}
	} else {
		n.broadcast()
	}

	return &a, nil
}

// Result returns the current status of this ReplicaNode's auction
func (n *ReplicaNode) Result(_ context.Context, _ *service.RequestStatus) (*service.AuctionResults, error) {
	return &service.AuctionResults{
		Ongoing: n.auction.IsOnGoing(),
		NodeId:  n.auction.GetHighestBidderId(),
		Price:   int32(n.auction.GetHighestBid())}, nil
}

// broadcast a replicate request to all backup ReplicaNodes.
func (n *ReplicaNode) broadcast() {
	wait := sync.WaitGroup{}

	for key, serviceClient := range n.replicaPeers {
		wait.Add(1)

		go func(key string, c service.ServiceClient) {
			defer wait.Done()
			n.logger.InfoPrintf("Replicating to %v.", key)

			_, err := c.PublishReplicate(context.Background(), &service.Replica{
				AuctionReplica: &service.AuctionReplica{
					HighestBidderId: n.auction.GetHighestBidderId(),
					HighestBid:      int32(n.auction.GetHighestBid()),
					OnGoing:         n.auction.IsOnGoing(),
					TimeRemaining:   int32(n.auction.GetTimeRemaining()),
				},
				Addresses: n.frontEndClients,
			})
			if err != nil {
				n.logger.ErrorPrintf("Could not replicate to %v", key)
			}
		}(key, serviceClient)
	}

	wait.Wait()
}

// PublishReplicate receives a replicate request from the primary ReplicaNode.
func (n *ReplicaNode) PublishReplicate(_ context.Context, replica *service.Replica) (*service.Ack, error) {
	n.replicate(replica)
	return &service.Ack{Status: service.Ack_SUCCESS}, nil
}

// replicate backs up the auction and other fields received from the primary ReplicaNode.
func (n *ReplicaNode) replicate(replica *service.Replica) {
	n.frontEndClients = replica.Addresses
	newAuction := auction.NewAuction()
	newAuction.SetHighestBidderId(replica.AuctionReplica.HighestBidderId)
	newAuction.SetHighestBid(int(replica.AuctionReplica.HighestBid))
	newAuction.SetIsOnGoing(replica.AuctionReplica.OnGoing)
	newAuction.SetRemainingTime(int(replica.AuctionReplica.TimeRemaining))
	n.auction = newAuction
}

// heartBeat continuously sends a service.HealthCheckRequest to the primary replica after a specific amount of time.
// If the message sending fails, a new election starts.
func (n *ReplicaNode) heartbeat() {
	go func() {
		for {
			heart.Wait()
			_, err := n.primaryReplica.Check(context.Background(), &service.HealthCheckRequest{Service: n.id})
			if err != nil {
				n.logger.InfoPrintf("%v has discovered that the primary server is down!", n.id)
				n.startElection()
			}
			time.Sleep(time.Duration(500) * time.Millisecond)
		}
	}()
}

// Check always returns a service.HealthCheckResponse. It returns an error if the connection is down.
// (This method is called on the primary replica)
func (n *ReplicaNode) Check(context.Context, *service.HealthCheckRequest) (*service.HealthCheckResponse, error) {
	return &service.HealthCheckResponse{Status: "Up"}, nil
}

// Election receives a service.ElectionRequest from a ReplicaNode, which has a lower id.
// Sends back this ReplicaNode's id to the caller.
func (n *ReplicaNode) Election(_ context.Context, er *service.ElectionRequest) (*service.ElectionAnswer, error) {
	n.logger.InfoPrintf("%v received an Election message from %v. Sending back answer...", n.id, er.Id)
	return &service.ElectionAnswer{Id: n.id}, nil
}

// Coordinator set the caller of the method as coordinator.
func (n *ReplicaNode) Coordinator(_ context.Context, cr *service.CoordinatorRequest) (*service.ElectionAnswer, error) {
	n.logger.InfoPrintf("%v received a Coordinator message from %v.", n.id, cr.Id)
	n.logger.InfoPrintf("%v is now the primary server", cr.Id)

	// Start heart
	heart.Done()

	// Set new primary server
	n.primaryReplica = client.CreateClient(cr.IpAddress, n.logger, grpc.WithInsecure())

	return &service.ElectionAnswer{}, nil
}

// broadCastElection broadcasts election messages to all processes with a higher id than this ReplicaNode.
func (n *ReplicaNode) broadcastElection() int {
	n.logger.InfoPrintf("%v is broadcasting Election messages to all other replica nodes...", n.id)
	doneBroadcasting := make(chan int)
	wait := sync.WaitGroup{}

	counter := utils.NewCounter()

	for key, serviceClient := range n.replicaPeers {
		// Only broadcast to higher processes (lower process ids).
		if key > n.id {
			continue
		}

		wait.Add(1)
		go func(key string, c service.ServiceClient) {
			defer wait.Done()

			n.logger.InfoPrintf("%v is broadcasting an Election messages to %v", n.id, key)

			answer, err := c.Election(context.Background(), &service.ElectionRequest{Id: n.id})

			if err != nil {
				//n.logger.InfoPrintf("%v received no answer. Declaring itself as winner...", n.id)
			} else {
				// Increment answer counter
				counter.Increment()

				// Received answer from higher process id => stop heart.
				if answer.Id < n.id {
					n.logger.InfoPrintf("%v received an answer from a higher process '%v'", n.id, answer.Id)
					n.logger.InfoPrintln("Stopping heart and broadcast...")
					heart.Add(1)
					doneBroadcasting <- 0
				}
			}
		}(key, serviceClient)
	}

	go func() {
		wait.Wait()
		close(doneBroadcasting)
	}()

	<-doneBroadcasting

	n.logger.InfoPrintf("%v is done broadcasting Election.", n.id)

	return counter.Value()
}

// broadcastVictory broadcasts the election victory of this ReplicaNode to all other ReplicaNodes.
func (n *ReplicaNode) broadcastVictory() {
	n.logger.InfoPrintf("%v is broadcasting Victory messages to all other replica nodes.", n.id)
	wait := sync.WaitGroup{}

	for key, serviceClient := range n.replicaPeers {
		wait.Add(1)
		go func(key string, c service.ServiceClient) {
			defer wait.Done()

			_, err := c.Coordinator(context.Background(), &service.CoordinatorRequest{Id: n.id, IpAddress: n.ipAddress.String()})
			if err != nil {
				//n.logger.ErrorPrintf("Error sending victory to %v. :: %v", key, err)
			}
		}(key, serviceClient)
	}

	wait.Wait()
	n.logger.InfoPrintf("%v is done broadcasting Victory.", n.id)
	n.reconnectToFrontEnd()
}

// reconnectToFrontEnd sends a request to all FrontEndNodes to connect to the new primary ReplicaNode.
func (n *ReplicaNode) reconnectToFrontEnd() {
	for _, s := range n.frontEndClients {
		c := client.CreateClient(s, n.logger, grpc.WithInsecure())
		_, err := c.Redial(context.Background(), &service.RequestPort{Name: n.id, Port: n.ipAddress.String()})
		if err != nil {
			n.logger.ErrorFatalf("Client could not connect to new Primary server %v", n.id)
		}
	}
}

// startElection starts a new election from this ReplicaNode.
func (n *ReplicaNode) startElection() {
	n.logger.InfoPrintf("%v is starting an Election.", n.id)

	// Check to see if I'm the highest process.
	highest := true
	for key := range n.replicaPeers {
		if key < n.id {
			highest = false
			break
		}
	}

	if highest {
		n.logger.InfoPrintf("%v is the highest process.", n.id)
		n.broadcastVictory()
	} else {
		n.logger.InfoPrintf("%v is NOT the highest process.", n.id)
		answers := n.broadcastElection()
		if answers == 0 {
			n.logger.InfoPrintf("%v received no answers. Declaring itself winner.", n.id)
			n.broadcastVictory()
		}
	}
}

// NewReplicaNode creates a new ReplicaNode with a given name and ip address.
func NewReplicaNode(name string, address string, port string) *ReplicaNode {
	logger := utils.NewLogger(name)
	ipAddress, err := net.ResolveTCPAddr("tcp", address+":"+port)
	if err != nil {
		logger.ErrorLogger.Fatalf("Could not resolve IP address %v:%v. :: %v", address, port, err)
	}

	return &ReplicaNode{
		id:              name,
		ipAddress:       ipAddress,
		lamport:         utils.NewLamport(),
		logger:          logger,
		replicaPeers:    make(map[string]service.ServiceClient),
		frontEndClients: make([]string, 0),
		auction:         auction.NewAuction(),
	}
}
