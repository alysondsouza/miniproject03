package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"server"
	"service"
	"utils"
)

var done = make(chan bool)

const PrimaryServer = "localhost:9000"

type FrontEndNode struct {
	primaryServer string
	id            string
	ipAddress     *net.TCPAddr
	logger        *utils.Logger
	lamport       *utils.Lamport
	service.UnimplementedServiceServer
}

func main() {
	// Flags: -name nodeX - address localhost -port 9000
	var name = flag.String("name", "nodeX", "The name of the replicaNode.")
	var address = flag.String("address", "localhost", "The address of the replicaNode.")
	var port = flag.String("port", "5000", "The number of the port.")
	flag.Parse()

	node := NewNode(*name, *address, *port)
	server.InitServer(node.ipAddress, node, node.logger)

	conn, err := grpc.Dial(node.primaryServer, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	//client connection
	c := service.NewServiceClient(conn)
	_, err = c.RegisterClientPortOnPrimaryServer(context.Background(), &service.RequestPort{Name: node.id, Port: node.ipAddress.String()})
	if err != nil {
		node.logger.ErrorFatalf("Error registering on primary server. :: %v", err)
	}
	fmt.Printf("%v, %v Logged in on Port %v", node.id, node.ipAddress, node.primaryServer)

	<-done
}

func NewNode(name, address, port string) *FrontEndNode {
	logger := utils.NewLogger(name)
	ipAddress, err := net.ResolveTCPAddr("tcp", address+":"+port)
	if err != nil {
		logger.ErrorLogger.Fatalf("Could not resolve IP address %v:%v. :: %v", address, port, err)
	}

	return &FrontEndNode{
		primaryServer: PrimaryServer,
		id:            name,
		ipAddress:     ipAddress,
		lamport:       utils.NewLamport(),
		logger:        logger,
	}
}

/*
func Bid (){
	log.Println("Give a bid: $")
	var bidInput int32
	fmt.Scanf("%s", &bidInput)

	//c error
	response, err := c.Bid(context.Background(), &bidInput)
	if err != nil {
		log.Fatalf("Error when calling Bid: %s", err)
	}
	log.Printf("Response from Server: %s", response.state)

}

*/
