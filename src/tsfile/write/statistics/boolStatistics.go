package statistics

/**
 * @Package Name: statistics
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-10 下午3:49
 * @Description:
 */

import (
)

type BoolStatistics struct {
	max		bool
	min 	bool
	first 	bool
	double 	int64
	sum 	int64
	last 	bool
	isEmpty	bool
}

//func (d *BoolStatistics) GetStatistics(t int64, tdt int, value interface{}) (string) {
//	return d.sensorId
//}
//
//func (d *BoolStatistics) Write(t int64, sd seriesWriter.SeriesWriter) (bool) {
//	if sd.GetTsDeviceId() == "" {
//		log.Info("give seriesWriter is null, do nothing and return.")
//		return false
//	}
//	sd.Write(t, d.value)
//	return true
//}


func NewBool() (*Statistics, error) {

	return &Statistics{
		isEmpty:true,
	},nil
}