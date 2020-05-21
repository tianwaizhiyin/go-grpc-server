package main

import (
	"github.com/tianwaizhiyin/go-grpc-server.git/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func main()  {
	creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server_no_password.key")
	if err != nil {
		log.Fatal(err)
	}
	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	list, _ := net.Listen("tcp", ":8080")
	rpcServer.Serve(list)
}

