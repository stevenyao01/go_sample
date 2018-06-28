package main

import (
	"fmt"
	"time"
	"strconv"
	"strings"
	"github.com/agent/log"
)

func main() {
	t := time.Now()

	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	fmt.Println(timestamp)
	retStr := subString(timestamp, 0, 13)
	fmt.Println(retStr)

	var str = "a\nb\nc"
	log.Info("str: %s", str)
	retArry := strings.Split(str, "\n")
	log.Info("retArray: %s", retArry[0])
}

func main1() {
	t := time.Now()
	//fmt.Println(t)
	//
	//fmt.Println(t.UTC().Format(time.UnixDate))
	//
	//fmt.Println(t.Unix())

	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	//timestamp := strconv.FormatInt(t.UTC()., 10)
	fmt.Println(timestamp)

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
