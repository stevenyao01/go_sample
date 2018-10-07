package metaData

/**
 * @Package Name: metaData
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-18 下午4:28
 * @Description:
 */

import (
	"bytes"
	"github.com/go_sample/src/tsfile/common/utils"
)

type TimeSeriesMetaData struct {
	sensorId							string
	tsDataType							int16
}

//func (t *TimeSeriesMetaData) WriteMagic()(int){
//
//	return 0
//}

func (t *TimeSeriesMetaData) Serialize (buf *bytes.Buffer) (int) {
	var byteLen int
	if t.sensorId == "" {
		n1, _ := buf.Write(utils.BoolToByte(false))
		byteLen += n1
	} else {
		n2, _ := buf.Write(utils.BoolToByte(true))
		byteLen += n2

		n3, _ := buf.Write(utils.Int32ToByte(int32(len(t.sensorId))))
		byteLen += n3
		n4, _ := buf.Write([]byte(t.sensorId))
		byteLen += n4
	}

	if t.tsDataType >= 0 && t.tsDataType <= 9 { // not empty
		n5, _ := buf.Write(utils.BoolToByte(true))
		byteLen += n5

		n6, _ := buf.Write(utils.Int16ToByte(t.tsDataType))
		byteLen += n6
	} else {
		n7, _ := buf.Write(utils.BoolToByte(false))
		byteLen += n7
	}

	return byteLen
}

func NewTimeSeriesMetaData(sid string, tdt int16) (*TimeSeriesMetaData, error) {

	return &TimeSeriesMetaData{
		sensorId:sid,
		tsDataType:tdt,
	},nil
}