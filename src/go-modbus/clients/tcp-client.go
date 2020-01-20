package main

import (
	"flag"
	"fmt"
	"github.com/dpapathanasiou/go-modbus"
	"log"
)

func main() {

	// get device host (url or ip address) and port from the command line
	var (
		host string
		port int
	)
	const (
		defaultHost = "127.0.0.1"
		defaultPort = modbusclient.MODBUS_PORT
	)

	flag.StringVar(&host, "host", defaultHost, "Slave device host (url or ip address)")
	flag.IntVar(&port, "port", defaultPort, fmt.Sprintf("Slave device port (the default is %d)", defaultPort))
	flag.Parse()

	// turn on the debug trace option, to see what is being transmitted
	trace := true

	conn, cerr := modbusclient.ConnectTCP(host, port)
	if cerr != nil {
		log.Println(fmt.Sprintf("Connection error: %s", cerr))
	} else {

		// attempt to read one (0x01) holding registers starting at address 200
		readData := make([]byte, 3)
		readData[0] = byte(200 >> 8)   // (High Byte)
		readData[1] = byte(200 & 0xff) // (Low Byte)
		readData[2] = 0x01

		// make this read request transaction id 1, with a 300 millisecond tcp timeout
		readResult, readErr := modbusclient.TCPRead(conn, 300, 1, modbusclient.FUNCTION_READ_HOLDING_REGISTERS, false, 0x00, readData, trace)
		if readErr != nil {
			log.Println(readErr)
		}
		log.Println(readResult)

		// attempt to write to a single coil at address 300
		writeData := make([]byte, 3)
		writeData[0] = byte(300 >> 8)   // (High Byte)
		writeData[1] = byte(300 & 0xff) // (Low Byte)
		writeData[2] = 0xff             // 0xff turns the coil on; 0x00 turns the coil off
		// make this read request transaction id 2, with a 300 millisecond tcp timeout
		writeResult, writeErr := modbusclient.TCPWrite(conn, 300, 2, modbusclient.FUNCTION_WRITE_SINGLE_COIL, false, 0x00, writeData, trace)
		if writeErr != nil {
			log.Println(writeErr)
		}
		log.Println(writeResult)

		modbusclient.DisconnectTCP(conn)
	}
}
