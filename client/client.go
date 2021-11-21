package main

import (
	"context"
	"fmt"
	"log"
	service "miniproject03/service"
	"os"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	//client connection
	c := service.NewServiceClient(conn)

	//how to read input from console //flags
	UserInput := os.Args[1:]
	ReadPort := UserInput[0]
	ReadName := UserInput[1]

	connect(ReadName, ReadPort, c)

}

func connect(name string, port string, serviceClient service.ServiceClient) {
	response, err := serviceClient.JoinClientToServer(context.Background(), &service.RequestConnection{Name: name, Port: port})
	if err != nil {
		fmt.Printf("Error:%v", err)
	}
	fmt.Printf("Join Response, %v", response.JoinStatus)
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
