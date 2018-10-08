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

type LongDataPoint struct {
	sensorId			string
	tsDataType			int16
	value 				int64
}

func NewLong(sId string, tdt int, val int64) (*DataPoint, error) {
	return &DataPoint{
		sensorId:sId,
		tsDataType:tdt,
		value:val,
	},nil
}