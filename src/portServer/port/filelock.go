package port

import (
	"syscall"
	"os"
)

/**
 * @Package Name: port
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-11-30 下午6:01
 * @Description:
 */

type FileLock struct {
	file *os.File
}

func (fileLock *FileLock) LockFile(name string, truncate bool) (*os.File, error) {
	flag := os.O_RDWR | os.O_CREATE
	if truncate {
		flag |= os.O_TRUNC
	}
	f, err := os.OpenFile(name, flag, 0666)
	if err != nil {
		return nil, err
	}
	fileLock.file = f
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		f.Close()
		return nil, err
	}
	return f, nil
}

func (fileLock *FileLock) UnLockFile() (error) {
	if err := syscall.Flock(int(fileLock.file.Fd()), syscall.LOCK_UN); err != nil {
		return err
	}
	fileLock.file.Close()
	return nil
}

func NewFileLock() (*FileLock, error) {
	return &FileLock{
	}, nil
}
