package utils

import (
	"encoding/binary"
	"math"
)

// always stored in big-endian order, where the high bytes come first
func U4(b []byte) int32 {
	bits := binary.BigEndian.Uint32(b)
	n := int32(bits)
	return n
}

func U1(b []byte) int32 {
	b4 := []byte{0, 0, 0, b[0]}
	bits := binary.BigEndian.Uint32(b4)
	n := int32(bits)
	return n
}

// always stored in big-endian order, where the high bytes come first
func U2(b []byte) int32 {
	b4 := []byte{0, 0, b[0], b[1]}
	bits := binary.BigEndian.Uint32(b4)
	n := int32(bits)
	return n
}

// always stored in big-endian order, where the high bytes come first
func Float(b []byte) float32 {
	bits := binary.BigEndian.Uint32(b)
	return math.Float32frombits(bits)
}

// always stored in big-endian order, where the high bytes come first
func Long(b []byte) int64 {
	bits := binary.BigEndian.Uint64(b)
	n := int64(bits)
	return n
}

// always stored in big-endian order, where the high bytes come first
func Double(b []byte) float64 {
	bits := binary.BigEndian.Uint64(b)
	return math.Float64frombits(bits)
}
