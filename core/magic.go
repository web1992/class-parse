package core

// Magic class file Magic num
type Magic struct {
	Hex
}

func (m *Magic) ByteLen() int {
	return U4_L
}

func MagicNew() *Magic {
	return &Magic{}
}

func (m *Magic) ReadObj(bytes []byte) int {
	bs := bytes[0:u4]
	m.Hex = HexByte(bs)
	return u4
}
