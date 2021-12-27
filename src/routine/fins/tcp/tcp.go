package tcp

import (
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"
)

const metricBufferSize = 32 * 1024

type mTcp struct {
	conn  	net.Conn
	mutex 	sync.Mutex
	mapInstance		map[int]*myInstance
}

func (m *mTcp) StartMTcp() error {

	log.Println("StartMTcp() in tcp.go")
	return nil
}

func (m *mTcp) Disconnect() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.conn != nil {
		m.conn.Close()
		m.conn = nil
	}
}

func (m *mTcp) Connect() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.conn == nil {
		conn, err := net.Dial("tcp", "127.0.0.1:6555")
		//conn, err := Dial("tcp", m.options.Broker)
		if err != nil {
			return err
		}
		m.conn = conn
	}

	return nil
}

func (m *mTcp) ReceiveData(buf []byte) (n int, err error) {
	err = m.Connect()
	if err != nil {
		return 0, err
	}

	n, err = m.conn.Read(buf)
	if err != nil {
		m.Disconnect()
	}

	return n, err
}

func (m *mTcp) byteToString(B []byte) (S string) {
	for i := 0; i < len(B); i++ {
		if B[i] == 0 {
			return string(B[0:i])
		}
	}
	return string(B)
}

func (m *mTcp) byteValid(B []byte) (b []byte) {
	for i := 0; i < len(B); i++ {
		if B[i] == 0 {
			return B[0:i]
		}
	}
	return B
}

func (m *mTcp) ReceiveMessage() (string, error) {
	data := make([]byte, metricBufferSize)
	n, err := m.ReceiveData(data)
	if err != nil {
		log.Println("err: ", err)
		return "", err
	}

	sData := m.byteToString(data)
	log.Println("receive ", n, " byte, data: ", sData)

	var config = new(Config)
	if err = json.Unmarshal(m.byteValid(data), config); err != nil {
		log.Println("unmarshal err: ", err)
	}
	log.Println("config-protocolId: ", config.Msg.ProtocolId)

	if config.MsgType == "START" {
		go m.StartInstance(config.MsgType, config.MsgVersion, config.Msg.ProtocolId)
	} else if config.MsgType == "STOP" {
		go m.StopInstance(config.MsgType, config.MsgVersion, config.Msg.ProtocolId)
	}


	return sData, nil
}

func (m *mTcp) StartInstance(command string, pVersion string, pId int) {
	mIns := MyInstanceNew()
	m.mapInstance[pId] = mIns
	mIns.StartMyInstance(command, pVersion, pId)
}

func (m *mTcp) StopInstance(command string, pVersion string, pId int) {
	m.mapInstance[pId].StopMyInstance(command, pVersion, pId)
}

func (m *mTcp) Receiving() {
	for {
		msg, err := m.ReceiveMessage()
		if err != nil {
			// log.Println("ReceiveMessage error:%s", err.Error())
			time.Sleep(11 * time.Second)
			continue
		}
		log.Println("msg: ", msg)
	}
}

func MTcpNew() *mTcp {
	return &mTcp{
		mapInstance: make(map[int]*myInstance),
	}
}
