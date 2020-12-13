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
