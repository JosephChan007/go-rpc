package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type MathUtil struct {
}

func (m *MathUtil) CircleArea(radius float32, area *float32) error {
	*area = 3.14 * radius * radius
	fmt.Printf("Caculate a circle area, radius: %f, result: %f\n", radius, *area)
	return nil
}

func (m *MathUtil) Response() error {
	fmt.Printf("Response a pong....")
	return nil
}

func main() {
	m := new(MathUtil)
	err := rpc.Register(m)
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
