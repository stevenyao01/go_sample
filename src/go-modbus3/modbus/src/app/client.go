//package app
//author: Lubia Yang
//create: 2013-10-21
//about: www.lubia.me

package app

import (
	"log"
	"github.com/go_sample/src/go-modbus3/modbus/src/modbusrtu"
	"github.com/go_sample/src/go-modbus3/modbus/src/modbustcp"
	"os"
)

func RtuClient() {
	fd, err := os.Open("/dev/ttyUSB0")
	if err != nil {
		log.Println("unable to open rs485")
		return
	}
	//b, err := modbusrtu.Read(fd, 0x03, 1, 3, 1)
	b, err := modbusrtu.Read(fd, 0x450, 3, 1, 1)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(b)
	err = modbusrtu.Write(fd, 0x03, 1, 3, 1, []byte{0, 1})
	if err != nil {
		log.Println(err.Error())
	}
}

func TcpClient() {
	mt := new(modbustcp.MbTcp)
	mt.Addr = 1
	mt.Code = 0x03
	mt.Data = []byte{0, 1}
	res, err := mt.Send("127.0.0.1:80")
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}
