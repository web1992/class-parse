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

const (
	u1 = U1_L
	u2 = U2_L
	u4 = U4_L
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
	Bytes
}

func CpClassNew() *CpClass {
	return &CpClass{
		Tag: TAG_CONSTANT_Class,
	}
}

func (class *CpClass) ReadObj(bytes []byte) int {
	class.NameIndex = NameIndex(U2(bytes[u1 : u1+u2]))
	class.Bytes = bytes
	return 0
}

func (class *CpClass) ObjLen() int {

	return u1 + u2
}

func (class *CpClass) View() interface{} {

	return "Class"
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
	Bytes
}

func CpFieldRefNew() *CpFieldRef {
	return &CpFieldRef{Tag: TAG_CONSTANT_Fieldref}
}

func (f *CpFieldRef) ReadObj(bytes []byte) int {
	f.Bytes = bytes
	f.ClassIndex = ClassIndex(U2(bytes[u1 : u1+u2]))
	f.NameAndTypeIndex = NameAndTypeIndex(U2(bytes[u1+u2 : u1+u2+u2]))
	return 0
}

func (f *CpFieldRef) ObjLen() int {
	return u1 + u2 + u2
}

func (f *CpFieldRef) View() interface{} {

	return "Fieldref"
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
	Bytes
}

func CpMethodRefNew() *CpMethodRef {
	return &CpMethodRef{
		Tag: TAG_CONSTANT_Methodref,
	}
}

func (method *CpMethodRef) ReadObj(bytes []byte) int {
	method.ClassIndex = ClassIndex(U2(bytes[u1 : u1+u2]))
	method.NameAndTypeIndex = NameAndTypeIndex(U2(bytes[u1+u2 : u1+u2+u2]))
	method.Bytes = bytes
	return 0
}

func (method *CpMethodRef) ObjLen() int {
	return u1 + u2 + u2
}

func (method *CpMethodRef) View() interface{} {
	return "Methodref"
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
	ClassIndex
	NameAndTypeIndex
	Bytes
}

func CpInterfaceMethodRefNew() *CpInterfaceMethodRef {
	return &CpInterfaceMethodRef{
		Tag: TAG_CONSTANT_InterfaceMethodref,
	}
}

func (im *CpInterfaceMethodRef) ReadObj(bytes []byte) int {
	im.ClassIndex = ClassIndex(U2(bytes[u1 : u1+u2]))
	im.NameAndTypeIndex = NameAndTypeIndex(U2(bytes[u1+u2 : u1+u2+u2]))
	im.Bytes = bytes
	return 0
}
func (im *CpInterfaceMethodRef) ObjLen() int {
	return u1 + u2 + u2
}

/*
tag =8
CONSTANT_String_info {
u1 tag;
u2 string_index;
}*/

type CpString struct {
	Tag
	StringIndex
	Bytes
}

func CpStringNew() *CpString {
	return &CpString{
		Tag: TAG_CONSTANT_String,
	}
}

func (s *CpString) ReadObj(bytes []byte) int {
	s.StringIndex = StringIndex(U2(bytes[u1 : u1+u2]))
	s.Bytes = bytes
	return 0
}
func (s *CpString) ObjLen() int {
	return u1 + u2
}

func (s *CpString) View() interface{} {
	return "String"
}

/*
tag=3
CONSTANT_Integer_info {
u1 tag;
u4 bytes;
}*/
// The bytes item of the CONSTANT_Integer_info structure represents the value of the int constant.
// The bytes of the value are stored in big-endian (high byte first) order.
type CpInteger struct {
	Tag
	Bytes
	Integer // bytes to integer
	Hex
}

func CpIntegerNew() *CpInteger {
	return &CpInteger{
		Tag: TAG_CONSTANT_Integer,
	}
}

func (i *CpInteger) ReadObj(bytes []byte) int {
	b := bytes[u1 : u1+u4]
	i.Bytes = b
	i.Integer = Integer(U4(b))
	i.Hex = HexByte(i.Bytes)
	return 0
}
func (i *CpInteger) ObjLen() int {
	return u1 + u4
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
	Bytes
	Float // bytes to float
	Hex
}

func CpFloatNew() *CpFloat {
	return &CpFloat{
		Tag: TAG_CONSTANT_Float,
	}
}

func (f *CpFloat) ReadObj(bytes []byte) int {
	b := bytes[u1 : u1+u4]
	f.Bytes = b
	f.Float = Float(Byte2Float(b))
	f.Hex = HexByte(b)
	return 0
}
func (f *CpFloat) ObjLen() int {
	return u1 + u4
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
	Tag   // tag =5
	Bytes // high_bytes + low_bytes
	Long  // bytes to long
	Hex   // Hex to view
}

func CpLongNew() *CpLong {
	return &CpLong{
		Tag: TAG_CONSTANT_Long,
	}
}

func (l *CpLong) ReadObj(bytes []byte) int {
	b := bytes[u1 : u1+u4+u4]
	l.Bytes = b
	l.Long = Long(Byte2Long(b))
	l.Hex = HexByte(b)
	return 0
}
func (l *CpLong) ObjLen() int {
	return u1 + u4 + u4
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
	Tag    // tag =6
	Bytes  // high_bytes + low_bytes
	Double // bytes to long
	Hex
}

func CpDoubleNew() *CpDouble {
	return &CpDouble{
		Tag: TAG_CONSTANT_Double,
	}
}
func (d *CpDouble) ReadObj(bytes []byte) int {
	b := bytes[u1 : u1+u4+u4]
	d.Bytes = b
	d.Double = Double(Byte2Double(b))
	d.Hex = HexByte(b)
	return 0
}
func (d *CpDouble) ObjLen() int {
	return u1 + u4 + u4
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
	Bytes
}

func CpNameAndTypeNew() *CpNameAndType {

	return &CpNameAndType{
		Tag: TAG_CONSTANT_NameAndType,
	}
}

func (cnat *CpNameAndType) ReadObj(bytes []byte) int {
	cnat.NameIndex = NameIndex(U2(bytes[u1 : u1+u2]))
	cnat.DescriptorIndex = DescriptorIndex(U2(bytes[u1+u2 : u1+u2+u2]))
	cnat.Bytes = bytes
	return 0
}

func (cnat *CpNameAndType) ObjLen() int {
	return u1 + u2 + u2
}

func (cnat *CpNameAndType) View() interface{} {

	return "NameAndType"
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
	len int32
	Bytes
	String // bytes to string

}

func CpUTF8New() *CpUTF8 {
	return &CpUTF8{
		Tag: TAG_CONSTANT_Utf8,
	}
}

func (u *CpUTF8) ReadObj(bytes []byte) int {

	l := U2(bytes[u1 : u1+u2])
	u.len = l
	bs := bytes[u1+u2 : u1+u2+l]
	u.Bytes = bs
	u.String = String(bs)

	return int(l)

}

func (u *CpUTF8) ObjLen() int {
	return u1 + u2
}

func (u *CpUTF8) View() interface{} {
	return u.String
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
	ReferenceKind
	ReferenceIndex
	Bytes
}

func CpMethodHandleNew() *CpMethodHandle {
	return &CpMethodHandle{
		Tag: TAG_CONSTANT_MethodHandle,
	}
}

func (mh *CpMethodHandle) ReadObj(bytes []byte) int {
	mh.Bytes = bytes
	mh.ReferenceKind = ReferenceKind(U1(bytes[u1 : u1+u2]))
	mh.ReferenceIndex = ReferenceIndex(U2(bytes[u1+u2 : u1+u2+u2]))
	return 0
}

func (mh *CpMethodHandle) ObjLen() int {
	return u1 + u1 + u2
}

/*tag=16
CONSTANT_MethodType_info {
u1 tag;
u2 descriptor_index;
}*/

type CpMethodType struct {
	Tag
	DescriptorIndex
	Bytes
}

func CpMethodTypeNew() *CpMethodType {
	return &CpMethodType{
		Tag: TAG_CONSTANT_MethodType,
	}
}

func (mt *CpMethodType) ReadObj(bytes []byte) int {
	mt.Bytes = bytes
	mt.DescriptorIndex = DescriptorIndex(U2(bytes[u1 : u1+u2]))
	return 0
}

func (mt *CpMethodType) ObjLen() int {
	return u1 + u2
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
	BootstrapMethodAttrIndex
	NameAndTypeIndex
	Bytes
}

func CpInvokeDynamicNew() *CpInvokeDynamic {
	return &CpInvokeDynamic{
		Tag: TAG_CONSTANT_InvokeDynamic,
	}
}

func (id *CpInvokeDynamic) ReadObj(bytes []byte) int {

	id.Bytes = bytes
	id.BootstrapMethodAttrIndex = BootstrapMethodAttrIndex(U2(bytes[u1 : u1+u2]))
	id.NameAndTypeIndex = NameAndTypeIndex(U2(bytes[u1+u2 : u1+u2+u2]))
	return 0
}

func (id *CpInvokeDynamic) ObjLen() int {
	return u1 + u2 + u2
}
