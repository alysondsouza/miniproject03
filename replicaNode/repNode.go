package main

import (
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
}

func main() {
	// Flags: -name nodeX - address localhost -port 9000 -ips 9001:9002
	var name = flag.String("name", "nodeX", "The name of the replicaNode.")
	var address = flag.String("address", "localhost", "The address of the replicaNode.")
	var port = flag.String("port", "9000", "The number of the port.")
	var ips = flag.String("ips", "", "The other ReplicaNode's ip addresses")
	flag.Parse()

	n := NewNode(*name, *address, *port)
	n.Init(*ips)

	<-done
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

func (n *ReplicaNode) GetName(_ context.Context, reqname *service.InfoRequest) (*service.InfoResponse, error) {
	fmt.Printf("%v is requesting name from %v\n", reqname, n.id)

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
			_, err := c.PublishClientAddress(context.Background(), &service.ReplicateClientAddress{Addresses: n.frontEndClients})
			if err != nil {
				// TODO IMPLEMENT ERROR HANDLING
				n.logger.ErrorFatalf("Error:Not replicated: %s", err)
			}
		}(key, serviceClient)
	}

	wait.Wait()
}

func (n *ReplicaNode) PublishClientAddress(_ context.Context, replic *service.ReplicateClientAddress) (*service.Ack, error){
	n.replicate(replic.Addresses)
	return &service.Ack{Status: service.Ack_SUCCESS}, nil
}

func (n *ReplicaNode) replicate(addresses []string) {
	for _, address := range addresses {
		n.frontEndClients = append(n.frontEndClients, address)
	}
}

func NewNode(name string, address string, port string) *ReplicaNode {
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
	}
}
