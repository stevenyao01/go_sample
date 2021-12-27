package main

import (
	"fmt"
	"github.com/go_sample/src/routine/process"
	"time"
)

func main() {
	pro, _ := process.ProcessNew("abc", "agent.log")
	_ = pro.StartFins()

	for i := 0; i < 10000; i++ {
		fmt.Println("sleep %d seconds in routine.go", i)
		time.Sleep(100 * time.Second)
	}
}
