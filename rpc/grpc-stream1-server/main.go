package main

import (
	"github.com/JosephChan007/go-rpc/rpc/message"
	"github.com/JosephChan007/go-rpc/rpc/service"
	"github.com/JosephChan007/go-rpc/rpc/util"
	"google.golang.org/grpc"
	"log"
	"net"
)

/**
 * 服务端stream
 */
func main() {

	// 双向数字签名证书
	creds := util.ServerCertHelper()

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
