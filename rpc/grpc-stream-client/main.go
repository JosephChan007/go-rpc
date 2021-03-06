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
 * 服务端stream
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

	stream, err := c.GetOrderListByServerStream(ctx, &OrderStatusRequest{
		Status: []OrderStatus{OrderStatus_UnPay, OrderStatus_Paied, OrderStatus_Refunded},
	})
	if err != nil {
		log.Fatalf("Order client could not invoke: %v", err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Order client stream could not read: %v", err)
		}
		for _, order := range response.Orders {
			log.Printf("stream data is: %v", order)
		}
	}

}
