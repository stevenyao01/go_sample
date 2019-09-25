package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Test{
    int  *ptr;
    char **strs;
}Test;
Test newTest() {
    Test test;

    test.ptr = malloc(sizeof(int) * 10);
    test.strs = malloc(sizeof(char*) * 5);
    for (int i=0; i<5; i++) {
        test.strs[i] = malloc(sizeof(char) * 1024);
        memset(test.strs[i], 0, sizeof(char) * 1024);
    }

    memset(test.ptr, 0, sizeof(int) * 10);

    return test;
}
void printTest(Test test) {
    printf("arr:");
    for (int i=0; i<10; i++) {
        printf("%d ", test.ptr[i]);
    }
    printf("\n");
    for (int i=0; i<5; i++) {
        printf("strs[%d]:%s\n", i, test.strs[i]);
    }
}
*/
import "C"
import (
 "unsafe"
)

func main() {
 t := C.newTest()

 C.printTest(t)

 // 取到C结构体中ptr指针
 addr := uintptr(unsafe.Pointer(t.ptr))

 // 模拟数组操作
 for i := uintptr(0); i < 10; i++ {
  *(*int32)(unsafe.Pointer(addr + i*4)) = int32(i)
 }

 addr = uintptr(unsafe.Pointer(t.strs))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*0)), C.CString("hello,world-1"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*1)), C.CString("hello,world-2"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*2)), C.CString("hello,world-3"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*3)), C.CString("hello,world-4"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*4)), C.CString("hello,world-5"))

 C.printTest(t)
}
