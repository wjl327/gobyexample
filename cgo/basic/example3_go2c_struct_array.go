package main

/*
#include <stdio.h>
#include <stdlib.h>

typedef struct User {
	char *name;
    int age;
} User;

void PrintUser(const User* user) {
	printf("User name is :%s, age is %d.\n", user->name, user->age);
}

void PrintUsers(const User users[], int nums) {
	for (int i = 0; i < nums; i++) {
		printf("User[%d] name is :%s, age is %d.\n", i, users[i].name, users[i].age);
	}
}

*/
import "C"
import (
	"strconv"
	"unsafe"
)

func main() {
	callPrintUser()
	callPrintUsers()
}

func callPrintUser() {
	user := C.User{}
	user.name = C.CString("ZhangSan")
	user.age = C.int(18)
	defer C.free(unsafe.Pointer(user.name))
	C.PrintUser(&user)
}

func callPrintUsers() {
	users := make([]C.User, 0)
	nums := 3
	for i := 0; i < nums; i++ {
		user := C.User{}
		user.name = C.CString("user-" + strconv.Itoa(i))
		user.age = C.int(20 + i)
		users = append(users, user)
	}
	C.PrintUsers((*C.User)(unsafe.Pointer(&users[0])), C.int(nums))
	for i := 0; i < nums; i++ {
		C.free(unsafe.Pointer(users[i].name))
	}
}
