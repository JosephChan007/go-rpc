package main

import (
	"github.com/JosephChan007/go-rpc/rpc/message"
	"github.com/JosephChan007/go-rpc/rpc/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	grpcServer := grpc.NewServer()
	message.RegisterOrderServiceServer(grpcServer, new(service.OrderServiceImpl))

	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("Order service failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Order service failed to serve: %v", err)
	}
}
