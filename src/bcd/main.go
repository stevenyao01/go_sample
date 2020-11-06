package main

import (
	"fmt"
	"github.com/go_sample/src/bcd/bcd"
)

/**
 * @Project: go_sample
 * @Package Name: main
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/9/29 下午6:03
 * @Description:
 */

func main() {
	fmt.Println("Uint32: ", bcd.ToUint32([]byte{0x9, 0x9, 0x9, 0x9}))
	fmt.Println("BCD: ", bcd.FromUint32(9090909))
}