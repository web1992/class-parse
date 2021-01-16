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
	u1 = U1_L
	u2 = U2_L
	u4 = U4_L
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

	return strings.Join(str, " ")
}

func (cf ClassFile) ThisDesc() string {
	var str []string
	if cf.AccessFlag.HasPublic() {
		str = append(str, "public")
	}
	if cf.AccessFlag.HasAbstract() {
		str = append(str, "abstract")
	}
	str = append(str, "class")
	str = append(str, cf.ThisClass.String)
	return strings.Join(str, " ")
}

func (cf ClassFile) SuperDesc() string {
	var str []string
	if cf.AccessFlag.HasSuper() {
		if !cf.SuperClass.isJavaLangObject() {
			str = append(str, "extends")
			str = append(str, cf.SuperClass.String)
		}
	}
	return strings.Join(str, " ")
}

func (cf ClassFile) InterfaceDesc() string {
	var str []string
	ifc := cf.InterfacesCount.Count
	var interfaceDesc string
	if ifc > 0 {
		str = append(str, "implements")
		var s []string
		for _, v := range cf.Interfaces {
			s = append(s, v.NameString)
		}
		interfaceDesc = strings.Join(s, ",")
		str = append(str, interfaceDesc)
	}

	return strings.Join(str, " ")
}
