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

type StringStatistics struct {
	max		string
	min 	string
	first 	string
	sum 	string
	last 	string
	isEmpty	bool
}

func NewString() (*Statistics, error) {

	return &Statistics{
		isEmpty:true,
	},nil
}