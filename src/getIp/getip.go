package main

import (
	"fmt"
	"net"
)

func GetIpAddress() (ip string, ipList []string) {
	ma := make(map[int]string)


	var ss = make([]string, 0)
	address, err := net.InterfaceAddrs()
	for k, addr := range address {
		ip, _ := addr.(*net.IPNet)
		if ip.IP.To4() != nil {
			ma[k] = ip.IP.String()
		}
	}
	fmt.Println("ss000: ", ma)

	for _, addr := range address {
		ip, _ := addr.(*net.IPNet)
		if ip.IP.To4() != nil {
			ss = append(ss, ip.IP.String())
		}
	}
	fmt.Println("ss: ", ss)

	conn, err := net.Dial("udp", "baidu.com:80")
	if err == nil {
		defer conn.Close()
		return conn.LocalAddr().(*net.UDPAddr).IP.String(), ss
	}

	address, err = net.InterfaceAddrs()
	if err == nil {
		for _, addr := range address {
			if ip, ok := addr.(*net.IPNet); ok && ip.IP.IsGlobalUnicast() {
				return ip.IP.String(), ss
			}
		}
	}

	return "0.0.0.0", ss
}

func main() {
	ip, ipList := GetIpAddress()
	fmt.Println("ip: ", ip)
	fmt.Println("ipList: ", ipList)
}
