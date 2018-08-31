package dataPoint

/**
 * @Package Name: dataPoint
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-28 下午4:27
 * @Description:
 */

import (
//"github.com/go_sample/src/tsfile/common/log"
	"time"
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

func (d *DataPoint) Write(t time.Time, sd seriesWriter.SeriesWriter) (bool) {
	if sd.GetTsDeviceId() == "" {
		log.Info("give seriesWriter is null, do nothing and return.")
		return false
	}
	sd.Write(t, d.value)
	//tdt := sd.GetTsDataType()
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