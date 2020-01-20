package main

import (
	"fmt"
	"sync/atomic"
	"runtime"
	"syscall"
)

/**
 * @Package Name: atomic
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-11-28 下午9:03
 * @Description:
 */

func main() {
	var Atomicvalue  atomic.Value
	syscall.Dup2()
	Atomicvalue.Store([]int{1,2,3,4,5})
	anotherStore(&Atomicvalue)
	fmt.Println("main: ",Atomicvalue)
}

func anotherStore(Atomicvalue *atomic.Value)  {
	Atomicvalue.Store([]int{6,7,8,9,10})
	fmt.Println("anotherStore: ",Atomicvalue)
}
