package main

import "github.com/golang/snappy"

func main() {
	// deCompress
	traceBuf, _ := snappy.Decode(nil, buf)

	//compress
	snappy.Encode(nil, b)
}
