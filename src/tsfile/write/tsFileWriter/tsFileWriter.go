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
)

const (
	fileName = "test.ts"
)

type TsFileWriter struct {
	tsFile 	*os.File
}

func (t *TsFileWriter) AddMeasurement(md sensorDescriptor.SensorDescriptor) ([]byte) {
 	log.Info("enter tsFileWriter->AddMeasurement()")
 	return nil
}

func (t *TsFileWriter) Write(v []byte) ([]byte,error) {
	// todo write data here

	log.Info("enter tsFileWriter->Write()")
	t.tsFile.Write(v)
	return nil,nil
}

func (t *TsFileWriter) Close() (bool) {
	// finished write file, and write magic string at file tail
	WriteMagic(t.tsFile)
	t.tsFile.Write([]byte("\n"))
	t.tsFile.Close()
	return true
}


func NewIoWriter() (*TsFileWriter, error) {
	newFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Info("open file:%s failed.", fileName)
	}
	WriteMagic(newFile)
 return &TsFileWriter{tsFile:newFile},nil
}