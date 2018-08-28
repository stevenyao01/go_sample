package rowGroupWriter

/**
 * @Package Name: rouGroupWriter
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-28 下午6:19
 * @Description:
 */

import (
//"github.com/go_sample/src/tsfile/common/log"
	"github.com/go_sample/src/tsfile/write/seriesWriter"
)

type RowGroupWriter struct {
	deviceId			string
	dataSeriesWriters	map[string]seriesWriter.SeriesWriter
}

//func (s *RouGroupWriter) Write(v []byte) ([]byte,error) {
//	return nil,nil
//}
//
//func (s *RouGroupWriter) Close() (bool) {
//	return true
//}


func New(dId string) (*RowGroupWriter, error) {
	// todo do measurement init and memory check

	return &RowGroupWriter{
		deviceId:dId,
	},nil
}