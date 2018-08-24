package main

import (
	"github.com/go_sample/src/tsfile/common/log"
	"github.com/go_sample/src/tsfile/write/tsFileWriter"
	"github.com/go_sample/src/tsfile/common/header"
	"github.com/go_sample/src/tsfile/write/sensorDescriptor"
)


func main(){

	tfWriter, tfwErr := tsFileWriter.NewIoWriter()
	if tfwErr != nil {
		log.Info("init tsFileWriter error = %s", tfwErr)
	}
	md, mdErr := sensorDescriptor.New("cpu_utility", header.FLOAT, header.TS_2DIFF)
	if mdErr != nil {
		log.Info("init sensorDescriptor error = %s", mdErr)
	}
	tfWriter.AddMeasurement(*md)
	tfWriter.Write([]byte("&TsFileData&"))
	tfWriter.Close()
}
