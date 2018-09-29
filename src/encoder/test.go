package main

import (
	"bytes"
	"fmt"
	"github.com/go_sample/src/tsfile/common/encoder"
)

/**
 * @Package Name: encoder
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-28 下午6:25
 * @Description:
 */

 func main () () {
	buf := bytes.NewBuffer([]byte{})
 	//enc := Encoder().Encode(1, buf)
 	enc, _ := enc.
 	tt := enc.
 	tt := enc.eo.Encode(1, buf)
 	fmt.Println("tt: %s", tt)
 }