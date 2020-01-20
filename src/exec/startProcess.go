package main

import (
	"os"
	"fmt"
)

func main() {
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}, //其他变量如果不清楚可以不设定
	}
	p, err := os.StartProcess("./bin/new_exec", []string{"./bin/new_exec", "-c /home/steven/code/code_sync/go/src/github.com/agent/bin/worker/6/worker6_input.conf -d /home/steven/code/code_sync/go/src/github.com/agent/bin/worker/6/worker6_output.conf"}, attr) //vim 打开tmp.txt文件
	if err != nil {
		fmt.Println(err)
	}
	//go func() {
	//	p.Signal(os.Kill) //kill process
	//}()

	pstat, err := p.Wait()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pstat) //signal: killed
}
