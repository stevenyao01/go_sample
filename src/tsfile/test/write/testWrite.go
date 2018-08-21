package main

import (
	"github.com/go_sample/src/tsfile/write"
	"github.com/go_sample/src/tsfile/common/log"
)

func main() {
	log.Info("in test write..")
	write.Write()
}
