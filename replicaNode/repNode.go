package main

import (
	"auction"
	"client"
	"context"
	"flag"
	"fmt"
	"net"
	"server"
	"service"
	"strings"
	"sync"
	"utils"
)

var done = make(chan bool)

type ReplicaNode struct {
	id              string
	ipAddress       *net.TCPAddr
	logger          *utils.Logger
	lamport         *utils.Lamport
	replicaPeers    map[string]service.ServiceClient
	frontEndClients []string
	service.UnimplementedServiceServer
	currentAuction *auction.Auction
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
	n.Init(*ips)
	n.StartAuction()

	<-done
}

func (n *ReplicaNode) StartAuction() {
	n.currentAuction.Start()
}

func (n *ReplicaNode) Bid(_ context.Context, rb *service.RequestBid) (*service.Ack, error) {
	//TODO
	valid := n.currentAuction.CheckBid(rb.GetNodeId(), int(rb.MyBid))
	var a service.Ack

	a.Status = service.Ack_SUCCESS

	if !valid {
		a.Status = service.Ack_FAIL
	}

	return &a, nil
}

func (n *ReplicaNode) Result(_ context.Context, rs *service.RequestStatus) (*service.AuctionResults, error) {
	return &service.AuctionResults{
		Ongoing: n.currentAuction.IsOnGoing(),
		NodeId:  n.currentAuction.GetHighestBidderId(),
		Price:   int32(n.currentAuction.GetHighestBid())}, nil
}

func (n *ReplicaNode) Init(ips string) {
	server.InitServer(n.ipAddress, n, n.logger)
	n.setUpIps(ips)
}

func (n *ReplicaNode) setUpIps(ips string) {
	var addresses = strings.Split(ips, ":")

	for _, address := range addresses {
		replicaClient := client.CreateClient(":"+address, n.logger)
		response, err := replicaClient.GetName(context.Background(), &service.InfoRequest{
			RequestName: n.id,
		})

		if err != nil {
			n.logger.ErrorLogger.Fatalf("Could not find the name from IP-address %v :: %v", address, err)
		}

		n.replicaPeers[response.ResponseName] = replicaClient
	}
}

func (n *ReplicaNode) GetName(_ context.Context, reqName *service.InfoRequest) (*service.InfoResponse, error) {
	fmt.Printf("%v is requesting name from %v\n", reqName, n.id)

	return &service.InfoResponse{
		ResponseName: n.id,
	}, nil
}

func (n *ReplicaNode) RegisterClientPortOnPrimaryServer(_ context.Context, rc *service.RequestPort) (*service.ReturnPort, error) {
	n.frontEndClients = append(n.frontEndClients, rc.Port)
	n.broadcast()
	return &service.ReturnPort{Address: fmt.Sprintf("%v registered with port %v in primary server", rc.Name, rc.Port)}, nil
}

func (n *ReplicaNode) broadcast() {
	wait := sync.WaitGroup{}

	for key, serviceClient := range n.replicaPeers {
		wait.Add(1)

		go func(key string, c service.ServiceClient) {
			defer wait.Done()
			n.logger.InfoPrintf("Replicating to %v.", key)

			_, err := c.PublishReplicate(context.Background(), &service.Replica{
				AuctionReplica: &service.AuctionReplica{
					HighestBidderId: n.currentAuction.GetHighestBidderId(),
					HighestBid:      int32(n.currentAuction.GetHighestBid()),
					OnGoing:         n.currentAuction.IsOnGoing(),
					TimeRemaining:   int32(n.currentAuction.GetTimeRemaining()),
				},
				Addresses: n.frontEndClients,
			})
			if err != nil {
				// TODO IMPLEMENT ERROR HANDLING
				n.logger.ErrorFatalf("Error:Not replicated: %s", err)
			}
		}(key, serviceClient)
	}

	wait.Wait()
}

func (n *ReplicaNode) PublishReplicate(_ context.Context, replica *service.Replica) (*service.Ack, error) {
	n.replicate(replica)
	return &service.Ack{Status: service.Ack_SUCCESS}, nil
}

func (n *ReplicaNode) replicate(replica *service.Replica) {
	n.frontEndClients = replica.Addresses
	newAuction := auction.NewAuction()
	newAuction.SetHighestBidderId(replica.AuctionReplica.HighestBidderId)
	newAuction.SetHighestBid(int(replica.AuctionReplica.HighestBid))
	newAuction.SetIsOnGoing(replica.AuctionReplica.OnGoing)
	newAuction.SetRemainingTime(int(replica.AuctionReplica.TimeRemaining))
	n.currentAuction = newAuction
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
		currentAuction:  auction.NewAuction(),
	}
}
