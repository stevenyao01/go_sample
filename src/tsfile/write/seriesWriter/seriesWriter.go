package seriesWriter

/**
 * @Package Name: seriesWriter
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-28 下午8:28
 * @Description:
 */

import (
//"github.com/go_sample/src/tsfile/common/log"
	"github.com/go_sample/src/tsfile/write/sensorDescriptor"
	"database/sql/driver"
	"time"
	"github.com/go_sample/src/tsfile/write/pageWriter"
)

type SeriesWriter struct {
	deviceId			string
	dataSeriesWriters	map[string]SeriesWriter

	desc				sensorDescriptor.SensorDescriptor
	tsDataType 			int
	pageWriter			pageWriter.PageWriter
	/* page size threshold 	*/
	psThres				int
	pageCountUpperBound	int
	/* value writer to encode data*/
	valueWriter			dataValueWriter.DataValueWriter
	/* value count on a page. It will be reset agter calling */
	valueCount 			int
	valueCountForNextSizeCheck	int
	/*statistics on a page. It will be reset after calling */
	pageStatistics		statistics.Statistics
	seriesStatistics	statistics.Statistics
	time				time.Time
	minTimestamp		int64
	sensorDescriptor	sensorDescriptor.SensorDescriptor
}

func (s *SeriesWriter) GetTsDataType() (int) {
	return s.tsDataType
}

func (s *SeriesWriter) GetTsDeviceId() (string) {
	return s.deviceId
}

func (s *SeriesWriter) Write(t time.Time, value interface{}) (bool) {
	s.time = t
	s.valueCount = s.valueCount + 1
	s.da

	return true
}


func New(dId string, d sensorDescriptor.SensorDescriptor, pw pageWriter.PageWriter, pst int) (*SeriesWriter, error) {
	// todo do measurement init and memory check

	return &SeriesWriter{
		deviceId:dId,
		desc:d,
		pageWriter:pw,
		psThres:pst,
	},nil
}