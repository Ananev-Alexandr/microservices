FROM golang:1.22.0-alpine AS builder

COPY . /github.com/Ananev-Alexandr/microservices
WORKDIR /github.com/Ananev-Alexandr/microservices

RUN go mod download
RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Ananev-Alexandr/microservices/bin/crud_server .

CMD ["./crud_server"]