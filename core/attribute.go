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
type Attribute struct {
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
	Code struct {
		Opcode int32
		Desc   string
	}
	ExceptionTableLength
	ExceptionTable []struct {
		//u2 start_pc
		//u2 end_pc
		//u2 handler_pc
		//u2 catch_type
	}
	AttributeCount
	Attributes
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

/*
Deprecated_attribute {
u2 attribute_name_index;
u4 attribute_length;
}*/
type DeprecatedAttribute struct {
	AttributeNameIndex
	AttributeLength
}
