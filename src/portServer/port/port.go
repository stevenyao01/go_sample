package port

import (
	"net"
	"fmt"
	"sync"
)

/**
 * @Package Name: port
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-11-30 下午3:58
 * @Description:
 */

var lock sync.Mutex

type ListenPort struct {
	port 	int
}


func (l *ListenPort) StartListen(){
	listener,err := net.Listen("tcp", "0.0.0.0:")
	if err != nil{
		fmt.Println("Listen Error: ", err.Error())
	}

	ip := listener.Addr().(*net.TCPAddr).IP
	port := listener.Addr().(*net.TCPAddr).Port

	fmt.Println("Metrics Listened ip: ", ip)
	fmt.Println("Metrics Listened port: ", port)

	for{
		conn,err := listener.Accept()
		if err!=nil{
			fmt.Println("Accept Error: ", err.Error())
			return
		}
		go l.doProcess(conn)
	}
}

func (l *ListenPort) doProcess(conn net.Conn)(){
	fmt.Println("hello doProcess!")
}

func NewListenPort() (*ListenPort, error) {
	return &ListenPort{
	},nil
}