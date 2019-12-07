package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {

	addr := "0.0.0.0:8000"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("listen tcp: %s err:%v", addr, err)
		return
	}

	server := &http.Server{
		ReadTimeout:  time.Duration(2 * time.Second),
		WriteTimeout: time.Duration(2 * time.Second),
	}

	runner := &Runner{}
	server.Handler = runner
	if err = server.Serve(l); err != nil {
		log.Printf("listen l:%v server:%v err:%v", l, server, err)
		return
	}

}

type Runner struct{}

func (run *Runner) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recv uri:", r.RequestURI)
}
