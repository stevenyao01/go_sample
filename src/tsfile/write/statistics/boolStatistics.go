package statistics

/**
 * @Package Name: statistics
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-10 下午3:49
 * @Description:
 */

import (
)

type BoolStatistics struct {
	max		bool
	min 	bool
	first 	bool
	sum 	int64
	last 	bool
	isEmpty	bool
}

func NewBool() (*Statistics, error) {

	return &Statistics{
		isEmpty:true,
	},nil
}