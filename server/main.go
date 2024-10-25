package main

import (
	pb "Chitty-Chat/grpc"
	"Chitty-Chat/server/util"
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedChatServiceServer
}

var mu sync.Mutex

func (s Server) AddClient(cs_bcs pb.ChatService_AddClientServer) error {
	errCh := make(chan error)
	messageobj, err := util.GetUserJoinedMessage(cs_bcs)

	if err != nil {
		log.Printf("Could not recieve initial message %v", err)
		errCh <- err
	}

	mu.Lock()
	util.LamportTimestamp++
	mu.Unlock()

	client := util.Client{ ClientName: messageobj.Sender, Stream: cs_bcs, ErrCh: &errCh }
	util.AddClientToMap(client)

	// Start a goroutine for the newly connected client
	// that listens for incomming messages from that client
	go util.RecieveMessages(client)

	// Broadcast new "client joined" message to all clients
	util.BroadcastMessage(messageobj.Message)

	return <-errCh
}

func (s Server) LeaveChat(ctx context.Context, in *pb.ClientName) (*pb.Empty, error) {
	util.RemoveClientFromMap(in.ClientName)
	util.LamportTimestamp = max(util.LamportTimestamp, in.Timestamp) + 1
	message := fmt.Sprintf("[%s has left the chat]", in.ClientName)
	util.BroadcastMessage(message)

	return &pb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	log.Printf("Server listening on %v\n", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &Server{})
	
	if err:= s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}