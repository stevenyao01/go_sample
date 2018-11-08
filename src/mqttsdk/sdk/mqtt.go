package mqtt

/**
 * @Package Name: sdk
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-11-1 上午10:46
 * @Description:
 */

import (
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/workerSDK/workerr"
	"github.com/workerSDK/log"
	"encoding/json"
	"net/http"
	"strings"
	"errors"
	"time"
	"fmt"
	"os"
	"io/ioutil"
	"strconv"
)

type Mqtt struct {
	mqttParams   *MqParams
	mqttClient   mqtt.Client
	mqttRequest  MqRequest
	mqttResponse MqReponse
}

type MqRequest struct {
	//CompanyId  int    `json:"company_id"`
	DeviceSk  string `json:"device_sk"`
	DeviceId   string `json:"device_id"`
	DeviceDesc string `json:"device_desc"`
}

type MqReponse struct {
	CompanyId    int    `json:"company_id"`
	ClientId     string `json:"client_id"`
	Broker       string `json:"broker"`
	Port         int    `json:"port"`
	DeviceId     string `json:"device_id"`
	SecretKey    string `json:"secret_key"`
	//Group        string `json:"group"`
	KeepAlive    int    `json:"keep_alive"`
	CleanSession bool   `json:"clean_session"`
	TimeStamp    int64  `json:"timestamp"`
}

type MqParams struct {
	companyId    int
	broker       string
	port         int
	clientId     string
	deviceId     string
	secretKey    string
	group        string
	keepAlive    int
	cleanSession bool
	timestamp    int64
}

type Callback func(client mqtt.Client, msg mqtt.Message) ()

func (m *Mqtt) Init(deviceId string, hostName string) (bool) {
	m.doInit(deviceId, hostName)
	return true
}

func (m *Mqtt) Uninit() () {

}

func (m *Mqtt) SendMessage(data []byte) (bool) {
	if (data == nil) {
		log.D("message is empty");
		return false;
	}

	if (!m.mqttClient.IsConnected()) {
		log.D("mqtt is disconnected\n");
		return false;
	}

	var result = false
	var publish_client_id = m.mqttParams.clientId;
	if (publish_client_id == "") {
		log.D("publish_device_id is nil\n");
		return false
	}

	//String topic = String.format("$SDK/device/%s/data/%s", publish_device_id, receiver_device_id);
	var topic = fmt.Sprintf("$DATA/WORKER/%s/data", publish_client_id);
	if (!m.validTopic(topic)) {
		log.D("receiver_uuid is invalid\n");
		return false;
	}
	log.D("sendtopic == %s", topic)
	log.D("sendData == %s", string(data))
	token := m.mqttClient.Publish(topic, 0, false, string(data))
	token.Wait()
	result = true
	return result;
}

func (m *Mqtt) ReceiveMessage(deviceId string, qos byte, channel string, callback Callback) (bool) {
//func (m *Mqtt) ReceiveMessage() ([]byte) {
//	if (!m.mqttClient.IsConnected()) {
//		log.D("mqtt is disconnected\n");
//		return false;
//	}
//
//	var result = false
//	var publish_client_id = m.mqttParams.clientId;
//	if (publish_client_id == "") {
//		log.D("publish_device_id is nil\n")
//		return false
//	}
//
//	//String topic = String.format("$SDK/device/%s/data/%s", publish_device_id, receiver_device_id);
//	var topic = fmt.Sprintf("$DATA/WORKER/%s/data", publish_client_id);
//	if (!m.validTopic(topic)) {
//		log.D("receiver_uuid is invalid\n");
//		return false;
//	}
//	log.D("receivetopic == %s", topic)
//	////token := m.mqttClient.Publish(topic, 0, false, string(data))
//	m.mqttClient.Subscribe(topic, qos, callback(client mqtt.Client, msg mqtt.Message) {
//	//m.mqttClient.Subscribe(topic, 0, callback(client mqtt.Client, msg mqtt.Message) {
//		// Determine the actual topic
//		fmt.Printf("Success SubscribeUplink with msg:%s\n", msg.Payload())
//	})
	return true
}

//func (m *Mqtt) UnReceive(deviceId string) () {
func (m *Mqtt) UnReceive() () {

}



func (m *Mqtt) register(deviceId string, hostName string) (*MqReponse, error, int) {
	////////////////////
	companySk := m.readFile("device.sk")
	////////////////////////////
	tmpRequest, _ := NewMqRequest(companySk, deviceId)
	m.mqttRequest = *tmpRequest
	data, _ := json.Marshal(m.mqttRequest)

	client := &http.Client{}

	log.D("url == %s\n", "http://"+hostName+"/device/register2.url")
	log.D("req == %s\n", string(data))

	reqest, err := http.NewRequest("POST", "http://"+hostName+"/device/register2.url", strings.NewReader(string(data)))
	reqest.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.D("req err == %s\n", (err))
		return nil, err, workerr.WORKER_REQUEST_ERROR.Code()
	}
	response, err := client.Do(reqest)
	if err != nil {
		log.E("request server error, hostname:%s, post-json:%s", hostName, data)
		return nil, err, workerr.WORKER_REQUEST_ERROR.Code()
	}
	defer response.Body.Close()

	status := response.StatusCode
	if (status != 200) {
		return nil, errors.New(fmt.Sprintf("http code == %d", status)), workerr.WORKER_RESPOND_ERROR.Code()
	}
	body, _ := ioutil.ReadAll(response.Body)

	log.D("http response == %s\n", string(body)) //, _ := ioutil.ReadAll(resp.Body))
	tmpResponse, _ := NewMqResponse()

	err11 := json.Unmarshal(body, &tmpResponse)
	if err11 != nil {
		log.Info("err11: %s", err11)
	}
	return tmpResponse, nil, workerr.WORKER_SUCCESS_EXIT.Code()
}

func (m *Mqtt) getMqttClient() mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker(m.mqttParams.broker).SetClientID(m.mqttParams.clientId)
	opts.SetUsername(m.mqttParams.deviceId)
	opts.SetPassword(m.mqttParams.secretKey)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.E("getMqttClient err == %s\n", token.Error());
		return nil
	}
	return c
}

func (m *Mqtt) doInit(deviceId string, hostName string) (bool) {
	for {
		response, err, errno := m.register(deviceId, hostName)
		if (err != nil) {
			log.E("register error:%s", err)
			time.Sleep(1 * time.Second)
			//continue
			os.Exit(errno)
		}
		if response != nil {
			m.mqttResponse = *response
		}
		break;
	}
	m.mqttParams, _ = NewMqParams(m.mqttResponse)
	m.mqttParams.broker = m.mqttParams.broker + ":" + strconv.Itoa(m.mqttParams.port)
	log.D("broker == %s\n", m.mqttParams.broker);
	m.mqttClient = m.getMqttClient()
	return true;
}

func (m *Mqtt) validTopic(topic string) (bool) {
	return true;
}

func (m *Mqtt) readFile(fileName string) string {
	fi, err := os.Open(fileName)
	if err != nil {
		log.Info("open file error, filename: %s", fileName)
		return ""
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	//log.Info(string(fd))
	return string(fd)
}

func NewMqtt() (*Mqtt, error) {
	return &Mqtt{
	}, nil
}

func NewMqRequest(cs string, did string) (*MqRequest, error) {
	return &MqRequest{
		//CompanyId:  ci,
		DeviceSk:  cs,
		DeviceId:   did,
		DeviceDesc: "",
	}, nil
}

func NewMqResponse() (*MqReponse, error) {
	return &MqReponse{
	}, nil
}

func NewMqParams(response MqReponse) (*MqParams, error) {
	return &MqParams{
		companyId:    response.CompanyId,
		clientId:     response.ClientId,
		deviceId:     response.DeviceId,
		broker:       response.Broker,
		port:         response.Port,
		secretKey:    response.SecretKey,
		//group:        response.Group,
		keepAlive:    response.KeepAlive,
		cleanSession: response.CleanSession,
		timestamp:    response.TimeStamp,
	}, nil
}
