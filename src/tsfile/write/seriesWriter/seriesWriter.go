package seriesWriter

/**
 * @Package Name: seriesWriter
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-28 下午8:28
 * @Description:
 */

import (
//"github.com/go_sample/src/tsfile/common/log"
)

type SeriesWriter struct {
	deviceId			string
	dataSeriesWriters	map[string]SeriesWriter
}

//func (s *SeriesWriter) Write(v []byte) ([]byte,error) {
//	return nil,nil
//}
//
//func (s *SeriesWriter) Close() (bool) {
//	return true
//}


func New(dId string) (*SeriesWriter, error) {
	// todo do measurement init and memory check

	return &SeriesWriter{
		deviceId:dId,
	},nil
}