package main

import (
	"fmt"
	"github.com/rfyiamcool/grpcall"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"sync"
	"time"
)

func main() {
	//CallSayHello()
	CallSayHelloStream()
}

func CallSayHello() {
	grpcall.SetProtoSetFiles("${home}/developer/go_workspace/opensource/gobyexample/grpc/client/hello.protoset")
	err := grpcall.InitDescSource()
	if err != nil {
		panic(err.Error())
	}

	var handler = DefaultEventHandler{}
	grpcEnter, err := grpcall.New(
		grpcall.SetHookHandler(&handler),
	)
	grpcEnter.Init()

	body := `{"name": "Jarvis"}`
	res, err := grpcEnter.Call("127.0.0.1:1234", "hello.HelloService", "SayHello", body)
	if err != nil {
		log.Println("reply err: ", err)
		return
	}
	log.Printf("res data:%v \n", res.Data)
}

func CallSayHelloStream() {
	grpcall.SetProtoSetFiles("${home}/developer/go_workspace/opensource/gobyexample/grpc/client/hello.protoset")
	err := grpcall.InitDescSource()
	if err != nil {
		panic(err.Error())
	}

	var handler = DefaultEventHandler{}
	grpcEnter, err := grpcall.New(
		grpcall.SetHookHandler(&handler),
	)
	grpcEnter.Init()

	body := `{"name": "Jarvis"}`
	res, err := grpcEnter.Call("127.0.0.1:1234", "hello.HelloService", "SayHelloStream", body)
	if err != nil {
		log.Println("reply err: ", err)
		return
	}

	close := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			body := fmt.Sprintf(`{"name": "Jarvis %s"}`, time.Now().Format("2006-01-02 15:04:05"))
			res.SendChan <- []byte(body)
			select {
			case <-close:
				fmt.Println("client send done chan: ", err)
				wg.Done()
				return
			default:
				time.Sleep(2 * time.Second)
			}
		}
	}()

	go func() {
		for {
			select {
			case msg, ok := <-res.ResultChan:
				fmt.Println("client recv data:", msg, ok)
			case err := <-res.DoneChan:
				close <- true
				fmt.Println("client recv done chan: ", err)
				return
			}
		}
	}()

	wg.Wait()
}

type DefaultEventHandler struct {
	sendChan chan []byte
}

func (h *DefaultEventHandler) OnReceiveData(md metadata.MD, resp string, respErr error) {
}

func (h *DefaultEventHandler) OnReceiveTrailers(stat *status.Status, md metadata.MD) {
}
