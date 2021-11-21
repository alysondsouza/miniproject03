package client

import (
	"context"
	"fmt"
	"log"
	service "miniproject03/service"
	"os"

	"google.golang.org/grpc"
)

type Client struct {
	connectingPort 	string
	name 			string
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("9000", grpc.WithInsecure())
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

	client := Client{connectingPort: ReadPort, name: ReadName}

	connect(&client, c)



}

func connect(client *Client, serviceClient service.serviceClient){
	response, _:= serviceClient.JoinClientToServer(context.Background(), service.RequestConnection{name: client.name, port: client.connectingPort})
	fmt.Printf("Join Response, %s", response)
}

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