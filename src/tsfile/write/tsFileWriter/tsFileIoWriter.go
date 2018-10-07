package tsFileWriter

/**
 * @Package Name: tsFileWriter
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-13 上午11:24
 * @Description:
 */

import (
	"github.com/go_sample/src/tsfile/common/log"
	"os"
	"github.com/go_sample/src/tsfile/common/tsFileConf"
	"github.com/go_sample/src/tsfile/write/sensorDescriptor"
	"bytes"
	"github.com/go_sample/src/tsfile/write/statistics"
	"github.com/go_sample/src/tsfile/write/metaData"
	"github.com/go_sample/src/tsfile/common/header"
	"github.com/go_sample/src/tsfile/write/fileSchema"
	"github.com/go_sample/src/tsfile/common/utils"
)

type TsFileIoWriter struct {
	tsIoFile 						*os.File
	memBuf					 		*bytes.Buffer
	currentRowGroupMetaData			*metaData.RowGroupMetaData
	currentChunkMetaData			*metaData.TimeSeriesChunkMetaData
	rowGroupMetaDataSli				[]*metaData.RowGroupMetaData
	rowGroupHeader					*header.RowGroupHeader
	chunkHeader						*header.ChunkHeader
}

const(
	MAXVALUE="max_value"
	MINVALUE="min_value"
	FIRST="first"
	SUM="sum"
	LAST="last"
)

func (t *TsFileIoWriter) GetTsIoFile () (*os.File) {
	return t.tsIoFile
}

func (t *TsFileIoWriter) GetPos () (int64) {
	currentPos, _ := t.tsIoFile.Seek(0, os.SEEK_CUR)
	return currentPos
}

func (t *TsFileIoWriter) EndChunk (size int64, totalValueCount int64) () {
	// todo set currentChunkMetaData
	t.currentChunkMetaData.SetTotalByteSizeOfPagesOnDisk(size)
	t.currentChunkMetaData.SetNumOfPoints(totalValueCount)
	t.currentRowGroupMetaData.AddTimeSeriesChunkMetaData(t.currentChunkMetaData)
	// log.Info("end Chunk: %v, totalvalue: %v", t.currentChunkMetaData, totalValueCount)
	t.currentChunkMetaData = nil

	return
}

func (t *TsFileIoWriter) EndRowGroup (memSize int64) () {
	t.currentRowGroupMetaData.SetTotalByteSize(memSize)
	t.rowGroupMetaDataSli = append(t.rowGroupMetaDataSli, t.currentRowGroupMetaData)
	log.Info("end row group: %v", t.currentRowGroupMetaData)
	//t.currentRowGroupMetaData = nil
}

func (t *TsFileIoWriter) EndFile (fs fileSchema.FileSchema) () {
	timeSeriesMap := fs.GetTimeSeriesMetaDatas()
	// log.Info("get time series map: %v", timeSeriesMap)
	tsDeviceMetaDataMap := make(map[string]*metaData.TsDeviceMetaData)
	for _, v := range t.rowGroupMetaDataSli {
		if v == nil {
			continue
		}
		currentDevice := v.GetDeviceId()
		if _, contain := tsDeviceMetaDataMap[currentDevice]; !contain {
			tsDeviceMetaData, _ := metaData.NewTimeDeviceMetaData()
			tsDeviceMetaDataMap[currentDevice] = tsDeviceMetaData
		}
		tdmd := tsDeviceMetaDataMap[currentDevice]
		tdmd.AddRowGroupMetaData(v)
		tsDeviceMetaDataMap[currentDevice] = tdmd
		// tsDeviceMetaDataMap[currentDevice].AddRowGroupMetaData(*v)
	}

	for _, tsDeviceMetaData := range tsDeviceMetaDataMap {
		startTime := int64(1) 	//int64(0x7fffffffffffffff)
		endTime := int64(1)		//int64(0x8000000000000000)
		for _, rowGroup := range tsDeviceMetaData.GetRowGroups() {
			for _, timeSeriesChunkMetaData := range rowGroup.GetTimeSeriesChunkMetaDataSli() {
				startTime = min(startTime, timeSeriesChunkMetaData.GetStartTime())
				endTime = max(endTime, timeSeriesChunkMetaData.GetEndTime())
			}
		}
		tsDeviceMetaData.SetStartTime(startTime)
		tsDeviceMetaData.SetEndTime(endTime)
	}
	tsFileMetaData, _ := metaData.NewTsFileMetaData(tsDeviceMetaDataMap, timeSeriesMap, tsFileConf.CurrentVersion)
	footerIndex := t.GetPos()
	log.Info("start to flush meta, file pos: %d", footerIndex)
	size := tsFileMetaData.SerializeTo(t.memBuf)
	// log.Info("t.memBuf: %s", t.memBuf)
	log.Info("finish flushing meta %v, file pos: %d", tsFileMetaData, t.GetPos())
	t.memBuf.Write(utils.Int32ToByte(int32(size)))
	t.memBuf.Write([]byte(tsFileConf.MAGIC_STRING))


	// flush mem-filemeta to file
	t.WriteBytesToFile(t.memBuf)
	log.Info("file pos: %d", t.GetPos())
}

func max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}


func (t *TsFileIoWriter) WriteMagic()(int){
	n, err := t.tsIoFile.Write([]byte(tsFileConf.MAGIC_STRING))
	if err != nil {
		log.Error("write start magic to file err: ", err)
	}
	return n
}

func (t *TsFileIoWriter) StartFlushRowGroup(deviceId string, rowGroupSize int64, seriesNumber int32)(int){
	log.Info("start flush rowgroup.")
	// timeSeriesChunkMetaDataMap := make(map[string]metaData.TimeSeriesChunkMetaData)
	timeSeriesChunkMetaDataSli := make([]*metaData.TimeSeriesChunkMetaData, 0)
	t.currentRowGroupMetaData, _ = metaData.NewRowGroupMetaData(deviceId, 0, t.GetPos(), timeSeriesChunkMetaDataSli)
	t.currentRowGroupMetaData.RecalculateSerializedSize()
	rowGroupHeader, _ := header.NewRowGroupHeader(deviceId, rowGroupSize, seriesNumber)
	log.Info("rowGroupHeader: %v", rowGroupHeader)
	rowGroupHeader.RowGroupHeaderToMemory(t.memBuf)
	t.rowGroupHeader = rowGroupHeader
	// rowgroup header bytebuffer write to file
	t.WriteBytesToFile(t.memBuf)
	// truncate bytebuffer to empty
	t.memBuf.Reset()

	return header.GetRowGroupSerializedSize(deviceId)
}

func (t *TsFileIoWriter) StartFlushChunk(sd *sensorDescriptor.SensorDescriptor, compressionType int16,
								tsDataType int16, encodingType int16, statistics statistics.Statistics,
								maxTimestamp int64, minTimestamp int64, pageBufSize int, numOfPages int)(int){
	t.currentChunkMetaData, _ = metaData.NewTimeSeriesChunkMetaData(sd.GetSensorId(), t.GetPos(), minTimestamp, maxTimestamp)
	chunkHeader, _ := header.NewChunkHeader(sd.GetSensorId(), pageBufSize, tsDataType, compressionType, encodingType, numOfPages, 0)
	chunkHeader.ChunkHeaderToMemory(t.memBuf)
	t.chunkHeader = chunkHeader
	// chunk header bytebuffer write to file
	t.WriteBytesToFile(t.memBuf)
	// truncate bytebuffer to empty
	t.memBuf.Reset()
	// todo set tsdigest
	tsDigest, _ := metaData.NewTsDigest()
	statisticsMap := make(map[string]bytes.Buffer)
	//var max bytes.Buffer
	//max.Write(statistics.GetMaxByte(tsDataType))
	//statisticsMap[MAXVALUE] = max
	//var min bytes.Buffer
	//min.Write(statistics.GetMinByte(tsDataType))
	//statisticsMap[MINVALUE] = min
	//var first bytes.Buffer
	//first.Write(statistics.GetFirstByte(tsDataType))
	//statisticsMap[FIRST] = first
	//var sum bytes.Buffer
	//sum.Write(statistics.GetSumByte(tsDataType))
	//statisticsMap[SUM] = sum
	//var last bytes.Buffer
	//last.Write(statistics.GetLastByte(tsDataType))
	//statisticsMap[LAST] = last

	var min bytes.Buffer
	min.Write(statistics.GetMinByte(tsDataType))
	statisticsMap[MINVALUE] = min
	var last bytes.Buffer
	last.Write(statistics.GetLastByte(tsDataType))
	statisticsMap[LAST] = last
	var sum bytes.Buffer
	sum.Write(statistics.GetSumByte(tsDataType))
	statisticsMap[SUM] = sum
	var first bytes.Buffer
	first.Write(statistics.GetFirstByte(tsDataType))
	statisticsMap[FIRST] = first
	var max bytes.Buffer
	max.Write(statistics.GetMaxByte(tsDataType))
	statisticsMap[MAXVALUE] = max

	tsDigest.SetStatistics(statisticsMap)
	t.currentChunkMetaData.SetDigest(*tsDigest)
	return header.GetChunkSerializedSize(sd.GetSensorId())
}

func (t *TsFileIoWriter) WriteBytesToFile (buf *bytes.Buffer) () {
	//声明一个空的slice,容量为timebuf的长度
	timeSlice := make([]byte, buf.Len())
	//把buf的内容读入到timeSlice内,因为timeSlice容量为timeSize,所以只读了timeSize个过来
	buf.Read(timeSlice)
	t.tsIoFile.Write(timeSlice)
	return
}

func NewTsFileIoWriter(file string) (*TsFileIoWriter, error) {
	// todo
	newFile, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Info("open file:%s failed.", file)
	}

	return &TsFileIoWriter{
		tsIoFile:newFile,
		memBuf:bytes.NewBuffer([]byte{}),
		rowGroupMetaDataSli:make([]*metaData.RowGroupMetaData, 0),
	},nil
}