package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	desc_chat "github.com/Ananev-Alexandr/microservices/pkg/chat_api"

	desc_user "github.com/Ananev-Alexandr/microservices/pkg/auth_api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = ":50051"

type user_api_server struct {
	desc_user.UnimplementedUserAPIServer
}

type chat_api_server struct {
	desc_chat.UnimplementedChatAPIServer
}

func (s *user_api_server) Get(ctx context.Context, req *desc_user.GetRequest) (*desc_user.GetResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *user_api_server) Create(ctx context.Context, req *desc_user.CreateRequest) (*desc_user.CreateResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *user_api_server) Update(ctx context.Context, req *desc_user.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *user_api_server) Delete(ctx context.Context, req *desc_user.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// ---------------------------------------------------------------------------------------------------------------------------
func (s *chat_api_server) Create(ctx context.Context, req *desc_chat.CreateRequest) (*desc_chat.CreateResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *chat_api_server) Delete(ctx context.Context, req *desc_chat.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *chat_api_server) SendMessage(ctx context.Context, req *desc_chat.SendMessageRequest) (*emptypb.Empty, error) {
	return nil, errors.New("not implemented")
}

// ---------------------------------------------------------------------------------------------------------------------------

func main() {
	lis, err := net.Listen("tcp", grpcPort)

	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc_user.RegisterUserAPIServer(s, &user_api_server{})
	desc_chat.RegisterChatAPIServer(s, &chat_api_server{})
	fmt.Println("server listening at", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
