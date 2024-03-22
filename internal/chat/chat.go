package chat

import (
	"context"
	"errors"

	desc_chat "github.com/Ananev-Alexandr/microservices/pkg/chat_api"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	desc_chat.UnimplementedChatAPIServer
	db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{db: db}
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
