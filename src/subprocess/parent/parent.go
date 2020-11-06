package main

import (
	"github.com/go_sample/src/log"
	"os"
	"os/exec"
)

/**
 * @Project: go_sample
 * @Package Name: parent
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/7/9 上午11:15
 * @Description:
 */

func main() {
	log.Info("parent pid: %d", os.Getpid())
	cmd := exec.Command("./upgrade")
	cmdIn, _ := cmd.StdinPipe()

	_, _ = cmdIn.Write([]byte("hello upgrade\ngoodbye upgrade\n"))
	_ = cmdIn.Close()

	err := cmd.Run()
	if err != nil {
		log.Info("err: %s", err.Error())
	}
}
