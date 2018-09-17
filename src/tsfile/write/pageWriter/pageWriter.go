package pageWriter

/**
 * @Package Name: pageWriter
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-31 下午3:51
 * @Description:
 */

import (
	"github.com/go_sample/src/tsfile/write/sensorDescriptor"
	"bytes"
	"github.com/go_sample/src/tsfile/write/statistics"
	"github.com/go_sample/src/tsfile/common/header"
	"github.com/go_sample/src/tsfile/common/log"
	"github.com/go_sample/src/tsfile/write/tsFileWriter"
	"os"
	"github.com/golang/net/html/atom"
)

type PageWriter struct {
	compressor 			string //Compressor
	desc 				sensorDescriptor.SensorDescriptor
	// todo this buf should change to compert type
	pageBuf 			*bytes.Buffer
	totalValueCount		int64
	maxTimestamp		int64
	minTimestamp		int64
}

func (p *PageWriter) WritePageHeaderAndDataIntoBuff(dataBuffer *bytes.Buffer, valueCount int, sts statistics.Statistics, maxTimestamp int64, minTimestamp int64) (int) {
	//this uncompressedSize should be calculate from timeBuf and valueBuf
	uncompressedSize := dataBuffer.Len()
	compressedSize := 0
	pageHeader, pageHeaderErr := header.NewPageHeader(int32(uncompressedSize), int32(compressedSize), int32(valueCount), sts, maxTimestamp, minTimestamp, p.desc.GetTsDataType())
	if pageHeaderErr != nil {
		log.Error("init pageHeader error: ", pageHeaderErr)
	}
	// write pageheader to pageBuf
	pageHeader.PageHeaderToMemory(*p.pageBuf)

	// write pageData to pageBuf
	//声明一个空的slice,容量为timebuf的长度
	timeSlice := make([]byte, dataBuffer.Len())
	//把buf的内容读入到timeSlice内,因为timeSlice容量为timeSize,所以只读了timeSize个过来
	dataBuffer.Read(timeSlice)
	p.pageBuf.Write(timeSlice)
	p.totalValueCount += int64(valueCount)
	return 0
}

func (p *PageWriter) WriteAllPagesOfSeriesToTsFile (tsFileIoWriter *tsFileWriter.TsFileIoWriter, seriesStatistics statistics.Statistics, numOfPage int) (int64) {
	if p.minTimestamp == -1 {
		log.Error("Write page error, minTime: %s, maxTime: %s")
	}
	chunkHeaderSize :=  tsFileIoWriter.StartFlushChunk(p.desc, header.UNCOMPRESSED, p.desc.GetTsDataType(), p.desc.GetTsEncoding(), seriesStatistics, p.maxTimestamp, p.minTimestamp, p.pageBuf.Len(), numOfPage)
	preSize := tsFileIoWriter.GetPos()
	tsFileIoWriter.WriteBytesToFile(p.pageBuf)
	dataSize := tsFileIoWriter.GetPos() - preSize
	chunkSize := int64(chunkHeaderSize) + dataSize
	tsFileIoWriter.EndTrunk(chunkSize, p.totalValueCount)
	return chunkSize
}

func (p *PageWriter) Reset () () {
	p.minTimestamp = -1
	p.pageBuf.Reset()
	p.totalValueCount = 0
	return
}


func New(sd sensorDescriptor.SensorDescriptor) (*PageWriter, error) {
	// todo do measurement init and memory check

	return &PageWriter{
		desc:sd,
		compressor:sd.GetCompressor(),
		pageBuf:bytes.NewBuffer([]byte{}),
	},nil
}