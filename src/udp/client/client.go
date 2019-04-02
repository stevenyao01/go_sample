package main

import (
	"fmt"
	"net"
)

/**
 * @Package Name: client
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-3-29 下午2:09
 * @Description:
 */


func main() {
	// 创建连接
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(127,0,0,1),
		Port: 23452,
	})
	if err != nil {
		fmt.Println("连接失败!", err)
		return
	}
	defer socket.Close()
	// 发送数据
	senddata := []byte("hello server!")
	_, err = socket.Write(senddata)
	if err != nil {
		fmt.Println("发送数据失败!", err)
		return
	}
	// 接收数据
	data := make([]byte, 4096)
	read, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("读取数据失败!", err)
		return
	}
	fmt.Println(read, remoteAddr)
	fmt.Printf("%s\n", string(data))
}