package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/wjl327/gobyexample/grpc/pb3/hi"
	"github.com/wjl327/gobyexample/grpc/server_http/service"
	"log"
	"net/http"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	//直接在grpc-gateway注册HiService，好处是简单，不过只有http端口。
	err := pb.RegisterHiServiceHandlerServer(ctx, mux, new(service.HiService))
	if err != nil {
		log.Println("grpc-gateway server error: ", err)
		return
	}

	log.Println("grpc-gateway server started!!!")
	http.ListenAndServe(":8081", mux)
}
