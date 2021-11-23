package server

import (
	"net"
	"utils"

	"service"

	"google.golang.org/grpc"
)

func InitServer(ipAddress *net.TCPAddr, s service.ServiceServer, logger *utils.Logger) {
	go func() {
		lis, err := net.Listen("tcp", ipAddress.String())
		if err != nil {
			logger.ErrorFatalf("Failed to listen on port %v. :: %v", ipAddress.Port, err)
		}

		grpcServer := grpc.NewServer()
		service.RegisterServiceServer(grpcServer, s)

		if err := grpcServer.Serve(lis); err != nil {
			logger.ErrorFatalf("Failed to serve gRPC server over port 9000: %v", err)
		}
	}()
}
