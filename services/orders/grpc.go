package main

import (
	"log"
	"net"

	handler "github.com/nitheshkv96/kitchen-microservice/services/orders/handler/orders"
	service "github.com/nitheshkv96/kitchen-microservice/services/orders/service"
	grpc "google.golang.org/grpc"
)

//*********Pattern***********

// Struct
type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("Failed to lisen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// register our grpc services
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}
