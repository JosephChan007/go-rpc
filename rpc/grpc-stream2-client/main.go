package main

import (
	"context"
	. "github.com/JosephChan007/go-rpc/rpc/message"
	"github.com/JosephChan007/go-rpc/rpc/util"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

/**
 * 双向stream
 */
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	status := []OrderStatus{1, 2, 0, 4}

	stream, err := c.GetOrderListByDoubleStream(ctx)
	if err != nil {
		log.Fatalf("Order client could not invoke1: %v", err)
	}

	for _, state := range status {

		// 发送请求
		req := &OrderStatusRequest{Status: []OrderStatus{state}}
		err := stream.Send(req)
		if err != nil {
			log.Printf("Order client could not invoke2: %v", err)
		}

		// 接受请求
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Order client could not invoke2: %v", err)
		}
		for _, order := range response.Orders {
			log.Printf("stream data is: %v", order)
		}
	}

}
