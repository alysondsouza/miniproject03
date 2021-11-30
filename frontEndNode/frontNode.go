package main

import (
	"bufio"
	"client"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"server"
	"service"
	"strconv"
	"strings"
	"utils"
)

var primClient service.ServiceClient
const PrimaryServer = "localhost:9000"
var done = make(chan bool)

type FrontEndNode struct {
	primaryServer string
	id            string
	ipAddress     *net.TCPAddr
	logger        *utils.Logger
	lamport       *utils.Lamport
	service.UnimplementedServiceServer
}

func main() {
	// Flags: -name nodeX -port 9000
	var name = flag.String("name", "nodeX", "The name of the replicaNode.")
	var address = flag.String("address", "localhost", "The address of the replicaNode.")
	var port = flag.String("port", "5000", "The number of the port.")
	flag.Parse()

	node := NewFrontEndNode(*name, *address, *port)

	// Host front end node on server
	server.InitServer(node.ipAddress, node, node.logger)

	//client connection to primary server
	primClient = client.CreateClient(node.primaryServer, node.logger, grpc.WithInsecure())

	_, err := primClient.RegisterClientPortOnPrimaryServer(context.Background(), &service.RequestPort{Name: node.id, Port: node.ipAddress.String()})
	if err != nil {
		node.logger.ErrorFatalf("Error registering on primary server. :: %v", err)
	}

	node.logger.InfoPrintf("%v registered itself on the primary server %v", node.id, node.primaryServer)

	go node.commandShell()

	<-done
}

// commandShell starts the command interface used to query on the auction hosted by a ReplicaNode.
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
				f.bid(myBid)
			}

		case "result":
			f.result()

		default:
			fmt.Println("Use one of the two commandos: 'bid' or 'request': \n" +
				"Example (To bid with 25kr): bid 25 \n" +
				"Example (To request the status of the Auction): result")
		}
	}
}

// bid on the ReplicaNode hosting the auction.
func (f *FrontEndNode) bid(bid int) {
	f.lamport.Increment()

	ack, err := primClient.Bid(context.Background(), &service.RequestBid{NodeId: f.id, MyBid: int32(bid), Lamport: f.lamport.Value()})
	if err != nil {
		f.logger.WarningPrintf("Connection was lost. :: %v", err)
	}

	if ack.GetStatus() == service.Ack_SUCCESS {
		f.logger.InfoPrintf("Bid status: %v", ack.GetStatus())
		f.result()
	} else if ack.GetStatus() == service.Ack_FAIL {
		f.logger.InfoPrintf("Bid status: %v", ack.GetStatus())
		f.logger.WarningPrintln("The bid was too low")
		f.result()
	} else if ack.GetStatus() == service.Ack_EXCEPTION {
		f.logger.WarningPrintf("Bid status: %v", ack.GetStatus())
		f.logger.WarningPrintln("The auction has ended.")
		f.result()
	}
}

// result queries the ReplicaNode hosting the auction.
func (f *FrontEndNode) result() {
	auctionResult, err := primClient.Result(context.Background(), &service.RequestStatus{})

	if err != nil {
		f.logger.ErrorPrintf("Connection was lost. :: ", err)
	}

	fmt.Println("---------------------")
	fmt.Printf("Auction still going: %v\n", auctionResult.Ongoing)
	fmt.Printf("Highest bidder: %v\n", auctionResult.NodeId)
	fmt.Printf("Highest bid: %v\n", auctionResult.Price)
	fmt.Println("---------------------")
}

// Redial dials up a new connection to the new primary ReplicaNode.
func (f *FrontEndNode) Redial(_ context.Context, rp *service.RequestPort) (*service.ReturnPort, error) {
	f.logger.InfoPrintln("Connecting to auction...")
	f.primaryServer = rp.Port
	primClient = client.CreateClient(rp.Port, f.logger, grpc.WithInsecure())
	f.logger.InfoPrintln("Connection established.")
	return &service.ReturnPort{Address: f.ipAddress.String()}, nil
}

// NewFrontEndNode creates and returns a new FrontEndNode with the specified name, and ip address.
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
