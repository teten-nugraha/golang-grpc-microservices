package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"user-service/protobuf/compiled"
)

type server struct {
	compiled.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *compiled.UserRequest) (*compiled.UserResponse, error) {
	return &compiled.UserResponse{
		UserId: req.UserId,
		Name:   fmt.Sprintf("User %d", req.UserId),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Gagal mendengarkan : %v", err)
	}

	grpcServer := grpc.NewServer()
	compiled.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("User Service berjalan pada port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
