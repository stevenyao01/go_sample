package metaData

/**
 * @Package Name: metaData
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-13 下午9:00
 * @Description:
 */

import (
	"github.com/go_sample/src/tsfile/common/log"
	"os"
	"github.com/go_sample/src/tsfile/common/tsFileConf"
)

type TimeSeriesChunkMetaData struct {
	sensorId							string
	fileOffsetOfCorrespondingData		int64
	startTime							int64
	endTime								int64
}

func (t *TimeSeriesChunkMetaData) WriteMagic()(int){

	return 0
}

func NewTimeSeriesChunkMetaData(sid string, fOffset int64, sTime int64, eTime int64) (*TimeSeriesChunkMetaData, error) {
	// todo

	return &TimeSeriesChunkMetaData{
		sensorId:sid,
		fileOffsetOfCorrespondingData:fOffset,
		startTime:sTime,
		endTime:eTime,
	},nil
}