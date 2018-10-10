package encoder

import (
	"github.com/go_sample/src/tsfile/common/log"
	"bytes"
	"github.com/go_sample/src/tsfile/common/header"
)

/**
 * @Package Name: encoder
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-28 下午5:55
 * @Description:
 */


type IntEncoder struct {
	tsDataType		int16
	//eo 				encoderOperation
}

func (i *IntEncoder) Encode (value interface{}, buffer *bytes.Buffer) () {
	log.Info("enter intEncoder!!")
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

func (i *IntEncoder) EncBool (value bool, buffer *bytes.Buffer) () {
	log.Info("final enc ok!")
	return
}
func (i *IntEncoder) EncInt32 (value int32, buffer *bytes.Buffer) () {
	log.Info("final enc ok! input int value: %d", value)
	return
}

func (i *IntEncoder) EncFloat32 (value float32, buffer *bytes.Buffer) () {
	log.Info("final enc ok!")
	return
}

func NewIntEncoder(tdt int16) (*IntEncoder, error) {
	log.Info("init intEncoder!!")
	return &IntEncoder{
		tsDataType:tdt,
	},nil
}

