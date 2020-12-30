package core

// ConstantPoolCount constant_pool_count
type ConstantPoolCount struct {
	Bytes
	Count int
}

func (cp *ConstantPoolCount) ByteLen() int {
	return u2
}

func ConstantPoolCountNew() *ConstantPoolCount {
	return &ConstantPoolCount{}
}

func (cpc *ConstantPoolCount) ReadObj(bytes []byte) int {
	cpc.Bytes = bytes[0:u2]
	cpc.Count = int(Byte2U2(bytes))
	return u2
}
