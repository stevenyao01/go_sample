package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
	//"strings"
	//"github.com/mqtt/utils/log"
)

func main2() {
	for {
		//时间戳
		t := time.Now().Unix()
		fmt.Println(t)

		//时间戳到具体显示的转化
		fmt.Println(time.Unix(t, 0).String())

		//带纳秒的时间戳
		t = time.Now().UnixNano()
		fmt.Println(t)
		fmt.Println("------------------")

		//基本格式化的时间表示
		fmt.Println(time.Now().String())

		fmt.Println(time.Now().Format("2006year 01month 02day"))

		time.Sleep(2 * time.Second)
	}



	//t := time.Now()
	//
	//fmt.Println(t)
	//fmt.Println(t.UTC().Format(time.UnixDate))
	//fmt.Println(t.Unix())
	//


	//timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	//fmt.Println(timestamp)
	//retStr := subString(timestamp, 0, 13)
	//
	//fmt.Println(retStr)


	//
	//var str = "a\nb\nc"
	//log.Info("str: %s", str)
	//retArry := strings.Split(str, "\n")
	//log.Info("retArray: %s", retArry[0])
}

func main() {

	/////////////////////
	sli:=make([]int ,0)
	for i := 0; i<10;i++  {
		sli=append(sli, 1)
	}
	//for i := 0; i<15;i++  {
	//	sli[i] = i
	//}
	fmt.Println("sli: ", sli)
	slif := sli[:11]
	fmt.Println("slif: ", slif)
	///////////////////

	var bBuf bytes.Buffer
	fmt.Println("bBuf size: ", bBuf.Len())

	t := time.Now()
	//fmt.Println(t)
	//
	//fmt.Println(t.UTC().Format(time.UnixDate))
	//
	//fmt.Println(t.Unix())

	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	//timestamp := strconv.FormatInt(t.UTC()., 10)
	retStr := subString(timestamp, 0, 13)

	fmt.Println(retStr)
	//fmt.Println(timestamp)


	//timestamp = timestamp[:10]
	//fmt.Println(timestamp)
}

func subString(source string, start int, end int) string {

	var substring = ""
	var pos = 0
	for _, c := range source {
		if pos < start {
			pos++
			continue
		}
		if pos >= end {
			break
		}
		pos++
		substring += string(c)
	}

	return substring
}
