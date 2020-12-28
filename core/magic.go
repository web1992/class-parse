package core

// Magic class file Magic num
type Magic struct {
	Bytes
	Hex
}

func (m *Magic) ByteLen() int {
	return U4_L
}

func MagicNew() *Magic {
	return &Magic{}
}

func (magic *Magic) ReadObj(bytes []byte) int {
	magic.Bytes = bytes[0:u4]
	magic.Hex = HexByte(magic.Bytes)
	return u4
}
