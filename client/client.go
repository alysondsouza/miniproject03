package client

import (
	"google.golang.org/grpc"
	"service"
	"utils"
)

func CreateClient(ipAddress string, logger *utils.Logger) service.ServiceClient {
	conn, err := grpc.Dial(ipAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.ErrorFatalf("Could not connect to %v. :: %v", ipAddress, err)
	}

	return service.NewServiceClient(conn)
}
