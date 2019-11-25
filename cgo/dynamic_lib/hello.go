package main

/*
#cgo CFLAGS: -I./src
#cgo LDFLAGS: -L${SRCDIR}/libs -lhello
#include "hello.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

//Go调用C 动态库例子
//备注：CFLAGS：设置编译参数。-I指定编译头文件目录
//备注：LDFLAGS：设置链接参数。-L指定链接库文件的目录。-l表示需要链接的库。
//编译动态库：gcc src/main.c -shared -fPIC -o libs/libhello.so
//编译Go程序：go build -o out

/**
  本例主要展示Go调用C动态库的例子。
  备注：CFLAGS：设置编译参数。-I指定编译头文件目录
       LDFLAGS：设置链接参数。-L指定链接库文件的目录。-l表示需要链接的库。

  工作目录：dynamic_lib
  编译动态库：gcc src/hello.c -shared -fPIC -o libs/libhello.so
  编译Go程序：go build -o out
*/
func main() {
	cs := C.CString("HelloLib CGO!\n")
	defer C.free(unsafe.Pointer(cs))
	C.SayHello(cs)
}
