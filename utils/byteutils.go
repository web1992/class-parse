package utils

func U4(b []byte) int32 {

	return 0
}

func U1(b []byte) int32 {
	return int32(b[0] << 0)
}

func U2(b []byte) int32 {
	return int32((b[0] << 8) + (b[1] << 0))
}
