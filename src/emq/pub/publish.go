package main

import (
	"fmt"
	"github.com/go_sample/src/emq/mqttClient"
)

const (
	broker = "tcp://10.111.103.251:1883"
	clientId = "123456654321"
	userName = "my_test"
	pubTopic = "lenovo_ub"
	subTopic = "lenovo_ub"
	qos = "0"
)

func main() {
	client, err := mqttClient.NewMqttClient(broker, clientId, userName, pubTopic, subTopic, qos)
	if err != nil {
		fmt.Println("new mqttClient err: ", err.Error())
	}
	if client == nil {
		fmt.Println("client is nil.")
	}
	client.Init()
	//client.Subscribe()
	client.Publish()
	//client.Unsubscribe()
	client.Destroy()

}
