package core

// ConstantPoolCount constant_pool_count
type ConstantPoolCount struct {
	Bytes
}

func (cp *ConstantPoolCount) ByteLen() int {
	return U2_L
}

// always stored in big-endian order, where the high bytes come first
func (cp *ConstantPoolCount) View() interface{} {
	// 0xFF = 1111 1111
	b := cp.Bytes
	n := ((b[0] & 0xFF) << 8) + ((b[1]) << 0)
	return int(n)
}
