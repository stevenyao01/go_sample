package pageWriter

/**
 * @Package Name: pageWriter
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-31 下午3:51
 * @Description:
 */

import (
//"github.com/go_sample/src/tsfile/common/log"
	"github.com/go_sample/src/tsfile/write/sensorDescriptor"
	"container/list"
)

type PageWriter struct {
	compressor 			string //Compressor
	desc 				sensorDescriptor.SensorDescriptor
	// todo this buf should change to compert type
	buf 				*list.List
	totalValueCount		int64
	maxTimestamp		int64
	minTimestamp		int64
}

func (s *PageWriter) WritePageHeaderAndDataIntoBuff() (int) {
	//
	return 0
}
//
//func (s *PageWriter) Close() (bool) {
//	return true
//}


func New(sd sensorDescriptor.SensorDescriptor) (*PageWriter, error) {
	// todo do measurement init and memory check

	return &PageWriter{
		desc:sd,
		compressor:sd.GetCompressor(),
		buf:list.New(),
	},nil
}