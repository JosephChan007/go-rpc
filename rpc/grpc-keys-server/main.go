package main

import (
	"github.com/JosephChan007/go-rpc/rpc/message"
	"github.com/JosephChan007/go-rpc/rpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func main() {

	// 单向数字签名证书
	creds, err := credentials.NewServerTLSFromFile("keys-server/server.crt", "keys-server/server_no_password.key")
	if err != nil {
		log.Fatal(err)
	}

	// 用证书建立服务端
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	message.RegisterOrderServiceServer(grpcServer, new(service.OrderServiceImpl))

	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("Order service failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Order service failed to serve: %v", err)
	}
}
