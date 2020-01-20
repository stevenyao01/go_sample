package main

import (
	"github.com/go_sample/src/portServer/port"
	"fmt"
)

/**
 * @Package Name: portServer
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-11-30 下午3:56
 * @Description:
 */

func main() () {
	fileLock, _ := port.NewFileLock()
	fileLock.LockFile("../../device.sk", false)
	//m := port.NewMutex()
	//if m.Lock() {
	//	fmt.Println("get lock")
	//}
	//if m.IsLocked() {
	//	fmt.Println("main加锁")
	//}
	var i = 0
	for true {
		if i < 100 {
			i += 1
		} else {
			break
		}
		go mylock(fileLock, i)
	}

	fileLock.UnLockFile()
	fmt.Println("main解锁")

	p, _ := port.NewListenPort()
	p.StartListen()
}

//func main() {
//	m := port.NewMutex()
//	if m.Lock() {
//		fmt.Println("get lock")
//	}
//	if m.IsLocked(){
//		fmt.Println("main加锁")
//	}
//	var i = 0
//	for true {
//		if i < 100 {
//			i += 1
//		} else {
//			break
//		}
//		go mylock(m, i)
//	}
//
//	m.Unlock()
//	fmt.Println("main解锁")
//	p, _ := port.NewListenPort()
//	p.StartListen()
//}
//
//func mylock(m *port.Mutex, num int) () {
//	m.Lock()
//	fmt.Println("协程", num, "加锁")
//	m.Unlock()
//	fmt.Println("协程", num, "解锁")
//}

func mylock(f *port.FileLock, num int) () {
	f.LockFile("../../device.sk", false)
	fmt.Println("协程", num, "加锁")
	f.UnLockFile()
	fmt.Println("协程", num, "解锁")
}
