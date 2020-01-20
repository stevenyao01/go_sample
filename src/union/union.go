package main

import (
	"unsafe"
	"fmt"
	"reflect"
)

/**
 * @Package Name: union
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-27 下午6:17
 * @Description:
 */


// ----- union begin ---------------------------------------------------
type IorBUnion interface {
	toB() *B
	toI() *I
}

type I struct {
	a int32
}

func (i *I) toB() *B {
	return (*B)(unsafe.Pointer(i))
}

func (i *I) toI() *I {
	return i
}

type B struct {
	c [34]int16
}

func (b *B) toB() *B {
	return b
}

func (b *B) toI() *I {
	return (*I)(unsafe.Pointer(b))
}

// ------- union end -------------------------------------------------

type myStruct struct {
	iOrB IorBUnion
	aaa  int
}

func main() {
	a := &I{0x060302}
	mystruct := myStruct{a, 33}
	b := (*B)(unsafe.Pointer(a))
	fmt.Printf("%x, %d\n", a.a, a.a)
	fmt.Printf("%v\n", b.c)
	fmt.Printf("%v\n", b.c[0])
	b.c[9] = 0x0008
	fmt.Printf("%v\n", b.c)
	fmt.Printf("%x, %d\n", a.a, a.a)
	fmt.Printf("%d\n", reflect.TypeOf(b.c).Size())
	fmt.Printf("%d\n", reflect.TypeOf(a).Size())
	fmt.Println(b.toB())
	fmt.Println(b.toI())
	fmt.Println(b.toI().toB())
	fmt.Println(a.toI().toB())
	fmt.Println(mystruct.iOrB.toI().toB())
}