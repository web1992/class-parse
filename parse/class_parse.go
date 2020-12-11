package parse

import (
	"class-parse/core"
	"class-parse/utils"
	"errors"
	"strings"
)

const fileSuffix = ".class"

var filePathInvalid = errors.New("file patch is invalid")
var fileSuffixInvalid = errors.New("file must has suffix " + fileSuffix)

type ClassParse struct {
	filePath string
	fileByes []byte
	p        int
}

func (cp *ClassParse) Name() string {
	return cp.filePath
}

func (cp *ClassParse) Bytes() []byte {
	return cp.fileByes
}

func (cp *ClassParse) String() string {
	return cp.filePath
}
func (cp *ClassParse) parseFile(filePath string) error {

	cp.filePath = filePath

	if len(filePath) == 0 {
		return filePathInvalid
	}

	checkBol := strings.HasSuffix(filePath, fileSuffix)

	if !checkBol {
		return fileSuffixInvalid
	}

	bytes, e := utils.ReadClassFile(filePath)

	if e != nil {
		return e
	}

	cp.fileByes = bytes
	return nil
}

func (cp *ClassParse) Magic() core.Magic {
	bytes := cp.Bytes()
	var m core.Magic
	magic := bytes[cp.p:m.ByteLen()]
	cp.p += m.ByteLen()
	m.Bytes = magic
	return m
}

func (cp *ClassParse) MinorVersion() core.MinorVersion {

	cp.Magic()

	bytes := cp.Bytes()

	var mv core.MinorVersion

	p := cp.p
	mvb := bytes[p : p+mv.ByteLen()]
	mv.Bytes = mvb

	cp.p += mv.ByteLen()

	return mv
}

func (cp *ClassParse) MajorVersion() core.MajorVersion {

	cp.MinorVersion()
	var mv core.MajorVersion
	bytes := cp.Bytes()

	p := cp.p
	mvb := bytes[p : p+mv.ByteLen()]
	mv.Bytes = mvb

	cp.p += mv.ByteLen()
	return mv
}

func (cp *ClassParse) ConstantPoolCount() core.ConstantPoolCount {

	cp.MajorVersion()
	var cpool core.ConstantPoolCount
	bytes := cp.Bytes()

	p := cp.p
	mvb := bytes[p : p+cpool.ByteLen()]
	cpool.Bytes = mvb

	cp.p += cpool.ByteLen()
	return cpool
}
