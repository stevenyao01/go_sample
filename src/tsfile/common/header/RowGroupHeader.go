package header


type RpwGroupHeader struct {
	deltaObjectId		string
	dataSize			uint64
	numOfChunks			int
	serializedSize		int
}
