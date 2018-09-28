package main

import (
	"github.com/go_sample/src/tsfile/common/log"
	"github.com/go_sample/src/tsfile/write/tsFileWriter"
	"github.com/go_sample/src/tsfile/common/header"
	"github.com/go_sample/src/tsfile/write/sensorDescriptor"
	"time"
	"github.com/go_sample/src/tsfile/common/compress"
)

const (
	fileName = "test.ts"
)


func main(){

	//decompressor := new(compress.SnappyEncompressor)
	//aSlice := make([]byte, 0)
	//enc := decompressor.Encompress(aSlice, []byte("hello moto"))
	//dec, _ := decompressor.Decompress(enc)
	//log.Info("dec: %s", dec)

	enCompress := new(compress.Encompress)
	aSlice := make([]byte, 0)
	enc := enCompress.GetEncompressor(header.SNAPPY).Encompress(aSlice, []byte("hello moto"))
	dec, _ := enCompress.DeCompress(enc)
	log.Info("dec: %s", dec)


	// init tsFileWriter
	tfWriter, tfwErr := tsFileWriter.NewTsFileWriter(fileName)
	if tfwErr != nil {
		log.Info("init tsFileWriter error = %s", tfwErr)
	}

	// init sensorDescriptor
	sd, sdErr := sensorDescriptor.New("sensor_1", header.FLOAT, header.TS_2DIFF)
	if sdErr != nil {
		log.Info("init sensorDescriptor error = %s", sdErr)
	}
	sd2, sdErr2 := sensorDescriptor.New("sensor_2", header.FLOAT, header.TS_2DIFF)
	if sdErr2 != nil {
		log.Info("init sensorDescriptor error = %s", sdErr2)
	}

	// add sensorDescriptor to tfFileWriter
	tfWriter.AddSensor(*sd)
	tfWriter.AddSensor(*sd2)

	// create a tsRecord
	ts := time.Now()
	//timestamp := strconv.FormatInt(ts.UTC().UnixNano(), 10)
	//fmt.Println(timestamp)
	tr, trErr := tsFileWriter.NewTsRecord(ts, "device_1")
	if trErr != nil {
		log.Info("init tsRecord error.")
	}

	// create two data points
	idp, iDpErr := tsFileWriter.NewInt("sensor_1", header.INT32,20)
	if iDpErr != nil {
		log.Info("init int data point error.")
	}
	fdp, fDpErr := tsFileWriter.NewFloat("sensor_1", header.FLOAT,90.0)
	if fDpErr != nil {
		log.Info("init float data point error.")
	}

	// add data points to ts record
	tr.AddTuple(*idp)
	tr.AddTuple(*fdp)

	// todo write tsRecord to file
	tfWriter.Write(*tr)
	//tfWriter.Write([]byte("&TsFileData&"))



	tr1, trErr1 := tsFileWriter.NewTsRecord(ts, "device_1")
	if trErr1 != nil {
		log.Info("init tsRecord error.")
	}

	// create two data points
	idp1, iDpErr1 := tsFileWriter.NewInt("sensor_2", header.INT32,20)
	if iDpErr1 != nil {
		log.Info("init int data point error.")
	}
	fdp1, fDpErr1 := tsFileWriter.NewFloat("sensor_2", header.FLOAT,90.0)
	if fDpErr1 != nil {
		log.Info("init float data point error.")
	}

	// add data points to ts record
	tr1.AddTuple(*idp1)
	tr1.AddTuple(*fdp1)

	// todo write tsRecord to file
	tfWriter.Write(*tr1)
	//tfWriter.Write([]byte("&TsFileData&"))



	// close file descriptor
	tfWriter.Close()
}
