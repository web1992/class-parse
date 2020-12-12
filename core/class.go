package core

type ClassReader interface {
	Magic() Magic
	MinorVersion() MinorVersion
	MajorVersion() MajorVersion
	CpCount() ConstantPoolCount
	CpInfo() CpInfo
}
