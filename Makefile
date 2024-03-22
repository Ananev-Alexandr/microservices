include .env
LOCAL_BIN:=$(CURDIR)/bin
PROTOC_GEN_GO:=bin\protoc-gen-go-grpc.exe
PROTOC_GEN_GO_GRPC:=bin\protoc-gen-go.exe

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

check-path:
	@echo %PATH%

generate:
	make generate-auth-api
	make generate-chat-api

generate-auth-api:
	mkdir pkg\\auth_api
	protoc --proto_path=api/auth \
	--go_out=pkg/auth_api --go_opt=paths=source_relative \
	--plugin=protoc-gen-go="$(PROTOC_GEN_GO)" \
	--go-grpc_out=pkg/auth_api --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc="$(PROTOC_GEN_GO_GRPC)" \
	api/auth/user_api.proto

generate-chat-api:
	mkdir pkg\\chat_api
	protoc --proto_path=api/chat-server \
	--go_out=pkg/chat_api --go_opt=paths=source_relative \
	--plugin=protoc-gen-go="$(PROTOC_GEN_GO)" \
	--go-grpc_out=pkg/chat_api --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc="$(PROTOC_GEN_GO_GRPC)" \
	api/chat-server/chat_api.proto

build:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/grpc_server/main.go

copy-to-server:
	scp service_linux root@10.10.19.111:


docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t .../test-server:v0.0.1 .
	docker login -u token -p ... ...
	docker push .../test-server:v0.0.1


local-migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} down -v