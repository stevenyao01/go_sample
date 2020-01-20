package main

import (
	"fmt"
	"github.com/go_sample/src/emq/mqttClient"
)

const (
	broker = "tcp://10.111.103.251:1883"
	clientId = "123456654321"
	userName = "my_test"
	qos = 0
)

func main() {
	client, err := mqttClient.NewMqttClient(broker, clientId, userName, qos)
	if err != nil {
		fmt.Println("new mqttClient err: ", err.Error())
	}
	if client == nil {
		fmt.Println("client is nil.")
	} else {
		errInit := client.Init()
		if errInit != nil {
			fmt.Println("connect error: ", errInit.Error())
		}
		defer client.UnInit()

		n, errSendMessage := client.SendMessage("demo", 0, []byte("hello world!!!"))
		if errSendMessage != nil {
			fmt.Println("error: send message error.")
		}
		fmt.Println("send message success, send length: ", n)
	}
}
