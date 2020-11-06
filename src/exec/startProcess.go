package main

import (
	"fmt"
	"os"
)

func main() {
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}, //其他变量如果不清楚可以不设定
	}

	var p[10] *os.Process
	var err error
	for i := 0; i < 10; i++  {
		p[i], err = os.StartProcess("./modbus", []string{"./modbus", "-c input.conf -d output.conf -o other.conf -debug"}, attr) //vim 打开tmp.txt文件
		if err != nil {
			fmt.Println(err)
		}
	}


	//go func() {
	//	p.Signal(os.Kill) //kill process
	//}()

	for j := 0; j < 10; j++  {
		pasta, err := p[j].Wait()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("wait j:%d return: %s", j, pasta) //signal: killed
	}



}


//func main() {
//	attr := &os.ProcAttr{
//		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}, //其他变量如果不清楚可以不设定
//	}
//	p, err := os.StartProcess("./bin/new_exec", []string{"./bin/new_exec", "-c /home/steven/code/code_sync/go/src/github.com/agent/bin/worker/6/worker6_input.conf -d /home/steven/code/code_sync/go/src/github.com/agent/bin/worker/6/worker6_output.conf"}, attr) //vim 打开tmp.txt文件
//	if err != nil {
//		fmt.Println(err)
//	}
//	//go func() {
//	//	p.Signal(os.Kill) //kill process
//	//}()
//
//	pstat, err := p.Wait()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println(pstat) //signal: killed
//}
