package statistics

/**
 * @Package Name: statistics
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-10 下午3:55
 * @Description:
 */

import (
)

type IntStatistics struct {
	max		int
	min 	int
	first 	int
	double 	int
	sum 	int64
	last 	int
}

func (i *IntStatistics)GetHeaderSize()(int){
	return 5 * 4 + 8 *1
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


func NewInt() (*Statistics, error) {

	return &Statistics{
	},nil
}