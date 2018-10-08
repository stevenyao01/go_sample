package tsFileWriter

import (
)

/**
 * @Package Name: DoubleDataPoint
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-27 下午3:19
 * @Description:
 */

type DoubleDataPoint struct {
	sensorId			string
	tsDataType			int16
	value 				int64
}

func NewDouble(sId string, tdt int, val int64) (*DataPoint, error) {
	return &DataPoint{
		sensorId:sId,
		tsDataType:tdt,
		value:val,
	},nil
}