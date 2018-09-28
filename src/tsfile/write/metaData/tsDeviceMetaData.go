package metaData

/**
 * @Package Name: metaData
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-20 上午11:42
 * @Description:
 */

import (
	"bytes"
	"github.com/go_sample/src/tsfile/common/utils"
	"github.com/go_sample/src/tsfile/common/log"
)

type TsDeviceMetaData struct {
	sensorId							string
	tsDataType							int16
	rowGroupmetaDataSli					[]RowGroupMetaData
	sizeOfRowGroupMetaDataSli			int
	serializedSize						int
	startTime							int64
	endTime								int64
}

func (t *TsDeviceMetaData) AddRowGroupMetaData (rgmd RowGroupMetaData) () {
	if len(t.rowGroupmetaDataSli) == 0 {
		t.rowGroupmetaDataSli = make([]RowGroupMetaData, 0)
	}
	t.rowGroupmetaDataSli = append(t.rowGroupmetaDataSli, rgmd)
	t.sizeOfRowGroupMetaDataSli += 1
	t.serializedSize += rgmd.GetserializedSize()
}

func (t *TsDeviceMetaData) GetRowGroups () ([]RowGroupMetaData) {
	return t.rowGroupmetaDataSli
}

func (t *TsDeviceMetaData) GetStartTime () (int64) {
	return t.startTime
}

func (t *TsDeviceMetaData) SetStartTime (time int64) () {
	t.startTime = time
}

func (t *TsDeviceMetaData) GetEndTime () (int64) {
	return t.endTime
}

func (t *TsDeviceMetaData) SetEndTime (time int64) () {
	t.endTime = time
}

func (t *TsDeviceMetaData) SerializeTo (buf *bytes.Buffer) (int) {
	if t.sizeOfRowGroupMetaDataSli != len(t.rowGroupmetaDataSli) {
		t.ReCalculateSerializedSize()
	}
	for _, v := range t.rowGroupmetaDataSli {
		log.Info("yaohp: %p", t.rowGroupmetaDataSli[0])
		log.Info("yaohp1: %p", v)
	}
	var byteLen int

	n1, _ := buf.Write(utils.Int64ToByte(t.startTime))
	byteLen += n1
	n2, _ := buf.Write(utils.Int64ToByte(t.endTime))
	byteLen += n2

	if len(t.rowGroupmetaDataSli) == 0 {
		n3, _ := buf.Write(utils.Int32ToByte(0))
		byteLen += n3
	} else {
		n4, _ := buf.Write(utils.Int32ToByte(int32(len(t.rowGroupmetaDataSli))))
		byteLen += n4
		for _, v := range t.rowGroupmetaDataSli {
			// serialize RowGroupMetaData
			byteLen += v.SerializeTo(buf)
		}
	}
	return byteLen
}

func (t *TsDeviceMetaData) ReCalculateSerializedSize () () {
	t.serializedSize = 2 * 8 + 1 * 4
	if t.rowGroupmetaDataSli != nil {
		for _, v := range t.rowGroupmetaDataSli {
			t.serializedSize += v.GetserializedSize()
		}
		t.sizeOfRowGroupMetaDataSli = len(t.rowGroupmetaDataSli)
	}
	t.sizeOfRowGroupMetaDataSli = 0
}

func NewTimeDeviceMetaData() (*TsDeviceMetaData, error) {

	return &TsDeviceMetaData{
		rowGroupmetaDataSli:make([]RowGroupMetaData, 0),
		serializedSize:2 * 8 + 1 * 4,
	},nil
}