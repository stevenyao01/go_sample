package process

import (
	"encoding/json"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

const metricBufferSize = 32 * 1024

type mTcp struct {
	conn              map[string]net.Conn
	listener          net.Listener
	buffer            chan []byte
	mutex             sync.Mutex
}

func (m *mTcp) store(conn net.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.conn[conn.RemoteAddr().String()] = conn
}

func (m *mTcp) delete(conn net.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.conn, conn.RemoteAddr().String())
}

func (m *mTcp) receive(conn net.Conn) {
	m.store(conn)
	buffer := make([]byte, metricBufferSize)
	for {
		n,err := conn.Read(buffer)
		if err != nil {
			log.Println("receive error:%s", err.Error())
			break
		}

		m.buffer <- buffer[:n]
	}
	conn.Close()
	m.delete(conn)
}

func (m *mTcp) accept(listener net.Listener) {
	for true {
		conn,err := listener.Accept()
		if err != nil {
			log.Println("accept error:%s", err.Error())
			break
		}

		log.Println("new connected:%s", conn.RemoteAddr().String())
		conf := new(Config)
		conf.Msg.ProtocolId = 110
		conf.Msg.ProtocolMd5 = "7bdd67dfa9db0fe92ce8c75819eaf33a"
		conf.MsgVersion = "1.1"
		//conf.IpConf = "protocol_110_input.conf"
		//conf.OpConf = "protocol_111_output.conf"
		//conf.OtConf = "protocol_112_other.conf"



		for i := 0; i < 10; i++ {
			conf.MsgType = "START"
			conf.Msg.ProtocolId = 100 + i
			confCmd,_ := json.Marshal(conf)
			log.Println("confCmd", string(confCmd))
			_, _ = conn.Write(confCmd)
			time.Sleep(1 * time.Second)
		}

		time.Sleep(10 * time.Second)

		for i := 0; i < 10; i++ {
			conf.MsgType = "STOP"
			conf.Msg.ProtocolId = 100 + i
			confCmd,_ := json.Marshal(conf)
			log.Println("confCmd", string(confCmd))
			_, _ = conn.Write(confCmd)
			time.Sleep(1 * time.Second)
		}


		go m.receive(conn)
	}
}

func (m *mTcp) Receiving() {
	go m.accept(m.listener)

	for true {
		data := <- m.buffer

		log.Println("data: ", data)

		//var mm MetricMessage
		//err := json.Unmarshal(data, &mm)
		//if err != nil {
		//	log.Error("receiving json unmarshal error:%s", err.Error())
		//	log.Error("receive json:%s", string(data))
		//	continue
		//}
		//switch mm.Type {
		//case metricCount:
		//	go m.SendCount([]byte(mm.Data))
		//case metricAlarm:
		//	go m.SendAlarm([]byte(mm.Data))
		//case metricBroadcast:
		//	go m.SendBroadcast([]byte(mm.Data))
		//default:
		//	log.Error("receiving error:%s", invalidMessageType.Error())
		//}
	}

	return
}

func MTcpNew() *mTcp {
	var addr = ":"
	addr += strconv.Itoa(6555)

	mListener,err := net.Listen("tcp", addr)
	//listener,err := netutil.Listen("tcp", addr)
	if err != nil {
		log.Println("Listen error: ",err)
	}
	log.Println("jjjjj: ", mListener.Addr().String())
	return &mTcp{
		listener:	mListener,
		conn:		make(map[string]net.Conn),
	}
}