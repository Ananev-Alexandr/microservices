package user

import (
	"context"
	"errors"

	desc_user "github.com/Ananev-Alexandr/microservices/pkg/auth_api"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	desc_user.UnimplementedUserAPIServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Get(ctx context.Context, req *desc_user.GetRequest) (*desc_user.GetResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *Service) Create(ctx context.Context, req *desc_user.CreateRequest) (*desc_user.CreateResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *Service) Update(ctx context.Context, req *desc_user.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Service) Delete(ctx context.Context, req *desc_user.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
