package main

import (
	"fmt"
	"github.com/go_sample/src/emq/mqttClient"
)

func main() {
	c, err := mqttClient.NewMqttClient()
	if err != nil {
		fmt.Println("new mqttClient err: ", err.Error())
	}
	c.
}
