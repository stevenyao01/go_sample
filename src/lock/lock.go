package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

/**
 * @Project: go_sample
 * @Package Name: lock
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/9/12 下午3:46
 * @Description:
 */

var locked_file string
var number int

//文件锁
type FileLock struct {
	dir string
	f   *os.File
}

func New(dir string) *FileLock {
	return &FileLock{
		dir: dir,
	}
}

//加锁
func (l *FileLock) Lock() error {
	f, err := os.Open(l.dir)
	if err != nil {
		return err
	}
	l.f = f
	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		return fmt.Errorf("cannot get flock: directory %s - %s", l.dir, err)
	} else {
		fmt.Println("get flock by directory %s", l.dir)
	}
	return nil
}

//释放锁
func (l *FileLock) Unlock() error {
	defer l.f.Close()
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN)
}

func test() {
	//flock := New(locked_file)
	err := flock.Lock()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("output : %d\n", number)
	time.Sleep(2 * time.Second)
	flock.Unlock()
	number++
	return
}

func main() {
	test_file_path, _ := os.Getwd()
	locked_file = test_file_path

	flock := New(locked_file)

	for i := 0; i < 10; i++ {
		go test()
		time.Sleep(1 * time.Second)
	}
	time.Sleep(2 * time.Second)

}

////var lock *os.File
////var err error
//
////func test(){
////	err = syscall.Flock(int(lock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
////	if err != nil {
////		fmt.Println("上一个任务未执行完成，暂停执行")
////		time.Sleep(1 * time.Second)
////	} else {
////		fmt.Println("上一个程序已经执行完毕")
////		time.Sleep(20 * time.Second)
////	}
////}
//
//func main() () {
//	lockFile := "./lock.pid"
//	lock, err := os.Create(lockFile)
//	if err != nil {
//		fmt.Println("创建文件锁失败", err)
//	}
//	defer os.Remove(lockFile)
//	defer lock.Close()
//
//	err = syscall.Flock(int(lock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
//	if err != nil {
//		fmt.Println("上一个任务未执行完成，暂停执行")
//		time.Sleep(1 * time.Second)
//	} else {
//		fmt.Println("上一个程序已经执行完毕")
//		//time.Sleep(20 * time.Second)
//	}
//
//	err = syscall.Flock(int(lock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
//	if err != nil {
//		fmt.Println("上一个任务未执行完成，暂停执行")
//		time.Sleep(1 * time.Second)
//	} else {
//		fmt.Println("上一个程序已经执行完毕")
//		//time.Sleep(20 * time.Second)
//	}
//
//	//for {
//	//	go test()
//	//	time.Sleep(100 * time.Millisecond)
//	//}
//}
