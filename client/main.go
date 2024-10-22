package main

import (
	pb "Chitty-Chat/ChatService"
	"Chitty-Chat/client/util"
	"context"
	"log"
)

func main() {
	conn, err := util.CreateClientServerConnection()
	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}

	client := pb.NewChatServiceClient(conn)
	clientName := util.StartApp()

	stream, err := client.AddClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to start stream: %v\n", err)
	}

	err = util.SendInitialMessage(stream, clientName)
	if err != nil {
		log.Fatalf("Failed to send initial message: %v\n", err)
	}

	go util.SendMessages(stream, clientName, client)
	go util.RecieveMessages(stream)

	//blocker
	bl := make(chan bool)
	<-bl
}