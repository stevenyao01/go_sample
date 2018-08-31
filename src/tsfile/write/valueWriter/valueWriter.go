package valueWriter

/**
 * @Package Name: valueWriter
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-31 下午4:51
 * @Description:
 */

import (
//"github.com/go_sample/src/tsfile/common/log"
//	"encoding/gob"
//	"github.com/elastic/beats/filebeat/harvester/reader"
	"time"
	"bytes"
)

type ValueWriter struct {
	// time
	timeEncoder		Encoder

	// value
	valueEncoder 	Encoder

	buf 			*bytes.Buffer
	//buf := bytes.NewBuffer([]byte{})
}

func (s *ValueWriter) Write(t time.Time, value interface{}) (string) {

	return s.sensorId
}


//// todo the return type should be Compressor, after finished Compressor we should modify it.
//func (s *ValueWriter) GetCompressor() (string) {
//	return s.compressor
//}
//
//func (s *ValueWriter) Close() (bool) {
//	return true
//}


func New() (*ValueWriter, error) {

	return &ValueWriter{
		//sensorId:sId,
		buf:bytes.NewBuffer([]byte{}),
	},nil
}