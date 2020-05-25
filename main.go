package main

import (
	"github.com/tianwaizhiyin/go-grpc-server.git/helper"
	"github.com/tianwaizhiyin/go-grpc-server.git/services"
	"google.golang.org/grpc"
	"net"
)

func main()  {
	//creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server_no_password.key")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//tls证书认证
	//cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	//certPool := x509.NewCertPool()
	//ca, _ := ioutil.ReadFile("cert/ca.pem")
	//certPool.AppendCertsFromPEM(ca)
	//
	//creds := credentials.NewTLS(&tls.Config{
	//	Certificates: []tls.Certificate{cert}, //服务端证书
	//	ClientAuth: tls.RequireAndVerifyClientCert,
	//	ClientCAs: certPool,
	//})


	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService)) //商品服务
	services.RegisterOrderServiceServer(rpcServer,new(services.OrdersService)) //订单服务
	services.RegisterUserServiceServer(rpcServer, new(services.UserService)) //用户服务

	list, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(list)

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Println(request)
	//	rpcServer.ServeHTTP(writer,request)
	//})
	//httpServer := &http.Server{
	//	Addr:":8080",
	//	Handler:mux,
	//}
	//httpServer.ListenAndServeTLS("keys/server.crt", "keys/server_no_password.key")
}