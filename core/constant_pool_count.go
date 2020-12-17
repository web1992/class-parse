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
	n := U2(b)
	return int(n)
}

func ConstantPoolCountNew() *ConstantPoolCount {
	return &ConstantPoolCount{}
}

func (cpc *ConstantPoolCount) ReadObj(bytes []byte) int {
	cpc.Bytes = bytes[0:u2]
	return 0
}

func (cpc *ConstantPoolCount) ObjLen() int {
	return u2
}
