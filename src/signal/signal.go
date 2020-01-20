package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// start goroutine to listen some signal
	go signalListen()
	// main loop
	for {
		time.Sleep(30 * time.Second)
		fmt.Println("main loop.")
	}
}

func signalListen() {
	// init os.signal channel
	c := make(chan os.Signal)
	// define catch signal
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	for {
		// wait channel
		sig := <-c
		// when receive signal,then notify channel,and print the follow info.
		fmt.Println("receive signal:", sig)
	}
}