package user

import (
	"context"
	"errors"
	"log"

	desc_user "github.com/Ananev-Alexandr/microservices/pkg/auth_api"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	desc_user.UnimplementedUserAPIServer
	db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{db: db}
}

func (s *Service) Get(ctx context.Context, req *desc_user.GetRequest) (*desc_user.GetResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *Service) Create(ctx context.Context, req *desc_user.CreateRequest) (*desc_user.CreateResponse, error) {
	// log.Printf("Create called with: %v", req)
	if req.Name == "" || req.Email == "" {
		return nil, errors.New("have empty field")
	}

	if InvalidPass(req.Password, req.PasswordConfirm) {
		return nil, errors.New("other password")
	}
	hash_password, err := HashPassword(req.Password)
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
	return &emptypb.Empty{}, nil
}

func (s *Service) Delete(ctx context.Context, req *desc_user.DeleteRequest) (*emptypb.Empty, error) {
	// log.Printf("Delete called with: %v", req)
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

	// Возвращаем пустой ответ после успешного удаления
	return &emptypb.Empty{}, nil
}

func InvalidPass(password, password_confirm string) bool {
	return password != password_confirm
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
