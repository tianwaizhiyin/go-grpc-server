package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/tianwaizhiyin/go-grpc-server.git/helper"
	"github.com/tianwaizhiyin/go-grpc-server.git/services"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main()  {
	gwmux := runtime.NewServeMux()
	gRpcEndPoint := "localhost:8081"
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}
	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(),
		gwmux, gRpcEndPoint, opt)
	if err != nil {
		log.Fatal(err)
	}

	err = services.RegisterOrderServiceHandlerFromEndpoint(context.Background(),
		gwmux, gRpcEndPoint, opt)
	if err != nil {
		log.Fatal(err)
	}


	httpServer := &http.Server{
		Addr:":8080",
		Handler:gwmux,
	}
	httpServer.ListenAndServe()
}