package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"syscall"
	"time"
)

/**
 * @Project: go_sample
 * @Package Name: main
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/7/9 上午11:30
 * @Description:
 */

const (
	fileName = "upgrade.log"
)

func stopProcess(pid int) error {
	pro, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	err = pro.Signal(syscall.SIGINT)
	if err != nil {
		return err
	}
	_ = writeFile(fileName, []byte(fmt.Sprint("send signal to parent: ", os.Getppid(), "\n")), 0644)
	return nil
}

func writeFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

func main() {

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_ = writeFile(fileName, []byte(err.Error()), 0644)

	} else {
		_ = writeFile(fileName, data, 0644)
	}

	_ = writeFile(fileName, []byte(fmt.Sprint("current pid: ", os.Getpid(), "\n")), 0644)
	_ = writeFile(fileName, []byte(fmt.Sprint("get parent pid: ", os.Getppid(), "\n")), 0644)

	_ = stopProcess(os.Getppid())

	time.Sleep(time.Duration(1) * time.Second)

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Duration(3) * time.Second)
		_ = writeFile(fileName, []byte(fmt.Sprint("mloop: ", os.Getppid(), "\n")), 0644)
	}
}
