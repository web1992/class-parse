package core

import "class-parse/utils"

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

const (
	u1 = U1_L
	u2 = U2_L
)

// The constant_pool table is indexed from 1 to constant_pool_count - 1.
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
type CpInfos []interface{}

/**
tag =7
CONSTANT_Class_info {
u1 tag;
u2 name_index;
}
*/

type CpClass struct {
	Tag
	NameIndex
}

func CpClassNew() *CpClass {
	return &CpClass{
		Tag: TAG_CONSTANT_Class,
	}
}

func (class *CpClass) ReadObj(bytes []byte) {
	class.NameIndex = NameIndex(utils.U2(bytes[u1 : u1+u2]))
}

func (class *CpClass) ObjLen() int {

	return u1 + u2
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
	ClassIndex
	NameAndTypeIndex
}

func CpFieldRefNew() *CpFieldRef {
	return &CpFieldRef{Tag: TAG_CONSTANT_Fieldref}
}

func (f *CpFieldRef) ReadObj(bytes []byte) {
	f.ClassIndex = ClassIndex(utils.U2(bytes[u1 : u1+u2]))
	f.NameAndTypeIndex = NameAndTypeIndex(utils.U2(bytes[u1+u2 : u1+u2+u2]))
}

func (f *CpFieldRef) ObjLen() int {
	return u1 + u2 + u2
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
	ClassIndex
	NameAndTypeIndex
}

func CpMethodRefNew() *CpMethodRef {
	return &CpMethodRef{
		Tag: TAG_CONSTANT_Methodref,
	}
}

func (method *CpMethodRef) ReadObj(bytes []byte) {
	method.ClassIndex = ClassIndex(utils.U2(bytes[u1 : u1+u2]))
	method.NameAndTypeIndex = NameAndTypeIndex(utils.U2(bytes[u1+u2 : u1+u2+u2]))
}

func (method *CpMethodRef) ObjLen() int {
	return u1 + u2 + u2
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
	NameIndex
	DescriptorIndex
}

func CpNameAndTypeNew() *CpNameAndType {

	return &CpNameAndType{
		Tag: TAG_CONSTANT_NameAndType,
	}
}

func (cnat *CpNameAndType) ReadObj(bytes []byte) {
	cnat.NameIndex = NameIndex(utils.U2(bytes[u1 : u1+u2]))
	cnat.DescriptorIndex = DescriptorIndex(utils.U2(bytes[u1+u2 : u1+u2+u2]))
}

func (cnat *CpNameAndType) ObjLen() int {
	return u1 + u2 + u2
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
