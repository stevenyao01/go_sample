package statistics

/**
 * @Package Name: statistics
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-10 下午2:32
 * @Description:
 */

import (
	"github.com/go_sample/src/tsfile/common/utils"
	"bytes"
)

type StatisticsOperate interface {
	write()
}

type Statistics struct {
	max		interface{}
	min 	interface{}
	first 	interface{}
	//double 	interface{}
	sum 	int64
	last 	interface{}
	tsDataType int16
	isEmpty	bool
}

func (s *Statistics)Serialize (buffer *bytes.Buffer) () {
	buffer.Write(s.GetMaxByte(s.tsDataType))
	buffer.Write(s.GetMinByte(s.tsDataType))
	buffer.Write(s.GetFirstByte(s.tsDataType))
	buffer.Write(s.GetLastByte(s.tsDataType))
	buffer.Write(s.GetSumByte(s.tsDataType))
}

func GetStatistics(tdt int16) (*Statistics) {
	switch tdt {
	case 0:
		// bool
		newStatistics, _ := NewBool()
		newStatistics.tsDataType = tdt
		newStatistics.max = 0
		newStatistics.min = 0
		newStatistics.sum = 0
		newStatistics.first = 0
		newStatistics.last = 0
		return newStatistics
	case 1:
		//int32
		newStatistics, _ := NewInt()
		newStatistics.tsDataType = tdt
		newStatistics.max = int32(0)
		newStatistics.min = int32(0)
		newStatistics.sum = 0
		newStatistics.first = int32(0)
		newStatistics.last = int32(0)
		return newStatistics
	case 2:
		//int64

	case 3:
		//float
		newStatistics, _ := NewFloat()
		newStatistics.tsDataType = tdt
		newStatistics.max = float32(0)
		newStatistics.min = float32(0)
		newStatistics.sum = int64(0)
		newStatistics.first = float32(0)
		newStatistics.last = float32(0)
		return newStatistics
	case 4:
		//double , float64 in golang as double in c
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
		newStatistics, _ := NewInt()
		newStatistics.tsDataType = tdt
		return newStatistics
	}
	return nil
}

func (s *Statistics)GetHeaderSize(tdt int16)(int){
	switch tdt {
	case 0:
		// bool
		return 5 * 1 + 8 * 1
	case 1:
		//int32
		return 5 * 4 + 8 *1
	case 2:
		//int64

	case 3:
		//float
		return 8 * 6
	case 4:
		//double , float64 in golang as double in c
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
		return 5 * 4 + 8 *1
	}
	return 0
}

func (s *Statistics)GetMaxByte(tdt int16)([]byte){
	switch tdt {
	case 0:
		// bool
		return utils.BoolToByte(s.max.(bool))
	case 1:
		//int32
		return utils.Int32ToByte(s.max.(int32))
	case 2:
		//int64
		return utils.Int64ToByte(s.max.(int64))
	case 3:
		//float
		return utils.Float32ToByte(s.max.(float32))
	case 4:
		//double , float64 in golang as double in c
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
		return utils.Int32ToByte(s.max.(int32))
	}
	return nil
}

func (s *Statistics)GetMinByte(tdt int16)([]byte){
	switch tdt {
	case 0:
		// bool
		return utils.BoolToByte(s.min.(bool))
	case 1:
		//int32
		return utils.Int32ToByte(s.min.(int32))
	case 2:
		//int64
		return utils.Int64ToByte(s.min.(int64))
	case 3:
		//float
		return utils.Float32ToByte(s.min.(float32))
	case 4:
		//double , float64 in golang as double in c
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
		return utils.Int32ToByte(s.min.(int32))
	}
	return nil
}

func (s *Statistics)GetFirstByte(tdt int16)([]byte){
	switch tdt {
	case 0:
		// bool
		return utils.BoolToByte(s.first.(bool))
	case 1:
		//int32
		return utils.Int32ToByte(s.first.(int32))
	case 2:
		//int64
		return utils.Int64ToByte(s.first.(int64))
	case 3:
		//float
		return utils.Float32ToByte(s.first.(float32))
	case 4:
		//double , float64 in golang as double in c
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
		return utils.Int32ToByte(s.first.(int32))
	}
	return nil
}

func (s *Statistics)GetLastByte(tdt int16)([]byte){
	switch tdt {
	case 0:
		// bool
		return utils.BoolToByte(s.last.(bool))
	case 1:
		//int32
		return utils.Int32ToByte(s.last.(int32))
	case 2:
		//int64
		return utils.Int64ToByte(s.last.(int64))
	case 3:
		//float
		return utils.Float32ToByte(s.last.(float32))
	case 4:
		//double , float64 in golang as double in c
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
		return utils.Int32ToByte(s.last.(int32))
	}
	return nil
}

func (s *Statistics) GetSumByte(tdt int16)([]byte){
	return utils.Int64ToByte(s.sum)
}

func (s *Statistics) GetserializedSize (tdt int16) (int) {
	return len(s.GetMaxByte(tdt)) + len(s.GetMinByte(tdt)) + len(s.GetFirstByte(tdt)) + len(s.GetLastByte(tdt)) + 8
}

func (s *Statistics) UpdateStatBool (value bool) () {
	if s.isEmpty {
		s.max = value
		s.min = value
		s.first = value
		s.last = value
		s.sum = 0
	} else {
		s.UpdateBool(value, value, value, value, 0)
	}
}

func (s *Statistics) UpdateBool (max bool, min bool, first bool, last bool, sum int64) () {
	if max && !s.max.(bool) {
		s.max = max
	}
	if !min && s.min.(bool) {
		s.min = min
	}
	s.last = last
}

func (s *Statistics) UpdateStatInt32 (value int32) () {
	if s.isEmpty {
		s.max = value
		s.min = value
		s.first = value
		s.last = value
		s.sum = 0
	} else {
		s.UpdateInt32(value, value, value, value, value)
	}
}

func (s *Statistics) UpdateInt32 (max int32, min int32, first int32, last int32, sum int32) () {
	if max > s.max.(int32) {
		s.max = max
	}
	if min < s.min.(int32) {
		s.min = min
	}
	s.sum += int64(sum)
	s.last = last
}

func (s *Statistics) UpdateStatInt64 (value int64) () {
	if s.isEmpty {
		s.max = value
		s.min = value
		s.first = value
		s.last = value
		s.sum = 0
	} else {
		s.UpdateInt64(value, value, value, value, value)
	}
}

func (s *Statistics) UpdateInt64 (max int64, min int64, first int64, last int64, sum int64) () {
	if max > s.max.(int64) {
		s.max = max
	}
	if min < s.min.(int64) {
		s.min = min
	}
	s.sum += sum
	s.last = last
}

func (s *Statistics) UpdateStatFloat32 (value float32) () {
	if s.isEmpty {
		s.max = value
		s.min = value
		s.first = value
		s.last = value
		s.sum = 0
	} else {
		s.UpdateFloat32(value, value, value, value, value)
	}
}

func (s *Statistics) UpdateFloat32 (max float32, min float32, first float32, last float32, sum float32) () {
	if max > s.max.(float32) {
		s.max = max
	}
	if min < s.min.(float32) {
		s.min = min
	}
	s.sum += int64(sum)
	s.last = last
}

func (s *Statistics) UpdateStatFloat64 (value float64) () {
	if s.isEmpty {
		s.max = value
		s.min = value
		s.first = value
		s.last = value
		s.sum = 0
	} else {
		s.UpdateFloat64(value, value, value, value, value)
	}
}

func (s *Statistics) UpdateFloat64 (max float64, min float64, first float64, last float64, sum float64) () {
	if max > s.max.(float64) {
		s.max = max
	}
	if min < s.min.(float64) {
		s.min = min
	}
	s.sum += int64(sum)
	s.last = last
}

func (s *Statistics) UpdateStats (tdt int16, value interface{}) () {
	switch tdt {
	case 0:
		// bool
		if data, ok := value.(bool); ok {
			s.UpdateStatBool(data)
		}
	case 1:
		//int32
		if data, ok := value.(int32); ok {
			s.UpdateStatInt32(data)
		}
	case 2:
		//int64
		if data, ok := value.(int64); ok {
			s.UpdateStatInt64(data)
		}

	case 3:
		//float
		//if data, ok := value.(float32); ok {
		if data, ok := value.(float32); ok {
			s.UpdateStatFloat32(data)
		}
	case 4:
		//double , float64 in golang as double in c
		if data, ok := value.(float64); ok {
			s.UpdateStatFloat64(data)
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
	return
}


//func New(sId string, tdt int, te int) (*DataPoint, error) {
//	// todo do measurement init and memory check
//
//	return &DataPoint{
//		sensorId:sId,
//		tsDataType:tdt,
//		tsEncoding:te,
//	},nil
//}