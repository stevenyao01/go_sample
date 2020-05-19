package main

import (
	"github.com/goburrow/modbus"
	"time"
)

type RTUClientHandler struct {
	*modbus.RTUClientHandler
}

func (r *RTUClientHandler) SetSlaveId(id byte) {
	r.SlaveId = id
}

func (r *RTUClientHandler) Reconnect() error {
	var err error
	if err = r.Close(); err != nil {
		return err
	}

	return r.Connect()
}

func NewRTUClientHandler(config *Config) *RTUClientHandler {
	handler := modbus.NewRTUClientHandler(config.RTU.Device)
	handler.BaudRate = config.RTU.BaudRate
	handler.DataBits = config.RTU.DataBits
	handler.Parity = config.RTU.Parity
	handler.StopBits = config.RTU.StopBits

	handler.Timeout = time.Duration(config.ConnTimeout) * time.Millisecond
	handler.IdleTimeout = defaultQueryTimeout * time.Millisecond
	//handler.Logger = log.New(os.Stdout, "modbusRTU:", log.LstdFlags)

	return &RTUClientHandler{RTUClientHandler: handler}
}
