package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/wjl327/gobyexample/grpc/pb3/hi"
	"log"
	"net/http"
)

type HiService2 struct{}

//收到一个请求回复一个响应
func (s *HiService2) SayHi(ctx context.Context, req *pb.HiReq) (*pb.HiRsp, error) {
	log.Println("[HiService] SayHi received: ", req.GetName())
	return &pb.HiRsp{Message: "Hi," + req.GetName()}, nil
}

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	//方式一：单独启动grpcServerEndpoint实例，好处是可以同时对外提供grpc端口和http端口服务。
	//grpcServerEndpoint := "localhost:9090"
	//opts := []grpc.DialOption{grpc.WithInsecure()}
	//err := gw.RegisterHiServiceHandlerFromEndpoint(ctx, mux,  grpcServerEndpoint, opts)

	//方式二：直接在grpc-gateway注册HiService，好处是简单，不过只有http端口。
	err := pb.RegisterHiServiceHandlerServer(ctx, mux, new(HiService2))
	if err != nil {
		log.Println("[HttpServer] error: ", err)
		return
	}

	log.Println("grpc-gateway server started!!!")
	http.ListenAndServe(":8081", mux)
}
