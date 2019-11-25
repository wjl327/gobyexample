package main

/*
#cgo CFLAGS: -I./src
#cgo LDFLAGS: -L${SRCDIR}/libs -lhello
#include "hello.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

/**
  本例主要展示Go调用C静态库的例子。
  备注：CFLAGS：设置编译参数。-I指定编译头文件目录
       LDFLAGS：设置链接参数。-L指定链接库文件的目录。-l表示需要链接的库。

  工作目录：static_lib
  编译静态库：gcc -c src/hello.c
  编译静态库：ar -rcs libs/libhello.a hello.o
  编译Go程序：go build -o out
*/
func main() {
	cs := C.CString("HelloStaticLib CGO!\n")
	defer C.free(unsafe.Pointer(cs))
	C.SayHello(cs)
}
