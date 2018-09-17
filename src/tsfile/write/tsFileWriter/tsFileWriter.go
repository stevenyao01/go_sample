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
	tsFileIoWriter				*TsFileIoWriter
	schema 						*fileSchema.FileSchema
	recordCount					int64
	recordCountForNextMemCheck 	int64
	rowGroupSizeThreshold		int64
	primaryRowGroupSize			int64
	pageSize					int64
	oneRowMaxSize				int64
	groupDevices				map[string]rowGroupWriter.RowGroupWriter
}

//var groupDevices = make(map[string]rowGroupWriter.RowGroupWriter)

func (t *TsFileWriter) AddSensor(sd sensorDescriptor.SensorDescriptor) ([]byte) {
 	log.Info("enter tsFileWriter->AddSensor()")
 	if _, ok := t.schema.GetSensorDescriptiorMap()[sd.GetSensorId()]; !ok {
		t.schema.GetSensorDescriptiorMap()[sd.GetSensorId()] = sd
	}else{
		log.Info("the given sensor has exist!")
	}
	// todo flush rowgroup
	t.checkMemorySizeAndMayFlushGroup()
 	return nil
}

func (t *TsFileWriter)checkMemorySizeAndMayFlushGroup()(bool){
	if t.recordCount >= t.recordCountForNextMemCheck {
		// calculate all group size
		memSize := calculateMemSizeForAllGroup()
		if memSize > t.rowGroupSizeThreshold {
			log.Info("start_write_row_group, memory space occupy: %s", memSize)
			t.recordCountForNextMemCheck = t.rowGroupSizeThreshold / t.oneRowMaxSize
			return t.flushAllRowGroups(false)
		} else {
			t.recordCountForNextMemCheck = t.recordCount + (t.rowGroupSizeThreshold - memSize) / t.oneRowMaxSize
			return false
		}
	}
	return false
}

/**
   * flush the data in all series writers and their page writers to outputStream.
   * @param isFillRowGroup whether to fill RowGroup
   * @return true - size of tsfile or metadata reaches the threshold.
   * 		 false - otherwise. But this function just return false, the Override of IoTDB may return true.
   */
func (t *TsFileWriter)flushAllRowGroups(isFillRowGroup bool)(bool){
	// todo flush data to disk
	if t.recordCount > 0 {
		for _, v := range t.groupDevices {
			v.PreFlush()
		}
		for k, _ := range t.groupDevices {
			groupDevice := t.groupDevices[k]
			rowGroupSize := groupDevice.GetCurrentRowGroupSize()
			t.tsFileIoWriter.StartFlushRowGroup(k, rowGroupSize, groupDevice.GetSeriesNumber())
			groupDevice.FlushToFileWriter(t.tsFileIoWriter)
		}
	}
	return true
}

func (t *TsFileWriter) Write(tr tsRecord.TsRecord) (bool) {
	log.Info("enter tsFileWriter->Write()")
	// todo write data here
	if t.checkIsDeviceExist(tr, *t.schema) {
		t.groupDevices[tr.GetDeviceId()].Write(tr.GetTime(), tr.GetDataPointMap())
		t.recordCount = t.recordCount + 1
		return t.checkMemorySize()
	}

	///////////////////////////////////////////////
	//t.tsFile.Write(v)
	return false
}


func (t *TsFileWriter) Close() (bool) {
	// finished write file, and write magic string at file tail
	t.tsFileIoWriter.WriteMagic()
	t.tsFileIoWriter.tsIoFile.Write([]byte("\n"))
	t.tsFileIoWriter.tsIoFile.Close()

	return true
}

func (t *TsFileWriter)checkMemorySize() (bool) {
	if t.recordCount >= t.recordCountForNextMemCheck {
		memSize := calculateMemSizeForAllGroup()
		if memSize > t.rowGroupSizeThreshold {
			log.Info("start write rowGroup, memory space occupy:%s", memSize)
			t.recordCountForNextMemCheck = t.rowGroupSizeThreshold / t.oneRowMaxSize
			return t.flushAllRowGroups(false)
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

func (t *TsFileWriter) checkIsDeviceExist(tr tsRecord.TsRecord, schema fileSchema.FileSchema) bool {
	var groupDevice *rowGroupWriter.RowGroupWriter
	var err error
	// check device
	if _, ok := t.groupDevices[tr.GetDeviceId()]; !ok {
		// if not exist
		groupDevice, err = rowGroupWriter.New(tr.GetDeviceId())
		if err != nil {
			log.Info("rowGroupWriter init ok!")
		}
		t.groupDevices[tr.GetDeviceId()] = *groupDevice
	} else { // if exist
		*groupDevice = t.groupDevices[tr.GetDeviceId()]
	}
	schemaSensorDescriptorMap := schema.GetSensorDescriptiorMap()
	for k, v := range tr.GetDataPointMap() {
		if contain, _ := utils.MapContains(schemaSensorDescriptorMap, v.GetSensorId()); contain {
			groupDevice.AddSeriesWriter(schemaSensorDescriptorMap[v.GetSensorId()], tsFileConf.PageSizeInByte)
		} else {
			log.Error("input sensor is invalid: ", v.GetSensorId())
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


func NewTsFileWriter(file string) (*TsFileWriter, error) {
	// file schema
	fs, fsErr := fileSchema.New()
	if fsErr != nil {
		log.Error("init fileSchema failed.")
	}

	// tsFileIoWriter
	tfiWriter, tfiwErr := NewTsFileIoWriter(file)
	if tfiwErr != nil {
		log.Error("init tsFileWriter error = ", tfiwErr)
	}

	// write start magic
	tfiWriter.WriteMagic()

	// init rowGroupSizeThreshold
	var prgs int64 = 0
	rgst := tsFileConf.GroupSizeInByte - prgs

 return &TsFileWriter{
 	tsFileIoWriter:tfiWriter,
 	schema:fs,
 	recordCount:0,
 	recordCountForNextMemCheck:1,
 	primaryRowGroupSize:prgs,
 	rowGroupSizeThreshold:rgst,
 	groupDevices:make(map[string]rowGroupWriter.RowGroupWriter),
 	},nil
}