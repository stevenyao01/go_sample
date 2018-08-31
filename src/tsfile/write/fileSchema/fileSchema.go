package fileSchema

/**
 * @Package Name: fileSchema
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-28 下午8:55
 * @Description:
 */

import (
	"github.com/go_sample/src/tsfile/write/sensorDescriptor"
)

type FileSchema struct {
	sensorDescriptorMap		map[string]sensorDescriptor.SensorDescriptor
	additionalProperties	map[string]string
}

func (f *FileSchema) GetSensorDescriptiorMap() (map[string]sensorDescriptor.SensorDescriptor) {
	return f.sensorDescriptorMap
}

//func (f *FileSchema) Close() (bool) {
//	return true
//}


func New() (*FileSchema, error) {
	// todo do measurement init and memory check

	return &FileSchema{
		sensorDescriptorMap:make(map[string]sensorDescriptor.SensorDescriptor),
		additionalProperties:make(map[string]string),
	},nil
}