package utils

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
	return 0
}

// always stored in big-endian order, where the high bytes come first
func Long(b []byte) int64 {
	return 0
}

// always stored in big-endian order, where the high bytes come first
func Double(b []byte) float64 {
	return 0
}
