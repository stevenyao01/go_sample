package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

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
		return fmt.Errorf("cannot flock directory %s - %s", l.dir, err)
	}
	return nil
}

//释放锁
func (l *FileLock) Unlock() error {
	defer l.f.Close()
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN)
}

func main() {
	testFilePath, _ := os.Getwd()
	lockedFile := testFilePath

	fmt.Println("locked_file: ", lockedFile)

	flock := New(lockedFile)
	err := flock.Lock()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < 10000; i++ {
		fmt.Printf("output : %d\n", i)
		time.Sleep(2 * time.Second)
	}


}

//func main() {
//	test_file_path, _ := os.Getwd()
//	locked_file := test_file_path
//
//	wg := sync.WaitGroup{}
//
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func(num int) {
//			flock := New(locked_file)
//			err := flock.Lock()
//			if err != nil {
//				wg.Done()
//				fmt.Println(err.Error())
//				return
//			}
//			fmt.Printf("output : %d\n", num)
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//	time.Sleep(2 * time.Second)
//
//}