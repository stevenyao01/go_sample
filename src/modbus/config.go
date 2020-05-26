package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/goburrow/modbus"
	"io/ioutil"
	"strconv"
	"strings"
)

/**
 * @Project: go_sample
 * @Package Name: main
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/5/17 下午10:27
 * @Description:
 */

// default config
const (
	defaultType = "rtu"
	// connect timeout
	defaultConnTimeout = 3000
	// polling interval
	defaultPollInterval = 1000
	// read/write timeout
	defaultQueryTimeout = 1000
	// read/write interval
	defaultQueryInterval = 60

	defaultDevice   = "/dev/ttyUSB0"
	defaultBaudRate = 4800
	defaultDataBits = 8
	defaultParity   = "N"
	defaultStopBits = 1
)

var TYPE_HexData = "HexData"  //
var TYPE_Boolean = "Boolean"    //
var TYPE_Int16 = "Short"    // short   === int16
var TYPE_UInt16 = "Word"  // word  BCD   === uint16
var TYPE_UInt16_BCD = "BCD"   // BCD   === uint16
var TYPE_Int32 = "Long"    // long === int32
var TYPE_UInt32 = "DWord"   // dword  LBCD   === uint32
var TYPE_UInt32_LBCD = "LBCD"
var TYPE_Int64 = "int64"
var TYPE_Float32 = "Float"   // float   === float32
var TYPE_Float64 = "Double"    // Double  === float64
var TYPE_String = "String"       // char

var TYPE_Boolean_Array = "Boolean_ARRAY"  // ----delete
var TYPE_Int16_Array = "Short_ARRAY"    // short
var TYPE_UInt16_Array = "Word_ARRAY"   // word_array
var TYPE_UInt16_BCD_Array = "BCD_ARRAY"
var TYPE_Int32_Array = "Long_ARRAY"     // long_array
var TYPE_UInt32_Array = "DWord_ARRAY"    // Dword_array
var TYPE_UInt32_LBCD_Array = "LBCD_ARRAY"
var TYPE_Int64_Array = "int64_ARRAY"      // ----  delete
var TYPE_Float32_Array = "Float_ARRAY"    // float_array
var TYPE_Float64_Array = "Double_ARRAY"     // double array

var InputStepMap  = map[string]uint16{  // number of byte
	TYPE_HexData:2,
	TYPE_Boolean:2,
	TYPE_UInt16:2,
	TYPE_UInt16_BCD:2,
	TYPE_Int16:2,
	TYPE_Int32:4,
	TYPE_UInt32:4,
	TYPE_UInt32_LBCD:4,
	TYPE_Int64:8,
	TYPE_Float32:4,
	TYPE_Float64:8,
	TYPE_String:2,

	TYPE_Boolean_Array:2,
	TYPE_Int16_Array:2,
	TYPE_UInt16_Array:2,
	TYPE_UInt16_BCD_Array:2,
	TYPE_Int32_Array:4,
	TYPE_UInt32_Array:4,
	TYPE_UInt32_LBCD_Array:4,
	TYPE_Int64_Array:8,
	TYPE_Float32_Array:4,
	TYPE_Float64_Array:8,
}

type ClientHandler interface {
	modbus.ClientHandler
	Connect() error
	Close() error
	Reconnect() error
	SetSlaveId(id byte)
}

type Reader struct {
	Name string `json:"name"`
	//Offset   uint16 `json:"offset"`
	SlaveId  int         `json:"slave_id"`
	Address  interface{} `json:"address"`
	ValueType string     `json:"valueType"`
	Quantity uint16      `json:"quantity"`
	//FuncCode int    `json:"func_code"`
}

type RTU struct {
	Device   string `json:"device"`
	BaudRate int    `json:"baud_rate"`
	DataBits int    `json:"data_bits"`
	Parity   string `json:"parity"`
	StopBits int    `json:"stop_bits"`
}

type Config struct {
	// connection type(tcp/rtu/ascii)
	Type 			string `json:"type"`
	// interval/timeout config (ms)
	// read/poll interval, read/write connect timeout.
	//QueryInterval int `json:"query_interval"`
	PollInterval 	int `json:"poll_interval"`
	//QueryTimeout  int `json:"query_timeout"`
	ConnTimeout 	int `json:"conn_timeout"`
	// delay between request
	DpInterval 		int `json:"dp_interval"`
	// readers
	Readers   		[]Reader  `json:"readers"`
	RTU 			RTU `json:"rtu"`

	SlaveId  		int   `json:"slave_id"` // HID
	Quantity 		int   `json:"quantity"`
}

func (c *Config) fix() {
	c.Type = defaultType
	if c.ConnTimeout < 0 {
		c.ConnTimeout = defaultConnTimeout
	}
	if c.PollInterval < 0 {
		c.PollInterval = defaultPollInterval
	}

	if c.RTU.Device == "" {
		c.RTU.Device = defaultDevice
	}
	if c.RTU.BaudRate < 1 {
		c.RTU.BaudRate = defaultBaudRate
	}
	if c.RTU.DataBits < 1 {
		c.RTU.DataBits = defaultDataBits
	}
	if c.RTU.Parity == "" {
		c.RTU.Parity = defaultParity
	}
	if c.RTU.StopBits < 1 {
		c.RTU.StopBits = defaultStopBits
	}
}

func (c *Config) String() string {
	return string(c.Json())
}

func (c *Config) Json() []byte {
	j, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		panic(err)
	}

	return j
}

func (c *Config) ToFile(filename string) error {
	c.fix()

	return ioutil.WriteFile(filename, c.Json(), 0755)
}

func FillAddress(address interface{}) (interface{}) {
	switch addr := address.(type) {
	case string:
		length := len(addr)
		if length >= 6 {
			return address
		}
		var buffer bytes.Buffer
		buffer.WriteString(addr[:1])
		buffer.WriteString(strings.Repeat("0", (6 - length)))
		buffer.WriteString(addr[1:])
		return buffer.String()
	}
	return address
}

func parseAddress(address string) (function, offset uint16, err error) {
	if len(address) != 6 {
		return 0, 0, fmt.Errorf("address '%v' length not equal 6", address)
	}

	fun, err := strconv.Atoi(address[:1])
	if err != nil {
		return 0, 0, err
	}

	function = uint16(fun)
	off, err := strconv.Atoi(address[1:])
	if err != nil {
		return 0, 0, err
	}
	if off > int(^uint16(0)) {
		return 0, 0, fmt.Errorf("out of uint16(0 - %d)", ^uint16(0))
	}
	offset = uint16(off)

	return
}

func ParseAddress(address interface{}) (function, offset uint16, err error) {
	switch addr := address.(type) {
	case string:
		if len(addr) == 6 {
			return parseAddress(addr)
		}
		a := strings.TrimLeft(addr, "0 ")
		if len(a) <= 6 {
			return parseAddress(fmt.Sprintf("%06s", a))
		}
		return 0, 0, fmt.Errorf("invalid address '%v'", addr)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return parseAddress(fmt.Sprintf("%06d", address))
	}
	return 0, 0, fmt.Errorf("not support type '%T':%+v", address, address)
}
