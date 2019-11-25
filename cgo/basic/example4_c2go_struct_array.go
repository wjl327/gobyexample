package main

/*
#include <stdio.h>
#include <stdlib.h>

int ids[] = {1, 2, 3, 4};
const int* GetIds() {
	return &ids[0];
}

typedef struct User {
	int id;
	char *name;
} User;

typedef struct UserGroup {
	int num;
	User *users;
} UserGroup;

const UserGroup* GetUserGroup() {
	UserGroup *userGroup = (UserGroup*)malloc(sizeof(UserGroup));
	userGroup->num = 3;
	User *us = (User*)malloc(sizeof(User) * userGroup->num);
	userGroup->users = us;
	for (int i = 0; i < userGroup->num; i++){
		User user = {i, "YoYo"};
		us[i] = user;
	}
	return userGroup;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	printIds()
	printUsers()
}

func printIds() {
	cIds := C.GetIds()
	goIds := (*[1 << 31]C.int)(unsafe.Pointer(cIds))[0:4:4]
	fmt.Println(goIds)
}

func printUsers() {
	userGroup := C.GetUserGroup()
	userSlice := (*[1 << 31]C.User)(unsafe.Pointer(userGroup.users))[0:userGroup.num:userGroup.num]
	for _, user := range userSlice {
		fmt.Printf("The user id's:%d name:%s\n", user.id, C.GoString(user.name))
	}
	C.free(unsafe.Pointer(userGroup.users))
	C.free(unsafe.Pointer(userGroup))
}
