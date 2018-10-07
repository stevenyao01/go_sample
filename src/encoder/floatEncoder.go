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

type FloatEncoder struct {
	encoderType		int16
	//eo 				encoderOperation
}

	//encode (value interface{}, buffer *bytes.Buffer) ()
	//flush (buffer *bytes.Buffer) ()
	//GetOneItemMaxSize () (int)
	//GetMaxByteSize () (int64)

func (i *FloatEncoder) Encode (l float32, buffer *bytes.Buffer) (int) {
	return 222
}

//func NewFloatEncoder(et int16) (*FloatEncoder, error) {
//	return &FloatEncoder{
//		encoderType:et,
//	},nil
//}

