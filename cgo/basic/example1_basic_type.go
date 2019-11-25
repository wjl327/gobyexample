package main

/*
#include <stdio.h>
#include <stdlib.h>

void PrintUser(const char* name, int age, int tags[], int tagNum) {
	printf("User name is :%s, age is %d.\nHis tag is :\n", name, age);
	for (int i = 0; i < tagNum; i++){
		printf("%d\n", tags[i]);
	}
	printf("\n");
}
*/
import "C"
import "unsafe"

func main() {
	cName := C.CString("Jarvis")
	defer C.free(unsafe.Pointer(cName))
	cAge := C.int(20)
	tags := []int32{1, 2, 3}
	cTags := (*C.int)(unsafe.Pointer(&tags[0]))
	cTagNum := C.int32_t(len(tags))
	C.PrintUser(cName, cAge, cTags, cTagNum)
}
