package main

import (
	"encoding/json"
	"fmt"
	"github.com/goburrow/modbus"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
	"time"
)

/**
 * @Project: go_sample
 * @Package Name: modbus
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/5/17 下午10:12
 * @Description:
 */

const (
	configFileName = "demo.conf"
)

/* define new struct Modbus for marshal*/
type Modbus struct {
	filename  string
	handler   ClientHandler
	config    *Config
	client    modbus.Client
	mutex     sync.Mutex
}

/**
 * @Description:
 * @Params:
 * @return:
 * @Date: 2020/5/17 下午10:21
 */
func (m *Modbus) run() (data []string, err error) {
	//return data,nil
	time.Sleep(time.Millisecond * time.Duration(m.config.PollInterval))
	//var msg = make(map[string][]interface{})
	var valueType string
	var slave byte
	var quantity uint16

	myS1 := make([]string, len(m.config.Readers))
	var result []byte
	var i int32 = 0
	for _, reader := range m.config.Readers {
		name := reader.Name
		slave = byte(reader.SlaveId)
		elementLength := reader.Quantity  // todo json name will be change
		valueType = reader.ValueType

		elementStep, ok := InputStepMap[valueType]
		if ok != true{
			fmt.Println("can not recognise type :", valueType)
		}
		if quantity < 1 {
			quantity = 1
		}
		quantity = elementLength * elementStep / 2
		if slave < 1 {
			slave = 1
		}
		m.handler.SetSlaveId(slave)
		address := FillAddress(reader.Address)
		function, offset, err := ParseAddress(address)

		if err != nil {
			log.Println("reader <"+name+"> parse address error:", err.Error())
			continue
		}
		switch function {
		case 0, 1:
			function += 1
		}
		offset -= 1
		m.mutex.Lock()
		switch function {
		case modbus.FuncCodeReadCoils:
			result, err = m.client.ReadCoils(offset, quantity)
		case modbus.FuncCodeReadDiscreteInputs:
			result, err = m.client.ReadDiscreteInputs(offset, quantity)
		case modbus.FuncCodeReadInputRegisters: // exchange FuncCodeReadHoldingRegisters and FuncCodeReadHoldingRegisters
			result, err = m.client.ReadHoldingRegisters(offset, quantity)
		case modbus.FuncCodeReadHoldingRegisters:
			result, err = m.client.ReadInputRegisters(offset, quantity)
		default:
			result, err = nil, fmt.Errorf("not support function code <%d>", function)
		}

		if err != nil {
			m.mutex.Unlock()
			log.Println("reader <"+name+"> error:", err.Error())
			retryCount := 1
			for {
				log.Println("modbus reconnecting ", retryCount, "...")
				if err := m.handler.Reconnect(); err != nil {
					log.Println("modbus reconnect reader <"+name+"> error:", err.Error())
					time.Sleep(time.Second)
					continue
				}
				break
			}
			log.Println("modbus reconnected, continue to do next.")
			m.mutex.Lock()
		}

		//fmt.Println("请求间间隔时间： ", m.config.DpInterval, " 毫秒")
		time.Sleep(time.Duration(m.config.DpInterval) * time.Millisecond)
		m.mutex.Unlock()
		signStr := fmt.Sprintf("%x", result)
		val, _ := strconv.ParseInt(signStr, 16, 10)
		fmt.Println("msg is ", val)
		myS1[i] = signStr
		i++
	}

	return myS1, nil
}

/**
 * @Description:
 * @Params:
 * @return:
 * @Date: 2020/5/17 下午10:35
 */
func (m *Modbus) Destroy() error {
	return m.handler.Close()
}

/**
 * @Description:
 * @Params:
 * @return:
 * @Date: 2020/5/17 下午10:16
 */
func ModbusNew(filename string) (*Modbus, error) {
	var err error
	var config = new(Config)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config.fix()
	if err = json.Unmarshal(data, config); err != nil {
		fmt.Println("Unmarshal error in NewModbus.")
		return nil, err
	}

	var handler ClientHandler
	handler = NewRTUClientHandler(config)

	if err = handler.Connect(); err != nil {
		return nil, err
	}

	client := modbus.NewClient(handler)

	return &Modbus{
		handler:   handler,
		config:    config,
		client:    client,
		filename:  filename,
	}, nil
}


func main() () {
	var c = new(Config)
	c.fix()
	c.Type = "rtu"
	c.PollInterval = 5000
	c.SlaveId = 2
	c.Quantity = 2
	c.ConnTimeout = 5
	c.Readers = append(c.Readers, Reader{
		Name: "demo0",
		Quantity: 1,
		SlaveId:  2,
		Address: "400000",
		ValueType:TYPE_Int16,
	})
	c.Readers = append(c.Readers, Reader{
		Name: "demo1",
		Quantity: 1,
		SlaveId:  2,
		Address: "400001",
		ValueType:TYPE_Int16,
	})
	c.Readers = append(c.Readers, Reader{
		Name: "demo2",
		Quantity: 1,
		SlaveId:  2,
		Address: "400002",
		ValueType:TYPE_Int16,
	})
	c.Readers = append(c.Readers, Reader{
		Name: "demo3",
		Quantity: 1,
		SlaveId:  2,
		Address: "400003",
		ValueType:TYPE_Int16,
	})
	c.Readers = append(c.Readers, Reader{
		Name: "demo4",
		Quantity: 1,
		SlaveId:  2,
		Address: "400004",
		ValueType:TYPE_Int16,
	})
	c.Readers = append(c.Readers, Reader{
		Name: "demo5",
		Quantity: 1,
		SlaveId:  2,
		Address: "400005",
		ValueType:TYPE_Int16,
	})
	c.Readers = append(c.Readers, Reader{
		Name: "demo6",
		Quantity: 1,
		SlaveId:  2,
		Address: "400006",
		ValueType:TYPE_Int16,
	})
	c.Readers = append(c.Readers, Reader{
		Name: "demo7",
		Quantity: 1,
		SlaveId:  2,
		Address: "400007",
		ValueType:TYPE_Int16,
	})
	c.Readers = append(c.Readers, Reader{
		Name: "demo8",
		Quantity: 1,
		SlaveId:  2,
		Address: "400008",
		ValueType:TYPE_Int16,
	})
	c.Readers = append(c.Readers, Reader{
		Name: "demo9",
		Quantity: 1,
		SlaveId:  1,
		Address: "400009",
		ValueType:TYPE_Int16,
	})
	c.Readers = append(c.Readers, Reader{
		Name: "demo10",
		Quantity: 1,
		SlaveId:  1,
		Address: "400010",
		ValueType:TYPE_Int16,
	})
	c.Readers = append(c.Readers, Reader{
		Name: "demo11",
		Quantity: 1,
		SlaveId:  2,
		Address: "400011",
		ValueType:TYPE_Int16,
	})
	c.Readers = append(c.Readers, Reader{
		Name: "demo12",
		Quantity: 1,
		SlaveId:  2,
		Address: "400012",
		ValueType:TYPE_Int16,
	})



	if err := c.ToFile(configFileName); err != nil {
		fmt.Println("write to file error: ", err.Error())
	}

	mod, err := ModbusNew(configFileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer mod.Destroy()

	// polling read
	if mod != nil {
		for i := 0; i < 1; i++ {
			data, errRun := mod.run()
			if errRun != nil {
				fmt.Println("read error: ", errRun.Error())
			}
			fmt.Println("get data: ", data)
		}
	}
}


