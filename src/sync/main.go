package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once
	aFlag := false
	onceBody := func() {
		fmt.Println("Only once")
		aFlag = true
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	if aFlag {
		for i := 0; i < 10000; i++ {
			fmt.Println("loop: ", i)
			time.Sleep(time.Duration(1) * time.Second)
		}
	} else {
		fmt.Println("exit....")
		return
	}

	fmt.Println("end....")
}
