package user

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	auth "github.com/Ananev-Alexandr/microservices/internal/auth"
	desc_user "github.com/Ananev-Alexandr/microservices/pkg/auth_api"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	desc_user.UnimplementedUserAPIServer
	db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{db: db}
}

func (s *Service) Get(ctx context.Context, req *desc_user.GetRequest) (*desc_user.GetResponse, error) {
	builderSelectOne := sq.Select("id", "name", "email", "role", "created_at", "updated_at").
		From("user_").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.Id})

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Printf("failed to build select query: %v", err)
		return nil, err
	}

	var (
		id          int64
		name, email string
		role        int32
		createdAt   time.Time
		updatedAt   sql.NullTime
	)

	err = s.db.QueryRow(ctx, query, args...).Scan(&id, &name, &email, &role, &createdAt, &updatedAt)
	if err != nil {
		log.Printf("failed to select user: %v", err)
		return nil, err
	}

	user := desc_user.GetResponse{
		Id:        id,
		Name:      name,
		Email:     email,
		Role:      desc_user.Role(role),
		CreatedAt: timestamppb.New(createdAt),
	}

	if updatedAt.Valid {
		user.UpdatedAt = timestamppb.New(updatedAt.Time)
	}

	return &user, nil
}

func (s *Service) Create(ctx context.Context, req *desc_user.CreateRequest) (*desc_user.CreateResponse, error) {
	if req.Name == "" || req.Email == "" {
		return nil, errors.New("have empty field")
	}

	if auth.InvalidPass(req.Password, req.PasswordConfirm) {
		return nil, errors.New("other password")
	}
	hash_password, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("cannot hash password")
	}
	builderInsert := sq.Insert("user_").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "password", "role").
		Values(req.Name, req.Email, hash_password, req.Role).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var userID int64
	err = s.db.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		log.Fatalf("failed to insert user: %v", err)
	}

	log.Printf("inserted user with id: %d", userID)

	return &desc_user.CreateResponse{
		Id: userID,
	}, nil
}

func (s *Service) Update(ctx context.Context, req *desc_user.UpdateRequest) (*emptypb.Empty, error) {
	builderUpdateUser := sq.Update("user_").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.Id})

	if req.Name != nil && req.Name.Value != "" {
		builderUpdateUser = builderUpdateUser.Set("name", req.Name.Value)
	}
	if req.Email != nil && req.Email.Value != "" {
		builderUpdateUser = builderUpdateUser.Set("email", req.Email.Value)
	}

	builderUpdateUser = builderUpdateUser.Set("updated_at", time.Now())

	query, args, err := builderUpdateUser.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}
	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to execute update query: %v", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Service) Delete(ctx context.Context, req *desc_user.DeleteRequest) (*emptypb.Empty, error) {
	builderInsert := sq.Delete("user_").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.Id})

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to build delete query: %v", err)
		return nil, err
	}
	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to delete user: %v", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
