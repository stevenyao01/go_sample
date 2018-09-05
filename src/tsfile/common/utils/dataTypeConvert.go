package utils

import (
	"github.com/go_sample/src/tsfile/common/log"
	"math"
	"encoding/binary"
	"bytes"
)

/**
 * @Package Name: utils
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-3 下午5:04
 * @Description:
 */

 // bool
func BoolToByte(flag bool) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, flag)
	if err != nil {
		log.Error("BoolToByte error : %s", err)
		return nil
	}
	return buffer.Bytes()
}


// int
func Int64ToByte(num int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Error("Int64ToByte error : %s", err)
		return nil
	}
	return buffer.Bytes()
}

func Int32ToByte(num int32) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Error("Int32ToByte error : %s", err)
		return nil
	}
	return buffer.Bytes()
}


// float
func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

