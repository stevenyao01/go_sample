package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go_sample/src/emq/mqttClient"
	"time"
)

const (
	broker = "tcp://10.111.103.251:1883"
	clientId = "123456654321abc"
	userName = "my_tests"
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

		for {
			errRec := client.ReceiveMessage("demo", 0, func(topic string, msg mqtt.Message) {
				fmt.Println("yao Topic: ", topic)
				fmt.Println("yao Msg: ", string(msg.Payload()))
			})
			if errRec != nil {
				fmt.Println("error: ", errRec.Error())
			}
		}

		defer client.UnReceiveMessage("demo")
	}
	time.Sleep(5 * time.Second)
}
