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
	double 	interface{}
	sum 	int64
	last 	interface{}
	tsDataType int
}

func (s *Statistics)Serialize (buffer bytes.Buffer) () {
	buffer.Write(s.GetMaxByte(s.tsDataType))
	buffer.Write(s.GetMinByte(s.tsDataType))
	buffer.Write(s.GetFirstByte(s.tsDataType))
	buffer.Write(s.GetLastByte(s.tsDataType))
	buffer.Write(s.GetSumByte(s.tsDataType))
}

func GetStatistics(tdt int) (*Statistics) {
	switch tdt {
	case 0:
		// bool
		newStatistics, _ := NewBool()
		newStatistics.tsDataType = tdt
		return newStatistics
	case 1:
		//int32
		newStatistics, _ := NewInt()
		newStatistics.tsDataType = tdt
		return newStatistics
	case 2:
		//int64

	case 3:
		//float
		newStatistics, _ := NewFloat()
		newStatistics.tsDataType = tdt
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
		return utils.Int64ToByte(s.max.(int64))
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
		return utils.Int64ToByte(s.min.(int64))
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
		return utils.Int64ToByte(s.first.(int64))
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
		return utils.Int64ToByte(s.last.(int64))
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

func (s *Statistics)GetSumByte(tdt int16)([]byte){
	return utils.Int64ToByte(s.sum)
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