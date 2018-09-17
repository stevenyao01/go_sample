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
)

type TsFileIoWriter struct {
	tsIoFile 						*os.File
	memBuf					 		*bytes.Buffer
	currentRowGroupMetaData			*metaData.RowGroupMetaData
	currentChunkMetaData			*metaData.TimeSeriesChunkMetaData
}

func (t *TsFileIoWriter) GetTsIoFile () (*os.File) {
	return t.tsIoFile
}

func (t *TsFileIoWriter) GetPos () (int64) {
	currentPos, _ := t.tsIoFile.Seek(0, os.SEEK_CUR)
	return currentPos
}

func (t *TsFileIoWriter) EndTrunk (size int64, totalValueCount int64) () {
	// todo set currentChunkMetaData

	return
}

func (t *TsFileIoWriter) WriteMagic()(int){
	n, err := t.tsIoFile.Write([]byte(tsFileConf.MAGIC_STRING))
	if err == nil {
		log.Error("write start magic to file err: ", err)
	}
	return n
}

func (t *TsFileIoWriter) StartFlushRowGroup(deviceId string, rowGroupSize int64, seriesNumber int)(int){

	return 0
}

func (t *TsFileIoWriter) StartFlushChunk(sd sensorDescriptor.SensorDescriptor, compressionType int16,
								tsDataType int16, encodingType int16, statistics statistics.Statistics,
								maxTimestamp int64, minTimestamp int64, pageBufSize int, numOfPages int)(int){
	t.currentChunkMetaData, _ = metaData.NewTimeSeriesChunkMetaData(sd.GetSensorId(), t.GetPos(), minTimestamp, maxTimestamp)
	chunkHeader, _ := header.NewChunkHeader(sd.GetSensorId(), pageBufSize, tsDataType, compressionType, encodingType, numOfPages)
	chunkHeader.ChunkHeaderToMemory(*t.memBuf)
	// todo set tsdigest


	return 0
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
	},nil
}