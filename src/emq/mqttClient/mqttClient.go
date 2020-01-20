package mqttClient

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"time"
)

type MqttClient struct {
	// input param
	broker    string
	clientId  string
	userName  string
	subTopics map[string]sub
	pubTopic  string
	qos       byte

	// self param
	c    mqtt.Client
	opts *mqtt.ClientOptions

	// status
	isConnected bool
}

type sub struct {
	topic    string
	qos      byte
	callback mqtt.MessageHandler
}

const (
	connectTimeOut     = 10 * time.Second
	publishTimeOut     = 10 * time.Second
	subscribeTimeOut   = 10 * time.Second
	unsubscribeTimeOut = 10 * time.Second
)

type CbReceive func(topic string, msg mqtt.Message)

type Client interface {
	Init() error
	UnInit()
	IsConnect() bool
	SendMessage(topic string, qos byte, data []byte) (n int, err error)
	ReceiveMessage(topic string, qos byte, callback CbReceive) error
	UnReceiveMessage(topic string)
}

// interface for third application
/*

*/
func (m *MqttClient) Init() error {
	m.configure()
	err := m.connect()
	if err != nil {
		return err
	}
	return nil
}

func (m *MqttClient) UnInit() {
	err := m.destroy()
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}
	fmt.Println("destroy success.")
	return
}

func (m *MqttClient) IsConnected() bool {
	if m.c == nil {
		return false
	}
	opt := m.c.OptionsReader()
	if !opt.AutoReconnect() {
		return m.c.IsConnected()
	}
	return m.isConnected
}

func (m *MqttClient) SendMessage(topic string, qos byte, data []byte) (n int, err error) {
	return m.publish(topic, qos, false, data)
}

func (m *MqttClient) ReceiveMessage(topic string, qos byte, callback CbReceive) error {
	return m.receiveMessage(topic, qos, func(i mqtt.Client, message mqtt.Message) {
		callback(topic, message)
	})
}

func (m *MqttClient) UnReceiveMessage(topic string) {
	m.unSubscribe(topic)
	return
}

// nation method
func (m *MqttClient) receiveMessage(topic string, qos byte, callback mqtt.MessageHandler) error {
	return m.subscribe(topic, qos, nil) //callback)
}

//set callback function
//var callBack mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
//	fmt.Printf("yhp TOPIC: %s\n", msg.Topic())
//	fmt.Printf("yhp MSG: %s\n", msg.Payload())
//}

func newTLSConfigSingle(filename string) (*tls.Config, error) {
	cert := x509.NewCertPool()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	cert.AppendCertsFromPEM(data)

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: cert,
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
	}, err
}

func newTLSConfigDouble(filename, certFile, keyFile string) (*tls.Config, error) {
	// Import trusted certificates from CAfile.pem.
	// Alternatively, manually add CA certificates to
	// default openssl CA bundle.
	certPool := x509.NewCertPool()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	certPool.AppendCertsFromPEM(data)

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return nil, err
	}

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certPool,
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
	}, nil
}

func (m *MqttClient) destroy() error {
	m.c.Disconnect(250)
	return nil
}

func (m *MqttClient) unSubscribe(topic ...string) {
	token := m.c.Unsubscribe(topic...)
	if !token.WaitTimeout(unsubscribeTimeOut) {
		fmt.Println("error: mqtt unsubscribe timeout.")
	}

	if token.Error() == nil {
		for _, t := range topic {
			if _, ok := m.subTopics[t]; ok {
				delete(m.subTopics, t)
			}
		}
	} else {
		fmt.Println("error: token.Error: ", token.Error())
	}
	fmt.Println("unSubscribe topic ", topic, " ok.")
	return
}

func (m *MqttClient) subscribe(topic string, qos byte, callback mqtt.MessageHandler) error {
	token := m.c.Subscribe(topic, m.qos, callback)
	if !token.WaitTimeout(subscribeTimeOut) {
		fmt.Println("error: mqtt subscribe timeout.")
	}
	if token.Error() == nil {
		m.subTopics[topic] = sub{topic: topic, qos: qos, callback: callback}
	}
	return token.Error()
}

func (m *MqttClient) publish(topic string, qos byte, retained bool, data []byte) (n int, err error) {
	token := m.c.Publish(topic, qos, retained, data)
	if !token.WaitTimeout(publishTimeOut) {
		fmt.Println("error: mqtt publish timeout.")
	}

	return len(data), token.Error()
}

func (m *MqttClient) getSubTopics() map[string]sub {
	var topics = make(map[string]sub)
	for key, val := range m.subTopics {
		topics[key] = val
	}
	return topics
}

func (m *MqttClient) configure() () {
	// init options
	m.opts = mqtt.NewClientOptions()
	m.opts.AddBroker(m.broker)
	m.opts.SetClientID(m.clientId)
	m.opts.SetUsername(m.userName)
	m.opts.SetKeepAlive(2 * time.Second)
	m.opts.SetPingTimeout(1 * time.Second)

	// set auto reconnect
	m.opts.SetAutoReconnect(true)
	// set clean session
	m.opts.SetCleanSession(true)
	// set onConnect handler
	m.opts.SetOnConnectHandler(func(i mqtt.Client) {
		m.isConnected = true
		topics := m.getSubTopics()
		for len(topics) != 0 {
			fmt.Println("Resubscribe begin!")
			for key, val := range topics {
				if err := m.subscribe(val.topic, val.qos, val.callback); err != nil {
					fmt.Println("error: ReSubscribe error: ", err.Error())
				} else {
					fmt.Println("ReSubscribe success: ", val.topic)
					delete(topics, key)
				}
			}
		}
		fmt.Println("Connect Success!")
	})
	// set connect lost handler
	m.opts.SetConnectionLostHandler(func(i mqtt.Client, e error) {
		m.isConnected = false
		fmt.Println("Connect Lost: ", e.Error())
	})
	// set default handler for new receive msg
	//m.opts.SetDefaultPublishHandler(callBack)
	m.opts.SetDefaultPublishHandler(func(i mqtt.Client, message mqtt.Message) {
		fmt.Println("TOPIC: ", message.Topic())
		fmt.Println("MSG: ", string(message.Payload()))
	})
	// set mqtt ssl access
	tlsConfig, err := newTLSConfigSingle("server.crt")
	if err != nil {
		fmt.Println("newTLSConfigSingle error: ", err.Error())
		tlsConfig = &tls.Config{RootCAs: x509.NewCertPool(), ClientAuth: tls.NoClientCert, ClientCAs: nil, InsecureSkipVerify: true}
	}
	if tlsConfig != nil {
		m.opts.SetTLSConfig(tlsConfig)
	}

	// init client
	m.c = mqtt.NewClient(m.opts)

	return
}

func (m *MqttClient) connect() (err error) {
	token := m.c.Connect()
	if !token.WaitTimeout(connectTimeOut) {
		fmt.Println("error: mqtt connect timeout.")
	}
	return token.Error()
}

func NewMqttClient(b string, ci string, un string, q byte) (*MqttClient, error) {
	return &MqttClient{
		broker:      b,
		clientId:    ci,
		userName:    un,
		subTopics:   make(map[string]sub),
		qos:         q,
		isConnected: false,
	}, nil
}
