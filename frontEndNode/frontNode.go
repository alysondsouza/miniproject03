package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"server"
	"service"
	"strconv"
	"strings"
	"utils"
)

var primClient service.ServiceClient
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

	node := NewFrontEndNode(*name, *address, *port)
	server.InitServer(node.ipAddress, node, node.logger)

	conn, err := grpc.Dial(node.primaryServer, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	//client connection to primary server
	primClient = service.NewServiceClient(conn)
	_, err = primClient.RegisterClientPortOnPrimaryServer(context.Background(), &service.RequestPort{Name: node.id, Port: node.ipAddress.String()})
	if err != nil {
		node.logger.ErrorFatalf("Error registering on primary server. :: %v", err)
	}
	fmt.Printf("%v, %v Logged in on Port %v\n", node.id, node.ipAddress, node.primaryServer)

	go node.commandShell()

	<-done
}

func (f *FrontEndNode) commandShell() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Write \"bid\" to bid or \"result\" for the current status of the auction.")

	for scanner.Scan() {
		var input = strings.Split(strings.ToLower(scanner.Text()), " ")

		switch input[0] {
		case "bid":
			myBid, err := strconv.Atoi(input[1])
			if err != nil {
				f.logger.WarningPrintf("%v is not a valid bid (must be an integer)\n", myBid)
			} else {
				f.bidCase(myBid)
			}

		case "result":
			f.resultCase()

		default:
			fmt.Println("Use one of the two commandos: 'bid' or 'request': \n" +
				"Example (To bid with 25kr): bid 25 \n" +
				"Example (To request the status of the Auction): result")
		}
	}
}

func (f *FrontEndNode) bidCase(bid int) {
	f.lamport.Increment()

	request := &service.RequestBid{
		NodeId:  f.id,
		MyBid:   int32(bid),
		Lamport: f.lamport.Value(),
	}

	ack, err := primClient.Bid(context.Background(), request)
	if err != nil {
		f.logger.WarningPrintf("Connection was lost. :: %v", err)
	}

	if ack.GetStatus() == service.Ack_SUCCESS {
		fmt.Printf("Bid Status: %v\n", ack.GetStatus())
	} else if ack.GetStatus() == service.Ack_FAIL {
		fmt.Printf("Bid Status: %v\nThe bid was too low.\nWrite 'result' to see the highest bid.\n", ack.GetStatus())
	}
}

func (f *FrontEndNode) resultCase() {
	result := &service.RequestStatus{}

	auctionResult, err := primClient.Result(context.Background(), result)

	if err != nil {
		f.logger.WarningPrintf("Connection was lost. :: ", err)
	}

	fmt.Printf("Action stil going: %v, highest bider: %v, highest bid: %v\n",
		auctionResult.Ongoing, auctionResult.NodeId, auctionResult.Price)
}

func NewFrontEndNode(name, address, port string) *FrontEndNode {
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
