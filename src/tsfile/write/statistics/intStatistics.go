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

type IntStatistics struct {
	max		int
	min 	int
	first 	int
	double 	int
	sum 	int64
	last 	int
	isEmpty	bool
}

func (i *IntStatistics)GetIntHeaderSize()(int){
	return 5 * 4 + 8 * 1
}

func NewInt() (*Statistics, error) {

	return &Statistics{
		isEmpty:true,
	},nil
}