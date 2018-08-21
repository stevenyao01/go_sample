package header


type PageHeader struct {
	uncompressedSize		int
	compressedSize			int
	numOfValues				int
	max_timestamp			uint64
	min_timestamp			uint64
	// todo it need read tsfile code
	//	statistics
	serializedSize		int
}