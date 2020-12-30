package core

// u2             methods_count;
// method_info    methods[methods_count];
type MethodCount struct {
	Count int32
}

/*
method_info {
u2             access_flags;
u2             name_index;
u2             descriptor_index;
u2             attributes_count;
attribute_info attributes[attributes_count];
}
*/
type Method struct {
	AccessFlag
	NameIndex
	DescriptorIndex
	AttributeCount
	Attributes
	NameString       string
	DescriptorString string
	AccessFlagString string
}

type Methods []Method

func MethodCountNew() *MethodCount {
	return &MethodCount{}
}

func (tc *MethodCount) ReadObj(bytes []byte) int {
	bs := bytes[0:u2]
	tc.Count = Byte2U2(bs)
	return u2
}

func MethodNew() *Method {
	return &Method{}
}

func (tc *Method) ReadObj(bytes []byte) int {

	var af AccessFlag
	af.ReadObj(bytes[0:u2])
	af.FlagString = GetFlag(af)
	tc.AccessFlag = af
	tc.NameIndex = NameIndex(Byte2U2(bytes[u2 : u2+u2]))
	tc.DescriptorIndex = DescriptorIndex(Byte2U2(bytes[u2+u2 : u2+u2+u2]))

	var ac AttributeCount
	bs := bytes[u2+u2+u2 : u2+u2+u2+u2]
	ac.Count = Byte2U2(bs)
	tc.AttributeCount = ac
	return u2 * 4
}
