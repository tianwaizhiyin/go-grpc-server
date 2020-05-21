package main

import (
	"github.com/tianwaizhiyin/goTest.git/services"
	"google.golang.org/grpc"
	"net"
)

func main()  {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	list, _ := net.Listen("tcp", ":8080")
	rpcServer.Serve(list)
}

