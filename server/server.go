package main

import (
	"log"
	"net"

	service "miniproject03/service"
	//client "miniproject03/client"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type BidingServer struct {
	service.UnimplementedServiceServer
	clients []string
}

type Client struct {
	connectingPort 	string
	name 			string
}

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := BidingServer{}
	service.RegisterServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}

func (s *BidingServer) Bid(ctx context.Context, bid *service.RequestBid) (*service.Ack, error) {
	return service.Ack{state: service.Ack_FAIL}, nil
}

func (s *BidingServer) Broadcast(ctx context.Context, bid *service.RequestBid) (*service.Ack, error) {
	return service.Ack{state: service.Ack_FAIL}, nil
}

func (s *BidingServer) JoinClientToServer(ctx context.Context, requestConn *service.RequestConnection) (*service.JoinResponse, error) {
	s.clients = append(s.clients, requestConn.Name)
	return service.JoinResponse{response: "Joinned to Server"}, nil
}