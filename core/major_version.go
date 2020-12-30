package core

// MajorVersion major_version
type MajorVersion struct {
	Bytes
	Version int
}

func MajorVersionNew() *MajorVersion {
	return &MajorVersion{}
}

func (mv *MajorVersion) ReadObj(bytes []byte) int {
	mv.Bytes = bytes[0:u2]
	mv.Version = int(Byte2U2(bytes))
	return u2
}
