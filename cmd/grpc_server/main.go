package main

import (
	"fmt"
	"net"

	desc_user "github.com/Ananev-Alexandr/microservices/pkg/auth_api"
	desc_chat "github.com/Ananev-Alexandr/microservices/pkg/chat_api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Ananev-Alexandr/microservices/internal/chat"
	"github.com/Ananev-Alexandr/microservices/internal/user"
)

const grpcPort = ":50051"

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	s := grpc.NewServer()
	reflection.Register(s)

	userService := user.NewService()
	chatService := chat.NewService()

	desc_user.RegisterUserAPIServer(s, userService)
	desc_chat.RegisterChatAPIServer(s, chatService)

	fmt.Println("server listening at", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
