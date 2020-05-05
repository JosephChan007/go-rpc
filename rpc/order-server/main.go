package main

import (
	"github.com/JosephChan007/go-rpc/rpc/service"
	"net"
	"net/http"
	"net/rpc"
)

func main() {

	s := new(service.OrderServiceImpl)

	err := rpc.Register(s)
	if err != nil {
		panic(err)
	}

	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	http.Serve(listen, nil)

}
