package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct{
    int  *ptr;
    char **strs;
}Test;

typedef struct{
    Test *tt;
}Test1;

Test1 newTest() {
    Test1 test;

    test.tt->ptr = malloc(sizeof(int) * 10);
    test.tt->strs = malloc(sizeof(char*) * 5);
    for (int i=0; i<5; i++) {
        test.tt->strs[i] = malloc(sizeof(char) * 1024);
        memset(test.tt->strs[i], 0, sizeof(char) * 1024);
        printf("ptr[%d] = %p\n", i, test.tt->strs[i]);
    }

    memset(test.tt->ptr, 0, sizeof(int) * 10);

    return test;
}
void printTest(Test1 test) {
    printf("arr:");
    for (int i=0; i<10; i++) {
        printf("%d ", test.tt->ptr[i]);
    }
    printf("\n");
    for (int i=0; i<5; i++) {
        printf("strs[%d]:%s\n", i, test.tt->strs[i]);
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
 addr := uintptr(unsafe.Pointer(t.tt.ptr))

 // 模拟数组操作
 for i := uintptr(0); i < 10; i++ {
  *(*int32)(unsafe.Pointer(addr + i*4)) = int32(i)
 }

 addr = uintptr(unsafe.Pointer(t.tt.strs))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*0)), C.CString("hello,world-1"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*1)), C.CString("hello,world-2"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*2)), C.CString("hello,world-3"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*3)), C.CString("hello,world-4"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*4)), C.CString("hello,world-5"))

 C.printTest(t)
}
