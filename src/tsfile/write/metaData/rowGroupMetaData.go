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
	"os"
	"github.com/go_sample/src/tsfile/common/tsFileConf"
	"github.com/go_sample/src/tsfile/write/dataPoint"
)

type RowGroupMetaData struct {
	tsIoFile 						*os.File
	deviceId						string
	totalByteSize					int64
	fileOffsetOfCorrespondingData	int64
	TimeSeriesChunkMetaDataMap		map[string]TimeSeriesChunkMetaData
}

func (t *RowGroupMetaData) WriteMagic()(int){
	n, err := t.tsIoFile.Write([]byte(tsFileConf.MAGIC_STRING))
	if err == nil {
		log.Error("write start magic to file err: ", err)
	}
	return n
}

func NewRowGroupMetaData(dId string, tbs int64, foocd int64, tscmdm map[string]TimeSeriesChunkMetaData) (*RowGroupMetaData, error) {
	// todo

	return &RowGroupMetaData{
		deviceId:dId,
		totalByteSize:tbs,
		fileOffsetOfCorrespondingData:foocd,
		TimeSeriesChunkMetaDataMap:tscmdm,
	},nil
}