package core

/**
| Constant Type               | Value |
| --------------------------- | ----- |
| CONSTANT_Class              | 7     |
| CONSTANT_Fieldref           | 9     |
| CONSTANT_Methodref          | 10    |
| CONSTANT_InterfaceMethodref | 11    |
| CONSTANT_String             | 8     |
| CONSTANT_Integer            | 3     |
| CONSTANT_Float              | 4     |
| CONSTANT_Long               | 5     |
| CONSTANT_Double             | 6     |
| CONSTANT_NameAndType        | 12    |
| CONSTANT_Utf8               | 1     |
| CONSTANT_MethodHandle       | 15    |
| CONSTANT_MethodType         | 16    |
| CONSTANT_InvokeDynamic      | 18    |
*/

type CpInfo struct {
	CpClass
	CpFieldRef
	CpMethodRef
	CpInterfaceMethodRef
	CpString
	CpInteger
	CpFloat
	CpLong
	CpDouble
	CpNameAndType
	CpUTF8
	CpMethodHandle
	CpMethodType
	CpInvokeDynamic
}

// CpInfos is a array for CpInfo
type CpInfos []CpInfo

/**
tag =7
CONSTANT_Class_info {
u1 tag;
u2 name_index;
}
*/

type CpClass struct {
	Tag
}

/*
tag =9
CONSTANT_Fieldref_info {
u1 tag;
u2 class_index;
u2 name_and_type_index;
}*/
type CpFieldRef struct {
	Tag
}

/**
tag =10
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/

type CpMethodRef struct {
	Tag
}

/*
tag =11
CONSTANT_InterfaceMethodref_info {
u1 tag;
u2 class_index;
u2 name_and_type_index;
}*/

type CpInterfaceMethodRef struct {
	Tag
}

/*
tag =8
CONSTANT_String_info {
u1 tag;
u2 string_index;
}*/

type CpString struct {
	Tag
}

/*
tag=3
CONSTANT_Integer_info {
u1 tag;
u4 bytes;
}*/
type CpInteger struct {
	Tag
}

/**
tag =4
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
type CpFloat struct {
	Tag
}

/**
tag=5
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type CpLong struct {
	Tag
}

/**
tag =6
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/

type CpDouble struct {
	Tag
}

/*
tag =12
CONSTANT_NameAndType_info {
u1 tag;
u2 name_index;
u2 descriptor_index;
}*/

type CpNameAndType struct {
	Tag
}

/**
tag=1
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
type CpUTF8 struct {
	Tag
}

/*
tag=15
CONSTANT_MethodHandle_info {
u1 tag;
u1 reference_kind;
u2 reference_index;
}
*/

type CpMethodHandle struct {
	Tag
}

/*tag=16
CONSTANT_MethodType_info {
u1 tag;
u2 descriptor_index;
}*/

type CpMethodType struct {
	Tag
}

/*
tag=18
CONSTANT_InvokeDynamic_info {
u1 tag;
u2 bootstrap_method_attr_index;
u2 name_and_type_index;
}*/

type CpInvokeDynamic struct {
	Tag
}
