package fileLock

import (
	"fmt"
	"os"
	"syscall"
)

/**
 * @Package Name: main
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-4-9 下午9:00
 * @Description:
 */

// 定义一个 FileLock 的struct
type FileLock struct {
	file string   // 目录路径，例如 /home/XXX/go/src
	fd   *os.File // 文件描述符
}

// 新建一个 FileLock
func New(fileName string) *FileLock {
	return &FileLock{
		file: fileName,
	}
}

// 加锁操作
func (f *FileLock) Lock() error {
	newFile, err := os.Open(f.file) // 获取文件描述符
	if err != nil {
		fmt.Println("err: ", err.Error())
		return err
	}
	f.fd = newFile
	err = syscall.Flock(int(f.fd.Fd()), syscall.LOCK_EX|syscall.LOCK_NB) // 加上排他锁，当遇到文件加锁的情况直接返回 Error
	if err != nil {
		return fmt.Errorf("cannot flock directory %s - %s", f.file, err)
	}
	return nil
}

// 解锁操作
func (f *FileLock) Unlock() error {
	defer f.fd.Close()                                    // close 掉文件描述符
	return syscall.Flock(int(f.fd.Fd()), syscall.LOCK_UN) // 释放 Flock 文件锁
}
