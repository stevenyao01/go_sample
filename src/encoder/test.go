package main

import (
	"github.com/go_sample/src/encoder/encoder"
	"github.com/go_sample/src/tsfile/common/header"
	"bytes"
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

	 // encode int
	 var iValue int32
	 iValue = 1
	 en1 := encoder.GetEncoder(header.PLAIN, header.INT32)
	 en1.Encode(iValue, buf)

	// encode float
	 var fValue float32
	 fValue = 1.2
	 en2 := encoder.GetEncoder(header.PLAIN, header.FLOAT)
	 en2.Encode(fValue, buf)

 }