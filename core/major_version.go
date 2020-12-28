package core

// MajorVersion major_version
type MajorVersion struct {
	Bytes
	Version int32
}

func (mv *MajorVersion) ByteLen() int {
	return U2_L
}

func MajorVersionNew() *MajorVersion {
	return &MajorVersion{}
}

func (mv *MajorVersion) ReadObj(bytes []byte) int {
	mv.Bytes = bytes[0:u2]
	mv.Version = Byte2U2(bytes)
	return u2
}
