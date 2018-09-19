package dataPoint

/**
 * @Package Name: dataPoint
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-28 下午4:27
 * @Description:
 */

import (
	"github.com/go_sample/src/tsfile/write/seriesWriter"
	"github.com/go_sample/src/tsfile/common/log"
)

type DataPointOperate interface {
	write()
}

type DataPoint struct {
	sensorId			string
	tsDataType			int
	value 				interface{}
}

func (d *DataPoint) GetSensorId() (string) {
	return d.sensorId
}

func (d *DataPoint) Write(t int64, sd seriesWriter.SeriesWriter) (bool) {
	if sd.GetTsDeviceId() == "" {
		log.Info("give seriesWriter is null, do nothing and return.")
		return false
	}
	sd.Write(t, d.value)
	return true
}


//func New(sId string, tdt int, te int) (*DataPoint, error) {
//	// todo do measurement init and memory check
//
//	return &DataPoint{
//		sensorId:sId,
//		tsDataType:tdt,
//		tsEncoding:te,
//	},nil
//}