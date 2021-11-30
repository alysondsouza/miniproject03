package client

import (
	"google.golang.org/grpc"
	"service"
	"utils"
)

func CreateClient(ipAddress string, logger *utils.Logger, ops... grpc.DialOption) service.ServiceClient {
	conn, err := grpc.Dial(ipAddress, ops...)
	if err != nil {
		logger.ErrorFatalf("Could not connect to %v. :: %v", ipAddress, err)
	}

	return service.NewServiceClient(conn)
}

