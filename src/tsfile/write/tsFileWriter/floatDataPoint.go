package tsFileWriter

/**
 * @Package Name: dataPoint
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-28 下午4:52
 * @Description:
 */

import (
)

/**
 * @Package Name: dataPoint
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-27 下午3:19
 * @Description:
 */

type FloatDataPoint struct {
	sensorId			string
	tsDataType			int
	value 				float32
}

func NewFloat(sId string, tdt int, val float32) (*DataPoint, error) {
	return &DataPoint{
		sensorId:sId,
		tsDataType:tdt,
		value:val,
	},nil
}