package main

import (
	"fmt"
	"github.com/go_sample/src/routine/fins/tcp"
	"time"
)

func main() {
	tt := tcp.MTcpNew()
	go tt.Receiving()

	for i := 0; i < 10000; i++ {
		fmt.Println("loop in fins.go ", i)
		time.Sleep(1 * time.Second)
	}
}
