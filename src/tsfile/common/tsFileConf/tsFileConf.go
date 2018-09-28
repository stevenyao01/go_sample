package tsFileConf


const (
	MQTT_CONF = "mqtt.conf"
)



const (
	// Memory size threshold for flushing to disk or HDFS, default value is 128MB
	GroupSizeInByte = 128 * 1024 * 1024

	// The memory size for each series writer to pack page, default value is 64KB
	PageSizeInByte = 64 * 1024

	// The maximum number of data points in a page, defalut value is 1024 * 1024
	MaxNumberOfPointsInPage = 1024 * 1024

	// Data type for input timestamp, TsFile supports INT32 or INT64
	timeSeriesDataType = "INT64"

	// Max length limitation of input string
	maxStringLength = 128

	//Floating-point precision
	floatPrecision = 2

	//Encoder of time series, TsFile supports TS_2DIFF, PLAIN and RLE(run-length encoding) Default value is TS_2DIFF
	timeSeriesEncoder = "TS_2DIFF"

	// Encoder of value series. default value is PLAIN.
	// For int, long data type, TsFile also supports TS_2DIFF and RLE(run-length encoding).
 	// For float, double data type, TsFile also supports TS_2DIFF, RLE(run-length encoding) and GORILLA.
 	// For text data type, TsFile only supports PLAIN.
	alueEncoder = "PLAIN"

	// Default bit width of RLE encoding is 8
	rleBitWidth = 8
	RLE_MIN_REPEATED_NUM = 8
	RLE_MAX_REPEATED_NUM = 0x7FFF
	RLE_MAX_BIT_PACKED_NUM = 63

	// Gorilla encoding configuration
	FLOAT_LENGTH = 32
	FLAOT_LEADING_ZERO_LENGTH = 5
	FLOAT_VALUE_LENGTH = 6
	DOUBLE_LENGTH = 64
	DOUBLE_LEADING_ZERO_LENGTH = 6
	DOUBLE_VALUE_LENGTH = 7

	// Default block size of two-diff. delta encoding is 128
	deltaBlockSize = 128

	// Bitmap configuration
	BITMAP_BITWIDTH = 1

	// Default frequency type is SINGLE_FREQ
	freqType = "SINGLE_FREQ"

	// Default PLA max error is 100
	plaMaxError uint64 = 100

	// Default SDT max error is 100
	sdtMaxError uint64 = 100

	// Default DFT satisfy rate is 0.1
	//dftSatisfyRate uint64 = 0.1


	// Data compression method, TsFile supports UNCOMPRESSED or SNAPPY.
	// Default value is UNCOMPRESSED which means no compression
	compressor = "UNCOMPRESSED"

	// Line count threshold for checking page memory occupied size
	pageCheckSizeThreshold = 100

	// Current version is 3
	CurrentVersion = 3

	// Default endian value is LITTLE_ENDIAN
	endian = "LITTLE_ENDIAN"


	// String encoder with UTF-8 encodes a character to at most 4 bytes.
	BYTE_SIZE_PER_CHAR = 4

	STRING_ENCODING = "UTF-8"

	CONFIG_FILE_NAME = "tsfile-format.properties"

	// The default grow size of class DynamicOneColumnData
	dynamicDataSize = 1000

	MAGIC_STRING = "TsFilev0.8.0"
)

func GetMagic() string{
	return MAGIC_STRING
}
