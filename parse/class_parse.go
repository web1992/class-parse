package parse

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/web1992/goclass/core"
	"github.com/web1992/goclass/utils"
	"log"
	"path/filepath"
	"strings"
)

const fileSuffix = ".class"

var filePathInvalid = errors.New("file patch is invalid")
var fileSuffixInvalid = errors.New("file must has suffix " + fileSuffix)
var byteNoUsed = errors.New("byte not all used")

type ClassReader struct {
	filePath   string // file path
	fullPath   string // file path
	modTime    int64
	modTimeStr string
	bytes      []byte // files bytes
	pointer    int    // read index
}

func (cp *ClassReader) CpDesc(thisName string) string {
	s1 := fmt.Sprintf("Classfile %s", cp.fullPath)
	s2 := fmt.Sprintf("  Last modified %s; size %d bytes", cp.modTimeStr, len(cp.bytes))
	h := md5.New()
	h.Write(cp.bytes)
	s3 := fmt.Sprintf("  MD5 checksum %s", hex.EncodeToString(h.Sum(nil)))
	s4 := fmt.Sprintf("  Compiled from \"%s%s\"", thisName, ".java")

	return strings.Join([]string{s1, s2, s3, s4}, "\n") + "\n"
}
func (cp *ClassReader) IncrPointer(num int) {
	cp.pointer += num
}

func (cp *ClassReader) Read(r core.Reader) {
	bytes := cp.bytes
	readL := r.ReadObj(bytes[cp.pointer:])
	cp.IncrPointer(readL)
}

func (cp *ClassReader) String() string {
	return cp.filePath
}

func (cp *ClassReader) Parse(filePath string) error {
	return cp.parseFile(filePath)
}

func (cp *ClassReader) parseFile(filePath string) error {

	cp.filePath = filePath

	if len(filePath) == 0 {
		return filePathInvalid
	}

	checkBol := strings.HasSuffix(filePath, fileSuffix)

	if !checkBol {
		return fileSuffixInvalid
	}

	f, e := utils.GetFile(filePath)

	if e != nil {
		return e
	}
	defer f.Close()
	bytes, e := utils.ReadClassFile(f)

	if e != nil {
		return e
	}

	fi, err := f.Stat()
	if err != nil {
		log.Println("stat filei nfo error")
	}
	cp.modTime = fi.ModTime().Unix()
	cp.modTimeStr = fi.ModTime().Format("2006-1-02")

	abs, _ := filepath.Abs(filePath)
	cp.fullPath = abs

	cp.bytes = bytes
	return nil
}

// ClassFile
// parse bytes to core.ClassFile obj
func (cp *ClassReader) ClassFile() core.ClassFile {

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

	// check pointer and bytes len
	if cp.pointer != len(cp.bytes) {
		core.Error.Println(byteNoUsed)
	}
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

func (cp *ClassReader) magic() core.Magic {
	var m = core.MagicNew()
	cp.Read(m)
	return *m
}

func (cp *ClassReader) minorVersion() core.MinorVersion {

	var mv = core.MinorVersionNew()
	cp.Read(mv)
	return *mv
}

func (cp *ClassReader) majorVersion() core.MajorVersion {

	var mv = core.MajorVersionNew()
	cp.Read(mv)
	return *mv
}

func (cp *ClassReader) constantPoolCount() core.ConstantPoolCount {

	var cpPool = core.ConstantPoolCountNew()
	cp.Read(cpPool)
	return *cpPool
}

func (cp *ClassReader) accessFlag() core.AccessFlag {

	var af = core.AccessFlagNew()
	cp.Read(af)
	af.FlagString = core.GetFlag(*af)
	af.FlagDesc = core.GetFlagDesc(*af)
	return *af
}

func (cp *ClassReader) thisClass(cpInfos core.CpInfos) core.ThisClass {

	var tc = core.ThisClassNew()
	cp.Read(tc)

	ci := tc.ClassIndex
	s := core.GetCp(cpInfos, int(ci))
	tc.String = s

	return *tc
}

func (cp *ClassReader) superClass(cpInfos core.CpInfos) core.SuperClass {

	superClass := core.SuperClassNew()
	cp.Read(superClass)
	ci := superClass.ClassIndex
	superClass.String = core.GetCp(cpInfos, int(ci))

	return *superClass
}

func (cp *ClassReader) interfacesCount() core.InterfacesCount {
	interfacesCount := core.InterfacesCountNew()
	cp.Read(interfacesCount)
	return *interfacesCount
}
func (cp *ClassReader) interfaces(cpInfos core.CpInfos, count core.InterfacesCount) core.Interfaces {

	var fs core.Interfaces
	c := count.Count

	for i := 0; i < c; i++ {
		var f core.Interface
		cp.Read(&f)
		ci := f.ClassIndex
		s := core.GetCp(cpInfos, int(ci))
		f.NameString = s
		fs = append(fs, f)
	}

	return fs
}

func (cp *ClassReader) fieldsCount() core.FieldsCount {

	fc := core.FieldsCountNew()
	cp.Read(fc)

	return *fc
}

func (cp *ClassReader) fields(cpInfos core.CpInfos, count core.FieldsCount) core.Fields {

	var fields core.Fields

	c := int(count.Count)
	for i := 0; i < c; i++ {
		field := core.FieldNew()
		cp.Read(field)
		field.NameString = core.GetCp(cpInfos, int(field.NameIndex))
		field.DescriptorString = core.GetCp(cpInfos, int(field.DescriptorIndex))
		field.AccessFlagString = core.GetFlag(field.AccessFlag)
		field.Attributes = cp.attributes(cpInfos, field.AttributeCount)
		fields = append(fields, *field)
	}

	return fields
}

func (cp *ClassReader) methodCount() core.MethodCount {

	mc := core.MethodCountNew()
	cp.Read(mc)
	return *mc
}

func (cp *ClassReader) methods(cpInfos core.CpInfos, count core.MethodCount) core.Methods {

	var ms core.Methods
	c := int(count.Count)
	for i := 0; i < c; i++ {
		m := core.MethodNew()
		cp.Read(m)
		mn := core.GetCp(cpInfos, int(m.NameIndex))
		m.NameString = mn
		m.DescriptorString = core.GetCp(cpInfos, int(m.DescriptorIndex))
		m.AccessFlagString = core.GetFlag(m.AccessFlag)
		// parse method attribute
		//fmt.Println("method name is", mn)
		m.Attributes = cp.attributes(cpInfos, m.AttributeCount)
		//fmt.Println("********************")

		ms = append(ms, *m)
	}

	return ms
}
