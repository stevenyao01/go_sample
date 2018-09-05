package valueWriter

/**
 * @Package Name: valueWriter
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-31 下午4:51
 * @Description:
 */

import (
//"github.com/go_sample/src/tsfile/common/log"
//	"encoding/gob"
//	"github.com/elastic/beats/filebeat/harvester/reader"
	"time"
	"bytes"
	"github.com/go_sample/src/tsfile/common/utils"
	"github.com/go_sample/src/tsfile/write/dataPoint"
)

type ValueWriter struct {
	// time
	timeEncoder		Encoder

	// value
	valueEncoder 	Encoder

	// todo these buffer should be encoding
	timeBuf 			*bytes.Buffer
	valueBuf 			*bytes.Buffer
	//buf := bytes.NewBuffer([]byte{})
}

func (s *ValueWriter) GetCurrentMemSize()(int){
	return s.timeBuf.Len() + s.valueBuf.Len()
}

func (s *ValueWriter) GetByteBuffer()(*bytes.Buffer){
	timeSize := s.timeBuf.Len()
	encodeBuffer := bytes.NewBuffer([]byte{})

	encodeBuffer.Write(utils.Int32ToByte(int32(timeSize)))
	encodeBuffer.
	return s.timeBuf.Len() + s.valueBuf.Len()
}

func (s *ValueWriter) Write(t int64, tdt int, value interface{}) () {
	var timeByteData []byte
	var valueByteData []byte
	switch tdt {
	case 0:
		// bool
		if data, ok := value.(bool); ok {
			valueByteData = utils.BoolToByte(data)
		}
	case 1:
		//int32
		if data, ok := value.(int32); ok {
			valueByteData = utils.Int32ToByte(data)
		}
	case 2:
		//int64
		if data, ok := value.(int64); ok {
			valueByteData = utils.Int64ToByte(data)
		}

	case 3:
		//float
		if data, ok := value.(float32); ok {
			valueByteData = utils.Float32ToByte(data)
		}
	case 4:
		//double , float64 in golang as double in c
		if data, ok := value.(float64); ok {
			valueByteData = utils.Float64ToByte(data)
		}
	case 5:
		//text
	case 6:
		//fixed_len_byte_array
	case 7:
		//enums
	case 8:
		//bigdecimal
	default:
		// int32
	}
	// write time to byteBuffer
	timeByteData = utils.Int64ToByte(t)
	s.timeBuf.Write(timeByteData)
	// write value to byteBuffer
	s.valueBuf.Write(valueByteData)
	return
}

func (s *ValueWriter) Reset() () {
	s.timeBuf.Reset()
	s.valueBuf.Reset()
	return
}

//// todo the return type should be Compressor, after finished Compressor we should modify it.
//func (s *ValueWriter) GetCompressor() (string) {
//	return s.compressor
//}
//
//func (s *ValueWriter) Close() (bool) {
//	return true
//}


func New() (*ValueWriter, error) {

	return &ValueWriter{
		//sensorId:sId,
		timeBuf:bytes.NewBuffer([]byte{}),
		valueBuf:bytes.NewBuffer([]byte{}),
	},nil
}