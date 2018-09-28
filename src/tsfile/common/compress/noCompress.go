package compress

/**
 * @Package Name: compress
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-27 下午6:29
 * @Description:
 */

type NoEncompressor struct{}

func (n *NoEncompressor) GetEncompressedLength(srcLen int) (int) {
	return srcLen
}

func (n *NoEncompressor) Encompress(dst []byte, src []byte) ([]byte) {
	return src
}
