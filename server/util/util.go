package util

import (
	pb "Chitty-Chat/grpc"
	"fmt"
	"log"
	"sync"
)

var Clients = make(map[string]Client)
var mu sync.Mutex
var LamportTimestamp int64 = 0

func AddClientToMap(client Client) {
	mu.Lock()
	Clients[client.ClientName] = client
	mu.Unlock()
}

func RemoveClientFromMap(clientName string) {
	mu.Lock()
	delete(Clients, clientName)
	mu.Unlock()
}

func GetUserJoinedMessage(cs_bcs pb.ChatService_AddClientServer) (*pb.ChatMessageClient, error) {
	messageobj, err := cs_bcs.Recv()
	return messageobj, err
}

func RecieveMessages(client Client) {
	for {
		messageobj, err := client.Stream.Recv()

		if err != nil {
			log.Printf("Could not recieve message %v\n", err)
			*client.ErrCh <- err
		}

		mu.Lock()
		LamportTimestamp = max(LamportTimestamp, messageobj.Timestamp) + 1
		mu.Unlock()

		message := fmt.Sprintf("%s: %s", messageobj.Sender, messageobj.Message)
		BroadcastMessage(message)
	}
}

func BroadcastMessage(message string) {
	// Broadcast the received message to all clients
	mu.Lock()
	LamportTimestamp++;
	for _, client := range Clients {
		go SendMessage(client, message)
	}
	mu.Unlock()
}

func SendMessage(client Client, message string) {
	err := client.Stream.Send(&pb.ServerResponse{ Message: message, Timestamp: LamportTimestamp })
	if err != nil {
		log.Printf("Could not send message to %s. %v\n", client.ClientName, err)
	}
}