package metaData

/**
 * @Package Name: metaData
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-18 下午4:28
 * @Description:
 */

import (

)

type TimeSeriesMetaData struct {
	sensorId							string
	tsDataType							int16
}

//func (t *TimeSeriesMetaData) WriteMagic()(int){
//
//	return 0
//}
//
//func (t *TimeSeriesMetaData) SetDigest (tsDigest TsDigest) () {
//	t.valueStatistics = tsDigest
//}

func NewTimeSeriesMetaData(sid string, tdt int16) (*TimeSeriesMetaData, error) {

	return &TimeSeriesMetaData{
		sensorId:sid,
		tsDataType:tdt,
	},nil
}