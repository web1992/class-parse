package core

import (
	"encoding/binary"
	"fmt"
	"math"
	"strings"
)

func HexByte(b []byte) Hex {
	s := fmt.Sprintf("%x", b)
	return Hex(strings.ToUpper(s))
}

// U4 4 bytes to int32
// always stored in big-endian order, where the high bytes come first
func U4(b []byte) int32 {
	bits := binary.BigEndian.Uint32(b)
	n := int32(bits)
	return n
}

// U1 1 bytes to int32
func U1(b []byte) int32 {
	b4 := []byte{0, 0, 0, b[0]}
	bits := binary.BigEndian.Uint32(b4)
	n := int32(bits)
	return n
}

// U2 2 bytes to int32
// always stored in big-endian order, where the high bytes come first
func U2(b []byte) int32 {
	b4 := []byte{0, 0, b[0], b[1]}
	bits := binary.BigEndian.Uint32(b4)
	n := int32(bits)
	return n
}

// always stored in big-endian order, where the high bytes come first
func Byte2Float(b []byte) float32 {
	bits := binary.BigEndian.Uint32(b)
	return math.Float32frombits(bits)
}

// always stored in big-endian order, where the high bytes come first
func Byte2Long(b []byte) int64 {
	bits := binary.BigEndian.Uint64(b)
	n := int64(bits)
	return n
}

// always stored in big-endian order, where the high bytes come first
func Byte2Double(b []byte) float64 {
	bits := binary.BigEndian.Uint64(b)
	return math.Float64frombits(bits)
}
