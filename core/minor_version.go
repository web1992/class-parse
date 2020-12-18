package core

// MinorVersion  minor_version
type MinorVersion struct {
	Bytes
	Version int32
}

func (mv *MinorVersion) ByteLen() int {
	return u2
}

func MinorVersionNew() *MinorVersion {
	return &MinorVersion{}
}

func (mv *MinorVersion) ReadObj(bytes []byte) int {
	mv.Bytes = bytes[0:u2]
	mv.Version = Byte2U2(bytes)
	return 0
}

func (mv *MinorVersion) ObjLen() int {
	return u2
}
