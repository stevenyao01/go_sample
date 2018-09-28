package compress

import (
	"github.com/golang/snappy"
)

type SnappyEncompressor struct{}

func (s *SnappyEncompressor) GetEncompressedLength(srcLen int) (int) {
	return snappy.MaxEncodedLen(srcLen)
}

func (s *SnappyEncompressor) Encompress (dst []byte, src []byte) ([]byte) {
	return snappy.Encode(dst, src)
}


