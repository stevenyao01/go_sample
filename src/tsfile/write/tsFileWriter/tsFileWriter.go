package tsFileWriter

/**
 * @Package Name: write
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-24 下午5:41
 * @Description:
 */

import (
	"github.com/go_sample/src/tsfile/common/log"
	"github.com/go_sample/src/tsfile/write/sensorDescriptor"
	"os"
	"github.com/go_sample/src/tsfile/write/tsRecord"
	"github.com/go_sample/src/tsfile/write/rowGroupWriter"
	"container/list"
	"github.com/go_sample/src/tsfile/write/fileSchema"
)

type TsFileWriter struct {
	tsFile 	*os.File
	schema 	*fileSchema.FileSchema
}

var groupDevices = make(map[string]rowGroupWriter.RowGroupWriter)

func (t *TsFileWriter) AddSensor(sd sensorDescriptor.SensorDescriptor) ([]byte) {
 	log.Info("enter tsFileWriter->AddSensor()")
 	return nil
}

func (t *TsFileWriter) Write(tr tsRecord.TsRecord) ([]byte,error) {
	// todo write data here
	if(checkIsDeviceExist(tr)) {

	}


	///////////////////////////////////////////////
	log.Info("enter tsFileWriter->Write()")
	//t.tsFile.Write(v)
	return nil,nil
}


func (t *TsFileWriter) Close() (bool) {
	// finished write file, and write magic string at file tail
	WriteMagic(t.tsFile)
	t.tsFile.Write([]byte("\n"))
	t.tsFile.Close()
	return true
}

func checkIsDeviceExist(tr tsRecord.TsRecord) bool {
	var groupDevice *rowGroupWriter.RowGroupWriter
	var err error
	// check device
	if _, ok := groupDevices[tr.DeviceId]; !ok {
		// if not exist
		groupDevice, err = rowGroupWriter.New(tr.DeviceId)
		if err != nil {
			log.Info("rowGroupWriter init ok!")
		}
		groupDevices[tr.DeviceId] = *groupDevice
	} else { // if exist
		*groupDevice = groupDevices[tr.DeviceId]
	}
	var next *list.Element
	for e := tr.DataPointList.Front(); e != nil; e = next {
		//next = e.Next()
		//l.Remove(e)

	}
	return true
}


func NewIoWriter(file string) (*TsFileWriter, error) {
	// file schema
	fs, fsErr := fileSchema.New()
	if fsErr != nil {
		log.Info("init fileSchema failed.")
	}
	// io writer
	newFile, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Info("open file:%s failed.", file)
	}

	// magci string
	WriteMagic(newFile)

 return &TsFileWriter{
 	tsFile:newFile,
 	schema:fs,
 	},nil
}