package header

import (
	"bytes"
	"github.com/go_sample/src/tsfile/common/utils"
)

const (
	MQTT_CONF = "mqtt.conf"
)

// TsDataType
const (
	BOOLEAN = iota  		// 0
	INT32					// 1
	INT64					// 2
	FLOAT					// 3
	DOUBLE					// 4
	TEXT					// 5
	FIXED_LEN_BYTE_ARRAY	// 6
	ENUMS					// 7
	BIGDECIMAL				// 8
)

// CompressionType
const (
	UNCOMPRESSED = iota		// 0
	SNAPPY					// 1
	GZIP					// 2
	LZO						// 3
	SDT						// 4
	PAA						// 5
	PLA						// 6
)

// TsEncoding
const (
	PLAIN = iota			// 0
	PLAIN_DICTIONARY		// 1
	RLE						// 2
	DIFF					// 3
	TS_2DIFF				// 4
	BITMAP					// 5
	GORILLA					// 6
)



type ChunkHeader struct {
	sensorId			string
	dataSize			int
	tsDataType			int16
	compressionType		int16
	encodingType		int16
	numOfPages			int
	serializedSize		int
	maxTombstoneTime 	int64
}

func (c *ChunkHeader) ChunkHeaderToMemory(buffer *bytes.Buffer)(int32){
	// todo write chunk header to buffer
	buffer.Write(utils.Int32ToByte(int32(len(c.sensorId))))
	buffer.Write([]byte(c.sensorId))
	buffer.Write(utils.Int32ToByte(int32(c.dataSize)))
	buffer.Write(utils.Int16ToByte(c.tsDataType))
	buffer.Write(utils.Int32ToByte(int32(c.numOfPages)))
	buffer.Write(utils.Int16ToByte(c.compressionType))
	buffer.Write(utils.Int16ToByte(c.encodingType))
	buffer.Write(utils.Int64ToByte(c.maxTombstoneTime))
	return int32(c.serializedSize)
}

func (c *ChunkHeader) GetSerializedSize () (int) {
	return c.serializedSize
}

func (c *ChunkHeader) GetMaxTombstoneTime () (mtt int64) {
	return c.maxTombstoneTime
}

func (c *ChunkHeader) SetMaxTombstoneTime (mtt int64) () {
	c.maxTombstoneTime = mtt
}


func NewChunkHeader(sId string, pbs int, tdt int16, ct int16, et int16, nop int, mtt int64) (*ChunkHeader, error) {
	// todo
	ss := 3 * 4 + 3 * 2 + len(sId) + 8
	return &ChunkHeader{
		sensorId:sId,
		dataSize:pbs,
		tsDataType:tdt,
		compressionType:ct,
		encodingType:et,
		numOfPages:nop,
		serializedSize:ss,
		maxTombstoneTime:mtt,
	},nil
}