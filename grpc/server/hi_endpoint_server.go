package main

import (
	"context"
	pb "github.com/wjl327/gobyexample/grpc/pb3/hi"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HiService struct{}

//收到一个请求回复一个响应
func (s *HiService) SayHi(ctx context.Context, req *pb.HiReq) (*pb.HiRsp, error) {
	log.Println("[HiService] SayHi received: ", req.GetName())
	return &pb.HiRsp{Message: "Hi," + req.GetName()}, nil
}

func main() {

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHiServiceServer(grpcServer, new(HiService))
	grpcServer.Serve(listener)

}
