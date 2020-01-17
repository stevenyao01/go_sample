package mqttClient

import (
	"fmt"
	//"log"
	"os"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
)

type MqttClient struct {
	// input param
	broker   string
	clientId string
	userName string
	subTopic string
	pubTopic string
	qos      string
	// self param
	c    mqtt.Client
	opts *mqtt.ClientOptions
}

//set callback function
var callBack mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

//func main() {
//	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
//	//mqtt.ERROR = log.New(os.Stdout, "", 0)
//
//	//connect mqtt-server and set clientID
//	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("mqtt_client")
//
//	//set userName
//	opts.SetUsername("mqtt_test")
//	opts.SetKeepAlive(2 * time.Second)
//	opts.SetDefaultPublishHandler(callBack)
//	opts.SetPingTimeout(1 * time.Second)
//
//	//create object
//	c := mqtt.NewClient(opts)
//	if token := c.Connect(); token.Wait() && token.Error() != nil {
//		panic(token.Error())
//	}
//
//	//subscribe topic
//	if token := c.Subscribe("go-mqtt/sample", 0, nil); token.Wait() && token.Error() != nil {
//		fmt.Println(token.Error())
//		os.Exit(1)
//	}
//
//	//publish topic
//	for i := 0; i < 5; i++ {
//		text := fmt.Sprintf("this is msg #%d!", i)
//		token := c.Publish("go-mqtt/sample", 0, false, text)
//		token.Wait()
//	}
//
//	//unsubscribe topic
//	time.Sleep(6 * time.Second)
//
//	if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
//		fmt.Println(token.Error())
//		os.Exit(1)
//	}
//
//	c.Disconnect(250)
//
//	time.Sleep(1 * time.Second)
//}

func (m *MqttClient) Reload() () {

}
func (m *MqttClient) Destroy() () {
	m.c.Disconnect(250)
	return
}
func (m *MqttClient) Unsubscribe() () {
	if token := m.c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	return
}
func (m *MqttClient) Subscribe() () {
	if token := m.c.Subscribe(m.subTopic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	return
}
func (m *MqttClient) Publish() () {
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := m.c.Publish(m.pubTopic, 0, false, text)
		token.Wait()
	}
	return
}

func (m *MqttClient) Init() (bool) {
	// init options
	m.opts = mqtt.NewClientOptions()
	m.opts.AddBroker(m.broker)
	m.opts.SetClientID(m.clientId)
	m.opts.SetUsername(m.userName)
	m.opts.SetKeepAlive(2 * time.Second)
	m.opts.SetDefaultPublishHandler(callBack)
	m.opts.SetPingTimeout(1 * time.Second)
	// init client
	m.c = mqtt.NewClient(m.opts)
	if token := m.c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(" new mqtt client err: ", token.Error())
		return false
	}
	return true
}

func NewMqttClient(b string, ci string, un string, pt string, st string, q string) (*MqttClient, error) {
	return &MqttClient{
		broker:   b,
		clientId: ci,
		userName: un,
		pubTopic: pt,
		subTopic: st,
		qos:      q,
	}, nil
}
