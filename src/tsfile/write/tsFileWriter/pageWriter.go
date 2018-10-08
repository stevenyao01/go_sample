package tsFileWriter

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
	"github.com/go_sample/src/tsfile/common/compress"
)

type PageWriter struct {
	compressor 			*compress.Encompress
	desc 				*sensorDescriptor.SensorDescriptor
	pageBuf 			*bytes.Buffer
	totalValueCount		int64
	maxTimestamp		int64
	minTimestamp		int64
}

func (p *PageWriter) WritePageHeaderAndDataIntoBuff(dataBuffer *bytes.Buffer, valueCount int, sts statistics.Statistics, maxTimestamp int64, minTimestamp int64) (int) {
	//this uncompressedSize should be calculate from timeBuf and valueBuf
	uncompressedSize := dataBuffer.Len()

	// write pageData to pageBuf
	//声明一个空的slice,容量为dataBuffer的长度
	dataSlice := make([]byte, dataBuffer.Len())
	//把buf的内容读入到timeSlice内,因为timeSlice容量为timeSize,所以只读了timeSize个过来
	dataBuffer.Read(dataSlice)

	var compressedSize int
	var enc []byte
	if p.desc.GetCompresstionType() == header.UNCOMPRESSED {
		compressedSize = uncompressedSize
	} else {
		aSlice := make([]byte, 0)
		enc = p.compressor.GetEncompressor(p.desc.GetCompresstionType()).Encompress(aSlice, dataSlice)
		compressedSize = len(enc)
	}


	pageHeader, pageHeaderErr := header.NewPageHeader(int32(uncompressedSize), int32(compressedSize), int32(valueCount), sts, maxTimestamp, minTimestamp, p.desc.GetTsDataType())
	if pageHeaderErr != nil {
		log.Error("init pageHeader error: ", pageHeaderErr)
	}
	// write pageheader to pageBuf
	log.Info("start to flush a page header into buffer, buf pos: %d", p.pageBuf.Len())
	pageHeader.PageHeaderToMemory(p.pageBuf)
	log.Info("pageHeader: %v", pageHeader)
	log.Info("finished to flush a page header into buffer, buf pos: %d", p.pageBuf.Len())

	//// write pageData to pageBuf
	////声明一个空的slice,容量为dataBuffer的长度
	//dataSlice := make([]byte, dataBuffer.Len())
	////把buf的内容读入到timeSlice内,因为timeSlice容量为timeSize,所以只读了timeSize个过来
	//dataBuffer.Read(dataSlice)
	log.Info("start to flush a page data into buffer, buf pos: %d", p.pageBuf.Len())
	if p.desc.GetCompresstionType() == header.UNCOMPRESSED {
		p.pageBuf.Write(dataSlice)
	} else {
		p.pageBuf.Write(enc)
	}
	log.Info("finished to flush a page data into buffer, buf pos: %d", p.pageBuf.Len())
	p.totalValueCount += int64(valueCount)
	return 0
}

func (p *PageWriter) WriteAllPagesOfSeriesToTsFile (tsFileIoWriter *TsFileIoWriter, seriesStatistics statistics.Statistics, numOfPage int) (int64) {
	if p.minTimestamp == -1 {
		log.Error("Write page error, minTime: %s, maxTime: %s")
	}
	// write trunk header to file
	chunkHeaderSize :=  tsFileIoWriter.StartFlushChunk(p.desc, header.UNCOMPRESSED, p.desc.GetTsDataType(), p.desc.GetTsEncoding(), seriesStatistics, p.maxTimestamp, p.minTimestamp, p.pageBuf.Len(), numOfPage)
	preSize := tsFileIoWriter.GetPos()
	// write all pages to file
	tsFileIoWriter.WriteBytesToFile(p.pageBuf)
	//// after write page, reset pageBuf
	//p.pageBuf.Reset()
	dataSize := tsFileIoWriter.GetPos() - preSize
	chunkSize := int64(chunkHeaderSize) + dataSize
	tsFileIoWriter.EndChunk(chunkSize, p.totalValueCount)
	return chunkSize
}

func (p *PageWriter) Reset () () {
	p.minTimestamp = -1
	p.pageBuf.Reset()
	p.totalValueCount = 0
	return
}

func (p *PageWriter) EstimateMaxPageMemSize () (int) {
	pageSize := p.pageBuf.Len()
	pageHeaderSize := header.CalculatePageHeaderSize(p.desc.GetTsDataType())
	return pageSize + pageHeaderSize
}

func (p *PageWriter) GetCurrentDataSize () (int) {
	size := p.pageBuf.Len()
	return size
}


func NewPageWriter(sd *sensorDescriptor.SensorDescriptor) (*PageWriter, error) {
	return &PageWriter{
		desc:sd,
		compressor:sd.GetCompressor(),
		pageBuf:bytes.NewBuffer([]byte{}),
	},nil
}