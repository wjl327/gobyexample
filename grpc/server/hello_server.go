package main

import (
	"context"
	pb "github.com/wjl327/gobyexample/grpc/pb3/hello"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type HelloService struct{}

//收到一个请求回复一个响应
func (s *HelloService) SayHello(ctx context.Context, req *pb.HelloReq) (*pb.HelloRsp, error) {
	log.Println("[HelloService] SayHello received: ", req.GetName())
	return &pb.HelloRsp{Message: "Hello," + req.GetName()}, nil
}

//每收到一个请求回复三个响应。
func (s *HelloService) SayHelloStream(stream pb.HelloService_SayHelloStreamServer) error {
	for {
		helloReq, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		log.Println("[HelloService] SayHelloStream received: ", helloReq.GetName())
		greets := []string{"Hi,", "Hello,", "你好,"}
		for _, greet := range greets {
			rsp := &pb.HelloRsp{Message: greet + helloReq.GetName()}
			err = stream.Send(rsp)
			if err != nil {
				return err
			}
		}
	} //这里for循环，只有当客户端CloseSend()或者客户端进程退出直接断开才会结束。
	return nil
}

func main() {

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, new(HelloService))
	grpcServer.Serve(listener)

}
