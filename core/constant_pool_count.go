package core

// ConstantPoolCount constant_pool_count
type ConstantPoolCount struct {
	Count int
}

func (cp *ConstantPoolCount) ByteLen() int {
	return u2
}

func ConstantPoolCountNew() *ConstantPoolCount {
	return &ConstantPoolCount{}
}

func (cpc *ConstantPoolCount) ReadObj(bytes []byte) int {
	cpc.Count = int(Byte2U2(bytes[0:u2]))
	return u2
}
