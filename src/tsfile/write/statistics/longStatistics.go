package statistics

/**
 * @Package Name: statistics
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-10 下午3:55
 * @Description:
 */

import (
)

type LongStatistics struct {
	max		int64
	min 	int64
	first 	int64
	sum 	int64
	last 	int64
	isEmpty	bool
}

func (i *IntStatistics)GetLongHeaderSize()(int){
	return 5 * 8
}

func NewLong() (*Statistics, error) {

	return &Statistics{
		isEmpty:true,
	},nil
}