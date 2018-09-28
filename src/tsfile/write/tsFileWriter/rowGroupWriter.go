package tsFileWriter

/**
 * @Package Name: rowGroupWriter
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-28 下午6:19
 * @Description:
 */

import (
	"github.com/go_sample/src/tsfile/write/sensorDescriptor"
	"github.com/go_sample/src/tsfile/common/utils"
	"github.com/go_sample/src/tsfile/common/log"
)

type RowGroupWriter struct {
	deviceId			string
	dataSeriesWriters	map[string]SeriesWriter
}

func (r *RowGroupWriter) AddSeriesWriter(sd sensorDescriptor.SensorDescriptor, pageSize int) () {
	if contain, _ := utils.MapContains(r.dataSeriesWriters, sd.GetSensorId()); !contain {
		// todo new pagewrite
		pw, _ := NewPageWriter(sd)

		// new serieswrite
		sw, _ := NewSeriesWriter(r.deviceId, sd, *pw, pageSize)

		r.dataSeriesWriters[sd.GetSensorId()] = *sw
	} else {
		log.Error("given sensor has exist, need not add to series writer again.")
	}
	return
}

func (r *RowGroupWriter) Write(t int64, data map[string]DataPoint) () {
	for _, v := range data {
		if ok, _ := utils.MapContains(r.dataSeriesWriters, v.GetSensorId()); ok {
			v.Write(t, r.dataSeriesWriters[v.GetSensorId()])
		} else {
			log.Error("time: %d, sensor id %s not found! ", t, v.GetSensorId())
		}
	}
	return
}

func (r *RowGroupWriter) FlushToFileWriter (tsFileIoWriter *TsFileIoWriter) () {
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

func (r *RowGroupWriter) GetSeriesNumber() (int32) {
	return int32(len(r.dataSeriesWriters))
}

func (r *RowGroupWriter) Close() (bool) {
	return true
}


func NewRowGroupWriter(dId string) (*RowGroupWriter, error) {
	// todo do measurement init and memory check

	return &RowGroupWriter{
		deviceId:dId,
		dataSeriesWriters:make(map[string]SeriesWriter),
	},nil
}