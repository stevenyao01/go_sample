package tsFileWriter

/**
 * @Package Name: valueWriter
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-31 下午4:51
 * @Description:
 */

import (
	"bytes"
	"github.com/go_sample/src/tsfile/common/utils"
	"github.com/go_sample/src/tsfile/write/sensorDescriptor"
	"github.com/go_sample/src/tsfile/common/tsFileConf"
	"github.com/go_sample/src/tsfile/common/log"
)

type ValueWriter struct {
	//// time
	//timeEncoder		Encoder
	//
	//// value
	//valueEncoder 	Encoder

	// todo these buffer should be encoding
	timeBuf 			*bytes.Buffer
	valueBuf 			*bytes.Buffer
	desc	 			*sensorDescriptor.SensorDescriptor
	//buf := bytes.NewBuffer([]byte{})
}

func (s *ValueWriter) GetCurrentMemSize()(int){
	return s.timeBuf.Len() + s.valueBuf.Len()
}

func (s *ValueWriter) GetByteBuffer()(*bytes.Buffer){
	timeSize := s.timeBuf.Len()
	encodeBuffer := bytes.NewBuffer([]byte{})

	//// write timeBuf size
	//encodeBuffer.Write(utils.Int32ToByte(int32(timeSize)))

	//声明一个空的slice,容量为timebuf的长度
	timeSlice := make([]byte, timeSize)
	//把buf的内容读入到timeSlice内,因为timeSlice容量为timeSize,所以只读了timeSize个过来
	s.timeBuf.Read(timeSlice)
	encodeBuffer.Write(timeSlice)

	//声明一个空的value slice,容量为valuebuf的长度
	valueSlice := make([]byte, s.valueBuf.Len())
	//把buf的内容读入到timeSlice内,因为timeSlice容量为timeSize,所以只读了timeSize个过来
	s.valueBuf.Read(valueSlice)
	encodeBuffer.Write(valueSlice)

	return encodeBuffer
}

func (s *ValueWriter) Write(t int64, tdt int16, value interface{}) () {
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
		//if data, ok := value.(float32); ok {
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
	encodeCount := s.desc.GetTimeCount()
	log.Info("s.timeBuf size1: %d", s.timeBuf.Len())
	if encodeCount == -1 {
		s.timeBuf.Write(utils.BoolToByte(true))
		log.Info("s.timeBuf size2: %d", s.timeBuf.Len())
		s.timeBuf.Write(timeByteData)
		log.Info("s.timeBuf size3: %d", s.timeBuf.Len())
		s.timeBuf.Write(timeByteData)
		log.Info("s.timeBuf size4: %d", s.timeBuf.Len())
		s.timeBuf.Write(timeByteData)
		log.Info("s.timeBuf size5: %d", s.timeBuf.Len())
		s.desc.SetTimeCount(encodeCount + 1)
	}
	if s.desc.GetTimeCount() == tsFileConf.DeltaBlockSize {
		s.timeBuf.Write(timeByteData)
		s.timeBuf.Write(timeByteData)
		s.timeBuf.Write(timeByteData)
		s.desc.SetTimeCount(0)
	}
	log.Info("s.timeBuf size6: %d", s.timeBuf.Len())

	//timeByteData = utils.Int64ToByte(t)
	//s.timeBuf.Write(timeByteData)
	// write value to byteBuffer
	s.valueBuf.Write(valueByteData)
	log.Info("s.valueBuf size1: %d", s.valueBuf.Len())
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


func NewValueWriter(d *sensorDescriptor.SensorDescriptor) (*ValueWriter, error) {

	return &ValueWriter{
		//sensorId:sId,
		timeBuf:bytes.NewBuffer([]byte{}),
		valueBuf:bytes.NewBuffer([]byte{}),
		desc:d,
	},nil
}