package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("test..")
	loop := 0
	flag := 0
	lock := 0
	lock1 := 0
	lock2 := 0
	unlock := 0
	var mutex sync.Mutex
	for {
		mutex.Lock()
		fmt.Println("lock here: ", lock)
		lock++
		fmt.Println("first for")
		if loop < 50 {
			for i := 0; i < 2; i++ {
				fmt.Println("loop begin...")
				if i < 6 {
					fmt.Println("number i: ", i)
					continue
				} else {
					fmt.Println("num i: ", i)
				}
			}
		} else {
			fmt.Println("loop > 20")
			flag++
			if flag < 50 {
				mutex.Unlock()
				fmt.Println("unlock1 here: ", lock1)
				lock1++
				continue
			} else {
				mutex.Unlock()
				fmt.Println("unlock2 here: ", lock2)
				lock2++
				fmt.Println("return at here.")
				return
			}
		}
		fmt.Println("unlock here: ", unlock)
		unlock++
		mutex.Unlock()
		loop++
	}
}
func main1() {
	aa := "kill -9 " + "1234"
	retStr := fmt.Sprintf("%#q", aa)
	fmt.Println("retStr: ", retStr)
	//cmd := exec.Command("/bin/bash", "-c", `kill -9 28184`)
	////创建获取命令输出管道
	//stdout, err := cmd.StdoutPipe()
	//if err != nil {
	//	fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
	//	return
	//}
	////执行命令
	//if err := cmd.Start(); err != nil {
	//	fmt.Println("Error:The command is err,", err)
	//	return
	//}
	////读取所有输出
	//bytes, err := ioutil.ReadAll(stdout)
	//if err != nil {
	//	fmt.Println("ReadAll Stdout:", err.Error())
	//	return
	//}
	//if err := cmd.Wait(); err != nil {
	//	fmt.Println("wait:", err.Error())
	//	return
	//}
	//fmt.Printf("stdout:\n\n %s", bytes)
}
