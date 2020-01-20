package encoder

import (
	"github.com/go_sample/src/tsfile/common/header"
	"github.com/go_sample/src/tsfile/common/log"
	"bytes"
)

/**
 * @Package Name: encoder
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-10-10 下午2:12
 * @Description:
 */

type PlainEncoder struct {
	tsDataType		int16
	//eo 				encoderOperation
}

func (i *PlainEncoder) Encode (value interface{}, buffer *bytes.Buffer) () {
	log.Info("enter PlainEncoder!!")
	switch {
	case i.tsDataType == header.BOOLEAN:
		if data, ok := value.(bool); ok {
			i.EncBool(data, buffer)
		}
	case i.tsDataType == header.INT32:
		if data, ok := value.(int32); ok {
			i.EncInt32(data, buffer)
		}
	case i.tsDataType == header.FLOAT:
		if data, ok := value.(float32); ok {
			i.EncFloat32(data, buffer)
		}
	}
	return
}

func (i *PlainEncoder) EncBool (value bool, buffer *bytes.Buffer) () {
	log.Info("final enc ok!")
	return
}
func (i *PlainEncoder) EncInt32 (value int32, buffer *bytes.Buffer) () {
	log.Info("final enc ok! input int value: %d", value)
	return
}

func (i *PlainEncoder) EncFloat32 (value float32, buffer *bytes.Buffer) () {
	log.Info("final enc ok! yao...")
	return
}

func NewPlainEncoder(tdt int16) (*PlainEncoder, error) {
	log.Info("init plainEncoder!!")
	return &PlainEncoder{
		tsDataType:tdt,
	},nil
}