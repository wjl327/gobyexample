package main

import (
	pb "github.com/wjl327/gobyexample/grpc/pb3/hi"
	"github.com/wjl327/gobyexample/grpc/server_http/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHiServiceServer(grpcServer, new(service.HiService))

	log.Println("grpc-endpoint server started!!!")
	grpcServer.Serve(listener)

}
