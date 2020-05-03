package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	client, err := rpc.DialHTTP("tcp", "hdfs-host3:9090")
	if err != nil {
		fmt.Printf("Client dial error: %v", err)
	}

	radius := 3.02
	var result float32
	err = client.Call("MathUtil.CircleArea", radius, &result)
	if err != nil {
		fmt.Printf("Client call error: %v", err)
	}

	var result1 float32
	clientCall := client.Go("MathUtil.CircleArea", radius, &result1, nil)
	if reply := <-clientCall.Done; reply.Error != nil {
		fmt.Printf("Async Client call error: %v", reply.Error)
	}

	fmt.Printf("Sync Circle area is %f\n", result)
	fmt.Printf("Async Circle area is %f\n", result1)
}
