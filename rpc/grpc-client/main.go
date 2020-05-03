package main

import (
	"context"
	"github.com/JosephChan007/go-rpc/rpc/message"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial("hdfs-host3:9091", grpc.WithInsecure(), grpc.WithBlock())
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
