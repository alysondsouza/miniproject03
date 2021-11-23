package main

import (
	"context"
	"fmt"
	"log"

	"os"
	"service"

	"google.golang.org/grpc"
)

var doneCha = make(chan bool)
func main() {
	//how to read input from console //flags
	UserInput := os.Args[1:]
	ReadPort := UserInput[0]
	ReadName := UserInput[1]

	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	//client connection
	c := service.NewServiceClient(conn)
	connect(ReadName, ReadPort, c)
	<-doneCha
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
