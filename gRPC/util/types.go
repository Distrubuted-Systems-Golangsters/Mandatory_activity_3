package util

import (
	pb "Chitty-Chat/ChatService"
)

type Client struct {
	ClientName string
	Stream pb.ChatService_AddClientServer
	ErrCh *chan error
}