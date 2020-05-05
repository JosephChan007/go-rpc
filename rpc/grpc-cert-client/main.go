package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/JosephChan007/go-rpc/rpc/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"time"
)

func main() {

	// 双向数字签名证书
	cert, err := tls.LoadX509KeyPair("cert-client/client.pem", "cert-client/client.key")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("cert-client/ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "hdfs-host3",
		RootCAs:      certPool,
	})

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
		OrderId: "20200501001",
	})
	if err != nil {
		log.Fatalf("Order client could not invoke: %v", err)
	}
	log.Printf("Order info is: %v", r)
}
