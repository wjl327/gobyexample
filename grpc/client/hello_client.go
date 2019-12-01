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
	stream.CloseSend() //客户端发完数据要关闭!(即发送EOF)!

	for {
		rsp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		log.Println("Client SayHelloStream Rece rsp: ", rsp.GetMessage())
	}

}
