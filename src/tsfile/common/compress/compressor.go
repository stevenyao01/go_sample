package compress

import (
	"github.com/go_sample/src/tsfile/common/log"
	"github.com/golang/snappy"
)

type Encompress struct{}

type Encompressor interface {
	GetEncompressedLength(srcLen int) (int)
	Encompress(dst []byte, src []byte) ([]byte)
}

func (e *Encompress) DeCompress(compressed []byte) ([]byte, error) {
	return snappy.Decode(nil, compressed)
}

func (e *Encompress) GetEncompressor(tsCompressionType int16) Encompressor {
	var encompressor Encompressor
	switch tsCompressionType {
	case 0:
		encompressor = new(NoEncompressor)
	case 1:
		encompressor = new(SnappyEncompressor)
	//case 2:
	//	encompressor = new(NoEncompressor)
	//case 3:
	//	encompressor = new(SnappyEncompressor)
	//case 4:
	//	encompressor = new(NoEncompressor)
	//case 5:
	//	encompressor = new(SnappyEncompressor)
	default:
		encompressor = new(NoEncompressor)
		log.Info("Encompressor not found, use default.")
	}
	return encompressor
}
