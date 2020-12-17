package core

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
	af.Count = U2(bs)
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
}
