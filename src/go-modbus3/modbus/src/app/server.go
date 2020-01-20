package app

import (
	"log"
	"github.com/go_sample/src/go-modbus3/modbus/src/modbustcp"
)

var h *handler

type handler struct {
}

func (h *handler) Server(req []byte) []byte {
	return []byte{}
}

func (h *handler) Fault(detail string) {

}

func TcpServer() {
	modbustcp.SetHandler(h)
	err := modbustcp.ServerCreate(80)
	if err != nil {
		log.Println(err.Error())
	}
}
