package main

import (
	"context"
	pb "github.com/wjl327/gobyexample/grpc/pb3/hello"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client := pb.NewHelloServiceClient(conn)
	rsp, err := client.SayHello(ctx, &pb.HelloReq{Name: "Jarvis"})
	if err != nil {
		log.Fatalf("Client SayHello error: %v", err)
	}
	log.Printf("Client SayHello Rece rsp: %s\n\n", rsp.GetMessage())

	stream, err := client.SayHelloStream(context.Background())
	stream.Send(&pb.HelloReq{Name: "Jarvis"})
	stream.CloseSend() //客户端关闭发送通道，即发送EOF，告诉服务端如果发完数据也可以结束链接。

	for {
		rsp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("Client recv server eof.")
				break
			}
			log.Fatalf("Client recv server err: %v", err)
		}
		log.Println("Client SayHelloStream Rece rsp: ", rsp.GetMessage())
	}

}
