package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

/**
 * @Project: go_sample
 * @Package Name: main
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/11/6 下午2:56
 * @Description:
 */

//func getip6() string {
//
//}

func getip2() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		return ""
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagBroadcast) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil &&  !strings.HasPrefix(ipnet.IP.String(), "169.254.") {
						fmt.Println(ipnet.IP.String())
						//return ipnet.IP.String()
					}
				}
			}
		}
	}

	return ""
}

func getip() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				gInnerIP := ipnet.IP.String()
				fmt.Println("gInnerIP: ", gInnerIP)
				//return gInnerIP
			}
		}
	}
	return ""
}

func GetIpAddress() string {
	//conn, err := net.Dial("udp", "baidu.com:80")
	//if err == nil {
	//	defer conn.Close()
	//	return conn.LocalAddr().(*net.UDPAddr).IP.String()
	//}

	address, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range address {
			if ip, ok := addr.(*net.IPNet); ok && ip.IP.IsGlobalUnicast() {
				fmt.Println("ipaa: ", ip.IP.String())
				//return ip.IP.String()
			}
		}
	}

	return "0.0.0.0"
}

func main2() {
	ip := getip()
	fmt.Println("ip: ", ip)
	ip2 := getip2()
	fmt.Println("ip2: ", ip2)
	ip3 := GetIpAddress()
	fmt.Println("ip3: ", ip3)
	return
}

func main1() {
	//if len(os.Args) != 2 {
	//	fmt.Println("Usage: ", os.Args[0], " host:port")
	//	os.Exit(1)
	//}
	service := "127.0.0.1:8888"
	//service := "127.0.0.1:41681"
	//service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	//_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	//_, err = conn.Write([]byte("{\"data\":\"tcp_data\"}"))
	//checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	//write(date)
	fmt.Println("asdjfsakldjf")
	os.Exit(0)
}

func main() {
	service := "127.0.0.1:8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println("readAll error: ", err.Error())
	}
	//fmt.Println(string(result))
	return
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1)
	}
}
