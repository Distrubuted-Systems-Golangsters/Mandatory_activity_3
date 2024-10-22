package gRPC_chat

import (
	pb "Chitty-Chat/ChatService"
	"Chitty-Chat/gRPC/util"
	"context"
	"fmt"
	"log"
)

type Server struct {
	pb.UnimplementedChatServiceServer
}

func (s Server) AddClient(cs_bcs pb.ChatService_AddClientServer) error {
	errCh := make(chan error)
	messageobj, err := util.GetUserJoinedMessage(cs_bcs)

	if err != nil {
		log.Printf("Could not recieve initial message %v", err)
		errCh <- err
	}

	client := util.Client{ ClientName: messageobj.Sender, Stream: cs_bcs, ErrCh: &errCh }
	util.AddClientToMap(client)

	// Start a goroutine for the newly connected client
	// that listens for incomming messages from that client
	go util.RecieveMessages(client)

	// Broadcast new "client joined" message to all clients except the
	// client where the message was sent from.
	util.BroadcastChatMessage(messageobj.Message, messageobj.Sender)

	return <-errCh
}

func (s Server) LeaveChat(ctx context.Context, in *pb.ClientName) (*pb.Empty, error) {
	util.RemoveClientFromMap(in.ClientName)
	message := fmt.Sprintf("[%s has left the chat]", in.ClientName)
	util.BroadcastMessage(message)

	return &pb.Empty{}, nil
}