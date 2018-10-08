package metaData

/**
 * @Package Name: metaData
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-13 下午9:00
 * @Description:
 */

import (
	"bytes"
	"github.com/go_sample/src/tsfile/common/utils"
)

type TimeSeriesChunkMetaData struct {
	sensorId							string
	fileOffsetOfCorrespondingData		int64
	startTime							int64
	endTime								int64
	valueStatistics						TsDigest
	totalByteSizeOfPagesOnDisk			int64
	numOfPoints							int64
}

func (t *TimeSeriesChunkMetaData) GetSerializedSize () (int) {
	if t.valueStatistics.sizeOfList <= 0 {
		return 1 * 4 + len(t.sensorId) + 5 * 8 + t.valueStatistics.GetNullDigestSize()
	}
	return 1 * 4 + len(t.sensorId) + 5 * 8 + t.valueStatistics.GetSerializedSize()
}

func (t *TimeSeriesChunkMetaData) SetDigest (tsDigest TsDigest) () {
	t.valueStatistics = tsDigest
}

func (t *TimeSeriesChunkMetaData) GetStartTime () (int64) {
	return t.startTime
}

func (t *TimeSeriesChunkMetaData) GetEndTime () (int64) {
	return t.endTime
}

func (t *TimeSeriesChunkMetaData) SetTotalByteSizeOfPagesOnDisk (size int64) () {
	t.totalByteSizeOfPagesOnDisk = size
}

func (t *TimeSeriesChunkMetaData) SetNumOfPoints (num int64) () {
	t.numOfPoints = num
}

func (t *TimeSeriesChunkMetaData) SerializeTo (buf *bytes.Buffer) (int) {
	var byteLen int

	n1, _ := buf.Write(utils.Int32ToByte(int32(len(t.sensorId))))
	byteLen += n1
	n2, _ := buf.Write([]byte(t.sensorId))
	byteLen += n2

	n3, _ := buf.Write(utils.Int64ToByte(t.fileOffsetOfCorrespondingData))
	byteLen += n3
	n4, _ := buf.Write(utils.Int64ToByte(t.numOfPoints))
	byteLen += n4
	n5, _ := buf.Write(utils.Int64ToByte(t.totalByteSizeOfPagesOnDisk))
	byteLen += n5
	n6, _ := buf.Write(utils.Int64ToByte(t.startTime))
	byteLen += n6
	n7, _ := buf.Write(utils.Int64ToByte(t.endTime))
	byteLen += n7

	if t.valueStatistics.sizeOfList <= 0 {
		byteLen += t.valueStatistics.GetNullDigestSize()
	} else {
		// tsDigest serializeTo
		byteLen += t.valueStatistics.serializeTo(buf)
	}

	return byteLen
}

func NewTimeSeriesChunkMetaData(sid string, fOffset int64, sTime int64, eTime int64) (*TimeSeriesChunkMetaData, error) {
	return &TimeSeriesChunkMetaData{
		sensorId:sid,
		fileOffsetOfCorrespondingData:fOffset,
		startTime:sTime,
		endTime:eTime,
		totalByteSizeOfPagesOnDisk:0,
		numOfPoints:0,
	},nil
}