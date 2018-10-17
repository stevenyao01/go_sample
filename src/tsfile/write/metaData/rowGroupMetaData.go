package metaData

/**
 * @Package Name: metaData
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-13 下午2:37
 * @Description:
 */

import (
	"github.com/go_sample/src/tsfile/common/log"
	"bytes"
	"github.com/go_sample/src/tsfile/common/utils"
)

type RowGroupMetaData struct {
	//tsIoFile 						*os.File
	deviceId						string
	totalByteSize					int64
	fileOffsetOfCorrespondingData	int64
	TimeSeriesChunkMetaDataSli		[]*TimeSeriesChunkMetaData
	serializedSize					int
	sizeOfChunkSli					int
}

func (r *RowGroupMetaData) AddTimeSeriesChunkMetaData (md *TimeSeriesChunkMetaData) () {
	if len(r.TimeSeriesChunkMetaDataSli) == 0 {
		r.TimeSeriesChunkMetaDataSli = make([]*TimeSeriesChunkMetaData, 0)
	}
	r.TimeSeriesChunkMetaDataSli = append(r.TimeSeriesChunkMetaDataSli, md)
	r.serializedSize += md.GetSerializedSize()
	r.sizeOfChunkSli += 1
}

func (r *RowGroupMetaData) SetTotalByteSize (ms int64) () {
	r.totalByteSize = ms
}

func (r *RowGroupMetaData) GetDeviceId () (string) {
	return r.deviceId
}

func (r *RowGroupMetaData) SerializeTo (buf *bytes.Buffer) (int) {
	if r.sizeOfChunkSli != len(r.TimeSeriesChunkMetaDataSli) {
		r.RecalculateSerializedSize()
	}
	var byteLen int

	n1, _ := buf.Write(utils.Int32ToByte(int32(len(r.deviceId))))
	byteLen += n1
	n2, _ := buf.Write([]byte(r.deviceId))
	byteLen += n2

	n3, _ := buf.Write(utils.Int64ToByte(r.totalByteSize))
	byteLen += n3
	n4, _ := buf.Write(utils.Int64ToByte(r.fileOffsetOfCorrespondingData))
	byteLen += n4

	n5, _ := buf.Write(utils.Int32ToByte(int32(len(r.TimeSeriesChunkMetaDataSli))))
	byteLen += n5
	for _, v := range r.TimeSeriesChunkMetaDataSli {
		byteLen += v.SerializeTo(buf)
	}

	return byteLen
}

func (r *RowGroupMetaData) GetTimeSeriesChunkMetaDataSli () ([]*TimeSeriesChunkMetaData) {
	//if r.TimeSeriesChunkMetaDataSli == nil {
	//	return nil
	//}
	return r.TimeSeriesChunkMetaDataSli
}

func (r *RowGroupMetaData) GetserializedSize () (int) {
	if r.sizeOfChunkSli != len(r.TimeSeriesChunkMetaDataSli) {
		r.RecalculateSerializedSize()
	}
	return r.serializedSize
}

func (r *RowGroupMetaData) RecalculateSerializedSize () () {
	r.serializedSize = 1 *4 + len(r.deviceId) + 2 * 8 + 1 * 4
	for _, v := range r.TimeSeriesChunkMetaDataSli {
		if &v != nil {
			r.serializedSize += v.GetSerializedSize()
			log.Info("timeSeriesChunkMetaDataSliaaaaaa: %s", v)
		}
	}
	r.sizeOfChunkSli = len(r.TimeSeriesChunkMetaDataSli)
	return
}

func NewRowGroupMetaData(dId string, tbs int64, foocd int64, tscmds []*TimeSeriesChunkMetaData) (*RowGroupMetaData, error) {
	return &RowGroupMetaData{
		deviceId:dId,
		totalByteSize:tbs,
		fileOffsetOfCorrespondingData:foocd,
		TimeSeriesChunkMetaDataSli:tscmds,
	},nil
}