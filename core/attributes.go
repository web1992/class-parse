package core

type AttributeNameIndex int32
type AttributeLength int32
type ConstantValueIndex int32
type MaxStack int32
type MaxLocals int32
type CodeBytesLength int32
type ExceptionTableLength int32
type NumberOfExceptions int32
type ExceptionIndexTable []int32
type LineNumberTableLength int32
type StartPc int32
type EndPc int32
type HandlerPc int32
type CatchType int32
type LineNumber int32

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

	l := Byte2U4(bytes[u2 : u2+u4])
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
	Name string
	AttributeNameIndex
	AttributeLength
	MaxStack
	MaxLocals
	CodeBytesLength
	CodeBytes []byte
	OpCodes
	ExceptionTableLength
	ExceptionTable []ExceptionTable
	AttributeCount
	AttributesBytes []byte
	Attributes
}

/**
  Exception table:
     from    to  target type
        64    72    83   Class java/lang/Exception
        64    72   104   any
        83    93   104   any
       104   106   104   any
*/
type ExceptionTable struct {
	StartPc
	EndPc
	HandlerPc
	CatchType
}

func (ca *CodeAttribute) ReadObj(bytes []byte) int {
	i := Byte2U2(bytes[0:u2])
	ca.AttributeNameIndex = AttributeNameIndex(i)
	l := Byte2U4(bytes[u2 : u2+u4])
	ca.AttributeLength = AttributeLength(l)

	ms := Byte2U2(bytes[u2+u4 : u2+u4+u2])
	ca.MaxStack = MaxStack(ms)

	ml := Byte2U2(bytes[u2+u4+u2 : u2+u4+u2+u2])
	ca.MaxLocals = MaxLocals(ml)

	cl := Byte2U4(bytes[u2+u4+u2+u2 : u2+u4+u2+u2+u4])
	ca.CodeBytesLength = CodeBytesLength(cl)

	bs := bytes[u2+u4+u2+u2+u4 : u2+u4+u2+u2+u4+cl]
	ca.CodeBytes = bs

	etl := Byte2U2(bytes[u2+u4+u2+u2+u4+cl : u2+u4+u2+u2+u4+cl+u2])
	ca.ExceptionTableLength = ExceptionTableLength(etl)

	var readLen = int(u2 + u4 + u2 + u2 + u4 + cl + u2)
	for i := 0; i < int(etl); i++ {
		var et ExceptionTable
		et.StartPc = StartPc(Byte2U2(bytes[readLen : readLen+u2]))
		readLen = readLen + u2
		et.EndPc = EndPc(Byte2U2(bytes[readLen : readLen+u2]))
		readLen = readLen + u2
		et.HandlerPc = HandlerPc(Byte2U2(bytes[readLen : readLen+u2]))
		readLen = readLen + u2
		et.CatchType = CatchType(Byte2U2(bytes[readLen : readLen+u2]))
		readLen = readLen + u2
		ca.ExceptionTable = append(ca.ExceptionTable, et)
	}

	ac := int(Byte2U2(bytes[readLen : readLen+u2]))
	ca.AttributeCount = AttributeCount{Count: int32(ac)}
	readLen = readLen + u2
	ca.AttributesBytes = bytes[readLen:]
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
	Name string
	AttributeNameIndex
	AttributeLength
	NumberOfExceptions
	ExceptionIndexTable
}

func (ea *ExceptionsAttribute) ReadObj(bytes []byte) int {
	i := Byte2U2(bytes[0:u2])
	ea.AttributeNameIndex = AttributeNameIndex(i)

	l := Byte2U4(bytes[u2 : u2+u4])
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

	l := Byte2U4(bytes[u2 : u2+u4])
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

	l := Byte2U4(bytes[u2 : u2+u4])
	da.AttributeLength = AttributeLength(l)

	return 0
}

func (da *DeprecatedAttribute) ObjLen() int {
	return u2 + u4
}

/*
SourceFile_attribute {
u2 attribute_name_index;
u4 attribute_length;
u2 sourcefile_index;
}
*/
type SourceFileAttribute struct {
	Name string
	AttributeNameIndex
	AttributeLength
	SourceFileIndex int32
	AttributeName   string
	SourceFileName  string
}

func (sfa *SourceFileAttribute) ReadObj(bytes []byte) int {
	i := Byte2U2(bytes[0:u2])
	sfa.AttributeNameIndex = AttributeNameIndex(i)

	l := Byte2U4(bytes[u2 : u2+u4])
	sfa.AttributeLength = AttributeLength(l)

	sfa.SourceFileIndex = Byte2U2(bytes[u2+u4 : u2+u4+u2])

	return 0
}

func (sfa *SourceFileAttribute) ObjLen() int {
	return u2 + u4 + u2
}

/*
InnerClasses_attribute {
u2 attribute_name_index;
u4 attribute_length;
u2 number_of_classes;
{   u2 inner_class_info_index;
u2 outer_class_info_index;
u2 inner_name_index;
u2 inner_class_access_flags;
} classes[number_of_classes];
}
*/

type InnerClassesAttribute struct {
	Name string
	AttributeNameIndex
	AttributeLength
	NumberOfClasses int32
	Classes         []struct {
		InnerClassInfoIndex   int32
		OuterClassInfoIndex   int32
		InnerNameIndex        int32
		InnerClassAccessFlags int32
	}
}

func (sfa *InnerClassesAttribute) ReadObj(bytes []byte) int {
	i := Byte2U2(bytes[0:u2])
	sfa.AttributeNameIndex = AttributeNameIndex(i)

	l := Byte2U4(bytes[u2 : u2+u4])
	sfa.AttributeLength = AttributeLength(l)

	return int(l)
}

func (sfa *InnerClassesAttribute) ObjLen() int {
	return u2 + u4
}

/*
BootstrapMethods_attribute {
u2 attribute_name_index;
u4 attribute_length;
u2 num_bootstrap_methods;
{
u2 bootstrap_method_ref;
u2 num_bootstrap_arguments;
u2 bootstrap_arguments[num_bootstrap_arguments];

} bootstrap_methods[num_bootstrap_methods];
}
*/

type BootstrapMethodsAttribute struct {
	Name string
	AttributeNameIndex
	AttributeLength
	NumBootstrapMethods int32
	BootstrapMethods    []BootstrapMethod
	AttributeName       string
}

type BootstrapMethod struct {
	BootstrapMethodRef    int32
	NumBootstrapArguments int32
	BootstrapArguments    []int32

	BootstrapMethodRefName string
	BootstrapArgumentName  []string
}

func (bma *BootstrapMethodsAttribute) ReadObj(bytes []byte) int {
	i := Byte2U2(bytes[0:u2])
	bma.AttributeNameIndex = AttributeNameIndex(i)

	l := Byte2U4(bytes[u2 : u2+u4])
	bma.AttributeLength = AttributeLength(l)

	m := Byte2U2(bytes[u2+u4 : u2+u4+u2])
	bma.NumBootstrapMethods = m

	bs := bytes[u2+u4+u2 : u2+u4+u2+l]

	mNum := int(m)
	for i := 0; i < mNum; i++ {
		base := i * u2
		var bm BootstrapMethod
		bm.BootstrapMethodRef = Byte2U2(bs[base : base+u2])
		nba := Byte2U2(bs[base+u2 : base+u2+u2])
		bm.NumBootstrapArguments = nba

		bs2 := bs[base+u2+u2 : base+u2+u2+u2]
		for j := 0; j < int(nba); j++ {
			base2 := j * u2
			nba := Byte2U2(bs2[base2 : base2+u2])
			bm.BootstrapArguments = append(bm.BootstrapArguments, nba)
		}
		bma.BootstrapMethods = append(bma.BootstrapMethods, bm)
	}
	return int(l)
}

func (bma *BootstrapMethodsAttribute) ObjLen() int {
	return u2 + u4
}
