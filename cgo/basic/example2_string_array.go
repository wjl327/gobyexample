package main

/*
#include <stdio.h>
#include <stdlib.h>

void PrintBook(char* books[], int num) {
	for (int i = 0; i < num; i++){
		printf("Book:《%s》\n", books[i]);
	}
}
*/
import "C"
import "unsafe"

func main() {
	goBooks := []string{"The Godfather", "The Old Man and the Sea", "War and Peace"}
	cBooks := make([]*C.char, len(goBooks))
	for i, book := range goBooks {
		cBooks[i] = C.CString(book)
	}
	defer func() {
		for i := range cBooks {
			C.free(unsafe.Pointer(cBooks[i]))
		}
	}()
	C.PrintBook((**C.char)(unsafe.Pointer(&cBooks[0])), C.int(len(cBooks)))
}
