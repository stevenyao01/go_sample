package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

/**
 * @Package Name: main
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-3 下午5:10
 * @Description:
 */

func main() {

	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, false)
	if err != nil {
		fmt.Println("BoolToByte error : %s", err)
	}
	fmt.Println(buffer.Bytes())
	fmt.Println("hello: ", buffer)

	b := []byte{0x00, 0x00, 0x03, 0xe8}

	b_buf :=  bytes .NewBuffer(b)

	var x int32

	binary.Read(b_buf, binary.BigEndian, &x)

	fmt.Println(x)



	fmt.Println(strings.Repeat("-", 100))



	x  =  1000

	b_buf  =  bytes .NewBuffer([]byte{})

	binary.Write(b_buf, binary.BigEndian, x)

	fmt.Println(b_buf.Bytes())

}