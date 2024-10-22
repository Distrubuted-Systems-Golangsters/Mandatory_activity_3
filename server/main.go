package main

import (
	pb "Chitty-Chat/ChatService"
	gRPC_chat "Chitty-Chat/gRPC"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	log.Printf("Server listening on %v\n", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &gRPC_chat.Server{})
	
	if err:= s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}