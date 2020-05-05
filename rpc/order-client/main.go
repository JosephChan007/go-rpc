package main

import (
	"fmt"
	"github.com/JosephChan007/go-rpc/rpc/message"
	"net/rpc"
)

func main() {

	client, err := rpc.DialHTTP("tcp", "hdfs-host3:9090")
	if err != nil {
		fmt.Printf("Client dial error: %v\n", err)
	}

	var res message.OrderInfo
	req := &message.OrderRequest{OrderId: "20200501009"}
	err = client.Call("OrderService.GetOrder", req, &res)
	if err != nil {
		fmt.Printf("Client call error: %v\n", err)
		return
	}

	fmt.Printf("Sync order info is: orderid=%s | ordername=%s | orderstatus=%s\n", res.OrderId, res.OrderName, res.Status)

}
