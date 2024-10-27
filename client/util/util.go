package util

import (
	pb "Chitty-Chat/grpc"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var LamportTimestamp int64 = 0

func StartApp() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Your Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read from console: %v\n", err)
	}
	name = strings.Trim(name, "\r\n")

	return name
}

func SendMessages(stream pb.ChatService_AddClientClient, clientName string, client pb.ChatServiceClient) {
	for {
		reader := bufio.NewReader(os.Stdin)
		enteredMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read from console: %v\n", err)
		}
		enteredMessage = strings.Trim(enteredMessage, "\r\n")

		if len(enteredMessage) > 128 || !utf8.ValidString(enteredMessage) {
			log.Println("!!! Invalid string !!!")
			continue
		} 

		LamportTimestamp++

		if enteredMessage == "quit" {
			log.Printf("Client leaving chat... { Timestamp: %d }\n", LamportTimestamp)
			client.LeaveChat(context.Background(), &pb.ClientName{ ClientName: clientName, Timestamp: LamportTimestamp })
			break
		}

		// Clear enteredMessage in console
		fmt.Print("\033[A\033[2K")
		log.Printf("Publishing message... { Timestamp: %d }\n", LamportTimestamp)
		sendErr := stream.Send(&pb.ChatMessageClient{ Sender: clientName, Message: enteredMessage, Timestamp: LamportTimestamp })
		if sendErr != nil {
			fmt.Printf("There was an error while sending message to server %v\n", sendErr)
		}
	}
}

func RecieveMessages(stream pb.ChatService_AddClientClient) {
	for {
		messageobj, err := stream.Recv()
		if err != nil {
			fmt.Printf("There was an error while receiving message: %v\n", err)
		}

		LamportTimestamp = max(LamportTimestamp, messageobj.Timestamp) + 1

		log.Printf("%s { Timestamp: %d }\n", messageobj.Message, LamportTimestamp)
	}
}

func CreateClientServerConnection() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:8080", opts...)
	return conn, err
}

func SendInitialMessage(stream grpc.BidiStreamingClient[pb.ChatMessageClient, pb.ServerResponse], clientName string) error {
	message := fmt.Sprintf("[%s has joined the chat]", clientName)
	LamportTimestamp++
	log.Printf("Client joining chat... { Timestamp: %d }", LamportTimestamp)
	err := stream.Send(&pb.ChatMessageClient{ Sender: clientName, Message: message, Timestamp: LamportTimestamp })

	return err
}