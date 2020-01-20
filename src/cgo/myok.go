package main


/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define NODE_NUMBER 100 * 1000

typedef struct Test{
 	char *ptr[NODE_NUMBER];
}Test;

Test newTest(int length) {
 	Test test;
    for(int i=0; i<length; i++){
        test.ptr[i] = malloc(sizeof(char) * 256);
        memset(test.ptr[i], 0, sizeof(char) * 256);
        printf("ptr[%d] = %p\n", i, test.ptr[i]);
    }
 	return test;
}

void printTest(Test test) {
 	for (int i=0; i<10; i++) {
  		printf("%s ", test.ptr[i]);
 	}
 	printf("\n");
}
*/
import "C"

import (
	"fmt"
)

/**
 * @Package Name: cgo
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-9-25 上午10:12
 * @Description:
 */

func main() {
	fmt.Println("new test.")
	t := C.newTest(10)

	fmt.Println("print test.")
	//C.printTest(t)

	// 取到C结构体中ptr指针
	//addr := uintptr(unsafe.Pointer(t.ptr))

	fmt.Println("copy data.")
	// 模拟数组操作
	for i := uintptr(0); i < 10; i++ {
		//*(*C.char)(unsafe.Pointer(addr + i*4)) = *C.CString("hello lenovo")
		//fmt.Println("addrM: ", addr + i*4)
        C.strcpy(t.ptr[i], C.CString("hello lenovo"))
		fmt.Println("addrN: ", t.ptr[i])
	}

	//for i := uintptr(0); i < 10; i++ {
	//	//*(*C.char)(unsafe.Pointer(addr + i*4)) = *C.CString("hello lenovo")
	//	fmt.Println("addrO: ", t.ptr[i])
	//}

	fmt.Println("print test again.")
	C.printTest(t)
	fmt.Println("program end.")
}
