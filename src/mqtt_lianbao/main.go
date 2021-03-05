package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/go_sample/src/mqtt/mqttSDK"
	"os"
)

/**
 * @Package Name: mqttSDK
 * @Author: steven yao steven
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2021/1/13 下午5:37
 * @Description:
 */

//const(
//	broker="172.17.170.234:4567"
//)

// response sub ex message struct
type SubResEx struct {
	// response edge time
	EdgeTime string `json:"agenttime"`
	// response status code, for example: 200, 400 (brain disabled design)
	Code string `json:"code"`
	// response executer result
	Message string `json:"message"`
	// response description (brain disabled design)
	Description string `json:"description"`
}

//// receive ex message struct
//type RecExMsg struct {
//	// receive ex message type, for ex: CMD
//	MsgType string `json:"msgtype"`
//	// receive ex message id
//	MsgId string `json:"msgid"`
//	// receive ex message version
//	MsgVersion string `json:"msgversion"`
//	// receive ex message job id
//	JobId string `json:"jobid"`
//	// receive ex message plat id
//	PlatId string `json:"platid"`
//	// receive ex message sub struct
//	Msg SubRecEx `json:"msg"`
//}

// response ex message struct
type ResExMsg struct {
	// response ex message type, such as: CMD
	MsgType string `json:"msgtype"`
	// response ex message id
	MsgId string `json:"msgid"`
	// response ex message version
	MsgVersion string `json:"msgversion"`
	// response ex message job id
	JobId string `json:"jobid"`
	// response ex message plat id
	PlatId string `json:"platid"`
	// response ex message sub struct
	Msg SubResEx `json:"msg"`
}


var broker = flag.String("broker", "172.17.170.234:4567", "create mqtt.conf and download device.sk.")
var sk = flag.String("sk", "./device.sk", "special device.sk.")
var agentId = flag.String("agentId", "", "agent uuid.")
var channel = flag.String("channel", "", "channel.")
var topic = flag.String("topic", "", "topic.")

func main() {
	flag.Parse()

	params := mqttSDK.NewSdkParams(*broker, *sk, "abc")
	client := mqttSDK.NewClientGO(params)
	if err := client.Init(); err != nil {
		fmt.Println("client init failed: " + err.Error())
		return
	}
	defer client.UnInit()

	if err := client.ReceiveBroadcast(*agentId + "/ex_to_server", 0, func(topic string, msg mqtt.Message) {
		fmt.Println("Topic:", topic)
		fmt.Println("Msg:", string(msg.Payload()))
		var res ResExMsg
		err := json.Unmarshal(msg.Payload(), &res)
		if err != nil {
			fmt.Println("res unmarshal err: ", err.Error())
		}
		fmt.Println("message: ", res.Msg.Message)
	}); err != nil {
		fmt.Println("ReceiveBroadcast err: ", err.Error())
	}

	//agent := ""
	topic := *agentId + "/ex_from_server"

	//format := `
	//	{
	//		"seq":"0",
	//		"cmd":"CMD",
	//		"workerid":"123",
	//		"executer":"%s"
	//	}
	//`
	format := `
		{
    		"msgtype": "CMD",
    		"msgid": "xxx",
    		"msgversion": "xxx",
    		"jobid": "xxx",
    		"msg": {
        		"cmd": "SHELL",
        		"executer": "%s",
        		"topic": "",
        		"reportinterval": "0"
    		}
		}
	`

	for {
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("readline err: ", err.Error())
		}
		if string(line) == "exit" {
			break
		}
		msg := fmt.Sprintf(format, line)
		fmt.Println(msg)
		n, err := client.Broadcast(topic, 0, []byte(msg))
		if err != nil {
			fmt.Println("Broadcast err: ", err.Error())
		}
		fmt.Println("broadcast success, send length:", n)
	}
}

func main1() {
	flag.Parse()

	fmt.Println("broker: ", *broker)
	fmt.Println("sk: ", *sk)
	fmt.Println("topic: ", *topic)
	fmt.Println("agentId: ", *agentId)
	fmt.Println("channel: ", *channel)

	fmt.Println("pls operate your edgeserver such as start or stop.")
	fmt.Println("")

	params := mqttSDK.NewSdkParams(*broker, *sk, "abc")
	client := mqttSDK.NewClientGO(params)
	if err := client.Init(); err != nil {
		fmt.Println("client init failed: " + err.Error())
		return
	}
	defer client.UnInit()

	for true {
		if *topic == "" {
			*topic = *agentId+"/"+*channel
		}
		err := client.ReceiveBroadcast(*topic, 0, func(topic string, msg mqtt.Message) {
			fmt.Println("************************sub Topic:", topic)
			fmt.Println("************************get Topic:", msg.Topic())
			fmt.Println("Msg:", string(msg.Payload()))
			fmt.Println("")
		})
		if  err != nil {
			fmt.Println("receive err: ", err.Error())
		}
	}
}