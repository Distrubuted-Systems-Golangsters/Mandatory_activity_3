package util

import (
	pb "Chitty-Chat/ChatService"
	"fmt"
	"log"
	"sync"
)

var Clients = make(map[string]Client)
var mu sync.Mutex

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

		message := fmt.Sprintf("%s: %s", messageobj.Sender, messageobj.Message)
		BroadcastChatMessage(message, messageobj.Sender)
	}
}

func BroadcastChatMessage(message string, senderName string) {
	// Broadcast the received message to all clients except 
	// the client that the message was sent from.
	mu.Lock()
	for _, client := range Clients {
		if client.ClientName != senderName {
			go SendMessage(client, message)
		}
	}
	mu.Unlock()
}

func BroadcastMessage(message string) {
	// Broadcast the received message to all clients
	mu.Lock()
	for _, client := range Clients {
		go SendMessage(client, message)
	}
	mu.Unlock()
}

func SendMessage(client Client, message string) {
	err := client.Stream.Send(&pb.ServerResponse{ Message: message })
	if err != nil {
		log.Printf("Could not send message to %s. %v\n", client.ClientName, err)
	}
}