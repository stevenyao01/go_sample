package tsFileWriter

import (
)

/**
 * @Package Name: StringDataPoint
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-27 下午3:19
 * @Description:
 */

type StringDataPoint struct {
	sensorId			string
	tsDataType			int16
	value 				string
}

func NewString(sId string, tdt int, val string) (*DataPoint, error) {
	return &DataPoint{
		sensorId:sId,
		tsDataType:tdt,
		value:val,
	},nil
}