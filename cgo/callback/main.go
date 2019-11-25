package main

/*
#cgo CFLAGS: -I./manager
#cgo LDFLAGS: -L${SRCDIR}/libs -lmanager
#include "registerCenter.h"
#include "registerCallback.h"
#include <stdlib.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

/**
  Go->C->Go 综合例子。
  Go调用C的方法注册一个回调函数，回调函数通过CGo又调用到Go。

  1、manager包下的属于C的部分，先编译成库libmanager.so。
  2、由于manager接口参数为函数指针，因此，在Go中调用它时也需要传个C函数，对应registerCallback.h和registerCallback.c部分。
  3、Go和CGo，则是main.go和easy_serve.go。

  step1:gcc manager/registerCenter.c -shared -fPIC -o libs/libmanager.so
  step2:go build -o out
*/
func main() {

	easy := C.CString("easy")
	defer C.free(unsafe.Pointer(easy))

	//先往调用C函数去注册回调
	f := C.NotifyFunc(C.EasyCallback)
	C.setNotifyFunc(easy, f)

	//调用C函数去触发回调通知
	C.callback(easy)

}
