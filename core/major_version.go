package core

// MajorVersion major_version
type MajorVersion struct {
	Version int
}

func MajorVersionNew() *MajorVersion {
	return &MajorVersion{}
}

func (mv *MajorVersion) ReadObj(bytes []byte) int {
	mv.Version = int(Byte2U2(bytes[0:u2]))
	return u2
}
