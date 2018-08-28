package fileSchema

/**
 * @Package Name: fileSchema
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-28 下午8:55
 * @Description:
 */

import (
//"github.com/go_sample/src/tsfile/common/log"
	"github.com/go_sample/src/tsfile/write/sensorDescriptor"
)

type FileSchema struct {
	sensorDescriptorMap		map[string]sensorDescriptor.SensorDescriptor
	additionalProperties	map[string]string
}

//func (s *FileSchema) Write(v []byte) ([]byte,error) {
//	return nil,nil
//}
//
//func (s *FileSchema) Close() (bool) {
//	return true
//}


func New() (*FileSchema, error) {
	// todo do measurement init and memory check

	return &FileSchema{
		sensorDescriptorMap:make(map[string]sensorDescriptor.SensorDescriptor),
		additionalProperties:make(map[string]string),
	},nil
}