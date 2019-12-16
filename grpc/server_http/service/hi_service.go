package service

import (
	"context"
	pb "github.com/wjl327/gobyexample/grpc/pb3/hi"
	"log"
)

type HiService struct{}

//收到一个请求回复一个响应
func (s *HiService) SayHi(ctx context.Context, req *pb.HiReq) (*pb.HiRsp, error) {
	log.Println("[HiService] SayHi received: ", req.GetName())
	return &pb.HiRsp{Message: "Hi," + req.GetName()}, nil
}
