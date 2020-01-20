package encoder

import (
	"github.com/go_sample/src/tsfile/common/log"
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

type FloatEncoder struct {
	tsDataType		int16
	//eo 				encoderOperation
}

	//encode (value interface{}, buffer *bytes.Buffer) ()
	//flush (buffer *bytes.Buffer) ()
	//GetOneItemMaxSize () (int)
	//GetMaxByteSize () (int64)

//func (i *FloatEncoder) Encode ([]byte) (int) {
//	log.Info("enter floatEncoder!!")
//	return 222
//}

func (f *FloatEncoder) Encode (value interface{}, buffer *bytes.Buffer) () {
	log.Info("enter floatEncoder!!")
	switch {
	case f.tsDataType == header.BOOLEAN:
		if data, ok := value.(bool); ok {
			f.EncBool(data, buffer)
		}
	case f.tsDataType == header.INT32:
		if data, ok := value.(int32); ok {
			f.EncInt32(data, buffer)
		}
	case f.tsDataType == header.FLOAT:
		if data, ok := value.(float32); ok {
			f.EncFloat32(data, buffer)
		}
	}
	return
}

func (i *FloatEncoder) EncBool (value bool, buffer *bytes.Buffer) () {
	log.Info("final enc ok!")
	return
}
func (i *FloatEncoder) EncInt32 (value int32, buffer *bytes.Buffer) () {
	log.Info("final enc ok!")
	return
}

func (i *FloatEncoder) EncFloat32 (value float32, buffer *bytes.Buffer) () {
	log.Info("final enc ok! input float value: %f", value)
	return
}

func NewFloatEncoder(tdt int16) (*FloatEncoder, error) {
	return &FloatEncoder{
		tsDataType:tdt,
	},nil
}

