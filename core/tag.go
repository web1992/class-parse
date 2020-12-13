package core

const (
	TAG_CONSTANT_Class              = 7
	TAG_CONSTANT_Fieldref           = 9
	TAG_CONSTANT_Methodref          = 10
	TAG_CONSTANT_InterfaceMethodref = 11
	TAG_CONSTANT_String             = 8
	TAG_CONSTANT_Integer            = 3
	TAG_CONSTANT_Float              = 4
	TAG_CONSTANT_Long               = 5
	TAG_CONSTANT_Double             = 6
	TAG_CONSTANT_NameAndType        = 12
	TAG_CONSTANT_Utf8               = 1
	TAG_CONSTANT_MethodHandle       = 15
	TAG_CONSTANT_MethodType         = 16
	TAG_CONSTANT_InvokeDynamic      = 18
)

type Tag int32
type ClassIndex int32
type NameAndTypeIndex int32
type NameIndex int32
type DescriptorIndex int32
