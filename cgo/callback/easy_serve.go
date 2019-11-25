package main

import "C"
import "fmt"

//export EasyCallbackGo
func EasyCallbackGo() {
	fmt.Println("Go EasyServe recevied c callback...")
}
