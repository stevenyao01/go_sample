package rowGroupWriter

/**
 * @Package Name: rowGroupWriter
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-28 下午6:19
 * @Description:
 */

import (
//"github.com/go_sample/src/tsfile/common/log"
	"github.com/go_sample/src/tsfile/write/seriesWriter"
	"github.com/go_sample/src/tsfile/write/sensorDescriptor"
	"github.com/go_sample/src/tsfile/write/dataPoint"
	"github.com/go_sample/src/tsfile/common/utils"
	"github.com/go_sample/src/tsfile/common/log"
	"github.com/go_sample/src/tsfile/write/pageWriter"
	"github.com/go_sample/src/tsfile/write/tsFileWriter"
)

type RowGroupWriter struct {
	deviceId			string
	dataSeriesWriters	map[string]seriesWriter.SeriesWriter
}

func (r *RowGroupWriter) AddSeriesWriter(sd sensorDescriptor.SensorDescriptor, pageSize int) () {
	if contain, _ := utils.MapContains(r.dataSeriesWriters, sd.GetSensorId()); contain {
		// todo new pagewrite
		pw, _ := pageWriter.New(sd)

		// new serieswrite
		sw, _ := seriesWriter.New(r.deviceId, sd, *pw, pageSize)

		r.dataSeriesWriters[sd.GetSensorId()] = *sw
	} else {
		log.Error("given sensor has exist!")
	}
	return
}

func (r *RowGroupWriter) Write(t int64, data map[string]dataPoint.DataPoint) () {
	for _, v := range data {
		if ok, _ := utils.MapContains(r.dataSeriesWriters, v.GetSensorId()); ok {
			v.Write(t, r.dataSeriesWriters[v.GetSensorId()])
		} else {
			log.Error("time: %s, measurement id %s not found! ", t, v.GetSensorId())
		}
	}
	return
}

func (r *RowGroupWriter) FlushToFileWriter (tsFileIoWriter *tsFileWriter.TsFileIoWriter) () {
	for _, v := range r.dataSeriesWriters {
		v.WriteToFileWriter(tsFileIoWriter)
	}
	return
}

func (r *RowGroupWriter) PreFlush()(){
	// todo should flush current pages to mem.

	return
}

func (r *RowGroupWriter) GetCurrentRowGroupSize() (int64) {
	// todo get current size

	return 128
}

func (r *RowGroupWriter) GetSeriesNumber() (int) {
	return len(r.dataSeriesWriters)
}

func (r *RowGroupWriter) Close() (bool) {
	return true
}


func New(dId string) (*RowGroupWriter, error) {
	// todo do measurement init and memory check

	return &RowGroupWriter{
		deviceId:dId,
		dataSeriesWriters:make(map[string]seriesWriter.SeriesWriter),
	},nil
}