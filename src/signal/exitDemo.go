package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//创建监听退出chan
var c = make(chan os.Signal)

func main()  {

	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go signalProcess()

	fmt.Println("start...")
	num := 0
	for {
		num++
		fmt.Println("seconds : ", num)
		time.Sleep(time.Second)
	}
}

func signalProcess() {
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("exit signal: ", s)
			ExitFunc()
		case syscall.SIGUSR1:
			fmt.Println("usr1", s)
		case syscall.SIGUSR2:
			fmt.Println("usr2", s)
		default:
			fmt.Println("other", s)
		}
	}
	return
}

func ExitFunc()  {
	fmt.Println("begin exit!!!")
	fmt.Println("clean!!!")
	fmt.Println("end exit!!!")
	os.Exit(0)
}
