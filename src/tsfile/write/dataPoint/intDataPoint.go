package dataPoint

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
	tsDataType			int
	value 				int
}

//func (d *DataPoint) Write(v []byte) ([]byte,error) {
//	return nil,nil
//}
//
//func (d *DataPoint) Close() (bool) {
//	return true
//}


func NewInt(sId string, tdt int, val int) (*DataPoint, error) {

	//switch tdt {
	//case 0:
	//	// bool
	//case 1:
	//	//
	//case 2:
	//	//
	//case 3:
	//	//
	//case 4:
	//	//
	//case 5:
	//	//
	//case 6:
	//	//
	//case 7:
	//	//
	//case 8:
	//	//
	//default:
	//	// int
	//}

	// todo
	return &DataPoint{
		sensorId:sId,
		tsDataType:tdt,
		value:val,
	},nil
}