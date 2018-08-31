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
	"github.com/go_sample/src/tsfile/write/fileSchema"
	"github.com/go_sample/src/tsfile/common/utils"
	"fmt"
	"github.com/go_sample/src/tsfile/common/tsFileConf"
)

type TsFileWriter struct {
	tsFile 						*os.File
	schema 						*fileSchema.FileSchema
	recordCount					int64
	recordCountForNextMemCheck 	int64
	rowGroupSizeThreshold		int64
	primaryRowGroupSize			int64
	pageSize					int64
	oneRowMaxSize				int64
}

var groupDevices = make(map[string]rowGroupWriter.RowGroupWriter)

func (t *TsFileWriter) AddSensor(sd sensorDescriptor.SensorDescriptor) ([]byte) {
 	log.Info("enter tsFileWriter->AddSensor()")
 	if _, ok := t.schema.GetSensorDescriptiorMap()[sd.GetSensorId()]; !ok {
		t.schema.GetSensorDescriptiorMap()[sd.GetSensorId()] = sd
	}else{
		log.Info("the given sensor has exist!")
	}
 	return nil
}

func (t *TsFileWriter) Write(tr tsRecord.TsRecord) (bool) {
	log.Info("enter tsFileWriter->Write()")
	// todo write data here
	if(checkIsDeviceExist(tr, *t.schema)) {
		groupDevices[tr.GetDeviceId()].Write(tr.GetTime(), tr.GetDataPointMap())
		t.recordCount = t.recordCount + 1
		return checkMemorySize(t)
	}

	///////////////////////////////////////////////
	//t.tsFile.Write(v)
	return false
}


func (t *TsFileWriter) Close() (bool) {
	// finished write file, and write magic string at file tail
	WriteMagic(t.tsFile)
	t.tsFile.Write([]byte("\n"))
	t.tsFile.Close()
	return true
}

func checkMemorySize(t *TsFileWriter) (bool) {
	if t.recordCount >= t.recordCountForNextMemCheck {
		memSize := calculateMemSizeForAllGroup()
		if memSize > t.rowGroupSizeThreshold {
			log.Info("start write rowGroup, memory space occupy:%s", memSize)
			t.recordCountForNextMemCheck = t.rowGroupSizeThreshold / t.oneRowMaxSize
			return flushRowGroup(false)
		} else {
			t.recordCountForNextMemCheck = t.recordCount + (t.rowGroupSizeThreshold - memSize) / t.oneRowMaxSize
			return false
		}
	}
	return false
}

func calculateMemSizeForAllGroup()(int64){
	// todo calculate all group memory size

	// return max size for write rowGroupHeader
	return 128 * 1024 *1024
}

/**
   * flush the data in all series writers and their page writers to outputStream.
   * @param isFillRowGroup whether to fill RowGroup
   * @return true - size of tsfile or metadata reaches the threshold.
   * 		 false - otherwise. But this function just return false, the Override of IoTDB may return true.
   */
func flushRowGroup(isFillRowGroup bool)(bool){
	// todo flush data to disk

	return true
}

func checkIsDeviceExist(tr tsRecord.TsRecord, schema fileSchema.FileSchema) bool {
	var groupDevice *rowGroupWriter.RowGroupWriter
	var err error
	// check device
	if _, ok := groupDevices[tr.GetDeviceId()]; !ok {
		// if not exist
		groupDevice, err = rowGroupWriter.New(tr.GetDeviceId())
		if err != nil {
			log.Info("rowGroupWriter init ok!")
		}
		groupDevices[tr.GetDeviceId()] = *groupDevice
	} else { // if exist
		*groupDevice = groupDevices[tr.GetDeviceId()]
	}
	schemaSensorDescriptorMap := schema.GetSensorDescriptiorMap()
	for k, v := range tr.GetDataPointMap() {
		if contain, _ := utils.MapContains(schemaSensorDescriptorMap, v.GetSensorId()); contain {
			groupDevice.AddSeriesWriter(schemaSensorDescriptorMap[v.GetSensorId()], tsFileConf.PageSizeInByte)
		} else {
			log.Error("input sensor is invalid: %s", v.GetSensorId())
		}
		fmt.Printf("k=%v, v=%v\n", k, v)
	}

	//var next *list.Element
	//for e := tr.DataPointList.Front(); e != nil; e = next {
	//	//next = e.Next()
	//	//l.Remove(e)
	//	var x dataPoint.DataPoint //x;// = e.Value
	//	x = e.Value;
	//	if utils.MapContains(shemaSensorDescriptorMap, e.Value.) {
	//
	//	}
	//
	//}
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

	// init rowGroupSizeThreshold
	var prgs int64 = 0
	rgst := tsFileConf.GroupSizeInByte - prgs

 return &TsFileWriter{
 	tsFile:newFile,
 	schema:fs,
 	recordCount:0,
 	recordCountForNextMemCheck:1,
 	primaryRowGroupSize:prgs,
 	rowGroupSizeThreshold:rgst,
 	},nil
}