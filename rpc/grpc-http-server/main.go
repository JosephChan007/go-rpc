package main

import (
	"github.com/JosephChan007/go-rpc/rpc/message"
	"github.com/JosephChan007/go-rpc/rpc/service"
	"google.golang.org/grpc"
	"net/http"
)

func main() {

	grpcServer := grpc.NewServer()
	message.RegisterOrderServiceServer(grpcServer, new(service.OrderServiceImpl))

	muxServer := http.NewServeMux()
	muxServer.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		grpcServer.ServeHTTP(writer, request)
	})
	httpServer := &http.Server{
		Addr:    ":9091",
		Handler: muxServer,
	}
	httpServer.ListenAndServe()
}
