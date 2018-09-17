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
)

type RowGroupMetaData struct {
	tsIoFile 						*os.File
}

func (t *RowGroupMetaData) WriteMagic()(int){
	n, err := t.tsIoFile.Write([]byte(tsFileConf.MAGIC_STRING))
	if err == nil {
		log.Error("write start magic to file err: ", err)
	}
	return n
}

func StartFlushRowgroup()(){

}

func NewRowGroupMetaData(file string) (*RowGroupMetaData, error) {
	// todo

	return &RowGroupMetaData{
	},nil
}