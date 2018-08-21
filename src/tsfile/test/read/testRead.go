package main

import (
	"github.com/go_sample/src/tsfile/read"
	"github.com/go_sample/src/tsfile/common/log"
)

func main() {
	log.Info("in test read..")
	read.Read()
}