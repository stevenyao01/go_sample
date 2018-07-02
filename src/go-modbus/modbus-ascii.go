// Package modbusclient provides modbus Serial Line/ASCII and TCP/IP access
// for client (master) applications to communicate with server (slave)
// devices. Logic specifically in this file implements the Serial Line/ASCII
// protocol.

package modbusclient

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/tarm/goserial"
	"io"
	"log"
	"time"
)

func Lrc(data []byte) uint8 {
	return lrc(data)
}

// Modbus ASCII does not use CRC, but Longitudinal Redundancy Check.
// lrc computes and returns the 2's compliment (-) of the sum of the given byte
// array modulo 256
func lrc(data []byte) uint8 {
	var sum uint8 = 0
	var lrc8 uint8 = 0
	for _, b := range data {
		sum += b
	}
	lrc8 = uint8(-int8(sum))
	return lrc8
}

// GenerateASCIIFrame is a method corresponding to a ASCIIFrame object which
// returns a byte array representing the associated serial line/ASCII
// application data unit (ADU)
func (frame *ASCIIFrame) GenerateASCIIFrame() []byte {

	packetLen := 7
	if len(frame.Data) > 0 {
		packetLen += len(frame.Data) + 1
		if packetLen > ASCII_FRAME_MAXSIZE {
			packetLen = ASCII_FRAME_MAXSIZE
		}
	}

	packet := make([]byte, packetLen)
	packet[0] = frame.SlaveAddress
	packet[1] = frame.FunctionCode
	packet[2] = byte(frame.StartRegister >> 8)       // (High Byte)
	packet[3] = byte(frame.StartRegister & 0xff)     // (Low Byte)
	packet[4] = byte(frame.NumberOfRegisters >> 8)   // (High Byte)
	packet[5] = byte(frame.NumberOfRegisters & 0xff) // (Low Byte)
	bytesUsed := 6

	for i := 0; i < len(frame.Data); i++ {
		packet[(bytesUsed + i)] = frame.Data[i]
	}
	bytesUsed += len(frame.Data)

	// add the lrc to the end
	packet_lrc := lrc(packet[:bytesUsed])
	packet[bytesUsed] = byte(packet_lrc)
	bytesUsed += 1

	// Convert raw bytes to ASCII packet
	ascii_packet := make([]byte, bytesUsed*2+3)
	hex.Encode(ascii_packet[1:], packet)

	asciiBytesUsed := bytesUsed*2 + 1

	// Frame the packet
	ascii_packet[0] = ':'                 // 0x3A
	ascii_packet[asciiBytesUsed] = '\r'   // CR 0x0D
	ascii_packet[asciiBytesUsed+1] = '\n' // LF 0x0A
	asciiBytesUsed += 2

	return bytes.ToUpper(ascii_packet[:asciiBytesUsed])
}

// ConnectASCII attempts to access the Serial Device for subsequent
// ASCII writes and response reads from the modbus slave device
func ConnectASCII(serialDevice string, baudRate int) (io.ReadWriteCloser, error) {
	conf := &serial.Config{Name: serialDevice, Baud: baudRate}
	ctx, err := serial.OpenPort(conf)
	return ctx, err
}

// DisconnectASCII closes the underlying Serial Device connection
func DisconnectASCII(ctx io.ReadWriteCloser) {
	ctx.Close()
}

// viaASCII is a private method which applies the given function validator,
// to make sure the functionCode passed is valid for the operation
// desired. If correct, it creates an ASCIIFrame given the corresponding
// information, attempts to open the serialDevice, and if successful, transmits
// it to the modbus server (slave device) specified by the given serial connection,
// and returns a byte array of the slave device's reply, and error (if any)
func viaASCII(connection io.ReadWriteCloser, fnValidator func(byte) bool, slaveAddress, functionCode byte, startRegister, numRegisters uint16, data []byte, timeOut int, debug bool) ([]byte, error) {
	if fnValidator(functionCode) {
		frame := new(ASCIIFrame)
		frame.TimeoutInMilliseconds = timeOut
		frame.SlaveAddress = slaveAddress
		frame.FunctionCode = functionCode
		frame.StartRegister = startRegister
		frame.NumberOfRegisters = numRegisters
		if len(data) > 0 {
			frame.Data = data
		}

		// generate the ADU from the ASCII frame
		adu := frame.GenerateASCIIFrame()
		if debug {
			log.Println(fmt.Sprintf("Tx: %x", adu))
			fmt.Println(fmt.Sprintf("Tx: %s", adu))
		}

		// transmit the ADU to the slave device via the
		// serial port represented by the fd pointer
		_, werr := connection.Write(adu)
		if werr != nil {
			if debug {
				log.Println(fmt.Sprintf("ASCII Write Err: %s", werr))
			}
			return []byte{}, werr
		}

		// allow the slave device adequate time to respond
		time.Sleep(time.Duration(frame.TimeoutInMilliseconds) * time.Millisecond)

		// then attempt to read the reply
		ascii_response := make([]byte, ASCII_FRAME_MAXSIZE)
		ascii_n, rerr := connection.Read(ascii_response)
		if rerr != nil {
			if debug {
				log.Println(fmt.Sprintf("ASCII Read Err: %s", rerr))
			}
			return []byte{}, rerr
		}

		// check the framing of the response
		if ascii_response[0] != ':' ||
			ascii_response[ascii_n-2] != '\r' ||
			ascii_response[ascii_n-1] != '\n' {
			if debug {
				log.Println("ASCII Response Framing Invalid")
				log.Println(fmt.Sprintf("%s", ascii_response))
			}
			return []byte{}, MODBUS_EXCEPTIONS[EXCEPTION_UNSPECIFIED]
		}

		// convert to raw bytes
		raw_n := (ascii_n - 3) / 2
		response := make([]byte, raw_n)
		hex.Decode(response, ascii_response[1:ascii_n-2])

		// check the validity of the response
		if response[0] != frame.SlaveAddress || response[1] != frame.FunctionCode {
			if debug {
				log.Println("ASCII Response Invalid")
			}
			if response[0] == frame.SlaveAddress && (response[1]&0x7f) == frame.FunctionCode {
				switch response[2] {
				case EXCEPTION_ILLEGAL_FUNCTION:
					return []byte{}, MODBUS_EXCEPTIONS[EXCEPTION_ILLEGAL_FUNCTION]
				case EXCEPTION_DATA_ADDRESS:
					return []byte{}, MODBUS_EXCEPTIONS[EXCEPTION_DATA_ADDRESS]
				case EXCEPTION_DATA_VALUE:
					return []byte{}, MODBUS_EXCEPTIONS[EXCEPTION_DATA_VALUE]
				case EXCEPTION_SLAVE_DEVICE_FAILURE:
					return []byte{}, MODBUS_EXCEPTIONS[EXCEPTION_SLAVE_DEVICE_FAILURE]
				}
			}
			return []byte{}, MODBUS_EXCEPTIONS[EXCEPTION_UNSPECIFIED]
		}

		// confirm the checksum (lrc)
		response_lrc := lrc(response[:raw_n-1])
		if response[raw_n-1] != response_lrc {
			// lrc failed (odd that there's no specific code for it)
			if debug {
				log.Println("ASCII Response Invalid: Bad Checksum")
			}
			// return the response bytes anyway, and let the caller decide
			return response, MODBUS_EXCEPTIONS[EXCEPTION_BAD_CHECKSUM]
		}

		// return only the number of bytes read
		return response, nil
	}

	return []byte{}, MODBUS_EXCEPTIONS[EXCEPTION_ILLEGAL_FUNCTION]
}

// ASCIIRead performs the given modbus Read function over ASCII to the given
// serialDevice, using the given frame data
func ASCIIRead(serialDeviceConnection io.ReadWriteCloser, slaveAddress, functionCode byte, startRegister, numRegisters uint16, timeOut int, debug bool) ([]byte, error) {
	return viaASCII(serialDeviceConnection, ValidReadFunction, slaveAddress, functionCode, startRegister, numRegisters, []byte{}, timeOut, debug)
}

// ASCIIWrite performs the given modbus Write function over ASCII to the given
// serialDevice, using the given frame data
func ASCIIWrite(serialDeviceConnection io.ReadWriteCloser, slaveAddress, functionCode byte, startRegister, numRegisters uint16, data []byte, timeOut int, debug bool) ([]byte, error) {
	return viaASCII(serialDeviceConnection, ValidWriteFunction, slaveAddress, functionCode, startRegister, numRegisters, data, timeOut, debug)
}
