package utils

import (
	"fmt"
	"math"
)

const (
	Byte = 1 << (iota * 10)
	KiByte
	MiByte
	GiByte
	TiByte
	PiByte
	EiByte
)

// SI Sizes.
const (
	IByte = 1
	KByte = IByte * 1000
	MByte = KByte * 1000
	GByte = MByte * 1000
	TByte = GByte * 1000
	PByte = TByte * 1000
	EByte = PByte * 1000
)

var bytesSizeTable = map[string]uint64{
	"b": Byte,
	"kib": KiByte,
	"kb": KByte,
	"mib": MiByte,
	"mb": MByte,
	"gib": GiByte,
	"gb": GByte,
	"tib": TiByte,
	"tb": TByte,
	"pib": PiByte,
	"pb": PByte,
	"eib": EiByte,
	"eb": EByte,
	// Without suffix
	"": Byte,
	"ki": KiByte,
	"k": KByte,
	"mi": MiByte,
	"m": MByte,
	"gi": GiByte,
	"g": GByte,
	"ti": TiByte,
	"t": TByte,
	"pi": PiByte,
	"p": PByte,
	"ei": EiByte,
	"e": EByte,
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := math.Floor(float64(s)/math.Pow(base, e)*10+0.5) / 10
	f := "%.0f %s"
	if val < 10 {
		f = "%.1f %s"
	}

	return fmt.Sprintf(f, val, suffix)
}

// Bytes produces a human readable representation of an SI size.
//
// See also: ParseBytes.
//
// Bytes(82854982) -> 83 MB
func Bytes(s uint64) string {
	sizes := []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(s, 1000, sizes)
}

// IBytes produces a human readable representation of an IEC size.
//
// See also: ParseBytes.
//
// IBytes(82854982) -> 79 MiB
func IBytes(s uint64) string {
	sizes := []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
	return humanateBytes(s, 1024, sizes)
}
