package header

import (
	"github.com/go_sample/src/tsfile/write/statistics"
	"bytes"
	"github.com/go_sample/src/tsfile/common/utils"
)

type PageHeader struct {
	uncompressedSize		int32
	compressedSize			int32
	numOfValues				int32
	max_timestamp			int64
	min_timestamp			int64
	statistics				statistics.Statistics
	serializedSize			int32
}

func (p *PageHeader)PageHeaderToMemory(buffer *bytes.Buffer)(int32){
	// write header to buffer
	buffer.Write(utils.Int32ToByte(p.uncompressedSize))
	buffer.Write(utils.Int32ToByte(p.compressedSize))
	buffer.Write(utils.Int32ToByte(p.numOfValues))
	buffer.Write(utils.Int64ToByte(p.max_timestamp))
	buffer.Write(utils.Int64ToByte(p.min_timestamp))
	p.statistics.Serialize(buffer)
	return p.serializedSize
}

func CalculatePageHeaderSize (tsDataType int16) (int) {
	pHeaderSize := 3 * 4 + 2 * 8
	statisticsSize := statistics.GetStatistics(tsDataType).GetserializedSize(tsDataType)
	return pHeaderSize + statisticsSize
}

func NewPageHeader(ucs int32, cs int32, nov int32, sts statistics.Statistics, max_t int64, min_t int64, tsDataType int16) (*PageHeader, error) {
	ss := 3 * 4 + 2 * 8 + sts.GetHeaderSize(tsDataType)
	return &PageHeader{
		uncompressedSize:ucs,
		compressedSize:cs,
		numOfValues:nov,
		max_timestamp:max_t,
		min_timestamp:min_t,
		statistics:sts,
		serializedSize:int32(ss),
	},nil
}