package main

import (
	"github.com/go_sample/src/go-modbus2/modbus"
	"log"
	"time"
	"bytes"
	"encoding/binary"
	"strings"
	"os"
	"io/ioutil"
	"strconv"
)

//
var gClient modbus.Client

// env map
var EnvMap = make(map[string]string)
var RtuDevice = "rtudevice"
var BaudRate = "baudrate"
var DataBits = "databits"
var Parity = "parity"
var StopBits = "stopbits"
var SlaveId = "slaveid"
var SerialTimeout = "serialtimeout"
var Address = "address"
var Len = "len"

// conf default value
const(
	DefaultRtuDevcie = "/dev/ttyUSB0"
	DefaultBaudRate = "19200"
	DefaultDataBits = "8"
	DefaultParity = "N"
	DefaultStopBits = "1"
	DefaultSlaveId = "1"
	DefaultSerialTimeout = "5"
	DefaultAddress = "1105"
	DefaultLen = "20"
)

const (
	workerConf = "modbus.conf"
)

func ReadFile(fileName string) string {
	fi,err := os.Open(fileName)
	if err != nil{
		log.Println("did not find config, use default.")
		return ""
	}
	defer fi.Close()
	fd,err := ioutil.ReadAll(fi)
	return string(fd)
}

func initConfMap() {
	strs := ReadFile(workerConf)
	if strs != "" {
		result := []string{}
		for _, lineStr := range strings.Split(strs, "\n") {
			lineStr = strings.TrimSpace(lineStr)
			if lineStr == "" {
				continue
			}
			result = strings.Split(lineStr, "=")
			k := result[0]
			v := result[1]
			EnvMap[k] = v
		}
	}

	// give a default value
	if EnvMap[RtuDevice] == "" {
		EnvMap[RtuDevice] = DefaultRtuDevcie
	}
	if EnvMap[BaudRate] == "" {
		EnvMap[BaudRate] = DefaultBaudRate
	}
	if EnvMap[DataBits] == "" {
		EnvMap[DataBits] = DefaultDataBits
	}
	if EnvMap[Parity] == "" {
		EnvMap[Parity] = DefaultParity
	}
	if EnvMap[StopBits] == "" {
		EnvMap[StopBits] = DefaultStopBits
	}
	if EnvMap[SlaveId] == "" {
		EnvMap[SlaveId] = DefaultSlaveId
	}
	if EnvMap[SerialTimeout] == "" {
		EnvMap[SerialTimeout] = DefaultSerialTimeout
	}
	if EnvMap[Address] == "" {
		EnvMap[Address] = DefaultAddress
	}
	if EnvMap[Len] == "" {
		EnvMap[Len] = DefaultLen
	}
}

func byteToInt(results []byte) int32 {
	bBuf := bytes.NewBuffer(results)
	var x int32
	binary.Read(bBuf, binary.BigEndian, &x)
	return x
}

func initTest() {
	initConfMap()
	handler := modbus.NewRTUClientHandler(EnvMap[RtuDevice])
	baudRate, _ := strconv.Atoi(EnvMap[BaudRate])
	handler.BaudRate = baudRate
	dataBits, _ := strconv.Atoi(EnvMap[DataBits])
	handler.DataBits = dataBits
	handler.Parity = EnvMap[Parity]
	stopBits, _ := strconv.Atoi(EnvMap[StopBits])
	handler.StopBits = stopBits
	slaveId, _ := strconv.Atoi(EnvMap[SlaveId])
	handler.SlaveId = byte(slaveId)
	serialTimeout, _ := strconv.Atoi(EnvMap[SerialTimeout])
	handler.Timeout = time.Duration(serialTimeout) * time.Second

	err := handler.Connect()
	if err != nil {
		log.Println("err: ", err)
	}
	defer handler.Close()

	client := modbus.NewClient(handler)

	gClient = client
}

func getModbusValue(client modbus.Client) []string {
	var i uint16
	inputLen, _ := strconv.ParseInt(EnvMap[Len], 10, 16)
	len := uint16(inputLen)
	inputAddr, _ := strconv.Atoi(EnvMap[Address])
	address := uint16(inputAddr)
	//log.Println("len: ", len)
	myS1 := make([]string, len)
	for i=0; i<len; i++ {
		results, err := gClient.ReadHoldingRegisters(address+i, 1)
		if err == nil {
			//log.Println("label = ", i)
		} else {
			log.Println("err: ", err)
		}

		a := byteToInt([]uint8{0, 0, 0, results[0]})
		b := byteToInt([]uint8{0, 0, 0, results[1]})
		var temperature = float64((a*256)+b) / 100
		//log.Println("temperature: ", temperature)
		myS1[i] = strconv.FormatFloat(temperature, 'f', 2, 64)
	}
	return myS1
}

func main() {

	initTest()

	array := getModbusValue(gClient)

	log.Println("array: ", array)
}
