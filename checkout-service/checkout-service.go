package main

import (
	"checkout-service/protobuf/compiled"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	compiled.UnimplementedCheckoutServiceServer
}

func (s *server) CreateOrder(ctx context.Context, req *compiled.OrderRequest) (*compiled.OrderResponse, error) {
	return &compiled.OrderResponse{
		OrderId: 12345,
		Status:  "Order Created",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Gagal mendengarkan: %v", err)
	}

	grpcServer := grpc.NewServer()
	compiled.RegisterCheckoutServiceServer(grpcServer, &server{})

	fmt.Println("Checkout Service berjalan pada port 50053...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
