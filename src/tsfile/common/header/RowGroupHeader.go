package header

import (
	"github.com/go_sample/src/tsfile/common/log"
	"bytes"
	"github.com/go_sample/src/tsfile/write/statistics"
)

type RowGroupHeader struct {
	deltaObjectId		string
	dataSize			uint64
	numOfChunks			int
	serializedSize		int
}

//func (p *PageHeader)HeadToMemory(buffer bytes.Buffer)(int32){
//	// todo write header to buffer
//	buffer.Write(utils.Int32ToByte(p.uncompressedSize))
//	buffer.Write(utils.Int32ToByte(p.compressedSize))
//	buffer.Write(utils.Int32ToByte(p.numOfValues))
//	buffer.Write(utils.Int64ToByte(p.max_timestamp))
//	buffer.Write(utils.Int64ToByte(p.min_timestamp))
//	p.statistics.Serialize(buffer)
//	return p.serializedSize
//}

func NewRowGroupHeader(ucs int32, cs int32, nov int32, sts statistics.Statistics, max_t int64, min_t int64, tsDataType int) (*RowGroupHeader, error) {
	// todo
	//ss := 3 * 4 + 2 * 8 + sts.GetHeaderSize(tsDataType)
	return &RowGroupHeader{
	},nil
}


//func NewRowGroupHeader() (*RowGroupHeader, error) {
// return &RowGroupHeader{},nil
//}