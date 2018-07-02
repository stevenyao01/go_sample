package main

import (
	"flag"
	"fmt"
	"github.com/dpapathanasiou/go-modbus"
	"log"
)

func main() {

	// get the device serial port from the command line
	var (
		serialPort    string
		slaveDevice   int
		startAddr     int
		numBytes      int
		baudRate      int
		responsePause int
	)

	const (
		defaultPort          = ""
		defaultSlave         = 1
		defaultStartAddress  = 3030
		defaultNumBytes      = 16
		defaultBaudRate      = 9600
		defaultResponsePause = 300
	)

	flag.StringVar(&serialPort, "serial", defaultPort, "Serial port (RS485) to use, e.g., /dev/ttyS0 (try \"dmesg | grep tty\" to find)")
	flag.IntVar(&slaveDevice, "slave", defaultSlave, fmt.Sprintf("Slave device number (default is %d)", defaultSlave))
	flag.IntVar(&startAddr, "start", defaultStartAddress, fmt.Sprintf("Start address (default is %d)", defaultStartAddress))
	flag.IntVar(&numBytes, "bytes", defaultNumBytes, fmt.Sprintf("Number of bytes to read from the start address (default is %d)", defaultNumBytes))
	flag.IntVar(&baudRate, "baud", defaultBaudRate, fmt.Sprintf("Baud Rate (default is %d)", defaultBaudRate))
	flag.IntVar(&responsePause, "pause", defaultResponsePause, fmt.Sprintf("Device Response Pause in milliseconds (default is %d)", defaultResponsePause))
	flag.Parse()

	if len(serialPort) > 0 {

		// turn on the debug trace option, to see what is being transmitted
		trace := true
		ctx, cerr := modbusclient.ConnectRTU(serialPort, baudRate)
		if cerr != nil {
			log.Println(fmt.Sprintf("RTU Connection Err: %s", cerr))
		} else {
			// attempt to read the [startAddr] address register on
			// slave device number [slaveDevice] via the [serialDevice]
			readResult, readErr := modbusclient.RTURead(ctx, byte(slaveDevice), modbusclient.FUNCTION_READ_HOLDING_REGISTERS, uint16(startAddr), uint16(numBytes), responsePause, trace)
			if readErr != nil {
				log.Println(readErr)
			} else {
				log.Println(fmt.Sprintf("Rx: %x", readResult))

				// Skip past the reply headers, and decode each of the data hi/lo byte pairs into integers
				var (
					sliceStart, sliceStop int
					data                  int16
					decodeErr             error
				)

				for i := 0; i < int(readResult[2]); i++ {
					// take the next two bytes, if available
					sliceStart = 3 + i
					sliceStop = 3 + i + 2

					// decode them into integers
					data, decodeErr = modbusclient.DecodeHiLo(readResult[sliceStart:sliceStop])
					if decodeErr != nil {
						log.Println(decodeErr)
					} else {
						log.Println(fmt.Sprintf("Decoded int = %d (from pair of data bytes: %x)", data, readResult[sliceStart:sliceStop]))
					}
				}
			}
			modbusclient.DisconnectRTU(ctx)
		}
		/*
			// attempt to read the [startAddr] address register on
			// slave device number [slaveDevice] via the [serialDevice]
			readResult, readErr := modbusclient.RTURead(serialPort, byte(slaveDevice), modbusclient.FUNCTION_READ_HOLDING_REGISTERS, uint16(startAddr), uint16(numBytes), baudRate, responsePause, trace)
			if readErr != nil {
				log.Println(readErr)
			}
			log.Println(fmt.Sprintf("Rx: %x", readResult))

			// Skip past the reply headers, and take a slice of the first data pair returned,
			// to decode their high/low bytes into the corresponding integer value
			firstInt, decodeErr := modbusclient.DecodeHiLo(readResult[3:5])
			if decodeErr != nil {
				log.Println(decodeErr)
			} else {
				log.Println(fmt.Sprintf("Decoded int = %d (from 1st pair of data bytes: %x)", firstInt, readResult[3:5]))
			}*/

	} else {

		// display the command line usage requirements
		flag.PrintDefaults()

	}

}
