package encoder1

import (
	"bytes"
)

/**
 * @Package Name: encoder
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-28 下午5:55
 * @Description:
 */


type IntEncoder struct {
	encoderType		int16
	//eo 				encoderOperation
}

//encode (value interface{}, buffer *bytes.Buffer) ()
//flush (buffer *bytes.Buffer) ()
//GetOneItemMaxSize () (int)
//GetMaxByteSize () (int64)

//func (i *IntEncoder) encode (1, buffer *bytes.Buffer) (int) {
//	fmt.println("hello")
//	return 11
//}

func (i *IntEncoder) Encode (l int, buffer *bytes.Buffer) (int) {
	return 111
}

//func NewIntEncoder(et int16) (*IntEncoder, error) {
//	return &IntEncoder{
//		encoderType:et,
//	},nil
//}

