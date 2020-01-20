package tsFileWriter

import (
)

/**
 * @Package Name: BoolDataPoint
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-27 下午3:19
 * @Description:
 */

type BoolDataPoint struct {
	sensorId			string
	tsDataType			int16
	value 				bool
}


func NewBool(sId string, tdt int, val bool) (*DataPoint, error) {
	return &DataPoint{
		sensorId:sId,
		tsDataType:tdt,
		value:val,
	},nil
}