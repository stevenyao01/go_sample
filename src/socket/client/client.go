package main

import (
	"net"
	"syscall"
)

/**
 * @Package Name: client
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-1-10 上午11:11
 * @Description:
 */

func main() {
	net.ResolveIPAddr()
	syscall.Socket()
}
