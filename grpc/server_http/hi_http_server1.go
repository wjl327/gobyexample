package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/wjl327/gobyexample/grpc/pb3/hi"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	//单独启动grpcServerEndpoint服务，然后再启动grpc-gateway服务，好处是可以同时对外提供grpc端口和http端口服务。
	grpcServerEndpoint := "localhost:9090"
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterHiServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Println("grpc-gateway server error: ", err)
		return
	}

	log.Println("grpc-gateway server started!!!")
	http.ListenAndServe(":8081", mux)
}
