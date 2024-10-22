package main

import (
	pb "Chitty-Chat/ChatService"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}

	client := pb.NewChatServiceClient(conn)
	clientName := startApp()

	stream, err := client.AddClient(context.Background())
	message := fmt.Sprintf("[%s has joined the chat]", clientName)
	stream.Send(&pb.ChatMessageClient{ Sender: clientName, Message: message })
	if err != nil {
		log.Fatalf("Failed to start stream connection: %v\n", err)
	}

	go sendMessage(stream, clientName, client)
	go recieveMessage(stream)

	//blocker
	bl := make(chan bool)
	<-bl
}

func startApp() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Your Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read from console: %v\n", err)
	}
	name = strings.Trim(name, "\r\n")

	return name
}

func sendMessage(stream pb.ChatService_AddClientClient, clientName string, client pb.ChatServiceClient) {
	for {
		reader := bufio.NewReader(os.Stdin)
		enteredMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read from console: %v\n", err)
		}
		enteredMessage = strings.Trim(enteredMessage, "\r\n")

		if enteredMessage == "quit" {
			client.LeaveChat(context.Background(), &pb.ClientName{ ClientName: clientName })
			break
		}

		// Clear enteredMessage in console
		fmt.Print("\033[A\033[2K")
		sendErr := stream.Send(&pb.ChatMessageClient{ Sender: clientName, Message: enteredMessage })
		if sendErr != nil {
			fmt.Printf("There was an error while sending message to server %v\n", sendErr)
		} else {
			fmt.Printf("Me: %s\n", enteredMessage)
		}
	}
}

func recieveMessage(stream pb.ChatService_AddClientClient) {
	for {
		messageobj, err := stream.Recv()
		if err != nil {
			fmt.Printf("There was an error while receiving message: %v\n", err)
		}

		fmt.Println(messageobj.Message)
	}
}