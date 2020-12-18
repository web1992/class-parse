package core

type AttributeNameIndex int32
type AttributeLength int32
type ConstantValueIndex int32
type MaxStack int32
type MaxLocals int32
type CodeLength int32
type ExceptionTableLength int32
type NumberOfExceptions int32
type ExceptionIndexTable []int32
type LineNumberTableLength int32
type StartPc int32
type EndPc int32
type HandlerPc int32
type CatchType int32
type LineNumber int32

//u2             attributes_count;
//attribute_info attributes[attributes_count];
type AttributeCount struct {
	Bytes
	Count int32
}

/*
Attribute

attribute_info {
u2 attribute_name_index;
u4 attribute_length;
u1 info[attribute_length];
}
*/
// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7

func AttributeNew() *Attribute {
	return &Attribute{}
}

type Attribute struct {
	AttributeNameIndex
	AttributeLength
	Name string
}

func (af *Attribute) ReadObj(bytes []byte) int {
	i := Byte2U2(bytes[0:u2])
	af.AttributeNameIndex = AttributeNameIndex(i)

	l := Byte2U4(bytes[u2 : u2+u4])
	af.AttributeLength = AttributeLength(l)

	return int(l)
}

func (af *Attribute) ObjLen() int {
	return u2 + u4
}

type Attributes []interface{}

func AttributeCountNew() *AttributeCount {
	return &AttributeCount{}
}

func (af *AttributeCount) ReadObj(bytes []byte) int {
	bs := bytes[0:u2]
	af.Bytes = bs
	af.Count = Byte2U2(bs)
	return 0
}

func (af *AttributeCount) ObjLen() int {
	return u2
}

/*
ConstantValue_attribute {
u2 attribute_name_index;
u4 attribute_length;
u2 constantvalue_index;
}
*/
type ConstantValueAttribute struct {
	AttributeNameIndex
	AttributeLength
	ConstantValueIndex
}

func (cva *ConstantValueAttribute) ReadObj(bytes []byte) int {
	i := Byte2U2(bytes[0:u2])
	cva.AttributeNameIndex = AttributeNameIndex(i)

	l := Byte2U2(bytes[u2 : u2+u4])
	cva.AttributeLength = AttributeLength(l)

	ii := Byte2U2(bytes[u2+u4 : u2+u4+l])
	cva.ConstantValueIndex = ConstantValueIndex(ii)
	return int(l)
}

func (cva *ConstantValueAttribute) ObjLen() int {
	return u2 + u4
}

/*
Code_attribute {
u2 attribute_name_index;
u4 attribute_length;
u2 max_stack;
u2 max_locals;
u4 code_length;
u1 code[code_length];
u2 exception_table_length;
{   u2 start_pc;
u2 end_pc;
u2 handler_pc;
u2 catch_type;
} exception_table[exception_table_length];
u2 attributes_count;
attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	AttributeNameIndex
	AttributeLength
	MaxStack
	MaxLocals
	CodeLength
	Code []struct {
		Opcode int32
		Desc   string
	}
	ExceptionTableLength
	ExceptionTable []struct {
		StartPc
		EndPc
		HandlerPc
		CatchType
	}
	AttributeCount
	Attributes
}

func (ca *CodeAttribute) ReadObj(bytes []byte) int {
	i := Byte2U2(bytes[0:u2])
	ca.AttributeNameIndex = AttributeNameIndex(i)

	l := Byte2U2(bytes[u2 : u2+u4])
	ca.AttributeLength = AttributeLength(l)

	//ms := Byte2U2(bytes[u2+u4 : u2+u4+u2])
	//ca.MaxStack = MaxStack(ms)
	//
	//ml := Byte2U2(bytes[u2+u4+u2 : u2+u4+u2+u2])
	//ca.MaxLocals=MaxLocals(ml)
	//
	//cl := Byte2U2(bytes[u2+u4+u2+u2 : u2+u4+u2+u2+u4])
	//ca.CodeLength=CodeLength(cl)

	return int(l)
}

func (ca *CodeAttribute) ObjLen() int {
	return u2 + u4
}

/*
Exceptions_attribute {
u2 attribute_name_index;
u4 attribute_length;
u2 number_of_exceptions;
u2 exception_index_table[number_of_exceptions];
}
*/

type ExceptionsAttribute struct {
	AttributeNameIndex
	AttributeLength
	NumberOfExceptions
	ExceptionIndexTable
}

func (ea *ExceptionsAttribute) ReadObj(bytes []byte) int {
	i := Byte2U2(bytes[0:u2])
	ea.AttributeNameIndex = AttributeNameIndex(i)

	l := Byte2U2(bytes[u2 : u2+u4])
	ea.AttributeLength = AttributeLength(l)

	return int(l)
}

func (ea *ExceptionsAttribute) ObjLen() int {
	return u2 + u4
}

/*
LineNumberTable_attribute {
u2 attribute_name_index;
u4 attribute_length;
u2 line_number_table_length;
{   u2 start_pc;
u2 line_number;
} line_number_table[line_number_table_length];
}
*/

type LineNumberTableAttribute struct {
	AttributeNameIndex
	AttributeLength
	LineNumberTableLength
	LineNumberTable []struct {
		StartPc
		LineNumber
	}
}

func (lnta *LineNumberTableAttribute) ReadObj(bytes []byte) int {
	i := Byte2U2(bytes[0:u2])
	lnta.AttributeNameIndex = AttributeNameIndex(i)

	l := Byte2U2(bytes[u2 : u2+u4])
	lnta.AttributeLength = AttributeLength(l)

	return int(l)
}

func (lnta *LineNumberTableAttribute) ObjLen() int {
	return u2 + u4
}

/*
Deprecated_attribute {
u2 attribute_name_index;
u4 attribute_length;
}*/
type DeprecatedAttribute struct {
	AttributeNameIndex
	AttributeLength
}

func (da *DeprecatedAttribute) ReadObj(bytes []byte) int {
	i := Byte2U2(bytes[0:u2])
	da.AttributeNameIndex = AttributeNameIndex(i)

	l := Byte2U2(bytes[u2 : u2+u4])
	da.AttributeLength = AttributeLength(l)

	return 0
}

func (da *DeprecatedAttribute) ObjLen() int {
	return u2 + u4
}
