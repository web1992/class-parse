package core

// MajorVersion major_version
type MajorVersion struct {
	Bytes
}

func (mv *MajorVersion) ByteLen() int {
	return U2_L
}

func (m *MajorVersion) View() interface{} {
	// 0xFF = 1111 1111
	b := m.Bytes
	//n := ((b[0] & 0xFF) << 8) + ((b[1]) << 0)
	n := (b[0] << 8) + ((b[1]) << 0)
	return int(n)
}

func MajorVersionNew() *MajorVersion {
	return &MajorVersion{}
}

func (mv *MajorVersion) ReadObj(bytes []byte) int {
	mv.Bytes = bytes[0:u2]
	return 0
}

func (mv *MajorVersion) ObjLen() int {
	return u2
}
