package main

import (
	pb "Chitty-Chat/ChatService"
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

type client struct {
	clientName string
	stream pb.ChatService_AddClientServer
	errCh *chan error
}

var clients = make(map[string]client)
var mu sync.Mutex

func (s *server) AddClient(cs_bcs pb.ChatService_AddClientServer) error {
	errCh := make(chan error)
	messageobj, err := cs_bcs.Recv()
	if err != nil {
		log.Printf("Could not recieve initial message %v", err)
		errCh <- err
	}

	client := client{clientName: messageobj.Sender, stream: cs_bcs, errCh: &errCh}
	mu.Lock()
	clients[client.clientName] = client
	mu.Unlock()
	
	// Start a goroutine for the newly connected client
	// that listens for incomming messages from that client
	go recieveMessage(client)

	// Broadcast new client joined  message to all clients except the 
	// client where the message was sent from. 
	broadcastChatMessage(messageobj.Message, messageobj.Sender)

	return <-errCh
}

func (s *server) LeaveChat(ctx context.Context, in *pb.ClientName) (*pb.Empty, error) {
	mu.Lock()
	delete(clients, in.ClientName)
	mu.Unlock()

	message := fmt.Sprintf("[%s has left the chat]", in.ClientName)
	broadcastMessage(message)

	return &pb.Empty{}, nil
}

func recieveMessage(client client) {
	for {
		messageobj, err := client.stream.Recv()
		if err != nil {
			log.Printf("Could not recieve message %v\n", err)
			*client.errCh <- err
		}

		message := fmt.Sprintf("%s: %s", messageobj.Sender, messageobj.Message)
		// Broadcast the received message to all clients except 
		// the client that the message was sent from.
		broadcastChatMessage(message, messageobj.Sender)
	}
}

func broadcastChatMessage(message string, senderName string) {
	mu.Lock()
	for _, client := range clients {
		if client.clientName != senderName {
			go sendMessage(client, message)
		}
	}
	mu.Unlock()
}

func broadcastMessage(message string) {
	mu.Lock()
	for _, client := range clients {
		go sendMessage(client, message)
	}
	mu.Unlock()
}

func sendMessage(client client, message string) {
	err := client.stream.Send(&pb.ServerResponse{ Message: message })
	if err != nil {
		log.Printf("Could not send message to %s. %v\n", client.clientName, err)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	log.Printf("Server listening on %v\n", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})
	
	if err:= s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}