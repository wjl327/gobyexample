package main

import "C"

//export SayHello
func SayHello(s string) *C.char {
	rsp := "hello, " + s
	return C.CString(rsp)
}

func main() {}
