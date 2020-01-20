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
	"github.com/go_sample/src/tsfile/common/header"
)

type RowGroupWriter struct {
	deviceId			string
	dataSeriesWriters	map[string]*SeriesWriter
}

func (r *RowGroupWriter) AddSeriesWriter(sd *sensorDescriptor.SensorDescriptor, pageSize int) () {
	if contain, _ := utils.MapContains(r.dataSeriesWriters, sd.GetSensorId()); !contain {
		// new pagewriter
		pw, _ := NewPageWriter(sd)

		// new serieswrite
		sw, _ := NewSeriesWriter(r.deviceId, sd, pw, pageSize)
		r.dataSeriesWriters[sd.GetSensorId()] = sw
	} else {
		log.Info("given sensor has exist, need not add to series writer again.")
	}
	return
}

func (r *RowGroupWriter) Write(t int64, data []*DataPoint) () {
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
	// flush current pages to mem.
	for _, v := range r.dataSeriesWriters {
		v.PreFlush()
	}
	return
}

func (r *RowGroupWriter) GetCurrentRowGroupSize() (int) {
	// get current size
	//size := int64(tfiw.rowGroupHeader.GetRowGroupSerializedSize())
	rowGroupHeaderSize := header.GetRowGroupSerializedSize(r.deviceId)
	size := rowGroupHeaderSize
	for k, v := range r.dataSeriesWriters {
		size += v.GetCurrentChunkSize(k)
	}

	return size
}

func (r *RowGroupWriter) GetSeriesNumber() (int32) {
	return int32(len(r.dataSeriesWriters))
}

func (r *RowGroupWriter) UpdateMaxGroupMemSize () (int64) {
	var bufferSize int64
	for _, v := range r.dataSeriesWriters {
		bufferSize += v.EstimateMaxSeriesMemSize()
	}
	return bufferSize
}

func (r *RowGroupWriter) Close() (bool) {
	return true
}


func NewRowGroupWriter(dId string) (*RowGroupWriter, error) {
	return &RowGroupWriter{
		deviceId:dId,
		dataSeriesWriters:make(map[string]*SeriesWriter),
	},nil
}