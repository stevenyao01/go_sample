package metaData

/**
 * @Package Name: metaData
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-18 上午10:36
 * @Description:
 */

import (
	"bytes"
)

type TsDigest struct {
	statistics 			map[string]bytes.Buffer
}

func (t *TsDigest) SetStatistics (statistics map[string]bytes.Buffer) () {
	t.statistics = statistics
	// todo recalculate serialized size
	//reCalculateSerializedSize()
}

func reCalculateSerializedSize () () {
	//todo calculate size again
}

func NewTsDigest() (*TsDigest, error) {
	return &TsDigest{
		statistics:make(map[string]bytes.Buffer),
	}, nil
}