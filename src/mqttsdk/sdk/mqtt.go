package sdk

/**
 * @Package Name: sdk
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-11-1 上午10:46
 * @Description:
 */

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"github.com/agent/log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	//"github.com/apex/log"
	"encoding/base64"
)

const timeOut = 30
var timeout = errors.New("mqtt timeout")

type CbReceive func(sender_device_id string, channel string, msg mqtt.Message) ()
type CbBroadCast func(topic string, msg mqtt.Message) ()

type Mqtt struct {
	mqttParams        *MqParams
	mqttClient        mqtt.Client
	mqttRequest       MqRequest
	mqttResponse      MqReponse
	//callBackReceive   mqtt.MessageHandler
	//callBackBroadCast mqtt.MessageHandler
}

type MqRequest struct {
	DeviceSk   string `json:"device_sk"`
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
	KeepAlive    int    `json:"keep_alive"`
	CleanSession bool   `json:"clean_session"`
	TimeStamp    int64  `json:"timestamp"`
}

type MqParams struct {
	companyId           int
	broker              string
	port                int
	clientId            string
	deviceId            string
	secretKey           string
	group               string
	keepAlive           int
	cleanSession        bool
	timestamp           int64
	mqttSubReceiveMap   map[string]*MqSub
	mqttSubBroadCastMap map[string]*MqSub
}

type Token interface {
	Wait() bool
	WaitTimeout(time.Duration) bool
	Error() error
}

type Client interface {
	Connect() error
	Disconnect()

	IsConnected() bool

	PublishUplink(topic string, msg string) Token
	SubscribeUplink(topic string) Token
}

var ctx log.Interface

func (m *Mqtt) Init(sdkParams SdkParams) (bool) {
	return m.doInit(sdkParams)
}

func (m *Mqtt) Uninit() (bool) {
	if !m.mqttClient.IsConnected() {
		fmt.Println("mqtt is disconnected")
		return false
	}
	m.mqttClient.Disconnect(100)
	return true
}

func (m *Mqtt) SendMessage(receive_device_id string, qos int, channel string, data []byte) (n int, err error) {
	if data == nil {
		fmt.Println("message is empty")
		return 0, errors.New("message is empty")
	}

	//m.checkConnect()

	var publish_client_id = m.mqttParams.clientId
	if publish_client_id == "" {
		fmt.Println("publish_device_id is nil")
		return 0, errors.New("publish_device_id is nil")
	}

	var topic = fmt.Sprintf("$LEAP/%d/%s/message/%s/%s", m.mqttParams.companyId, receive_device_id, m.mqttParams.deviceId, m.encodeTopic(channel))
	if !m.validTopic(topic) {
		fmt.Println("receiver_uuid is invalid")
		return 0, errors.New("receiver_uuid is invalid")
	}
	fmt.Println("sendtopic == " + topic)
	fmt.Println("sendData == " + string(data))
	token := m.mqttClient.Publish(topic, 2, false, string(data))
	if !token.WaitTimeout(timeOut * time.Second) {
		return 0,timeout
	}

	return len(data), token.Error()
}

func (m *Mqtt) ReceiveMessage(sender_device_id string, qos int, channel string, callback CbReceive) error {
	//m.checkConnect()

	var publish_client_id = m.mqttParams.clientId
	if publish_client_id == "" {
		fmt.Println("publish_device_id is nil\n")
		return errors.New("publish_device_id is nil")
	}

	var topic = fmt.Sprintf("$LEAP/%d/%s/message/%s/%s", m.mqttParams.companyId, sender_device_id, m.mqttParams.deviceId, m.encodeTopic(channel))
	if !m.validTopic(topic) {
		fmt.Println("receive topic is invalid")
		return errors.New("receive topic is invalid")
		//return false;
	}
	fmt.Println("receive topic == " + topic)
	if _, ok := m.mqttParams.mqttSubReceiveMap[topic]; !ok {
		mqSub, _ := NewMqSub(topic, byte(qos), callback, nil)
		m.mqttParams.mqttSubReceiveMap[topic] = mqSub
	}

	token := m.mqttClient.Subscribe(topic, byte(qos), m.CbReceive)
	if !token.WaitTimeout(timeOut * time.Second) {
		return timeout
	}

	return token.Error()
}

func (m *Mqtt) Broadcast(topic string, data []byte) (n int, err error) {
	if data == nil {
		fmt.Println("message is empty")
		return 0, errors.New("message is empty")
	}

	//m.checkConnect()

	var publish_client_id = m.mqttParams.clientId
	if publish_client_id == "" {
		fmt.Println("publish_device_id is nil")
		return 0, errors.New("publish_device_id is nil")
	}

	//fmt.Println("broadcast topic == " + topic)
	token := m.mqttClient.Publish(topic, 1, false, string(data))
	if !token.WaitTimeout(timeOut * time.Second) {
		log.Error("publish timeout.")
		return 0, timeout
	}
	if token.Error() != nil {
		fmt.Println("publish error!")
	}//else{
	//	fmt.Println("broadcast Data == " + string(data))
	//}

	return len(data), token.Error()
}

func (m *Mqtt) ReceiveBroadcast(topic string, callback CbBroadCast) error {
	//m.checkConnect()

	//var publish_client_id = m.mqttParams.clientId
	//if publish_client_id == "" {
	//	fmt.Println("publish_device_id is nil\n")
	//	return errors.New("publish_device_id is nil")
	//}

	fmt.Println("receive broadcast topic == " + topic)
	if _, ok := m.mqttParams.mqttSubBroadCastMap[topic]; !ok {
		mqSub, _ := NewMqSub(topic, 0, nil, callback)
		m.mqttParams.mqttSubBroadCastMap[topic] = mqSub
	}

	token := m.mqttClient.Subscribe(topic, 0, m.CbBroadCast)
	if !token.WaitTimeout(timeOut * time.Second) {
		return timeout
	}
	if token.Error() != nil {
		fmt.Println("subscribe error: ", token.Error().Error())
	}
	return token.Error()
}

func (m *Mqtt) CbReceive(client mqtt.Client, msg mqtt.Message) {
	var channel string
	result := strings.Split(msg.Topic(), "/")
	if len(result[5]) > 0 {
		channel = m.decodeTopic(result[5])
	}
	m.mqttParams.mqttSubReceiveMap[msg.Topic()].callbackReceive(result[1], channel, msg)
}

func (m *Mqtt) CbBroadCast(client mqtt.Client, msg mqtt.Message) {
	//fmt.Println("receive msg in callback: ", string(msg.Payload()))
	//fmt.Println("topic:", msg.Topic())
	m.mqttParams.mqttSubBroadCastMap[msg.Topic()].callbackBroadCast(msg.Topic(), msg)
}

var callBack mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) () {
	fmt.Println("receive msg in callback: ", string(msg.Payload()))
	//msg.Topic()
}

func (m *Mqtt) UnReceive(sender_device_id string, channel string) (bool) {
	//m.checkConnect()

	var topic = fmt.Sprintf("$LEAP/%d/%s/message/%s/%s", m.mqttParams.companyId, sender_device_id, m.mqttParams.deviceId, m.encodeTopic(channel))
	if (!m.validTopic(topic)) {
		fmt.Println("receive topic is invalid")
		return false
	}
	fmt.Println("UnReceive topic == " + topic)
	m.mqttClient.Unsubscribe(topic)
	return true
}

func (m *Mqtt) encodeTopic(channel string) (string) {
	encodeString := base64.StdEncoding.EncodeToString([]byte(channel))
	encodeString = strings.Replace(encodeString, "+", "-", -1)
	encodeString = strings.Replace(encodeString, "/", "_", -1)
	return encodeString
}

func (m *Mqtt) decodeTopic(channel string) (string) {
	channelString := channel
	channelString = strings.Replace(channelString, "_", "/", -1)
	channelString = strings.Replace(channelString, "-", "+", -1)
	decode, err := base64.StdEncoding.DecodeString(channel)
	if err == nil {
		return string(decode)
	}
	return ""
}

//func (m *Mqtt) checkConnect() () {
//	for !m.mqttClient.IsConnected() {
//		err := m.reConnect()
//		if err != nil {
//			fmt.Println("mqtt connect is disconnected, retry now!")
//			time.Sleep(1 * time.Second)
//			continue
//		}
//		fmt.Println("mqtt connect is resumed.")
//		break
//	}
//}
//
//func (m *Mqtt) reConnect() (error) {
//	if m.mqttClient != nil {
//		token := m.mqttClient.Connect()
//		token.WaitTimeout(timeOut * time.Second)
//		if token.Error() != nil {
//			fmt.Printf("Could not reConnect to MQTT\n")
//			return token.Error()
//		} else {
//			fmt.Printf("Success reConnect to MQTT\n")
//			if m.mqttClient.IsConnected() {
//				if len(m.mqttParams.mqttSubReceiveMap) > 0 {
//					for _, v := range m.mqttParams.mqttSubReceiveMap {
//						m.mqttClient.Subscribe(v.topic, v.qos, m.CbReceive)
//					}
//				}
//				if len(m.mqttParams.mqttSubBroadCastMap) > 0 {
//					for _, v := range m.mqttParams.mqttSubBroadCastMap {
//						m.mqttClient.Subscribe(v.topic, v.qos, m.CbBroadCast)
//					}
//				}
//			}
//			return nil
//		}
//	}
//	return errors.New("mqttClient is nil!")
//}

func (m *Mqtt) register(sdkParams SdkParams) (*MqReponse, error) {
	companySk := m.readFile(sdkParams.device_sk)
	tmpRequest, _ := NewMqRequest(companySk, sdkParams.device_id)
	m.mqttRequest = *tmpRequest
	data, _ := json.Marshal(m.mqttRequest)

	client := &http.Client{}

	fmt.Println("url == " + "http://" + sdkParams.server + "/device/register2.url\n")
	fmt.Println("req == " + string(data) + "\n")

	reqest, err := http.NewRequest("POST", "http://"+sdkParams.server+"/device/register2.url", strings.NewReader(string(data)))
	reqest.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Debug("req err == " + err.Error())
		return nil, err
	}
	response, err := client.Do(reqest)
	if err != nil {
		log.Error("request server error, hostname: " + sdkParams.server + "post-json: " + string(data))
		return nil, err
	}
	defer response.Body.Close()

	status := response.StatusCode
	if (status != 200) {
		return nil, errors.New(fmt.Sprintf("http code == %d", status))
	}
	body, _ := ioutil.ReadAll(response.Body)

	log.Debug("http response == " + string(body)) //, _ := ioutil.ReadAll(resp.Body))
	tmpResponse, _ := NewMqResponse()

	err11 := json.Unmarshal(body, &tmpResponse)
	if err11 != nil {
		log.Debug("err11: " + err11.Error())
	}
	return tmpResponse, nil
}

func (m *Mqtt) NewClient(ctx log.Interface, username, password string, brokers ...string) mqtt.Client {
	tlsconfig := m.newTLSConfigSingle()

	mqttOpts := mqtt.NewClientOptions()

	for _, broker := range brokers {
		mqttOpts.AddBroker(broker)
	}

	mqttOpts.SetClientID(m.mqttParams.clientId)
	mqttOpts.SetUsername(username)
	mqttOpts.SetPassword(password)

	// TODO: Some tuning of these values probably won't hurt:
	mqttOpts.SetKeepAlive(30 * time.Second)
	mqttOpts.SetPingTimeout(10 * time.Second)

	mqttOpts.SetAutoReconnect(true)
	// Usually this setting should not be used together with random ClientIDs, but
	// we configured The Things Network's MQTT servers to handle this correctly.
	mqttOpts.SetCleanSession(true)

	mqttOpts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		//ctx.WithField("message", msg).Warn("Received unhandled message")
		fmt.Println("message: ", msg)
	})
	mqttOpts.SetMaxReconnectInterval(100*time.Millisecond)

	mqttOpts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		fmt.Println("connect lost.")
		////ctx.WithError(err).Warn("Disconnected, reconnecting...")
		//fmt.Println("Disconnected, reconnecting.....")
		//for true {
		//	m.mqttClient.Disconnect(0)
		//	token := m.mqttClient.Connect()
		//	if !token.WaitTimeout(timeOut) {
		//		fmt.Println("reConnect timeout")
		//		continue
		//	}
		//
		//	if token.Error() != nil {
		//		fmt.Println("reConnect failed: ", token.Error().Error())
		//		continue
		//	}
		//	fmt.Println("reConnect ok!")
		//	break
		//}
	})

	mqttOpts.SetOnConnectHandler(func(client mqtt.Client) {
		fmt.Println("Connected")
		for true {
			token := m.mqttClient.Subscribe("$LEAP/localid", 0, m.CbBroadCast)
			token.WaitTimeout(timeOut)
			if token.Error() != nil {
				fmt.Println("reSubscribe failed")
				continue
			}
			fmt.Println("reSubscribe ok!")
			break
		}
	})

	mqttOpts.SetTLSConfig(tlsconfig)

	return mqtt.NewClient(mqttOpts)
}

func (m *Mqtt) newTLSConfigSingle() *tls.Config {
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile("server.crt")
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: true,
		// Certificates = list of certs client sends to server.
		//Certificates: []tls.Certificate{cert},
	}
}

func (m *Mqtt) newTLSConfigDouble() *tls.Config {
	// Import trusted certificates from CAfile.pem.
	// Alternatively, manually add CA certificates to
	// default openssl CA bundle.
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile("server.crt")
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}
	fmt.Println("0. resd pemCerts Success")

	// Import client certificate/key pair
	cert, err := tls.LoadX509KeyPair("server.crt", "server.crt")
	if err != nil {
		panic(err)
	}
	fmt.Println("1. resd cert Success")

	// Just to print out the client certificate..
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		panic(err)
	}
	fmt.Println("2. resd cert.Leaf Success")

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: true,
		// Certificates = list of certs client sends to server.
		Certificates: []tls.Certificate{cert},
	}
}

func (m *Mqtt) doInit(sdkParams SdkParams) (bool) {
	for {
		response, err := m.register(sdkParams)
		if err != nil {
			log.Error("register error: " + err.Error())
			time.Sleep(1 * time.Second)
			//continue
			return false
		}
		if response != nil {
			m.mqttResponse = *response
		}
		break
	}
	m.mqttParams, _ = NewMqParams(m.mqttResponse)
	m.mqttParams.broker = m.mqttParams.broker + ":" + strconv.Itoa(m.mqttParams.port)
	log.Debug("broker == " + m.mqttParams.broker)

	m.runMqttClient()
	return true
}

func (m *Mqtt) runMqttClient() {
	var logLevel = log.InfoLevel
	ctx = &log.Logger{
		Level: logLevel,
		//Handler: NewLogHanler(os.Stdout),
	}

	m.mqttClient = m.NewClient(
		ctx,
		m.mqttParams.deviceId,
		m.mqttParams.secretKey,
		m.mqttParams.broker,
	)

	token := m.mqttClient.Connect()
	token.WaitTimeout(timeOut * time.Second)
	if token.Error() != nil {
		//ctx.WithError(token.Error()).Fatal("Could not connect to MQTT")
		fmt.Printf("Could not connect to MQTT\n")
	} else {
		fmt.Printf("Success connect to MQTT\n")
	}
}

func (m *Mqtt) validTopic(topic string) (bool) {
	return true
}

func (m *Mqtt) readFile(fileName string) string {
	fi, err := os.Open(fileName)
	if err != nil {
		log.Debug("open file error, filename: " + fileName)
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
		DeviceSk:   cs,
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
		companyId:           response.CompanyId,
		clientId:            response.ClientId,
		deviceId:            response.DeviceId,
		broker:              response.Broker,
		port:                response.Port,
		secretKey:           response.SecretKey,
		keepAlive:           response.KeepAlive,
		cleanSession:        response.CleanSession,
		timestamp:           response.TimeStamp,
		mqttSubReceiveMap:   make(map[string]*MqSub),
		mqttSubBroadCastMap: make(map[string]*MqSub),
	}, nil
}
