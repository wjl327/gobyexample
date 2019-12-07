package main

import (
	"fmt"
	"net/http"
)

func main() {

	addr := "0.0.0.0:8000"

	// Use DefaultServeMux

	//方式一：
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Recv uri:%s, Welcome to Go!", r.RequestURI)))
	}
	http.Handle("/", http.HandlerFunc(fn))
	http.ListenAndServe(addr, nil)

	//方式二：
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("Recv uri:", r.RequestURI)
	//})
	//http.ListenAndServe(addr, nil)

	// Use Custom ServeMux
	//mux := http.NewServeMux()
	//mux.Handle("/hello", &GreetHandler{})
	//mux.Handle("/greet", &GreetHandler{})
	//mux.HandleFunc("/", nil)
	//http.ListenAndServe(addr, mux)

}

type GreetHandler struct{}

func (h *GreetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Recv uri:%s, Welcome to Go!", r.RequestURI)))
}
