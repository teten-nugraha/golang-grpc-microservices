package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"product-service/protobuf/compiled"
)

type server struct {
	compiled.UnimplementedProductServiceServer
}

func (s *server) GetProduct(ctx context.Context, req *compiled.ProductRequest) (*compiled.ProductResponse, error) {
	return &compiled.ProductResponse{
		ProductId: req.ProductId,
		Name:      fmt.Sprintf("Product %d", req.ProductId),
		Price:     99.99,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Gagal mendengarkan: %v", err)
	}

	grpcServer := grpc.NewServer()
	compiled.RegisterProductServiceServer(grpcServer, &server{})

	fmt.Println("Product Service berjalan pada port 50052...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
