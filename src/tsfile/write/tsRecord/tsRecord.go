package tsRecord

/**
 * @Package Name: tsRecord
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-27 下午2:36
 * @Description:
 */

import (
//"github.com/go_sample/src/tsfile/common/log"
	"time"
	"container/list"
	"sync"
)

type TsRecord struct {
	time				time.Duration
	deviceId			string
	dataPointList		*list.List
	m 					sync.Mutex
}


func (t *TsRecord) SetTime(time time.Duration) () {
	t.time = time
	return
}

func (t *TsRecord) addTuple(tuple DataPoint) () {
	pushBack(t, tuple)
	return
}

func pushBack(t *TsRecord, tuple DataPoint) {
	if tuple == nil {
		return
	}
	t.m.Lock()
	defer t.m.Unlock()
	t.dataPointList.PushBack(tuple)
	return
}

func front(t *TsRecord) *list.Element {
	t.m.Lock()
	defer t.m.Unlock()
	return t.dataPointList.Front()
}

func remove(t *TsRecord, element *list.Element) {
	if element == nil {
		return
	}
	t.m.Lock()
	defer t.m.Unlock()
	t.dataPointList.Remove(element)
}

// this remove has some issue, we cann't use as the follow:
//for e := l.Front(); e != nil; e = e.Next {
//	l.Remove(e)
//}

// because when we remove ,element.next == nil then the loop for element != nil is ok,then exit.
// so we must use as the following two ways:
//way1:
//	var next *list.Element
//	for element := list.Front(); element != nil; element = next {
//		next = element.Next()
//		list.remove(element)
//	}
//
//way2:
//	for {
//		element := list.Front()
//		if element == nil {
//			break
//		}
//		list.remove(element)
//	}

func len(t *TsRecord) int {
	t.m.Lock()
	defer t.m.Unlock()
	return t.dataPointList.Len()
}


func New(t time.Duration, dId string) (*TsRecord, error) {
	// todo

	return &TsRecord{
		time:t,
		deviceId:dId,
		dataPointList:list.New(),
	},nil
}