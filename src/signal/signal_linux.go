// +build linux

package main

import (
	"os"
	"os/signal"
	"syscall"
	"github.com/lenovo/common/workerr"
)

func SetCaptureParentSignal(){
	syscall.Syscall(syscall.SYS_PRCTL, syscall.PR_SET_PDEATHSIG, uintptr(syscall.SIGHUP), 0)
}

func CatchSignal(){
	signalChan := make(chan os.Signal, 1)
	go func (){
		<- signalChan
		os.Exit(workerr.WORKER_SUCCESS_EXIT.Code())
	}()
	signal.Notify(signalChan, syscall.SIGHUP)
}

func StartSignalCatch() {
	SetCaptureParentSignal()
	CatchSignal()
}