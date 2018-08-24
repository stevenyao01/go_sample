package utils

import "time"

/**
 * @Package Name: utils
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-24 下午8:37
 * @Description:
 */

 import (
 	"github.com/go_sample/src/tsfile/common/log"
 )

func calculateTime(){
	var d time.Duration
	t0 := time.Now()
	log.Info("my log.")
	t1 := time.Now()
	d = t1.Sub(t0)
	log.Info("cost time = %v\n", d)
}