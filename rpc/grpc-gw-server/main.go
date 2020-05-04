package main

func main() {
	/*
		// 以http形式对外提供接口
		// 由于以下gw文件创建命令报错，故将代码注释
		// protoc --grpc-gateway_out=logtostderr=true:. message.proto

		// 双向数字签名证书
		cred := util.ServerCertHelper()

		muxServer := runtime.NewServeMux()

		option := grpc.DialOption(grpc.WithTransportCredentials(cred))
		err := message.RegisterOrderServiceHandlerFromEndpoint(
			context.Background(),
			muxServer,

			// 注意: 此处必须用grpc服务端所在地址，若http接口服务端与grpc服务端部署在同一台机器，可以用:localhost:9090
			"hdfs-host3:9090",

			option
		)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}

		httpServer := &http.Server{
			Addr:    ":9091",	// 此处是http接口的端口
			Handler: muxServer,
		}
		httpServer.ListenAndServe()

	*/
}
