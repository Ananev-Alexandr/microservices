package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	desc_user "github.com/Ananev-Alexandr/microservices/pkg/auth_api"
	desc_chat "github.com/Ananev-Alexandr/microservices/pkg/chat_api"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Ananev-Alexandr/microservices/internal/chat"
	"github.com/Ananev-Alexandr/microservices/internal/config"
	"github.com/Ananev-Alexandr/microservices/internal/user"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "../../.env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	// Считываем переменные окружения
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	// Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)

	userService := user.NewService(pool)
	chatService := chat.NewService(pool)

	desc_user.RegisterUserAPIServer(s, userService)
	desc_chat.RegisterChatAPIServer(s, chatService)

	fmt.Println("server listening at", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
