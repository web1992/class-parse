package core

// MajorVersion major_version
type MajorVersion struct {
	Bytes
}

func (mv *MajorVersion) ByteLen() int {
	return U2_L
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
