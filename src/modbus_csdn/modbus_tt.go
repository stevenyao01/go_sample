package main

/**
 * @Project: modbus_csdn
 * @Package Name: modbus
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/5/18 下午10:12
 * @Description:
 */

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/goburrow/modbus"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

//
var gClient modbus.Client

// env map
var EnvMap = map[string]string {
	"rtudevice":"/dev/ttyUSB0",
	"baudrate":"9600",
	"databits":"8",
	"parity":"N",
	"stopbits":"1",
	"slaveid":"1",
	"serialtimeout":"10",
	"address":"0",
	"len":"20",
	"coils_address":"",
	"coils_operate":"",
	"coils_operate_len":"1",
	"inputregisters_address":"",
	"inputregisters_operate":"",
	"inputregisters_operate_len":"0",
	// config you holdingregisters here
	"holdingregisters_address":"13",
	"holdingregisters_operate":"read",
	"holdingregisters_operate_len":"1",
	"discreteinputs_address":"",
	"discreteinputs_operate":"",
	"discreteinputs_operate_len":"1",
}
var RtuDevice = "rtudevice"
var BaudRate = "baudrate"
var DataBits = "databits"
var Parity = "parity"
var StopBits = "stopbits"
var SlaveId = "slaveid"
var SerialTimeout = "serialtimeout"
var Address = "address"
var Len = "len"
var CoilsAddress = "coils_address"
var CoilsOperate = "coils_operate"
var CoilsOperateLen = "coils_operate_len"
var InputRegistersAddress = "inputregisters_address"
var InputRegistersOperate = "inputregisters_operate"
var InputRegistersOperateLen = "inputregisters_operate_len"
var HoldingRegistersAddress = "holdingregisters_address"
var HoldingRegistersOperate = "holdingregisters_operate"
var HoldingRegistersOperateLen = "holdingregisters_operate_len"
var DiscreteInputsAddress = "discreteinputs_address"
var DiscreteInputsOperate = "discreteinputs_operate"
var DiscreteInputsOperateLen = "discreteinputs_operate_len"



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
	strFile := ReadFile(workerConf)
	if strFile != "" {
		result := []string{""}
		for _, lineStr := range strings.Split(strFile, "\n") {
			lineStr = strings.TrimSpace(lineStr)
			if lineStr == "" {
				continue
			}
			result = strings.Split(lineStr, "=")
			k := result[0]
			v := result[1]
			if v != "" {
				EnvMap[k] = v
			}
		}
	}
}

func byteToInt(results []byte) int {
	bBuf := bytes.NewBuffer(results)
	var x int
	binary.Read(bBuf, binary.BigEndian, &x)
	return x
}

func initConfig() {
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

func getCoils() []string {
	var i uint16
	inputLen, _ := strconv.ParseInt(EnvMap[CoilsOperateLen], 10, 16)
	askLen := uint16(inputLen)
	inputAddr, _ := strconv.Atoi(EnvMap[CoilsAddress])
	address := uint16(inputAddr)
	//log.Println("len: ", len)
	myS1 := make([]string, askLen)
	for i=0; i<askLen; i++ {
		results, err := gClient.ReadCoils(address+i, 1)
		if err != nil {
			log.Println("err: ", err)
			return nil
		}
		myS1[i] = strconv.Itoa(byteToInt(results))
	}
	return myS1
}

func getInputRegisters() []string {
	var i uint16
	inputLen, _ := strconv.ParseInt(EnvMap[InputRegistersOperateLen], 10, 16)
	askLen := uint16(inputLen)
	inputAddr, _ := strconv.Atoi(EnvMap[InputRegistersAddress])
	address := uint16(inputAddr)
	//log.Println("len: ", len)
	myS1 := make([]string, askLen)
	for i=0; i<askLen; i++ {
		results, err := gClient.ReadInputRegisters(address+i, 1)
		if err != nil {
			log.Println("err: ", err)
			return nil
		}
		myS1[i] = strconv.Itoa(byteToInt(results))
	}
	return myS1
}

func getHoldingRegisters() []string {
	var i uint16
	inputLen, _ := strconv.ParseInt(EnvMap[HoldingRegistersOperateLen], 10, 16)
	askLen := uint16(inputLen)
	inputAddr, _ := strconv.Atoi(EnvMap[HoldingRegistersAddress])
	address := uint16(inputAddr)
	//log.Println("len: ", len)
	myS1 := make([]string, askLen)
	for i=0; i<askLen; i++ {
		results, err := gClient.ReadHoldingRegisters(address+i, 1)
		if err != nil {
			log.Println("err: ", err)
			return nil
		}
		signStr := fmt.Sprintf("%x", results)

		// val, _ := strconv.ParseInt(signStr, 16, 10)
		// log.Println("val: ", val)

		myS1[i] = signStr
		time.Sleep(100 * time.Millisecond)
	}
	return myS1
}

func getDiscreteInputs() []string {
	var i uint16
	inputLen, _ := strconv.ParseInt(EnvMap[DiscreteInputsOperateLen], 10, 16)
	askLen := uint16(inputLen)
	inputAddr, _ := strconv.Atoi(EnvMap[DiscreteInputsAddress])
	address := uint16(inputAddr)
	//log.Println("len: ", len)
	myS1 := make([]string, askLen)
	for i=0; i<askLen; i++ {
		results, err := gClient.ReadDiscreteInputs(address+i, 1)
		if err != nil {
			log.Println("err: ", err)
			return nil
		}
		myS1[i] = strconv.Itoa(byteToInt(results))
	}
	return myS1
}

func main() {

	initConfig()

	if EnvMap[CoilsAddress] != "" {
		log.Println("do process coils.")
		if EnvMap[CoilsOperate] == "read" {
			log.Println("read coils.")
			array := getCoils()
			if array == nil {
				log.Println("read nil from modbus!!!")
			} else {
				log.Println("array: ", array)
			}
		} else {
			log.Println("write coils.")
		}
	}
	if EnvMap[InputRegistersAddress] != "" {
		log.Println("do process InputRegisters.")
		if EnvMap[InputRegistersOperate] == "read" {
			log.Println("read InputRegisters.")
			array := getInputRegisters()
			if array == nil {
				log.Println("read nil from modbus!!!")
			} else {
				log.Println("array: ", array)
			}
		} else {
			log.Println("write InputRegisters.")
		}
	}
	if EnvMap[HoldingRegistersAddress] != "" {
		log.Println("do process HoldingRegisters.")
		if EnvMap[HoldingRegistersOperate] == "read" {
			log.Println("read HoldingRegisters.")
			array := getHoldingRegisters()
			if array == nil {
				log.Println("read nil from modbus!!!")
			} else {
				log.Println("array: ", array)
			}
		} else {
			log.Println("write HoldingRegisters.")
		}
	}

	if EnvMap[DiscreteInputsAddress] != "" {
		log.Println("do process DiscreteInputs.")
		if EnvMap[DiscreteInputsOperate] == "read" {
			log.Println("read DiscreteInputs.")
			array := getDiscreteInputs()
			if array == nil {
				log.Println("read nil from modbus!!!")
			} else {
				log.Println("array: ", array)
			}
		} else {
			log.Println("write DiscreteInputs.")
		}
	}
}