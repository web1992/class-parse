package core

import "fmt"

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

For all attributes, the attribute_name_index must be a valid unsigned
16-bit index into the constant pool of the class.
The constant_pool entry at attribute_name_index must be a
CONSTANT_Utf8_info structure (§4.4.7) representing the name of the attribute.
The value of the attribute_length item indicates the length of the subsequent information in bytes.
The length does not include the initial six bytes that
contain the attribute_name_index and attribute_length items.

*/

type CodeAttribute struct {
	Offset int
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
	u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];

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

func (et *ExceptionTable) getDesc(cpInfo CpInfos) string {

	if et.CatchType == 0 {
		return "any"
	}
	return GetCp(cpInfo, int(et.CatchType))
}
func (et *ExceptionTable) ReadObj(bytes []byte) int {

	readLen := 0

	et.StartPc = StartPc(Byte2U2(bytes[0:u2]))
	readLen += u2

	et.EndPc = EndPc(Byte2U2(bytes[readLen : readLen+u2]))
	readLen += u2

	et.HandlerPc = HandlerPc(Byte2U2(bytes[readLen : readLen+u2]))
	readLen = readLen + u2

	et.CatchType = Byte2U2(bytes[readLen : readLen+u2])
	readLen = readLen + u2

	return u2 * 4
}

func (ca *CodeAttribute) ReadObj(bytes []byte) int {
	readLen := 0
	i := Byte2U2(bytes[0:u2])
	readLen += u2
	ca.AttributeNameIndex = AttributeNameIndex(i)
	l := Byte2U4(bytes[readLen : readLen+u4])
	readLen += u4
	ca.AttributeLength = AttributeLength(l)

	ms := Byte2U2(bytes[readLen : readLen+u2])
	ca.MaxStack = MaxStack(ms)
	readLen += u2

	ml := Byte2U2(bytes[readLen : readLen+u2])
	ca.MaxLocals = MaxLocals(ml)
	readLen += u2

	cl := int(Byte2U4(bytes[readLen : readLen+u4]))
	ca.CodeBytesLength = CodeBytesLength(cl)
	readLen += u4
	offset := ca.Offset + readLen

	bs := bytes[readLen : readLen+cl]
	ca.CodeBytes = bs
	readLen += cl

	ca.OpCodes = ca.ParseOpCodes(offset, cl, bs)

	etl := Byte2U2(bytes[readLen : readLen+u2])
	ca.ExceptionTableLength = ExceptionTableLength(etl)
	readLen += u2

	for i := 0; i < int(etl); i++ {
		var et ExceptionTable
		rl := et.ReadObj(bytes[readLen:])
		et.TypeDesc = et.getDesc(ca.CpInfos)
		readLen += rl
		ca.ExceptionTable = append(ca.ExceptionTable, et)
	}

	ac := int(Byte2U2(bytes[readLen : readLen+u2]))
	ca.AttributeCount = AttributeCount{Count: int32(ac)}
	readLen = readLen + u2
	for i := 0; i < ac; i++ {
		attributeNameIndex := int(Byte2U2(bytes[readLen : readLen+u2]))
		readLen = readLen + u2
		attributeLen := int(Byte2U4(bytes[readLen : readLen+u4]))
		readLen += u4
		attr := CreateAttributeByIndex(attributeNameIndex, ca.Attribute.CpInfos)
		if obj, ok := attr.(Reader); ok {
			obj.ReadObj(bytes[readLen-u2-u4 : readLen+attributeLen])
		}
		readLen += attributeLen
		ca.Attributes = append(ca.Attributes, attr)
	}
	return u2 + u4 + int(l)
}

func (ca *CodeAttribute) ParseOpCodes(offset int, codeLength int, bs []byte) OpCodes {

	var ops OpCodes
	hadReadLen := 0
	for hadReadLen < codeLength {

		op := Byte2U1(bs[hadReadLen : hadReadLen+U1_L])
		_bs := bs[hadReadLen:]
		desc := GetOpDesc(int(op))
		opObj := CreateOpCode(op)
		if o, ok := opObj.(*OpCodeTableSwitch); ok {
			// The alignment required of the 4-byte operands of the tableswitch instruction guarantees 4-byte alignment of those operands if
			// and only if the method that contains the tableswitch starts on a 4-byte boundary.
			o.Offset = offset + hadReadLen + 1
			o.Base = int32(hadReadLen)
			o.LineNo = hadReadLen
			readLen := o.ReadObj(_bs)
			fmt.Printf("%d: %s \n", hadReadLen, GetTableSwitchDesc(*o, desc))
			hadReadLen = readLen + hadReadLen
			ops = append(ops, o)
			continue
		}
		if o, ok := opObj.(*OpCodeLookupSwitch); ok {
			o.Offset = offset + hadReadLen
			o.Base = int32(hadReadLen)
			o.LineNo = hadReadLen
			readLen := o.ReadObj(_bs)
			fmt.Printf("%d: %s \n", hadReadLen, GetLookupSwitchDesc(*o, desc))
			hadReadLen = readLen + hadReadLen
			ops = append(ops, o)
			continue
		}
		if o, ok := opObj.(Reader); ok {
			readLen := o.ReadObj(_bs)
			fmt.Printf("%d: %s \n", hadReadLen, desc)

			if opc, ok := opObj.(OpCoder); ok {
				opc.SetLineNo(hadReadLen)
			}

			hadReadLen = readLen + hadReadLen
			ops = append(ops, o)
			continue
		} else {
			panic(fmt.Sprintf("Error opObj %v", opObj))
		}
	}

	return ops
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
	Attribute
	Name string
	AttributeNameIndex
	AttributeLength
	LineNumberTableLength
	LineNumberTable []LineTable
}

type LineTable struct {
	StartPc
	LineNumber
}

func (lt *LineTable) ReadObj(bytes []byte) int {
	lt.StartPc = StartPc(Byte2U2(bytes[0:u2]))
	lt.LineNumber = LineNumber(Byte2U2(bytes[u2 : u2+u2]))
	return u2 + u2
}

func (lineTable *LineNumberTableAttribute) ReadObj(bytes []byte) int {
	readLen := 0
	i := Byte2U2(bytes[readLen : readLen+u2])
	lineTable.AttributeNameIndex = AttributeNameIndex(i)
	readLen += u2

	l := int(Byte2U4(bytes[readLen : readLen+u4]))
	lineTable.AttributeLength = AttributeLength(l)
	readLen += u4

	tableLen := int(Byte2U2(bytes[readLen : readLen+u2]))
	lineTable.LineNumberTableLength = LineNumberTableLength(tableLen)
	readLen += u2

	bs := bytes[readLen:]
	rl := 0
	for i := 0; i < tableLen; i++ {
		var lt LineTable
		rl += lt.ReadObj(bs[rl:])
		lineTable.LineNumberTable = append(lineTable.LineNumberTable, lt)
	}

	return u2 + u4 + l
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
	readLen := 0
	sfa.AttributeNameIndex = Byte2U2(bytes[0 : readLen+u2])
	readLen += u2
	l := Byte2U4(bytes[readLen : readLen+u4])
	sfa.AttributeLength = l
	readLen += u4

	sfa.SourceFileIndex = Byte2U2(bytes[readLen : readLen+u2])
	readLen += u2
	sfa.SourceFileName = GetCp(sfa.CpInfos, int(sfa.SourceFileIndex))

	return readLen
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
	readLen := 0
	i := int(Byte2U2(bytes[0 : readLen+u2]))
	readLen += u2
	bma.AttributeNameIndex = AttributeNameIndex(i)
	bma.AttributeName = GetCp(bma.CpInfos, i)

	l := int(Byte2U4(bytes[readLen : readLen+u4]))
	bma.AttributeLength = AttributeLength(l)
	readLen += u4

	m := Byte2U2(bytes[readLen : readLen+u2])
	bma.NumBootstrapMethods = m
	readLen += u2

	mNum := int(m)
	for i := 0; i < mNum; i++ {
		var bm BootstrapMethod
		bm.BootstrapMethodRef = Byte2U2(bytes[readLen : readLen+u2])
		bm.BootstrapMethodRefName = GetCp(bma.CpInfos, int(bm.BootstrapMethodRef))
		readLen += u2

		nba := Byte2U2(bytes[readLen : readLen+u2])
		bm.NumBootstrapArguments = nba
		readLen += u2

		for j := 0; j < int(nba); j++ {
			nba := Byte2U2(bytes[readLen : readLen+u2])
			readLen += u2
			bm.BootstrapArguments = append(bm.BootstrapArguments, nba)
			bm.BootstrapArgumentName = append(bm.BootstrapArgumentName, GetCp(bma.CpInfos, int(nba)))
		}
		bma.BootstrapMethods = append(bma.BootstrapMethods, bm)
	}
	return u2 + u4 + l
}

/**

RuntimeVisibleAnnotations_attribute {
    u2         attribute_name_index;
    u4         attribute_length;
    u2         num_annotations;
    annotation annotations[num_annotations];
}

annotation {
    u2 type_index;
    u2 num_element_value_pairs;
    {   u2            element_name_index;
        element_value value;
    } element_value_pairs[num_element_value_pairs];
}

*/
type RuntimeVisibleAnnotationsAttr struct {
	Attribute
	NumAnnotations int
	Annotations    []Annotation
}

func (rva *RuntimeVisibleAnnotationsAttr) ReadObj(bytes []byte) int {

	readLen := 0
	i := Byte2U2(bytes[0:u2])
	readLen += u2
	rva.AttributeNameIndex = i

	l := Byte2U4(bytes[readLen : readLen+u4])
	rva.AttributeLength = l
	readLen += u4

	rva.NumAnnotations = int(Byte2U2(bytes[readLen : readLen+u2]))
	readLen += u2

	for i := 0; i < rva.NumAnnotations; i++ {
		var ann Annotation
		ann.CpInfos = rva.Attribute.CpInfos
		readLen += ann.ReadObj(bytes[readLen:])
		rva.Annotations = append(rva.Annotations, ann)
	}

	return u2 + u4 + int(l)
}

/**

annotation {
    u2 type_index;
    u2 num_element_value_pairs;
    {   u2            element_name_index;
        element_value value;
    } element_value_pairs[num_element_value_pairs];
}
*/
type Annotation struct {
	CpInfos
	TypeIndex        int
	TypeName         string
	NumPairs         int
	ElementValuePair []ElementValuePair
}

func (ann *Annotation) ReadObj(bytes []byte) int {

	readLen := 0
	typeIndex := int(Byte2U2(bytes[readLen : readLen+u2]))
	readLen += u2
	ann.TypeIndex = typeIndex
	ann.TypeName = GetCp(ann.CpInfos, typeIndex)
	//fmt.Printf("Annotation name is %s \n", ann.TypeName)
	pairsNum := int(Byte2U2(bytes[readLen : readLen+u2]))
	readLen += u2
	ann.NumPairs = pairsNum

	for i := 0; i < pairsNum; i++ {
		var evp ElementValuePair
		elementNameIndex := int(Byte2U2(bytes[readLen : readLen+u2]))
		readLen += u2
		evp.ElementNameIndex = elementNameIndex
		evp.ElementName = GetCp(ann.CpInfos, elementNameIndex)

		var ev ElementValue
		ev.CpInfos = ann.CpInfos
		readLen += ev.ReadObj(bytes[readLen:])
		evp.ElementValue = ev

		ann.ElementValuePair = append(ann.ElementValuePair, evp)
	}

	return readLen
}

type ElementValuePair struct {
	ElementNameIndex int
	ElementName      string
	ElementValue
}

/*
tag Item	Type	value Item	Constant Type
B	byte	const_value_index	CONSTANT_Integer
C	char	const_value_index	CONSTANT_Integer
D	double	const_value_index	CONSTANT_Double
F	float	const_value_index	CONSTANT_Float
I	int		const_value_index	CONSTANT_Integer
J	long	const_value_index	CONSTANT_Long
S	short	const_value_index	CONSTANT_Integer
Z	boolean	const_value_index	CONSTANT_Integer
s	String	const_value_index	CONSTANT_Utf8
e	Enum type	enum_const_value	Not applicable
c	Class	class_info_index	Not applicable
@	Annotation type	annotation_value	Not applicable
[	Array type	array_value	Not applicable
*/
type ElementValue struct {
	CpInfos
	Tag
	TagDesc string
	Value
}

type Value struct {
	ConstValue
	EnumConstValue
	ClassInfoValue
	Annotation
	ArrayValue
}

type ConstValue struct {
	ConstValueIndex int
	Name            string
}

type ClassInfoValue struct {
	ClassInfoIndex int
	Name           string
}
type EnumConstValue struct {
	TypeNameIndex  int
	ConstNameIndex int
}

type ArrayValue struct {
	NumValues int
	Values    []ElementValue
}

/**
element_value {
    u1 tag;
    union {
        u2 const_value_index;

        {   u2 type_name_index;
            u2 const_name_index;
        } enum_const_value;

        u2 class_info_index;

        annotation annotation_value;

        {   u2            num_values;
            element_value values[num_values];
        } array_value;
    } value;
}
*/
func (ev *ElementValue) ReadObj(bytes []byte) int {

	readLen := 0
	ev.Tag = Tag(Byte2U1(bytes[readLen : readLen+u1]))
	readLen += u1
	cpInfos := ev.CpInfos
	ev.TagDesc = GetTagDesc(ev.Tag)

	switch ev.Tag {
	case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z', 's':
		{
			var v Value
			ev.Value = v
			v.ConstValueIndex = int(Byte2U2(bytes[readLen : readLen+u2]))
			v.ConstValue.Name = GetCp(cpInfos, v.ConstValueIndex)
			readLen += u2
		}
	case 'e':
		{
			var v Value
			ev.Value = v

			typeNameIndex := int(Byte2U2(bytes[readLen : readLen+u2]))
			readLen += u2
			constNameIndex := int(Byte2U2(bytes[readLen : readLen+u2]))
			readLen += u2

			var ecv EnumConstValue
			ecv.ConstNameIndex = constNameIndex
			ecv.TypeNameIndex = typeNameIndex
			v.EnumConstValue = ecv
		}
	case 'c':
		{
			var v Value
			ev.Value = v
			v.ClassInfoIndex = int(Byte2U2(bytes[readLen : readLen+u2]))
			v.ClassInfoValue.Name = GetCp(cpInfos, v.ClassInfoIndex)
			readLen += u2
		}
	case '@':
		{
			var v Value
			ev.Value = v
			var ann Annotation
			ann.CpInfos = cpInfos
			readLen += ann.ReadObj(bytes[readLen:])
			v.Annotation = ann
		}
	case '[':
		{
			var v Value
			ev.Value = v

			numValues := int(Byte2U2(bytes[readLen : readLen+u2]))
			readLen += u2

			var av ArrayValue
			v.ArrayValue = av
			av.NumValues = numValues

			for i := 0; i < numValues; i++ {
				var _ev ElementValue
				_ev.CpInfos = cpInfos
				readLen += _ev.ReadObj(bytes[readLen:])
				av.Values = append(av.Values, _ev)
			}
		}
	default:
		panic(fmt.Sprintf("error for tag %d", ev.Tag))
	}

	return readLen
}

func GetTagDesc(tag Tag) string {

	return string(tag)
}
