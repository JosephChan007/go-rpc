package main

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/JosephChan007/go-rpc/rpc/message"
	"github.com/JosephChan007/go-rpc/rpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
)

func main() {

	// 双向数字签名证书
	cert, err := tls.LoadX509KeyPair("cert-server/server.pem", "cert-server/server.key")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("cert-server/ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAnyClientCert,
		ClientCAs:    certPool,
	})

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
