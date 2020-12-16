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
	filePath string // file paths
	fileByes []byte // files bytes
	pointer  int    // []byte index
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

// ClassFile
// parse bytes to core.ClassFile obj
func (cp *ClassParse) ClassFile() core.ClassFile {

	magic := cp.magic()
	minorVersion := cp.minorVersion()
	majorVersion := cp.majorVersion()
	poolCount := cp.constantPoolCount()
	cpInfos := cp.cpInfos(poolCount)
	accessFlag := cp.accessFlag()
	thisClass := cp.thisClass(cpInfos)
	superClass := cp.superClass(cpInfos)
	interfacesCount := cp.interfacesCount()
	interfaces := cp.interfaces(cpInfos, interfacesCount)
	fieldsCount := cp.fieldsCount()
	fields := cp.fields(cpInfos, fieldsCount)
	methodCount := cp.methodCount()
	methods := cp.methods(cpInfos, methodCount)
	attributeCount := cp.attributeCount()
	attributes := cp.attributes(cpInfos, attributeCount)

	return core.ClassFile{
		Magic:             magic,
		MinorVersion:      minorVersion,
		MajorVersion:      majorVersion,
		ConstantPoolCount: poolCount,
		CpInfos:           cpInfos,
		AccessFlag:        accessFlag,
		ThisClass:         thisClass,
		SuperClass:        superClass,
		InterfacesCount:   interfacesCount,
		Interfaces:        interfaces,
		FieldsCount:       fieldsCount,
		Fields:            fields,
		MethodCount:       methodCount,
		Methods:           methods,
		AttributeCount:    attributeCount,
		Attributes:        attributes,
	}

}

func (cp *ClassParse) magic() core.Magic {
	var m = core.MagicNew()
	cp.Read(m)
	return *m
}

func (cp *ClassParse) minorVersion() core.MinorVersion {

	//cp.magic()
	var mv = core.MinorVersionNew()
	cp.Read(mv)
	return *mv
}

func (cp *ClassParse) majorVersion() core.MajorVersion {

	//cp.minorVersion()
	var mv = core.MajorVersionNew()
	cp.Read(mv)
	return *mv
}

func (cp *ClassParse) constantPoolCount() core.ConstantPoolCount {

	//cp.majorVersion()
	var cpPool = core.ConstantPoolCountNew()
	cp.Read(cpPool)
	return *cpPool
}

func (cp *ClassParse) accessFlag() core.AccessFlag {

	//cp.cpInfos()
	var af = core.AccessFlagNew()
	cp.Read(af)
	return *af
}

func (cp *ClassParse) thisClass(cpInfos core.CpInfos) core.ThisClass {

	//cp.accessFlag()
	var tc = core.ThisClassNew()
	cp.Read(tc)

	ci := tc.ClassIndex
	s := core.GetCp(cpInfos, int(ci))
	tc.String = s

	return *tc
}

func (cp *ClassParse) superClass(cpInfos core.CpInfos) core.SuperClass {

	return core.SuperClass{}
}

func (cp *ClassParse) interfacesCount() core.InterfacesCount {

	return core.InterfacesCount{}
}
func (cp *ClassParse) interfaces(cpInfos core.CpInfos, count core.InterfacesCount) core.Interfaces {

	return core.Interfaces{}
}

func (cp *ClassParse) fieldsCount() core.FieldsCount {

	return core.FieldsCount{}
}

func (cp *ClassParse) fields(cpInfos core.CpInfos, count core.FieldsCount) core.Fields {

	return core.Fields{}
}

func (cp *ClassParse) methodCount() core.MethodCount {

	return core.MethodCount{}
}

func (cp *ClassParse) methods(cpInfos core.CpInfos, count core.MethodCount) core.Methods {

	return core.Methods{}
}

func (cp *ClassParse) attributeCount() core.AttributeCount {

	return core.AttributeCount{}
}

func (cp *ClassParse) attributes(cpInfos core.CpInfos, attributeCount core.AttributeCount) core.Attributes {

	return core.Attributes{}
}
