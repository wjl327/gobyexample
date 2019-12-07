package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	Get()
	Post()
}

func Post() {
	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/hello/greet",
		strings.NewReader("name=xiaoMi&age=20"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	client := http.Client{}
	rsp, err := client.Do(req)
	defer rsp.Body.Close()

	headers := rsp.Header
	fmt.Println(headers)
	fmt.Printf("rsp code: %d\n", rsp.StatusCode)
	fmt.Printf("rsp content length: %d\n", rsp.ContentLength)
	content, _ := ioutil.ReadAll(rsp.Body)
	fmt.Printf("rsp content: %s\n", content)
}

func Get() {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000/hello/greet?name=abc", nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
	req.Header.Set("auth_key", "930rj2k393r2klk3rkdijfemnkkl323ksere8rh")

	client := http.Client{}
	rsp, err := client.Do(req)
	defer rsp.Body.Close()

	headers := rsp.Header
	fmt.Println(headers)
	fmt.Printf("rsp code: %d\n", rsp.StatusCode)
	fmt.Printf("rsp content length: %d\n", rsp.ContentLength)
	content, _ := ioutil.ReadAll(rsp.Body)
	fmt.Printf("rsp content: %s\n", content)
}

func SimpleGet() {
	rsp, err := http.Get("http://127.0.0.1:8000/hello/greet?name=abc")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer rsp.Body.Close()

	headers := rsp.Header
	fmt.Println(headers)
	fmt.Printf("rsp code: %d\n", rsp.StatusCode)
	fmt.Printf("rsp content length: %d\n", rsp.ContentLength)
	content, _ := ioutil.ReadAll(rsp.Body)
	fmt.Printf("rsp content: %s\n", content)
}
