package main

import (
	"github.com/go_sample/src/tsfile/common/log"
	"github.com/go_sample/src/tsfile/write/tsFileWriter"
	"github.com/go_sample/src/tsfile/common/header"
	"github.com/go_sample/src/tsfile/write/sensorDescriptor"
)


func main(){

	// init tsFileWriter
	tfWriter, tfwErr := tsFileWriter.NewIoWriter()
	if tfwErr != nil {
		log.Info("init tsFileWriter error = %s", tfwErr)
	}

	// init sensorDescriptor
	sd, sdErr := sensorDescriptor.New("cpu_utility", header.FLOAT, header.TS_2DIFF)
	if sdErr != nil {
		log.Info("init sensorDescriptor error = %s", sdErr)
	}

	// add sensorDescriptor to tfFileWriter
	tfWriter.AddSensor(*sd)

	// todo create a tsRecord

	// todo write tsRecord to file
	tfWriter.Write([]byte("&TsFileData&"))

	// close file descriptor
	tfWriter.Close()
}
