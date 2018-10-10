package encoder

import (
	"github.com/go_sample/src/tsfile/common/header"
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

type Encoder interface {
	Encode (value interface{}, buffer *bytes.Buffer) ()
}

func GetEncoder (et int16, tdt int16) (Encoder) {
	var encoder Encoder
	switch {
	case et == header.PLAIN:
		encoder, _ = NewPlainEncoder(tdt)
	case et == header.RLE:
		encoder, _ = NewFloatEncoder(tdt)
	}

	return encoder
}

