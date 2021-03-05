package main

import (
	"fmt"
	"time"
)

/**
 * @Project: go_sample
 * @Package Name: test
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/12/9 下午3:36
 * @Description:
 */



/* define new struct DataGet*/
type DataGet struct {
	syncInData   chan []byte
	aSyncInData  chan []byte
	syncFlagCh   chan bool
	syncDataFlag bool
	interval     int64
	tt			 *time.Timer
}

func (d *DataGet) SyncDataFlagUpdate() {
	sdfInterval := 24 * 60 * 60
	timer := time.NewTimer(time.Duration(sdfInterval) * time.Second)
	for {
		if !timer.Stop() {
			select {
			case <- timer.C:
			default:
			}
		}
		timer.Reset(time.Duration(sdfInterval) * time.Second)
		select {
		case val := <- d.syncFlagCh:
			d.syncDataFlag = val
			continue
		case <- timer.C:
			if d.syncDataFlag {
				d.syncDataFlag = false
			}
			continue
		}

	}
}

func (d *DataGet) getSyncInData() {
	var data []byte

	fmt.Println("source data 01.")
	d.syncFlagCh <- true

	fmt.Println("source data 02.")
	data = []byte("lenovo dibg.")
	time.Sleep(10)
	d.syncInData <- data


	fmt.Println("source data 03.")
	d.syncFlagCh <- false
	fmt.Println("source data 04.")
	return
}

func (d *DataGet) getData() (data []byte, err error) {
	fmt.Println("enter getdata.")
	if !d.syncDataFlag {
		fmt.Println("before getSyncInData.")
		go d.getSyncInData()
	}

	fmt.Println("after getSyncInData.")
	if !d.tt.Stop() {
		select {
		case <- d.tt.C:
		default:
		}
	}
	d.tt.Reset(time.Duration(1) * time.Second)
	select {
	case syncVal := <-d.syncInData:
		//fmt.Println("get data.")
		return syncVal, nil
	case aSyncVal := <-d.aSyncInData:
		return aSyncVal, nil
	case <- d.tt.C:
		fmt.Println("asdfjalskdfjalskdfjslakdjfslakdjfasldfjslkadjfaslkdjfskladjf")
		return nil, nil
	}
}

/**
 * @Description:
 * @Params:
 * @return:
 * @Date: 2020/12/9 下午3:41
 */
func dataNew() (*DataGet, error) {
	return &DataGet{
		syncInData:   make(chan []byte, 1024),
		aSyncInData:  make(chan []byte, 1024),
		syncFlagCh:   make(chan bool),
		syncDataFlag: false,
		interval: 1,
		tt: time.NewTimer(time.Duration(1) * time.Second),
	}, nil
}

func main() () {
	d, _ := dataNew()
	go d.SyncDataFlagUpdate()

	for i := 0; i < 10000; i++ {
		fmt.Println("...")
		data, _ := d.getData()
		fmt.Println("get ", i ," data: ", string(data))
	}
}
