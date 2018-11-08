package main

import "github.com/go_sample/src/mqttsdk/sdk"

/**
 * @Package Name: mqttsdk
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-11-1 上午10:46
 * @Description:
 */

func main() {
	mqtt, _ := mqtt.NewMqtt()
	//mqtt.Init(100000, "1234567890qwertyuiop", "8888", "172.17.170.179:8200")
	mqtt.Init("8888", "172.17.170.179:8200")
	mqtt.SendMessage([]byte("luck you."))
}
