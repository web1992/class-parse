package core

// MinorVersion  minor_version
type MinorVersion struct {
	Version int
}

func (mv *MinorVersion) ByteLen() int {
	return u2
}

func MinorVersionNew() *MinorVersion {
	return &MinorVersion{}
}

func (mv *MinorVersion) ReadObj(bytes []byte) int {
	bs := bytes[0:u2]
	mv.Version = int(Byte2U2(bs))
	return u2
}
