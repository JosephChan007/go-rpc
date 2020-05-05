package main

import (
	"context"
	. "github.com/JosephChan007/go-rpc/rpc/message"
	"github.com/JosephChan007/go-rpc/rpc/util"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {

	// 双向数字签名证书
	creds := util.ClientCertHelper()

	// 用证书调用服务端
	conn, err := grpc.Dial("hdfs-host3:9091", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.GetOrder(ctx, &OrderRequest{
		OrderId: "20200501002",
	})
	if err != nil {
		log.Fatalf("Order client could not invoke: %v", err)
	}
	log.Printf("Order info is: %v", res)
}
