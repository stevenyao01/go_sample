package main

import (
	"github.com/go_sample/src/log"
	"os/exec"
)

/**
 * @Project: go_sample
 * @Package Name: cmd
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/7/7 下午4:27
 * @Description:
 */

func main() () {
	c := exec.Command("cmd", "/C", "ipconfig")
	if err := c.Start(); err != nil {
		log.Println("Error: ", err)
	}
}
