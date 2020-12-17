package core

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

// U1_L 1 byte
// U2_L 2 bytes
// U4_L 4 bytes
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
type Tag int32
type ClassIndex int32
type NameAndTypeIndex int32
type NameIndex int32
type DescriptorIndex int32
type StringIndex int32
type Integer int32
type Float float32
type Long int64
type Double float64
type ReferenceKind int32
type ReferenceIndex int32
type BootstrapMethodAttrIndex int32
type String string

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
