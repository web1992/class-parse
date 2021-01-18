package core

const (
	U1 = 1 // 1 byte
	U2 = 2 // 2 bytes
	U4 = 4 // 4 bytes
)
const (
	u1             = U1
	u2             = U2
	u4             = U4
	NewLine        = "\n"
	OneSpace       = " "
	COMMA          = ","
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
