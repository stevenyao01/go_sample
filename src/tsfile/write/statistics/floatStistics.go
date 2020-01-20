package statistics

/**
 * @Package Name: statistics
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-10 下午3:56
 * @Description:
 */

import (
)

type FloatStatistics struct {
	max		int64
	min 	int64
	first 	int64
	sum 	int64
	last 	int64
	isEmpty bool
}

func (i *IntStatistics)GetFloatHeaderSize()(int){
	return 8 * 5
}

func NewFloat() (*Statistics, error) {

	return &Statistics{
		isEmpty:true,
		max:0,
		min:0,
		first:0,
		//double:0,
		sum:0,
		last:0,
	},nil
}