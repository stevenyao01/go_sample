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
	"github.com/go_sample/src/tsfile/write/valueWriter"
	"github.com/go_sample/src/tsfile/common/tsFileConf"
	"github.com/go_sample/src/tsfile/common/log"
	"github.com/go_sample/src/tsfile/write/statistics"
	"github.com/go_sample/src/tsfile/write/tsFileWriter"
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
	valueWriter			valueWriter.ValueWriter
	/* value count on a page. It will be reset agter calling */
	valueCount 			int
	valueCountForNextSizeCheck	int
	/*statistics on a page. It will be reset after calling */
	pageStatistics		statistics.Statistics
	seriesStatistics	statistics.Statistics
	time				int64
	minTimestamp		int64
	sensorDescriptor	sensorDescriptor.SensorDescriptor
	minimumRecordCountForCheck int
	numOfPages			int
}

func (s *SeriesWriter) GetTsDataType() (int) {
	return s.tsDataType
}

func (s *SeriesWriter) GetTsDeviceId() (string) {
	return s.deviceId
}

func (s *SeriesWriter) GetNumOfPages() (int) {
	return s.numOfPages
}

func (s *SeriesWriter) Write(t int64, value interface{}) (bool) {
	s.time = t
	s.valueCount = s.valueCount + 1
	s.valueWriter.Write(t, s.tsDataType, value)
	// todo statistics ignore here, if necessary, Statistics.java
	if s.minTimestamp == -1 {
		s.minTimestamp = t
	}
	// todo check page size and write page data to buffer
	s.checkPageSizeAndMayOpenNewpage()
	return true
}

func (s *SeriesWriter) WriteToFileWriter (tsFileIoWriter *tsFileWriter.TsFileIoWriter) () {
	s.pageWriter.WriteAllPagesOfSeriesToTsFile(tsFileIoWriter, s.seriesStatistics, s.numOfPages)
	s.pageWriter.Reset()
	// reset series_statistics
	s.seriesStatistics = *statistics.GetStatistics(s.tsDataType)
}

func (s *SeriesWriter)checkPageSizeAndMayOpenNewpage() () {
	if s.valueCount == tsFileConf.MaxNumberOfPointsInPage {
		log.Info("current line count reaches the upper bound, write page %s", s.sensorDescriptor)
		// todo write data to buffer
		s.WritePage()
	} else if s.valueCount >= s.valueCountForNextSizeCheck {
		currentColumnSize := s.valueWriter.GetCurrentMemSize()
		if currentColumnSize > s.psThres {
			// todo write data to buffer
			s.WritePage()
		} else {
			log.Info("not enough size now.")
		}
		// int * 1.0 / int 为float， 再乘以valueCount，得到下次检查的count
		s.valueCountForNextSizeCheck = s.psThres * 1.0 / currentColumnSize * s.valueCount
	}
}

func (s *SeriesWriter) WritePage()(){
	s.pageWriter.WritePageHeaderAndDataIntoBuff(s.valueWriter.GetByteBuffer(), s.valueCount, s.pageStatistics, s.time, s.minTimestamp)
	// todo pageStatistics
	s.numOfPages += 1

	//
	s.minTimestamp = -1
	s.valueCount = 0
	s.valueWriter.Reset()
	s.ResetPageStatistics()
	return
}

func (s *SeriesWriter) ResetPageStatistics()(){
	s.pageStatistics = *statistics.GetStatistics(s.tsDataType)
	return
}


func New(dId string, d sensorDescriptor.SensorDescriptor, pw pageWriter.PageWriter, pst int) (*SeriesWriter, error) {

	return &SeriesWriter{
		deviceId:dId,
		desc:d,
		pageWriter:pw,
		psThres:pst,
		pageCountUpperBound:tsFileConf.MaxNumberOfPointsInPage,
		minimumRecordCountForCheck:1,
		valueCountForNextSizeCheck:1,
		numOfPages:0,
		tsDataType:d.GetTsDataType(),
		seriesStatistics:*statistics.GetStatistics(d.GetTsDataType()),
	},nil
}