package main

import (
	"fmt"
	"bytes"
)

/**
 * @Package Name: write
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-31 下午5:05
 * @Description:
 */

//func (b *bytes.Buffer) Write(p []byte) (n int, err error) {
//	b.lastRead = opInvalid
//	m := b.grow(len(p))
//	return copy(b.buf[m:], p), nil
//}
//
//func (b *Buffer) Write(p []byte) (n int, err error){
//	b.
//}


func main2() (){

	fmt.Println("===========以下通过Write把swift写入Learning缓冲器尾部=========")
	newBytes1 := []byte("Learning")
	newBytes2 := []byte("swift")
	//创建一个内容Learning的缓冲器
	buf := bytes.NewBuffer([]byte{})
	//打印为Learning
	fmt.Println(buf.String())
	//将newBytes这个slice写到buf的尾部
	buf.Write(newBytes1)
	buf.Write(newBytes2)
	fmt.Println(buf.String())
}