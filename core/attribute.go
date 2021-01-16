package core

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7

// Attributes an array of attribute
type Attributes []interface{}

/*
Attribute

attribute_info {
u2 attribute_name_index;
u4 attribute_length;
u1 info[attribute_length];
}
*/
func AttributeNew(name string, cpInfo CpInfos) *Attribute {
	return &Attribute{
		Name:    name,
		CpInfos: cpInfo,
	}
}

// Attribute as default attribute_info
// Must read bytes from file,change the pointer
// make other bytes can read success
type Attribute struct {
	CpInfos
	Offset             int
	Name               string
	AttributeNameIndex int32
	AttributeLength    int32
}

func (af *Attribute) String() string {
	return af.Name
}

func (af *Attribute) ReadObj(bytes []byte) int {
	i := Byte2U2(bytes[0:u2])
	af.AttributeNameIndex = i

	l := Byte2U4(bytes[u2 : u2+u4])
	af.AttributeLength = l

	//af.Bytes = bytes[u2+u4: u2+u4+l]
	// just ignore left bytes
	return u2 + u4 + int(l)
}

//u2             attributes_count;
//attribute_info attributes[attributes_count];
type AttributeCount struct {
	Count int32
}

func AttributeCountNew() *AttributeCount {
	return &AttributeCount{}
}

func (af *AttributeCount) ReadObj(bytes []byte) int {
	bs := bytes[0:u2]
	af.Count = Byte2U2(bs)
	return u2
}
