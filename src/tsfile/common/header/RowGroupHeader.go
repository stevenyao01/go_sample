package header

import (
	"github.com/go_sample/src/tsfile/common/log"
)

type RpwGroupHeader struct {
	deltaObjectId		string
	dataSize			uint64
	numOfChunks			int
	serializedSize		int
}


//type Worker struct {
//
//}

func (r *RpwGroupHeader) OnProcess(v []byte) ([]byte) {
 log.Info("filebeat OnProcess get string:%s",v)
 return v
}

func (r *RpwGroupHeader) OnIn() ([]byte,error) {
 log.Info("filebeat OnIn string:%s","test")
 return nil,nil
}


func (r *RpwGroupHeader) OnOut(v []byte) (bool) {
 log.Info("filebeat OnOut get string:%s",v)
 return true
}


func New() (*RpwGroupHeader, error) {
 return &RpwGroupHeader{},nil
}