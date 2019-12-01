package main

import (
	"context"
	pb "github.com/wjl327/gobyexample/grpc/pb3/hello_bak"
	"google.golang.org/grpc"
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
	rsp, err := client.SayHello(ctx, &pb.HelloReq{Name: "WJL"})
	if err != nil {
		log.Fatalf("Client SayHello error: %v", err)
	}
	log.Println("Client SayHello Rece rsp: ", rsp.GetMessage())
}
