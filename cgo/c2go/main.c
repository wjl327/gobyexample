#include <stdio.h>
#include "hello.h"

/**
  本例主要展示Go导出实现给C语言调用Go例子。

  工作目录：c2go
  静态库用法：
       go build -o hello.a -buildmode=c-archive ./src
       gcc -o out main.c hello.a

  动态库用法：
       go build -o hello.so -buildmode=c-shared ./src
       gcc -o out main.c hello.so
*/
int main() {
    printf("I'm C program.\n");

    char cname[] = "Jarvis";
    int length = sizeof(cname) / sizeof(cname[0]);
    GoString name = {cname, length};
    printf("Rsp: %s\n", SayHello(name));
    return 0;
}