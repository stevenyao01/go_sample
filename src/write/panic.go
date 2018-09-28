package main

import "fmt"

/**
 * @Package Name: main
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-20 下午6:30
 * @Description:
 */

func main() {
	defer func() {     //必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err)    //这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()
	f()
}


func f() {
	fmt.Println("a")
	panic(55)
	fmt.Println("b")

	fmt.Println("f")
}