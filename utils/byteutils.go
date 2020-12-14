package utils

import (
	"encoding/binary"
	"math"
)

// always stored in big-endian order, where the high bytes come first
func U4(b []byte) int32 {
	return int32((b[0] << 32) + (b[1] << 16) + (b[2] << 8) + (b[3] << 0))
}

func U1(b []byte) int32 {
	return int32(b[0] << 0)
}

// always stored in big-endian order, where the high bytes come first
func U2(b []byte) int32 {
	return int32((b[0] << 8) + (b[1] << 0))
}

// always stored in big-endian order, where the high bytes come first
func Float(b []byte) float32 {
	bits := binary.BigEndian.Uint32(b)
	return math.Float32frombits(bits)
}

// always stored in big-endian order, where the high bytes come first
func Long(b []byte) int64 {
	binary.Varint(b)
	return int64((b[0] << 32) + (b[1] << 16) + (b[2] << 8) + (b[3] << 0) + (b[4] << 0) + (b[5] << 0) + (b[6] << 0) + (b[7] << 0))
}

// always stored in big-endian order, where the high bytes come first
func Double(b []byte) float64 {
	bits := binary.BigEndian.Uint64(b)
	return math.Float64frombits(bits)
}
