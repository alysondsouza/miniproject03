package main

import (
	"log"
	"net"

	"service"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type AuctionServer struct {
	service.UnimplementedServiceServer
	clients []Client
}

type Client struct {
	ConnectingPort string
	Name           string
}

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := &AuctionServer{clients: make([]Client, 0)}
	service.RegisterServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}

func (s *AuctionServer) Bid(ctx context.Context, bid *service.RequestBid) (*service.Ack, error) {
	return &service.Ack{Status: service.Ack_SUCCESS}, nil
}

func (s *AuctionServer) Broadcast(ctx context.Context, bid *service.RequestBid) (*service.Ack, error) {
	return &service.Ack{Status: service.Ack_SUCCESS}, nil
}

func (s *AuctionServer) JoinClientToServer(ctx context.Context, requestConn *service.RequestConnection) (*service.JoinResponse, error) {
	s.clients = append(s.clients, Client{ConnectingPort: requestConn.Port, Name: requestConn.Name})
	return &service.JoinResponse{JoinStatus: "Joinned to Server"}, nil
}
