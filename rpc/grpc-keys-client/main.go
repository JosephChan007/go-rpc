package main

import (
	"context"
	"github.com/JosephChan007/go-rpc/rpc/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

func main() {

	// 单向数字签名证书
	creds, err := credentials.NewClientTLSFromFile("keys-client/server.crt", "hdfs-host3")
	if err != nil {
		log.Fatal(err)
	}

	// 用证书调用服务端
	conn, err := grpc.Dial("hdfs-host3:9091", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := message.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetOrder(ctx, &message.OrderRequest{
		OrderId:   "20200501002",
		Timestamp: 0,
	})
	if err != nil {
		log.Fatalf("Order client could not invoke: %v", err)
	}
	log.Printf("Order info is: %v", r)
}
