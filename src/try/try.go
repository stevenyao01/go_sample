package main

/**
 * @Package Name: try
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-5 下午3:50
 * @Description:
 */

import (
	"reflect"
	"fmt"
)

// Try catches exception from f
func Try(f func()) *tryStruct {
	return &tryStruct{
		catches: make(map[reflect.Type]ExeceptionHandler),
		hold:    f,
	}
}

// ExeceptionHandler handle exception
type ExeceptionHandler func(interface{})

type tryStruct struct {
	catches map[reflect.Type]ExeceptionHandler
	hold    func()
}

func (t *tryStruct) Catch(e interface{}, f ExeceptionHandler) *tryStruct {
	t.catches[reflect.TypeOf(e)] = f
	return t
}

func (t *tryStruct) Finally(f func()) {
	defer func() {
		if e := recover(); nil != e {
			if h, ok := t.catches[reflect.TypeOf(e)]; ok {
				h(e)
			} else {
				f()
			}
		}
	}()

	t.hold()
}

func main() {
	Try(func() {
		panic(123.0)
	}).Catch(1, func(e interface{}) {
		fmt.Println("int", e)
	}).Catch("", func(e interface{}) {
		fmt.Println("string", e)
	}).Finally(func() {
		fmt.Println("finally")
	})
}