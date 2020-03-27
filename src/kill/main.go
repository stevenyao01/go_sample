package main

import (
	"fmt"
)

func main() {
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
