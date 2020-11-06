package main

import (
	"time"
)

/**
 * @Project: ticker
 * @Package Name: ticker
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/9/7 下午2:41
 * @Description:
 */

/* define new struct Tick*/
type Tick struct {
	queue       *Queue
	interval	int32
	readers		interface{}
}

/**
* @Description:
* @Params:
* @return:
* @Date: 2020/9/8 上午11:30
*/
func (t *Tick) Run() {
	tick := time.NewTicker(time.Duration(t.interval) * time.Second)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			//fmt.Println("Ticker:", time.Now().Format("2006-01-02 15:04:05"), "  interval:", t.interval)
			//str := t.readers.(string)
			//fmt.Println("str:", str)
			t.queue.Enqueue(t.readers)
		}
	}
}

/**
* @Description:
* @Params:
* @return:
* @Date: 2020/9/7 下午2:43
*/
func TickNew(q *Queue, ivl int32, data interface{}) (*Tick, error) {
	return &Tick{
		interval: ivl,
		readers:  data,
		queue:    q,
	}, nil
}