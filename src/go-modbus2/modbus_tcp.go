package main

import (
	"github.com/go_sample/src/go-modbus2/modbus"
	"time"
	"log"
	"os"
)

/**
 * @Package Name: main
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-10-17 下午6:04
 * @Description:
 */

func main1() () {
	handler := modbus.NewTCPClientHandler("localhost:502")
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 0xFF
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	if err != nil {
		log.Printf("connect err: %s", err)
	}
	defer handler.Close()

	client := modbus.NewClient(handler)
	results, err := client.ReadDiscreteInputs(15, 2)
	log.Printf("results1: %s", results)
	results, err = client.WriteMultipleRegisters(1, 2, []byte{0, 3, 0, 4})
	log.Printf("results2: %s", results)
	results, err = client.WriteMultipleCoils(5, 10, []byte{4, 3})
	log.Printf("results3: %s", results)
}
