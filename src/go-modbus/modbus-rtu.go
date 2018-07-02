// Package modbusclient provides modbus Serial Line/RTU and TCP/IP access
// for client (master) applications to communicate with server (slave)
// devices. Logic specifically in this file implements the Serial Line/RTU
// protocol.

package modbusclient

import (
	"fmt"
	"github.com/tarm/goserial"
	"io"
	"log"
	"time"
)

// crc computes and returns a cyclic redundancy check of the given byte array
func crc(data []byte) uint16 {
	var crc16 uint16 = 0xffff
	l := len(data)
	for i := 0; i < l; i++ {
		crc16 ^= uint16(data[i])
		for j := 0; j < 8; j++ {
			if crc16&0x0001 > 0 {
				crc16 = (crc16 >> 1) ^ 0xA001
			} else {
				crc16 >>= 1
			}
		}
	}
	return crc16
}

// GenerateRTUFrame is a method corresponding to a RTUFrame object which
// returns a byte array representing the associated serial line/RTU
// application data unit (ADU)
func (frame *RTUFrame) GenerateRTUFrame() []byte {

	packetLen := 8
	if len(frame.Data) > 0 {
		packetLen = RTU_FRAME_MAXSIZE
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

	// add the crc to the end
	packet_crc := crc(packet[:bytesUsed])
	packet[bytesUsed] = byte(packet_crc & 0xff)
	packet[(bytesUsed + 1)] = byte(packet_crc >> 8)
	bytesUsed += 2

	return packet[:bytesUsed]
}

// ConnectRTU attempts to access the Serial Device for subsequent
// RTU writes and response reads from the modbus slave device
func ConnectRTU(serialDevice string, baudRate int) (io.ReadWriteCloser, error) {
	conf := &serial.Config{Name: serialDevice, Baud: baudRate}
	ctx, err := serial.OpenPort(conf)
	return ctx, err
}

// DisconnectRTU closes the underlying Serial Device connection
func DisconnectRTU(ctx io.ReadWriteCloser) {
	ctx.Close()
}

// viaRTU is a private method which applies the given function validator,
// to make sure the functionCode passed is valid for the operation
// desired. If correct, it creates an RTUFrame given the corresponding
// information, attempts to open the serialDevice, and if successful, transmits
// it to the modbus server (slave device) specified by the given serial connection,
// and returns a byte array of the slave device's reply, and error (if any)
func viaRTU(connection io.ReadWriteCloser, fnValidator func(byte) bool, slaveAddress, functionCode byte, startRegister, numRegisters uint16, data []byte, timeOut int, debug bool) ([]byte, error) {
	if fnValidator(functionCode) {
		frame := new(RTUFrame)
		frame.TimeoutInMilliseconds = timeOut
		frame.SlaveAddress = slaveAddress
		frame.FunctionCode = functionCode
		frame.StartRegister = startRegister
		frame.NumberOfRegisters = numRegisters
		if len(data) > 0 {
			frame.Data = data
		}

		// generate the ADU from the RTU frame
		adu := frame.GenerateRTUFrame()
		if debug {
			log.Println(fmt.Sprintf("Tx: %x", adu))
		}

		// transmit the ADU to the slave device via the
		// serial port represented by the fd pointer
		_, werr := connection.Write(adu)
		if werr != nil {
			if debug {
				log.Println(fmt.Sprintf("RTU Write Err: %s", werr))
			}
			return []byte{}, werr
		}

		// allow the slave device adequate time to respond
		time.Sleep(time.Duration(frame.TimeoutInMilliseconds) * time.Millisecond)

		// then attempt to read the reply
		response := make([]byte, RTU_FRAME_MAXSIZE)
		n, rerr := connection.Read(response)
		if rerr != nil {
			if debug {
				log.Println(fmt.Sprintf("RTU Read Err: %s", rerr))
			}
			return []byte{}, rerr
		}

		// check the validity of the response
		if response[0] != frame.SlaveAddress || response[1] != frame.FunctionCode {
			if debug {
				log.Println("RTU Response Invalid")
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

		// confirm the checksum (crc)
		response_crc := crc(response[:(n - 2)])
		if response[(n-2)] != byte((response_crc&0xff)) ||
			response[(n-1)] != byte((response_crc>>8)) {
			// crc failed (odd that there's no specific code for it)
			if debug {
				log.Println("RTU Response Invalid: Bad Checksum")
			}
			// return the response bytes anyway, and let the caller decide
			return response[:n], MODBUS_EXCEPTIONS[EXCEPTION_BAD_CHECKSUM]
		}

		// return only the number of bytes read
		return response[:n], nil
	}

	return []byte{}, MODBUS_EXCEPTIONS[EXCEPTION_ILLEGAL_FUNCTION]
}

// RTURead performs the given modbus Read function over RTU to the given
// serialDevice, using the given frame data
func RTURead(serialDeviceConnection io.ReadWriteCloser, slaveAddress, functionCode byte, startRegister, numRegisters uint16, timeOut int, debug bool) ([]byte, error) {
	return viaRTU(serialDeviceConnection, ValidReadFunction, slaveAddress, functionCode, startRegister, numRegisters, []byte{}, timeOut, debug)
}

// RTUWrite performs the given modbus Write function over RTU to the given
// serialDevice, using the given frame data
func RTUWrite(serialDeviceConnection io.ReadWriteCloser, slaveAddress, functionCode byte, startRegister, numRegisters uint16, data []byte, timeOut int, debug bool) ([]byte, error) {
	return viaRTU(serialDeviceConnection, ValidWriteFunction, slaveAddress, functionCode, startRegister, numRegisters, data, timeOut, debug)
}
