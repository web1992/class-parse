package core

// u2             methods_count;
// method_info    methods[methods_count];
type MethodCount struct {
	Bytes
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

type Methods []Method

func MethodCountNew() *MethodCount {
	return &MethodCount{}
}

func (tc *MethodCount) ReadObj(bytes []byte) int {
	bs := bytes[0:u2]
	tc.Bytes = bs
	tc.Count = Byte2U2(bs)
	return 0
}

func (tc *MethodCount) ObjLen() int {
	return u2
}

func MethodNew() *Method {
	return &Method{}
}

func (tc *Method) ReadObj(bytes []byte) int {
	tc.Bytes = bytes

	var af AccessFlag
	af.ReadObj(bytes[0:u2])

	tc.AccessFlag = af
	tc.NameIndex = NameIndex(Byte2U2(bytes[u2 : u2+u2]))
	tc.DescriptorIndex = DescriptorIndex(Byte2U2(bytes[u2+u2 : u2+u2+u2]))

	var ac AttributeCount
	bs := bytes[u2+u2+u2 : u2+u2+u2+u2]
	ac.Count = Byte2U2(bs)
	ac.Bytes = bs
	tc.AttributeCount = ac
	return 0
}

func (tc *Method) ObjLen() int {
	return u2 * 4
}
