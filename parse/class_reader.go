package parse

import (
	"class-parse/core"
)

// ClassReader read class file
type ClassReader interface {
	Magic() core.Magic
	MinorVersion() core.MinorVersion
	MajorVersion() core.MajorVersion
	CpCount() core.ConstantPoolCount
	CpInfo() core.CpInfo
}

// CpReader read Constant poll form file
type CpReader interface {
	ReadCpInfo() core.CpInfo
}
