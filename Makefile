LOCAL_BIN:=$(CURDIR)/bin
PROTOC_GEN_GO:=bin\protoc-gen-go-grpc.exe
PROTOC_GEN_GO_GRPC:=bin\protoc-gen-go.exe

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

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
	
