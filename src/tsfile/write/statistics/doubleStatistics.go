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

type DoubleStatistics struct {
	max		int64
	min 	int64
	first 	int64
	sum 	int64
	last 	int64
	isEmpty	bool
}

func (i *DoubleStatistics)GetDoubleHeaderSize()(int){
	return 5 * 8
}

func NewDouble() (*Statistics, error) {

	return &Statistics{
		isEmpty:true,
	},nil
}