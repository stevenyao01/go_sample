package main

import (
	"fmt"
	"log"
	"time"
)

/**
 * @Project: go_sample
 * @Package Name: convert
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2021/1/19 下午12:26
 * @Description:
 */


func main() {
	ch1 := make(chan int)
	go fmt.Println(<-ch1)
	ch1 <- 5
	time.Sleep(1 * time.Second)
	log.Println("hello!!")
}

func main1() () {

	//interface 转string
	var a interface{}
	var str5 string
	a = "3432423"
	str5 = a.(string)
	fmt.Println(str5)

	//interface 转 []byte
	var b interface{}
	var str6 []byte
	b = []byte("3432423")
	str6 = b.([]byte)
	fmt.Println(string(str6))
}
