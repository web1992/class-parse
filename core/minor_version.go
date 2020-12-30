package core

// MinorVersion  minor_version
type MinorVersion struct {
	Bytes
	Version int
}

func (mv *MinorVersion) ByteLen() int {
	return u2
}

func MinorVersionNew() *MinorVersion {
	return &MinorVersion{}
}

func (mv *MinorVersion) ReadObj(bytes []byte) int {
	mv.Bytes = bytes[0:u2]
	mv.Version = int(Byte2U2(bytes))
	return u2
}
