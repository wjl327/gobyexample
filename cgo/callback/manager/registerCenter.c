#include "registerCenter.h"
#include <stdio.h>
#include <string.h>

NotifyFunc easyNotifyFunc;
NotifyFunc hardNotifyFunc;
NotifyFunc otherNotifyFunc;


void setNotifyFunc(const char* serveType, NotifyFunc notifyFunc) {
    if(strcmp(serveType, "easy") == 0) {
        easyNotifyFunc = notifyFunc;
    }
    else if(strcmp(serveType, "hard") == 0) {
        hardNotifyFunc = notifyFunc;
    }
    else {
        otherNotifyFunc = notifyFunc;
    }
    printf("C set NotifyFunc[%s] success!\n", serveType);
}

void callback(const char* serveType) {
    printf("C ready callback go NotifyFunc[%s]!\n", serveType);
    if(strcmp(serveType, "easy") == 0) {
        easyNotifyFunc();
    }
    else if(strcmp(serveType, "hard") == 0) {
        hardNotifyFunc();
    }
    else {
        otherNotifyFunc();
    }
}
