package core

type AttributeNameIndex int
type AttributeLength int
type ConstantValueIndex int
type MaxStack int
type MaxLocals int
type CodeBytesLength int
type ExceptionTableLength int
type NumberOfExceptions int
type ExceptionIndexTable []int
type LineNumberTableLength int
type StartPc int
type EndPc int
type HandlerPc int
type CatchType int
type LineNumber int

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
	//return int(l)
	return u2 + u4 + int(l)
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

attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}

*/

type LineNumberTableAttr struct {
	Attribute
}

func (lnt *LineNumberTableAttr) ReadObj(bytes []byte) int {

	return 0
}

type CodeAttribute struct {
	Attribute
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
	CatchType int32
	TypeDesc  string
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
		et.CatchType = Byte2U2(bytes[readLen : readLen+u2])
		readLen = readLen + u2
		ca.ExceptionTable = append(ca.ExceptionTable, et)
	}

	ac := int(Byte2U2(bytes[readLen : readLen+u2]))
	ca.AttributeCount = AttributeCount{Count: int32(ac)}
	readLen = readLen + u2
	for i := 0; i < ac; i++ {
		attributeNameIndex := int(Byte2U2(bytes[readLen : readLen+u2]))
		readLen = readLen + u2
		attributeLen := int(Byte2U4(bytes[readLen : readLen+u4]))
		readLen = readLen + (u4 + attributeLen)
		attr := CreateAttributeByIndex(attributeNameIndex, ca.Attribute.CpInfos)
		ca.Attributes = append(ca.Attributes, attr)
	}
	return u2 + u4 + int(l)
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
	Attribute
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

	return u2 + u4 + int(l)
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
	LineNumberTable []LineNumberTableAttr
}

func (lnta *LineNumberTableAttribute) ReadObj(bytes []byte) int {
	i := Byte2U2(bytes[0:u2])
	lnta.AttributeNameIndex = AttributeNameIndex(i)

	l := Byte2U4(bytes[u2 : u2+u4])
	lnta.AttributeLength = AttributeLength(l)

	return u2 + u4 + int(l)
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
	Attribute
	SourceFileIndex int32
	AttributeName   string
	SourceFileName  string
}

func (sfa *SourceFileAttribute) ReadObj(bytes []byte) int {
	sfa.AttributeNameIndex = Byte2U2(bytes[0:u2])

	l := Byte2U4(bytes[u2 : u2+u4])
	sfa.AttributeLength = l

	sfa.SourceFileIndex = Byte2U2(bytes[u2+u4 : u2+u4+u2])
	sfa.SourceFileName = GetCp(sfa.CpInfos, int(sfa.SourceFileIndex))

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
	Attribute
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

	return u2 + u4 + int(l)
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
	Attribute
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
	return u2 + u4 + int(l)
}
