package main

/*
#include <stdio.h>
#include <string.h>
#include "worker.h"
extern int _main(int argc, char *argv[]);
extern int call_in(int(*worker_in)(worker_t*, char*, int), worker_t *this, char *data, int len);
extern int call_process(int(*worker_process)(worker_t*, const char*, int, char*, int), worker_t *this, const char *in, int len_in, char *out, int len_out);
extern int call_out(int(*worker_out)(worker_t*, const char*, int), worker_t *this, const char *data, int len);
*/
import "C"
import (
	"os"
	"time"
	"fmt"
	"github.com/workerSDK"
	"unsafe"
	"errors"
)

func main() {
	l := len(os.Args)
	argv := make([]*C.char, 0, l)
	for _,v := range os.Args {
		p := (*C.char)(unsafe.Pointer(&v))
		argv = append(argv, p)
	}

	C._main(C.int(l), nil)
}

// implementation interface from workerSDK
type UserCb struct {
	workerSDK.UsercbIO
	user	C.worker_t
}

type OnIn func(*C.worker_t, *C.char, C.int)C.int

func NewCB(u *C.worker_t) UserCb {
	return UserCb{user:*u}
}

func (u UserCb) OnIn() (data []byte, err error) {
	time.Sleep(1 * time.Second)
	buf := make([]byte, 10240)
	//int(*worker_in)(worker_t *this, char *data, int len);

	//n := ((C.worker_in)(u.user.worker_in))(u.user, (*C.char)(unsafe.Pointer(&buf[0])), 10240)
	//n := OnIn(u.user.worker_in)(&u.user, (*C.char)(unsafe.Pointer(&buf[0])), 10240)
	n := C.call_in(u.user.worker_in, &u.user, (*C.char)(unsafe.Pointer(&buf[0])), 10240)
	//n := u.user.worker_in(u.user, &buf[0], 10240)
	if int(n) == -1 {
		return nil, errors.New("call C worker_in error")
	}

	return buf[:int(n)], err
}

func (u UserCb) OnOut(data []byte) (err error) {
	//n := u.user.worker_out(&u.user, (*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)))
	//
	//if int(n) == -1 {
	//	return errors.New("call c worker on out error")
	//}

	return err
}

func (u UserCb) OnProcess(data []byte) (afterData []byte, err error) {

	afterData = make([]byte, 1024)
	n := C.call_process(u.user.worker_process, &u.user, (*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)), (*C.char)(unsafe.Pointer(&afterData[0])), 1024)
	//n := u.user.worker_process(&u.user, (*C.char)(unsafe.Pointer(&data[0])), len(data), &afterData[0], 1024)
	if int(n) == -1 {
		return nil, errors.New("call c worker process error")
	}

	return afterData[:n], err
}

func (u UserCb) ReloadConfig() (err error) {
	return err
}

func (u UserCb) Destroy() (err error) {
	return err
}

//export GetLocalConfig
func GetLocalConfig(buf *C.char, length C.int) {
	conf,err := workerSDK.NewConfig(workerSDK.SOURCE)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(conf.LoFileName) < int(length) {
		length = C.int(len(conf.LoFileName))
	}

	C.strncpy(buf, C.CString(conf.LoFileName), C.ulonglong(length))
}

//export Run
func Run(user *C.worker_t) C.int {
	// 1. Create workerSDK config object.
	conf,err := workerSDK.NewConfig(int(user.worker_type))
	if err != nil {
		return C.int(-1)
	}

	// 2. Use workerSDK config create workerSDK object.
	worker,err := workerSDK.NewWorker(conf)
	if err != nil {
		return C.int(-1)
	}

	// 3. Create user callback object and set callback to workerSDK object.
	// SOURCE and TARGET workerSDK must set callback to workerSDK object.
	usercb := NewCB(user)
	worker.SetUserCallBack(usercb)

	// 4. Run workerSDK object, block run.
	// Start workerSDK object, non-block run, Can be recycled using Wait.
	if err := worker.Run(); err != nil {
		fmt.Printf("worker Run error:%s\n", err.Error())
		return C.int(-1)
	}

	return C.int(0)
}
