package core

// FieldsCount
// u2 fields_count
type FieldsCount struct {
	Bytes
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
	Bytes
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
	tc.Bytes = bs
	tc.Count = Byte2U2(bs)
	return u2
}

func FieldNew() *Field {
	return &Field{}
}

func (tc *Field) ReadObj(bytes []byte) int {
	tc.Bytes = bytes

	var af AccessFlag
	afBytes := bytes[0:u2]
	af.Bytes = afBytes
	af.Flag = Byte2U2(af.Bytes)
	af.FlagString = GetFlag(af)
	tc.AccessFlag = af

	tc.NameIndex = NameIndex(Byte2U2(bytes[u2 : u2*2]))
	tc.DescriptorIndex = DescriptorIndex(Byte2U2(bytes[u2*2 : u2*3]))

	var ac AttributeCount
	acBytes := bytes[u2*3 : u2*4]
	ac.Count = Byte2U2(acBytes)
	ac.Bytes = acBytes
	tc.AttributeCount = ac
	return u2 * 4
}
