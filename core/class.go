package core

import (
	"fmt"
	"strings"
)

const (
	U1_L = 1 // 1 byte
	U2_L = 2 // 2 bytes
	U4_L = 4 // 4 bytes
)
const (
	u1             = U1_L
	u2             = U2_L
	u4             = U4_L
	NewLine        = "\n"
	OneSpace       = " "
	P              = ","
	_class         = "class"
	_extends       = "extends"
	_implements    = "implements"
	_minor_version = "minor version"
	_major_version = "major version"
)

// Bytes ,binary Bytes
type Bytes []byte
type Hex string

// Constant pool
type Tag int
type ClassIndex int
type NameAndTypeIndex int
type NameIndex int
type DescriptorIndex int
type StringIndex int
type Integer int
type Float float32
type Long int64
type Double float64
type ReferenceKind int
type ReferenceIndex int
type BootstrapMethodAttrIndex int
type String string

/**
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
// ClassFile
type ClassFile struct {
	Magic
	MinorVersion
	MajorVersion
	ConstantPoolCount
	CpInfos
	AccessFlag
	ThisClass
	SuperClass
	InterfacesCount
	Interfaces
	FieldsCount
	Fields
	MethodCount
	Methods
	AttributeCount
	Attributes
}

func (cf ClassFile) String() string {
	var str []string
	for _, m := range cf.Methods {
		fmt.Println(m.NameString)
	}

	return strings.Join(str, "\n")
}

/**
public class Main extends AbstractMain<java.lang.String> implements InterfaceMain

*/
func (cf ClassFile) ClassDesc() string {

	// public class Main extends AbstractMain implements InterfaceMain
	var str []string
	str = append(str, cf.ThisDesc())
	str = append(str, cf.SuperDesc())
	str = append(str, cf.InterfaceDesc())
	str = append(str, NewLine)
	str = append(str, cf.VersionDesc())
	str = append(str, NewLine)
	str = append(str, cf.FlagDesc())
	str = append(str, NewLine)
	str = append(str, cf.ConstantPool())
	str = append(str, NewLine)
	str = append(str, cf.FieldDesc())
	str = append(str, NewLine)
	str = append(str, cf.MethodDesc())
	return strings.Join(str, OneSpace)
}

/**

  public Main();
    descriptor: ()V
    flags: ACC_PUBLIC
    Code:
      stack=1, locals=1, args_size=1
         0: aload_0
         1: invokespecial #1                  // Method AbstractMain."<init>":()V
         4: return
      LineNumberTable:
        line 9: 0
*/
func (cf ClassFile) MethodDesc() string {
	var str []string

	mc := int(cf.MethodCount.Count)
	if mc > 0 {
		for _, m := range cf.Methods {
			md := getMDesc(m)
			str = append(str, md)
		}
	}
	return strings.Join(str, NewLine)
}

func getMDesc(m Method) string {
	var str []string

	s1 := GetFlagDesc(m.AccessFlag)
	str = append(str, s1+" "+m.NameString)

	s2 := fmt.Sprintf("descriptor: %s", m.DescriptorString)
	str = append(str, s2)
	s3 := fmt.Sprintf("flags: %s", m.AccessFlag.String())
	str = append(str, s3)

	return strings.Join(str, NewLine) + NewLine
}

/**
  public static final java.lang.Integer INT_MAX;
    descriptor: Ljava/lang/Integer;
    flags: ACC_PUBLIC, ACC_STATIC, ACC_FINAL
    RuntimeVisibleAnnotations:
      0: #65()
*/
func (cf ClassFile) FieldDesc() string {
	var str []string
	fc := int(cf.FieldsCount.Count)
	if fc > 0 {
		str = append(str, "{")
		for _, f := range cf.Fields {
			fd := getFDesc(f)
			str = append(str, fd)
		}
		str = append(str, "}")
	}
	return strings.Join(str, NewLine)
}

func getFDesc(f Field) string {
	var str []string

	s1 := GetFlagDesc(f.AccessFlag)
	s2 := f.DescriptorString
	s3 := f.NameString
	s4 := fmt.Sprintf("%s %s %s;", s1, s2, s3)
	str = append(str, s4)

	s5 := fmt.Sprintf("descriptor: %s", s2)
	str = append(str, s5)

	s6 := fmt.Sprintf("flags: %s", f.AccessFlag.String())
	str = append(str, s6)

	ac := int(f.AttributeCount.Count)

	if ac > 0 {

	}

	return strings.Join(str, NewLine) + NewLine

}

func (cf ClassFile) ConstantPool() string {
	var str []string
	str = append(str, "Constant pool:"+NewLine)
	str = append(str, cf.CpInfos.String())
	return strings.Join(str, "")
}
func (cf ClassFile) FlagDesc() string {
	return "flags: " + cf.AccessFlag.FlagString
}

func (cf ClassFile) VersionDesc() string {
	var str []string
	v1 := cf.MinorVersion.Version
	v2 := cf.MajorVersion.Version

	vs1 := fmt.Sprintf("%13s: %d", _minor_version, v1)
	vs2 := fmt.Sprintf("%14s: %1d", _major_version, v2)
	str = append(str, vs1)
	str = append(str, vs2)
	return strings.Join(str, NewLine)

}

func (cf ClassFile) ThisDesc() string {
	var str []string
	afDesc := GetFlagDesc(cf.AccessFlag)
	str = append(str, afDesc)

	str = append(str, _class)
	str = append(str, cf.ThisClass.String)
	return strings.Join(str, OneSpace)
}

func (cf ClassFile) SuperDesc() string {
	var str []string
	if cf.AccessFlag.HasSuper() {
		if !cf.SuperClass.isJavaLangObject() {
			str = append(str, _extends)
			str = append(str, cf.SuperClass.String)
		}
	}
	return strings.Join(str, OneSpace)
}

func (cf ClassFile) InterfaceDesc() string {
	var str []string
	ifc := cf.InterfacesCount.Count
	var interfaceDesc string
	if ifc > 0 {
		str = append(str, _implements)
		var s []string
		for _, v := range cf.Interfaces {
			s = append(s, v.NameString)
		}
		interfaceDesc = strings.Join(s, P)
		str = append(str, interfaceDesc)
	}

	return strings.Join(str, OneSpace)
}
