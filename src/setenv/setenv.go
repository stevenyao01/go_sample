package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/**
 * @Package Name: setenv
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-1-2 下午3:14
 * @Description:
 */

func main() {
	err := os.Setenv("XIAO", "xiaochuanaaaa") //临时设置 系统环境变量
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(os.Getenv("XIAO")) //获取环境变量
	fmt.Println(os.Getenv("GOPATH"))
	for _, v := range os.Environ() { //获取全部系统环境变量 获取的是 key=val 的[]string
		str := strings.Split(v, "=")
		fmt.Printf("key=%s,val=%s \n", str[0], str[1])
	}
	time.Sleep(1000 * time.Second)
	{}
}
