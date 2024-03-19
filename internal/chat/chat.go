package chat

import (
	"context"
	"errors"

	desc_chat "github.com/Ananev-Alexandr/microservices/pkg/chat_api"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	desc_chat.UnimplementedChatAPIServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Create(ctx context.Context, req *desc_chat.CreateRequest) (*desc_chat.CreateResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *Service) Delete(ctx context.Context, req *desc_chat.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Service) SendMessage(ctx context.Context, req *desc_chat.SendMessageRequest) (*emptypb.Empty, error) {
	return nil, errors.New("not implemented")
}
