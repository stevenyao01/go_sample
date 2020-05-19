package main

/**
 * @Project: go_sample
 * @Package Name: panic
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/5/14 上午11:27
 * @Description:
 */

import (
	"fmt"
	"io/ioutil"
	"runtime/debug"
	"time"
)

const coreDumpFilename = "core.dump"


func myPanic() string {
	panic("my panic.")
}
func main() {
	defer func() {
		err := recover()
		_ = coreDump(coreDumpFilename, err)
	}()
	_ = myPanic()
	fmt.Println("cann't excute here, because panic.")
}

func coreDump(filename string, err interface{}) error {
	stack := debug.Stack()

	filename = filename + "." + fmt.Sprintf("%d", time.Now().Unix())
	//fmt.Println(filename)

	var content []byte
	if err != nil {
		content = []byte(fmt.Sprintf("%s\n%s", err, stack))
	} else {
		content = []byte(fmt.Sprintf("%s", stack))
	}

	return ioutil.WriteFile(filename, content, 0644)
}

func main2() {
f1()
f2()
f3()
return
}

func f1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%s\n", err)
		}
	}()
	fmt.Println("func1")
	return
}
func f2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%s\n", err)
		}
	}()
	panic("my panic.")
	fmt.Println("func2")
	return
}
func f3() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%s\n", err)
		}
	}()
	fmt.Println("func3")
	return
}
