package dataPoint

/**
 * @Package Name: dataPoint
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-28 下午4:27
 * @Description:
 */

import (
//"github.com/go_sample/src/tsfile/common/log"
)

//type DataPoint interface {
//	sensorId	string
//	tsDataType		int
//	value			unsafe.Pointer
//}

type DataPoint struct {
	sensorId			string
	tsDataType			int
	value 				interface{}
}



//func (s *DataPoint) Write(v []byte) ([]byte,error) {
//	return nil,nil
//}
//
//func (s *DataPoint) Close() (bool) {
//	return true
//}


//func New(sId string, tdt int, te int) (*DataPoint, error) {
//	// todo do measurement init and memory check
//
//	return &DataPoint{
//		sensorId:sId,
//		tsDataType:tdt,
//		tsEncoding:te,
//	},nil
//}