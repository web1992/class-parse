package core

// FieldsCount
// u2 fields_count
type FieldsCount struct {
	Count int32
}

// Field
/*
field_info {
u2             access_flags;
u2             name_index;
u2             descriptor_index;
u2             attributes_count;
attribute_info attributes[attributes_count];
}
*/
type Field struct {
	AccessFlag
	NameIndex
	DescriptorIndex
	AttributeCount
	Attributes
	NameString       string
	DescriptorString string
	AccessFlagString string
}

type Fields []Field

func FieldsCountNew() *FieldsCount {
	return &FieldsCount{}
}

func (tc *FieldsCount) ReadObj(bytes []byte) int {
	bs := bytes[0:u2]
	tc.Count = Byte2U2(bs)
	return u2
}

func FieldNew() *Field {
	return &Field{}
}

func (tc *Field) ReadObj(bytes []byte) int {

	readLen := 0
	var af AccessFlag
	af.Flag = int(Byte2U2(bytes[0 : readLen+u2]))
	readLen += u2
	af.FlagString = GetFlag(af)
	tc.AccessFlag = af

	tc.NameIndex = NameIndex(Byte2U2(bytes[readLen : readLen+u2]))
	readLen += u2
	tc.DescriptorIndex = DescriptorIndex(Byte2U2(bytes[readLen : readLen+u2]))
	readLen += u2

	var ac AttributeCount
	ac.Count = Byte2U2(bytes[readLen : readLen+u2])
	readLen += u2

	tc.AttributeCount = ac
	return u2 * 4
}

/*
field_info {
u2             access_flags;
u2             name_index;
u2             descriptor_index;
u2             attributes_count;
attribute_info attributes[attributes_count];
}
*/
