package main

import (
	"fmt"
	"time"
)

/**
 * @Project: go_sample
 * @Package Name: ticker
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/9/7 下午1:27
 * @Description:
 */


func main() {
	// Ticker 包含一个通道字段C，每隔时间段 d 就向该通道发送当时系统时间。
	// 它会调整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。
	// 如果d <= 0会触发panic。关闭该 Ticker 可以释放相关资源。

	//q := new(Queue)
	//q.Init()

	q, _ := QueueNew()

	t1, _ := TickNew(q,1, "t1")
	go t1.Run()

	t2, _ := TickNew(q,2, "t2")
	go t2.Run()

	t3, _ := TickNew(q,3, "t3")
	go t3.Run()

	t4, _ := TickNew(q,4, "t4")
	go t4.Run()

	//t5, _ := TickNew(q,1, "{\"key\":\"key1\",     \"value\":\"value1\" }")
	//go t5.Run()

	go dequeuePrint(q)

	time.Sleep(11 * time.Second)
	fmt.Println("ok")
}

func dequeuePrint(q *Queue) {
	for {
		if q.Size() > 0 {
			d := q.Dequeue()
			str := d.(string)
			fmt.Println("Dequeue str:", str)
		}
	}
}

///* define new struct Tick*/
//type Tick struct {
//	interval	int32
//}
//
///**
// * @Description:
// * @Params:
// * @return:
// * @Date: 2020/9/8 上午11:30
// */
//func (t *Tick) Run() {
//	tick := time.NewTicker(time.Duration(t.interval) * time.Second)
//	defer tick.Stop()
//	for {
//		select {
//		case <-tick.C:
//			fmt.Println("Ticker:", time.Now().Format("2006-01-02 15:04:05"), "  interval:", t.interval)
//		}
//	}
//}
//
///**
// * @Description:
// * @Params:
// * @return:
// * @Date: 2020/9/7 下午2:43
// */
//func TikNew(ivl int32) (*Tick, error) {
//	return &Tick{
//		interval: ivl,
//	}, nil
//}
