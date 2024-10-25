package main

import (
	clientUtil "Chitty-Chat/client/util"
	pb "Chitty-Chat/grpc"
	"context"
	"log"
)


func main() {
	conn, err := clientUtil.CreateClientServerConnection()
	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}

	client := pb.NewChatServiceClient(conn)
	clientName := clientUtil.StartApp()

	stream, err := client.AddClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to start stream: %v\n", err)
	}

	err = clientUtil.SendInitialMessage(stream, clientName)
	if err != nil {
		log.Fatalf("Failed to send initial message: %v\n", err)
	}

	go clientUtil.SendMessages(stream, clientName, client)
	go clientUtil.RecieveMessages(stream)

	//blocker
	bl := make(chan bool)
	<-bl
}