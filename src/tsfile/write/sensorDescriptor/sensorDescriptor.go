package sensorDescriptor

/**
 * @Package Name: measurementDescriptor
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-24 下午7:38
 * @Description:
 */

import (
	//"github.com/go_sample/src/tsfile/common/log"
)

type SensorDescriptor struct {
	sensorId			string
	tsDataType			int
	tsEncoding			int

	//typeConverter		TsDataTypeConverter
	//encodingConverter	TsEncodingConverter
	compressor		Compressor
	//conf 				TsFileConfig
	//props 				make(map[string]string)
}

func (s *SensorDescriptor) GetSensorId() (string) {
	return s.sensorId
}

// todo the return type should be Compressor, after finished Compressor we should modify it.
func (s *SensorDescriptor) GetCompressor() (string) {
	return s.compressor
}

func (s *SensorDescriptor) Close() (bool) {
	return true
}


func New(sId string, tdt int, te int) (*SensorDescriptor, error) {
	// todo do measurement init and memory check

	return &SensorDescriptor{
		sensorId:sId,
		tsDataType:tdt,
		tsEncoding:te,
		},nil
}