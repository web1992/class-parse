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

// Bytes ,binary Bytes
type Bytes []byte
type Hex string

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
