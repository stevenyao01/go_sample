package tsFileWriter

import (
)

/**
 * @Package Name: dataPoint
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-27 下午3:19
 * @Description:
 */

type IntDataPoint struct {
	sensorId			string
	tsDataType			int16
	value 				int32
}

//func (d *DataPoint) Write(v []byte) ([]byte,error) {
//	return nil,nil
//}
//
//func (d *DataPoint) Close() (bool) {
//	return true
//}


func NewInt(sId string, tdt int, val int32) (*DataPoint, error) {

	// todo
	return &DataPoint{
		sensorId:sId,
		tsDataType:tdt,
		value:val,
	},nil
}