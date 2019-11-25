#ifndef __REGISTERCENTER_H__
#define __REGISTERCENTER_H__

#ifdef __cplusplusc
extern "C" {
#endif

typedef void (*NotifyFunc)();
void setNotifyFunc(const char* serveType, NotifyFunc notifyFunc);
void callback(const char* serveType);

#ifdef __cplusplus
}
#endif

#endif