package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go_sample/src/emq/mqttClient"
	"time"
)

const (
	broker = "tcp://172.17.170.234:1883"
	agentId = "agent_7f5bbb00-b379-4e54-8541-a737ba288988"
	channel = "t3"

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
			errRec := client.ReceiveMessage(agentId + "/" + channel, 0, func(topic string, msg mqtt.Message) {
				fmt.Println("yao Topic: ", topic)
				fmt.Println("yao Msg: ", string(msg.Payload()))
			})
			if errRec != nil {
				fmt.Println("error: ", errRec.Error())
			}
		}

		defer client.UnReceiveMessage(agentId + "/" + channel)
	}
	time.Sleep(5 * time.Second)
}
