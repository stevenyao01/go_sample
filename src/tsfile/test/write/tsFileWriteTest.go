package main

import (
	"github.com/go_sample/src/tsfile/common/log"
	"github.com/go_sample/src/tsfile/write/tsFileWriter"
	"github.com/go_sample/src/tsfile/common/header"
	"github.com/go_sample/src/tsfile/write/sensorDescriptor"
	"github.com/go_sample/src/tsfile/write/dataPoint"
	"github.com/go_sample/src/tsfile/write/tsRecord"
	"time"
)

const (
	fileName = "test.ts"
)


func main(){

	// init tsFileWriter
	tfWriter, tfwErr := tsFileWriter.NewIoWriter(fileName)
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

	// create a tsRecord
	ts := time.Now()
	//timestamp := strconv.FormatInt(ts.UTC().UnixNano(), 10)
	//fmt.Println(timestamp)
	tr, trErr := tsRecord.New(ts, "suer1.thinkpad.T200")
	if trErr != nil {
		log.Info("init tsRecord error.")
	}

	// create two data points
	idp, iDpErr := dataPoint.NewInt("cpu_utility", header.INT32,20)
	if iDpErr != nil {
		log.Info("init int data point error.")
	}
	fdp, fDpErr := dataPoint.NewFloat("cpu_utility", header.FLOAT,90.0)
	if fDpErr != nil {
		log.Info("init float data point error.")
	}

	// add data points to ts record
	tr.AddTuple(*idp)
	tr.AddTuple(*fdp)



	// todo write tsRecord to file
	tfWriter.Write(*tr)
	//tfWriter.Write([]byte("&TsFileData&"))

	// close file descriptor
	tfWriter.Close()
}
