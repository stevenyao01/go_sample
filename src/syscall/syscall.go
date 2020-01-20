package main

import (
	"fmt"
	"time"
	"github.com/go_sample/src/syscall/fileLock"
)

/**
 * @Package Name: syscall
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-4-9 下午8:34
 * @Description:
 */
func main(){
	fl := fileLock.New("/dev/ttyUSB0")
	fl.Lock()
	for i := 0; i < 60; i++ {
		time.Sleep(time.Second * 1)
		fmt.Print(".")
	}

	fl.Unlock()
}

//func init() error {
//	f, err := os.Open("/dev/ttyUSB0")
//	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB) // 加上排他锁，当遇到文件加锁的情况直接返回 Error
//	if err != nil {
//		//return fmt.Errorf("cannot flock directory %s - %s", f.dir, err)
//		fmt.Println("file is locked.")
//		f.Close()
//		return
//	}else{
//		fmt.Println("file is not locked.")
//	}
//	for i := 0; i < 60; i++ {
//		time.Sleep(time.Second * 1)
//		fmt.Println(".")
//	}
//	defer f.Close() // close 掉文件描述符
//	return syscall.Flock(int(f.Fd()), syscall.LOCK_UN) // 释放 Flock 文件锁
//
//}
