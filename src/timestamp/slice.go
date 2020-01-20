package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**
 * @Package Name: main
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-4-17 下午1:42
 * @Description:
 */

func slicebytetostring(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: bh.Data,
		Len:  bh.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}

func main() {
	sli:=make([]int ,15)
	//for i := 0; i<10;i++  {
	//	sli=append(sli, 1)
	//}
	for i := 0; i<15;i++  {
		sli[i] = i
	}
	fmt.Println("sli: ", sli)
	slif := sli[:11]
	fmt.Println("slif: ", slif)
	

	//slif:=sli[:100000*500]
	//slib:=sli[100000*500:]
	//slif=append(slif, 10)
	//slif=append(slif, slib...)
	//fmt.Println("slice 的插入速度" + time.Now().Sub(t).String())
	//
	//var em *list.Element
	//len:=l.Len()
	//var i int
	//for e := l.Front(); e != nil; e = e.Next() {
	//	i++
	//	if i ==len/2 {
	//		em=e
	//		break
	//	}
	//}
	////忽略掉找中间元素的速度。
	//t = time.Now()
	//ef:=l.PushBack(2)
	//l.MoveBefore(ef,em)
	//fmt.Println("list: " + time.Now().Sub(t).String())
}
