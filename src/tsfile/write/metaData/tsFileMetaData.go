package metaData

/**
 * @Package Name: metaData
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-20 下午3:07
 * @Description:
 */

import (
	"bytes"
	"github.com/go_sample/src/tsfile/common/utils"
	"github.com/go_sample/src/tsfile/common/log"
)

type TsFileMetaData struct {
	deviceMap							map[string]TsDeviceMetaData
	timeSeriesMetaDataMap				map[string]TimeSeriesMetaData
	currentVersion						int
	createBy							string
	firstTimeSeriesMetadataOffset		int64
	lastTimeSeriesMetadataOffset		int64
	firstTsDeltaObjectMetadataOffset	int64
	lastTsDeltaObjectMetadataOffset		int64
}

//func (t *TsFileMetaData) GetRowGroups () ([]RowGroupMetaData) {
//	return t.rowGroupmetaDataSli
//}

func (t *TsFileMetaData) SerializeTo (buf *bytes.Buffer) (int) {
	var byteLen int
	if t.deviceMap == nil {
		n, _ := buf.Write(utils.Int32ToByte(0))
		byteLen += n
	} else {
		d1, _ := buf.Write(utils.Int32ToByte(int32(len(t.deviceMap))))
		byteLen += d1
		for k, v := range t.deviceMap {
			// write string tsDeviceMetaData key
			d2, _ := buf.Write(utils.Int32ToByte(int32(len(k))))
			byteLen += d2
			d3, _ := buf.Write([]byte(k))
			byteLen += d3
			// tsDeviceMetaData SerializeTo
			byteLen += v.SerializeTo(buf)
			log.Info("v: %s", v)
		}
	}
	if t.timeSeriesMetaDataMap == nil {
		e1, _ := buf.Write(utils.Int32ToByte(0))
		byteLen += e1
	} else {
		e2, _ := buf.Write(utils.Int32ToByte(int32(len(t.timeSeriesMetaDataMap))))
		byteLen += e2
		for _, vv := range t.timeSeriesMetaDataMap {
			// timeSeriesMetaData SerializeTo
			byteLen += vv.Serialize(*buf)
			log.Info("vv: %s", vv)
		}
	}
	f1, _ := buf.Write(utils.Int32ToByte(int32(t.currentVersion)))
	byteLen += f1
	if t.createBy == "" {
		// write flag for t.createBy
		f2, _ := buf.Write(utils.BoolToByte(true))
		byteLen += f2
	} else {
		// write flag for t.createBy
		f3, _ := buf.Write(utils.BoolToByte(false))
		byteLen += f3
		// write string t.createBy
		f4, _ := buf.Write(utils.Int32ToByte(int32(len(t.createBy))))
		byteLen += f4
		f5, _ := buf.Write([]byte(t.createBy))
		byteLen += f5
	}

	off1, _ := buf.Write(utils.Int64ToByte(t.firstTimeSeriesMetadataOffset))
	byteLen += off1
	off2, _ := buf.Write(utils.Int64ToByte(t.lastTimeSeriesMetadataOffset))
	byteLen += off2
	off3, _ := buf.Write(utils.Int64ToByte(t.firstTsDeltaObjectMetadataOffset))
	byteLen += off3
	off4, _ := buf.Write(utils.Int64ToByte(t.lastTsDeltaObjectMetadataOffset))
	byteLen += off4

	return byteLen
}



func NewTsFileMetaData(tdmd map[string]TsDeviceMetaData, tss map[string]TimeSeriesMetaData, version int) (*TsFileMetaData, error) {

	return &TsFileMetaData{
		deviceMap:tdmd,
		timeSeriesMetaDataMap:tss,
		currentVersion:version,
		createBy:"",
	},nil
}