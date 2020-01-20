package encoder

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

const(
	MAX_STRING_LENGTH = "max_string_length"
	MAX_POINT_NUMBER = "max_point_number"
)

type Encoder struct {
	encoderType		int16
	eo 				encoderOperation
}

type encoderOperation interface {
	encode (value interface{}, buffer *bytes.Buffer) ()
	flush (buffer *bytes.Buffer) ()
	GetOneItemMaxSize () (int)
	GetMaxByteSize () (int64)
}

func NewEncoder(et int16) (*Encoder, error) {
	return &Encoder{
		encoderType:et,
	},nil
}

