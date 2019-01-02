package port

import (
"sync"
"sync/atomic"
"unsafe"
)

/**
 * @Package Name: port
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-11-30 下午5:36
 * @Description:
 */


const (
	LockedFlag   int32 = 1
	UnlockedFlag int32 = 0
)

type Mutex struct {
	in     sync.Mutex
	status *int32
}

func NewMutex() *Mutex {
	status := UnlockedFlag
	return &Mutex{
		status: &status,
	}
}

func (m *Mutex) Lock() bool {
	m.in.Lock()
	return true
}

func (m *Mutex) Unlock() {
	m.in.Unlock()
	atomic.AddInt32(m.status, UnlockedFlag)
}

func (m *Mutex) TryLock() bool {
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.in)), UnlockedFlag, LockedFlag) {
		atomic.AddInt32(m.status, LockedFlag)
		return true
	}
	return false
}

func (m *Mutex) IsLocked() bool {
	if atomic.LoadInt32(m.status) == LockedFlag {
		return true
	}
	return false
}