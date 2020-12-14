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
	pointer  int
}

func (cp *ClassParse) IncrPointer(num int) {
	cp.pointer += num
}

func (cp *ClassParse) Read(r Reader) {

	bytes := cp.Bytes()

	l := r.ObjLen()
	b := bytes[cp.pointer : cp.pointer+l]
	readL := r.ReadObj(b)

	cp.IncrPointer(l + readL)
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
	var m = core.MagicNew()
	cp.Read(m)
	return *m
}

func (cp *ClassParse) MinorVersion() core.MinorVersion {

	cp.Magic()
	var mv = core.MinorVersionNew()
	cp.Read(mv)
	return *mv
}

func (cp *ClassParse) MajorVersion() core.MajorVersion {

	cp.MinorVersion()
	var mv = core.MajorVersionNew()
	cp.Read(mv)
	return *mv
}

func (cp *ClassParse) ConstantPoolCount() core.ConstantPoolCount {

	cp.MajorVersion()
	var cpPool = core.ConstantPoolCountNew()
	cp.Read(cpPool)
	return *cpPool
}

func (cp *ClassParse) AccessFlag() core.AccessFlag {

	cp.CpInfos()
	var af = core.AccessFlagNew()
	cp.Read(af)
	return *af
}
