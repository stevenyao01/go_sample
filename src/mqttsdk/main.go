package main

import (
	"github.com/eclipse/paho.mqtt.golang"
	"fmt"
	"github.com/go_sample/src/mqttsdk/sdk"
	"time"
)

/**
 * @Package Name: mqttsdk
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-11-1 上午10:46
 * @Description:
 */



func main() {
	sdkParams, _ := sdk.NewSdkParams("172.17.171.20:8200", "device.sk", "uuid")
	mqtt, _ := sdk.NewMqtt()
	mqtt.Init(*sdkParams)
	mqtt.ReceiveBroadcast("$LEAP/localid", callBackReceiveBroadCast)
	mqtt.Broadcast("$LEAP/localid", []byte("luck you."))

	sdkParams1, _ := sdk.NewSdkParams("172.17.171.20:8200", "device.sk", "edge01")
	mqtt1, _ := sdk.NewMqtt()
	mqtt1.Init(*sdkParams1)
	mqtt1.ReceiveMessage("edge02", 0, "/#channel+b", callBackReceiveMessage)
	mqtt1.SendMessage("edge02", 0, "/#channel+b", []byte("luck you."))


	sdkParams2, _ := sdk.NewSdkParams("172.17.171.20:8200", "device.sk", "edge02")
	mqtt2, _ := sdk.NewMqtt()
	mqtt2.Init(*sdkParams2)
	mqtt2.ReceiveMessage("edge01", 0, "/#channel+a中文", callBackReceiveMessage)
	mqtt2.SendMessage("edge01", 0, "/#channel+a中文", []byte("luck you."))

	time.Sleep(1*time.Second)

	//mqtt.Uninit()
	//mqtt1.Uninit()
	//mqtt2.Uninit()
}


var callBackReceiveMessage sdk.CbReceive = func(sender_device_id string, channel string, msg mqtt.Message) () {
	fmt.Println("sender_device_id: ", sender_device_id)
	fmt.Println("channel: ", channel)
	fmt.Println("receive msg in callback: sdfsadfsadf ", string(msg.Payload()))
}

var callBackReceiveBroadCast sdk.CbBroadCast = func(topic string, msg mqtt.Message) () {
	fmt.Println("topic: ", topic)
	fmt.Println("receiveBroadCast msg in callback: sdfsadfsadf ", string(msg.Payload()))
}


