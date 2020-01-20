package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct{
    int  *ptr;
    char **strs;
    char **strs2;
    char **strs3;
    char **strs4;
}Test;

typedef struct{
	Test *tt;
	char *ll;
}Test1;

Test1 newTest() {
    Test1 test;

    test.tt = malloc(sizeof(Test)+1);

    test.tt->ptr = malloc(sizeof(int) * 10);
    test.tt->strs = malloc(sizeof(char*) * 5);
    test.tt->strs2 = malloc(sizeof(char*) * 5);
    test.tt->strs3 = malloc(sizeof(char*) * 5);
    test.tt->strs4 = malloc(sizeof(char*) * 5);
    for (int i=0; i<5; i++) {
        test.tt->strs[i] = malloc(sizeof(char) * 13);
        memset(test.tt->strs[i], 0, sizeof(char) * 13);
        printf("str[%d] = %p\n", i, test.tt->strs[i]);
    }
    test.tt->strs2 = malloc(sizeof(char*) * 5);
    for (int i=0; i<5; i++) {
        test.tt->strs2[i] = malloc(sizeof(char) * 13);
        memset(test.tt->strs2[i], 0, sizeof(char) * 13);
        printf("str2[%d] = %p\n", i, test.tt->strs2[i]);
    }
    test.tt->strs3 = malloc(sizeof(char*) * 5);
    for (int i=0; i<5; i++) {
        test.tt->strs3[i] = malloc(sizeof(char) * 13);
        memset(test.tt->strs3[i], 0, sizeof(char) * 13);
        printf("str3[%d] = %p\n", i, test.tt->strs3[i]);
    }
    test.tt->strs4 = malloc(sizeof(char*) * 5);
    for (int i=0; i<5; i++) {
        test.tt->strs4[i] = malloc(sizeof(char) * 13);
        memset(test.tt->strs4[i], 0, sizeof(char) * 13);
        printf("str4[%d] = %p\n", i, test.tt->strs4[i]);
    }


    memset(test.tt->ptr, 0, sizeof(int) * 10);


    return test;
}
void printTest(Test1 test) {
    printf("arr:");
    for (int i=0; i<10; i++) {
        printf("ptr: %p\n", &test.tt->ptr[i]);
    }
    printf("\n");
    for (int i=0; i<5; i++) {
        printf("str[%d] = %p\n", i, test.tt->strs[i]);
        //printf("strs[%d]:%s\n", i, test.tt->strs[i]);
    }
     printf("\n");
    for (int i=0; i<5; i++) {
        printf("str2[%d] = %p\n", i, test.tt->strs2[i]);
        //printf("strs2[%d]:%s\n", i, test.tt->strs2[i]);
    }

     printf("\n");
    for (int i=0; i<5; i++) {
        printf("str3[%d] = %p\n", i, test.tt->strs3[i]);
        //printf("strs3[%d]:%s\n", i, test.tt->strs3[i]);
    }
     printf("\n");
    for (int i=0; i<5; i++) {
        printf("str4[%d] = %p\n", i, test.tt->strs4[i]);
        //printf("strs4[%d]:%s\n", i, test.tt->strs4[i]);
    }
}
*/
import "C"
import (
 "unsafe"
 "fmt"
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
 fmt.Println("addr: ", addr)
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*0)), C.CString("hello,world-1"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*1)), C.CString("hello,world-2"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*2)), C.CString("hello,world-3"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*3)), C.CString("hello,world-4"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*4)), C.CString("hello,world-5"))

 addr = uintptr(unsafe.Pointer(t.tt.strs2))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*0)), C.CString("hello,world-1"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*1)), C.CString("hello,world-2"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*2)), C.CString("hello,world-3"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*3)), C.CString("hello,world-4"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*4)), C.CString("hello,world-5"))

 addr = uintptr(unsafe.Pointer(t.tt.strs3))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*0)), C.CString("hello,world-1"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*1)), C.CString("hello,world-2"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*2)), C.CString("hello,world-3"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*3)), C.CString("hello,world-4"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*4)), C.CString("hello,world-5"))

 addr = uintptr(unsafe.Pointer(t.tt.strs4))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*0)), C.CString("hello,world-1"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*1)), C.CString("hello,world-2"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*2)), C.CString("hello,world-3"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*3)), C.CString("hello,world-4"))
 C.strcpy(*(**C.char)(unsafe.Pointer(addr + 8*4)), C.CString("hello,world-5"))

 C.printTest(t)
}
