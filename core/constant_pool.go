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

// The value of the reference_kind item must be in the range 1 to 9.
// The value denotes the kind of this method handle, which characterizes its bytecode behavior
// Kind	Description	Interpretation
// 1	REF_getField	getfield C.f:T
// 2	REF_getStatic	getstatic C.f:T
// 3	REF_putField	putfield C.f:T
// 4	REF_putStatic	putstatic C.f:T
// 5	REF_invokeVirtual	invokevirtual C.m:(A*)T
// 6	REF_invokeStatic	invokestatic C.m:(A*)T
// 7	REF_invokeSpecial	invokespecial C.m:(A*)T
// 8	REF_newInvokeSpecial	new C; dup; invokespecial C.<init>:(A*)V
// 9	REF_invokeInterface	invokeinterface C.m:(A*)T

const (
	_ = iota
	REF_getField
	REF_getStatic
	REF_putField
	REF_putStatic
	REF_invokeVirtual
	REF_invokeStatic
	REF_invokeSpecial
	REF_newInvokeSpecial
	REF_invokeInterface
)

var referenceMap = make(map[int]string)

func init() {
	referenceMap[REF_getField] = "REF_getField"
	referenceMap[REF_getStatic] = "REF_getStatic"
	referenceMap[REF_putField] = "REF_putField"
	referenceMap[REF_putStatic] = "REF_putStatic"
	referenceMap[REF_invokeVirtual] = "REF_invokeVirtual"
	referenceMap[REF_invokeStatic] = "REF_invokeStatic"
	referenceMap[REF_invokeSpecial] = "REF_invokeSpecial"
	referenceMap[REF_newInvokeSpecial] = "REF_newInvokeSpecial"
	referenceMap[REF_invokeInterface] = "REF_invokeInterface"
}

func getReferenceKind(referenceKind int32) string {

	return referenceMap[int(referenceKind)]
}

// CpInfos is a array for CpInfo
type CpInfos []interface{}

type ConstantIndexU2 int
type ConstantIndexU4 int

func (ci ConstantIndexU2) ReadObj(bytes []byte) int {

	return u2
}

func (ci ConstantIndexU4) ReadObj(bytes []byte) int {

	return u4
}

// ConstantPoolCount constant_pool_count
type ConstantPoolCount struct {
	Count int
}

func (cp *ConstantPoolCount) ByteLen() int {
	return u2
}

func ConstantPoolCountNew() *ConstantPoolCount {
	return &ConstantPoolCount{}
}

func (cpc *ConstantPoolCount) ReadObj(bytes []byte) int {
	cpc.Count = int(Byte2U2(bytes[0:u2]))
	return u2
}

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
		Tag: JVM_CONSTANT_Class,
	}
}

func (class *CpClass) ReadObj(bytes []byte) int {
	class.NameIndex = NameIndex(Byte2U2(bytes[u1 : u1+u2]))
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
	return &CpFieldRef{Tag: JVM_CONSTANT_Fieldref}
}

func (f *CpFieldRef) ReadObj(bytes []byte) int {
	f.ClassIndex = ClassIndex(Byte2U2(bytes[u1 : u1+u2]))
	f.NameAndTypeIndex = NameAndTypeIndex(Byte2U2(bytes[u1+u2 : u1+u2+u2]))
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
		Tag: JVM_CONSTANT_Methodref,
	}
}

func (method *CpMethodRef) ReadObj(bytes []byte) int {
	method.ClassIndex = ClassIndex(Byte2U2(bytes[u1 : u1+u2]))
	method.NameAndTypeIndex = NameAndTypeIndex(Byte2U2(bytes[u1+u2 : u1+u2+u2]))
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
	ClassIndex
	NameAndTypeIndex
}

func CpInterfaceMethodRefNew() *CpInterfaceMethodRef {
	return &CpInterfaceMethodRef{
		Tag: JVM_CONSTANT_InterfaceMethodref,
	}
}

func (im *CpInterfaceMethodRef) ReadObj(bytes []byte) int {
	im.ClassIndex = ClassIndex(Byte2U2(bytes[u1 : u1+u2]))
	im.NameAndTypeIndex = NameAndTypeIndex(Byte2U2(bytes[u1+u2 : u1+u2+u2]))
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
}

func CpStringNew() *CpString {
	return &CpString{
		Tag: JVM_CONSTANT_String,
	}
}

func (s *CpString) ReadObj(bytes []byte) int {
	s.StringIndex = StringIndex(Byte2U2(bytes[u1 : u1+u2]))
	return u1 + u2
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
	Integer // bytes to integer
	Hex
}

func CpIntegerNew() *CpInteger {
	return &CpInteger{
		Tag: JVM_CONSTANT_Integer,
	}
}

func (i *CpInteger) ReadObj(bytes []byte) int {
	b := bytes[u1 : u1+u4]
	i.Integer = Integer(Byte2U4(b))
	i.Hex = HexByte(b)
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
	Float // bytes to float
	Hex
}

func CpFloatNew() *CpFloat {
	return &CpFloat{
		Tag: JVM_CONSTANT_Float,
	}
}

func (f *CpFloat) ReadObj(bytes []byte) int {
	b := bytes[u1 : u1+u4]
	f.Float = Float(Byte2Float(b))
	f.Hex = HexByte(b)
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
	Tag  // tag =5
	Long // bytes to long
	Hex  // Hex to view
}

func CpLongNew() *CpLong {
	return &CpLong{
		Tag: JVM_CONSTANT_Long,
	}
}

func (l *CpLong) ReadObj(bytes []byte) int {
	b := bytes[u1 : u1+u4+u4]
	l.Long = Long(Byte2Long(b))
	l.Hex = HexByte(b)
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
	Double // bytes to long
	Hex
}

func CpDoubleNew() *CpDouble {
	return &CpDouble{
		Tag: JVM_CONSTANT_Double,
	}
}
func (d *CpDouble) ReadObj(bytes []byte) int {
	b := bytes[u1 : u1+u4+u4]
	d.Double = Double(Byte2Double(b))
	d.Hex = HexByte(b)
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
}

func CpNameAndTypeNew() *CpNameAndType {

	return &CpNameAndType{
		Tag: JVM_CONSTANT_NameAndType,
	}
}

func (cnat *CpNameAndType) ReadObj(bytes []byte) int {
	cnat.NameIndex = NameIndex(Byte2U2(bytes[u1 : u1+u2]))
	cnat.DescriptorIndex = DescriptorIndex(Byte2U2(bytes[u1+u2 : u1+u2+u2]))
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
	len    int32
	String // bytes to string

}

func CpUTF8New() *CpUTF8 {
	return &CpUTF8{
		Tag: JVM_CONSTANT_Utf8,
	}
}

func (u *CpUTF8) ReadObj(bytes []byte) int {

	l := Byte2U2(bytes[u1 : u1+u2])
	u.len = l
	bs := bytes[u1+u2 : u1+u2+l]
	u.String = String(bs)

	return u1 + u2 + int(l)

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
}

func CpMethodHandleNew() *CpMethodHandle {
	return &CpMethodHandle{
		Tag: JVM_CONSTANT_MethodHandle,
	}
}

func (mh *CpMethodHandle) ReadObj(bytes []byte) int {
	mh.ReferenceKind = ReferenceKind(Byte2U1(bytes[u1 : u1+u1]))
	mh.ReferenceIndex = ReferenceIndex(Byte2U2(bytes[u1+u1 : u1+u1+u2]))
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
}

func CpMethodTypeNew() *CpMethodType {
	return &CpMethodType{
		Tag: JVM_CONSTANT_MethodType,
	}
}

func (mt *CpMethodType) ReadObj(bytes []byte) int {
	mt.DescriptorIndex = DescriptorIndex(Byte2U2(bytes[u1 : u1+u2]))
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
}

func CpInvokeDynamicNew() *CpInvokeDynamic {
	return &CpInvokeDynamic{
		Tag: JVM_CONSTANT_InvokeDynamic,
	}
}

func (id *CpInvokeDynamic) ReadObj(bytes []byte) int {

	id.BootstrapMethodAttrIndex = BootstrapMethodAttrIndex(Byte2U2(bytes[u1 : u1+u2]))
	id.NameAndTypeIndex = NameAndTypeIndex(Byte2U2(bytes[u1+u2 : u1+u2+u2]))
	return u1 + u2 + u2
}
