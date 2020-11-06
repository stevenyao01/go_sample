package main

import (
	"github.com/go_sample/src/mlog/log"
)

/**
 * @Project: go_sample
 * @Package Name: mlog
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/7/2 下午3:16
 * @Description:
 */

func main() () {
	log.Println("test Println.")
	log.Info("test Info.")
	log.Debug("test Debug.")
	log.Error("test Error.")

	//logger, _ := logger.LoggerNew()
	//
	//logger.Info("info")
	//logger.Debug("debug")
	//logger.Error("error")
}
