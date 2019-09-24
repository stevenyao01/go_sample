package myNtp

/**
 * @Package Name: myNtp
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-7-31 上午11:29
 * @Description:
 */

import (
	"testing"
	"fmt"
	"net"
	"log"
)

const (
	NtpServer = "time.windows.com:123" // 阿里云NTP服务器 "ntp1.aliyun.com:123"
)

func TestNewClient(t *testing.T) {
	var (
		ntp    *Ntp
		buffer []byte
		err    error
		ret    int
	)
	conn, err := net.Dial("udp", NtpServer)
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
		conn.Close()
	}()
	ntp = NewNtp()
	conn.Write(ntp.GetBytes())
	buffer = make([]byte, 2048)
	ret, err = conn.Read(buffer)
	if err == nil {
		if ret > 0 {
			ntp.Parse(buffer, true)
			fmt.Println("time: ", ntp.TransmitTimestamp)
		}
	}
}