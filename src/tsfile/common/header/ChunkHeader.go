package header


const (
	MQTT_CONF = "mqtt.conf"
)

// TsDataType
const (
	BOOLEAN = iota  // 0
	INT32			// 1
	INT64			// 2
	FLOAT			// 3
	DOUBLE			// 4
	TEXT			// 5
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
	measurementId		string
	dataSize			int
	//tsDataType			TsDataType
	tsDataType			int
	//compressionType		CompressionType
	compressionType		int
	//encodingType		TsEncoding
	encodingType		int
	numOfPages			int
	serializedSize		int
}

//func deserialize() {
//
//}
//
//
//func serialize() {
//
//}

//func main() {
//	fmt.Println("test!!!")
//	fmt.Println("float=", FLOAT)
//	fmt.Println("SNAPPY=", SNAPPY)
//	fmt.Println("BITMAP=", BITMAP)
//}