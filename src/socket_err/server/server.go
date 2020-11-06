package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

/**
 * @Project: go_sample
 * @Package Name: socket_err
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/11/6 下午2:53
 * @Description:
 */

func main() {
	service := ":8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime)) // don't care about return value
		conn.Close()                // we're finished with this client
	}
}