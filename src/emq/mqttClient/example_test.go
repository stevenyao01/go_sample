package mqttClient

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"time"
)

//const (
//	broker = "tcp://10.111.103.251:1883"
//	clientIdSend = "123456654321"
//	clientIdReceive = "123456654321abc"
//	userName = "my_tests"
//	qos = 0
//)

func ExampleMqttClient_ReceiveMessage() {
	client, err := NewMqttClient(broker, clientIdReceive, userName, qos)
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

func ExampleMqttClient_SendMessage() {
	client, err := NewMqttClient(broker, clientIdSend, userName, qos)
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
