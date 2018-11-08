package cmd

/**
 * @Package Name: mqtt
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-11-2 下午4:13
 * @Description:
 */

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"github.com/apex/log"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var ctx log.Interface

const QoS = 0x02

func init() {
	fmt.Printf("init mqtt test\n")

}

func RunMqttClient() {
	fmt.Printf("Run mqtt test\n")
	var logLevel = log.InfoLevel
	ctx = &log.Logger{
		Level:   logLevel,
		//Handler: NewLogHanler(os.Stdout),
	}

	mqttClient := NewClient(
		ctx,
		"ttnhdl",
		"",
		"",
		fmt.Sprintf("ssl://%s", "192.168.195.201:8883"),
	)

	var err = mqttClient.Connect()
	if err != nil {
		ctx.WithError(err).Fatal("Could not connect to MQTT")
		fmt.Printf("Could not connect to MQTT\n")
	} else {
		fmt.Printf("Success connect to MQTT\n")
	}

	mqttClient.PublishUplink("test", "hello mqtt!")
	mqttClient.SubscribeUplink("test")

	for true {
		log.Info("hello!!!!")
	}
}

// Client connects to the MQTT server and can publish/subscribe on uplink, downlink and activations from devices
type Client interface {
	Connect() error
	Disconnect()

	IsConnected() bool

	// Uplink pub/sub
	PublishUplink(topic string, msg string) Token
	SubscribeUplink(topic string) Token
}

type Token interface {
	Wait() bool
	WaitTimeout(time.Duration) bool
	Error() error
}

type simpleToken struct {
	err error
}

// Wait always returns true
func (t *simpleToken) Wait() bool {
	return true
}

// WaitTimeout always returns true
func (t *simpleToken) WaitTimeout(_ time.Duration) bool {
	return true
}

// Error contains the error if present
func (t *simpleToken) Error() error {
	return t.err
}

type defaultClient struct {
	mqtt MQTT.Client
	ctx  log.Interface
}

func NewClient(ctx log.Interface, id, username, password string, brokers ...string) Client {
	tlsconfig := NewTLSConfig()

	mqttOpts := MQTT.NewClientOptions()

	for _, broker := range brokers {
		mqttOpts.AddBroker(broker)
	}

	mqttOpts.SetClientID("ypf_dewqfvcdeqfcdqwcdq")
	mqttOpts.SetUsername(username)
	mqttOpts.SetPassword(password)

	// TODO: Some tuning of these values probably won't hurt:
	mqttOpts.SetKeepAlive(30 * time.Second)
	mqttOpts.SetPingTimeout(10 * time.Second)

	// Usually this setting should not be used together with random ClientIDs, but
	// we configured The Things Network's MQTT servers to handle this correctly.
	mqttOpts.SetCleanSession(false)

	mqttOpts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		ctx.WithField("message", msg).Warn("Received unhandled message")
	})

	mqttOpts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
		ctx.WithError(err).Warn("Disconnected, reconnecting...")
	})

	mqttOpts.SetOnConnectHandler(func(client MQTT.Client) {
		ctx.Debug("Connected")
	})

	mqttOpts.SetTLSConfig(tlsconfig)

	return &defaultClient{
		mqtt: MQTT.NewClient(mqttOpts),
		ctx:  ctx,
	}
}

var (
	// ConnectRetries says how many times the client should retry a failed connection
	ConnectRetries = 10
	// ConnectRetryDelay says how long the client should wait between retries
	ConnectRetryDelay = time.Second
)

func (c *defaultClient) Connect() error {
	if c.mqtt.IsConnected() {
		return nil
	}
	var err error
	for retries := 0; retries < ConnectRetries; retries++ {
		token := c.mqtt.Connect()
		token.Wait()
		err = token.Error()
		if err == nil {
			break
		}
		<-time.After(ConnectRetryDelay)
	}
	if err != nil {
		return fmt.Errorf("Could not connect: %s", err)
	}
	return nil
}

func (c *defaultClient) Disconnect() {
	if !c.mqtt.IsConnected() {
		return
	}
	c.mqtt.Disconnect(25)
}

func (c *defaultClient) IsConnected() bool {
	return c.mqtt.IsConnected()
}

func (c *defaultClient) PublishUplink(topic string, msg string) Token {
	return c.mqtt.Publish(topic, QoS, false, msg)
}

func (c *defaultClient) SubscribeUplink(topic string) Token {
	return c.mqtt.Subscribe(topic, QoS, func(mqtt MQTT.Client, msg MQTT.Message) {
		// Determine the actual topic
		fmt.Printf("Success SubscribeUplink with msg:%s\n", msg.Payload())
	})
}

func NewTLSConfig() *tls.Config {
	// Import trusted certificates from CAfile.pem.
	// Alternatively, manually add CA certificates to
	// default openssl CA bundle.
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile("samplecerts/ca.pem")
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}
	fmt.Println("0. resd pemCerts Success")

	// Import client certificate/key pair
	cert, err := tls.LoadX509KeyPair("samplecerts/client-crt.pem", "samplecerts/client-key.pem")
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
