//package dataPoint
package dataPoint

import (
	"unsafe"
	"fmt"
	//"reflect"
)


/**
 * @Package Name: union
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-27 下午6:17
 * @Description:
 */

// ----- union begin ---------------------------------------------------
type DataPoint interface {
	toBool() *B
	toInt32() *I32
	toInt64() *I64
	toDouble() *D
	toText() *T
	toBigdecimal() *BD
}

// bool
type B struct {
	a int32
}

func (i *B) toBool() *B {
	return i
}

func (i *B) toInt32() *I32 {
	return (*I32)(unsafe.Pointer(i))
}

func (i *B) toInt64() *I64 {
	return (*I64)(unsafe.Pointer(i))
}

func (i *B) toDouble() *D {
	return (*D)(unsafe.Pointer(i))
}

func (i *B) toText() *T {
	return (*T)(unsafe.Pointer(i))
}

func (i *B) toBigdecimal() *BD {
	return (*BD)(unsafe.Pointer(i))
}

// int 32
type I32 struct {
	b int32
}

func (i *I32) toBool() *B {
	return (*B)(unsafe.Pointer(i))
}

func (i *I32) toInt32() *I32 {
	return i
}

func (i *I32) toInt64() *I64 {
	return (*I64)(unsafe.Pointer(i))
}

func (i *I32) toDouble() *D {
	return (*D)(unsafe.Pointer(i))
}

func (i *I32) toText() *T {
	return (*T)(unsafe.Pointer(i))
}

func (i *I32) toBigdecimal() *BD {
	return (*BD)(unsafe.Pointer(i))
}

// int64
type I64 struct {
	c int32
}

func (i *I64) toBool() *B {
	return (*B)(unsafe.Pointer(i))
}

func (i *I64) toInt32() *I32 {
	return (*I32)(unsafe.Pointer(i))
}

func (i *I64) toInt64() *I64 {
	return i
}

func (i *I64) toDouble() *D {
	return (*D)(unsafe.Pointer(i))
}

func (i *I64) toText() *T {
	return (*T)(unsafe.Pointer(i))
}

func (i *I64) toBigdecimal() *BD {
	return (*BD)(unsafe.Pointer(i))
}

// double
type D struct {
	d int32
}

func (i *D) toBool() *B {
	return (*B)(unsafe.Pointer(i))
}

func (i *D) toInt32() *I32 {
	return (*I32)(unsafe.Pointer(i))
}

func (i *D) toInt64() *I64 {
	return (*I64)(unsafe.Pointer(i))
}

func (i *D) toDouble() *D {
	return i
}

func (i *D) toText() *T {
	return (*T)(unsafe.Pointer(i))
}

func (i *D) toBigdecimal() *BD {
	return (*BD)(unsafe.Pointer(i))
}

// text
type T struct {
	e int32
}

func (i *T) toBool() *B {
	return (*B)(unsafe.Pointer(i))
}

func (i *T) toInt32() *I32 {
	return (*I32)(unsafe.Pointer(i))
}

func (i *T) toInt64() *I64 {
	return (*I64)(unsafe.Pointer(i))
}

func (i *T) toDouble() *D {
	return (*D)(unsafe.Pointer(i))
}

func (i *T) toText() *T {
	return i
}

func (i *T) toBigdecimal() *BD {
	return (*BD)(unsafe.Pointer(i))
}

// text
type BD struct {
	f int32
}

func (i *BD) toBool() *B {
	return (*B)(unsafe.Pointer(i))
}

func (i *BD) toInt32() *I32 {
	return (*I32)(unsafe.Pointer(i))
}

func (i *BD) toInt64() *I64 {
	return (*I64)(unsafe.Pointer(i))
}

func (i *BD) toDouble() *D {
	return (*D)(unsafe.Pointer(i))
}

func (i *BD) toText() *T {
	return (*T)(unsafe.Pointer(i))
}

func (i *BD) toBigdecimal() *BD {
	return i
}


// ------- union end -------------------------------------------------

type myStruct struct {
	dP DataPoint
	aaa  int
}

func main() {
	a := &I32{0x060302}
	mystruct := myStruct{a, 33}
	b := (*B)(unsafe.Pointer(a))
	fmt.Printf("%x, %d\n", a.b, a.b)
	fmt.Printf("%v\n", b.a)
	//fmt.Printf("%v\n", b.a[0])
	b.a = 0x0008
	fmt.Printf("%v\n", b.a)
	//fmt.Printf("%x, %d\n", a.a, a.a)
	//fmt.Printf("%d\n", reflect.TypeOf(b.c).Size())
	//fmt.Printf("%d\n", reflect.TypeOf(a).Size())
	//fmt.Println(b.toB())
	//fmt.Println(b.toI())
	//fmt.Println(b.toI().toB())
	//fmt.Println(a.toI().toB())
	//fmt.Println(mystruct.iOrB.toI().toB())
	fmt.Println(mystruct.dP.toInt32().toInt64())
}