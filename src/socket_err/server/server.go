package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
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

	data, err := ioutil.ReadFile("/home/steven/tmp/66.log")
	if err != nil {
		fmt.Println("read file err: ", err.Error())
	}

	fmt.Println("data size is: ", len(data))

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//daytime := time.Now().String()
		//conn.Write([]byte(daytime)) // don't care about return value
		n, err := conn.Write(data)
		if err != nil {
			fmt.Println("write err: ", err.Error())
		}else {
			fmt.Println("write ", n, " bytes.")
		}
		err = conn.Close()                // we're finished with this client
		if err != nil {
			fmt.Println("close err: ", err.Error())
		}
	}
}