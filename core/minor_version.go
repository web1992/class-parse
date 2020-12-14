package core

// MinorVersion  minor_version
type MinorVersion struct {
	Bytes
}

func (mv *MinorVersion) ByteLen() int {
	return U2_L
}

func (m *MinorVersion) View() interface{} {
	// 0xFF = 1111 1111
	b := m.Bytes
	//n := ((b[0] & 0xFF) << 8) + ((b[1]) << 0)
	n := (b[0] << 8) + ((b[1]) << 0)
	return int(n)
}

func MinorVersionNew() *MinorVersion {
	return &MinorVersion{}
}

func (mv *MinorVersion) ReadObj(bytes []byte) int {
	mv.Bytes = bytes[0:u2]
	return 0
}

func (mv *MinorVersion) ObjLen() int {
	return u2
}
