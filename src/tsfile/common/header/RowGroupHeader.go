package header

import (
	"bytes"
	"github.com/go_sample/src/tsfile/common/utils"
)

type RowGroupHeader struct {
	deviceId			string
	dataSize			int64
	numOfChunks			int32
	serializedSize		int32
}

func (r *RowGroupHeader) RowGroupHeaderToMemory (buffer *bytes.Buffer) (int32) {
	// todo write header to buffer
	buffer.Write(utils.Int32ToByte(int32(len(r.deviceId))))
	buffer.Write([]byte(r.deviceId))
	buffer.Write(utils.Int64ToByte(r.dataSize))
	buffer.Write(utils.Int32ToByte(r.numOfChunks))

	return r.serializedSize
}

func (r *RowGroupHeader) GetSerializedSize () (int32) {
	return r.serializedSize
}

func NewRowGroupHeader(dId string, rgs int64, sn int32) (*RowGroupHeader, error) {
	// todo
	ss := 1 * 4 + 1 * 8 + len(dId)
	return &RowGroupHeader{
		deviceId:dId,
		dataSize:rgs,
		numOfChunks:sn,
		serializedSize:int32(ss),
	},nil
}