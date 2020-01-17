package main

import (
	"fmt"
	"github.com/go_sample/src/emq/mqttClient"
	"time"
)

const (
	broker = "tcp://10.111.103.251:1883"
	clientId = "123456654321abc"
	userName = "my_tests"
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
	for {
		client.Subscribe()
		//client.Publish()
		time.Sleep(1 * time.Second)
	}

	client.Unsubscribe()
	client.Destroy()

}
