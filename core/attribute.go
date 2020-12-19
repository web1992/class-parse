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
func AttributeNew() *Attribute {
	return &Attribute{}
}

// Attribute as default attribute_info
// Must read bytes from file,change the pointer
// make other bytes can read success
type Attribute struct {
	Bytes
	AttributeNameIndex int32
	AttributeLength    int32
	Name               string
}

func (af *Attribute) ReadObj(bytes []byte) int {
	i := Byte2U2(bytes[0:u2])
	af.AttributeNameIndex = i

	l := Byte2U4(bytes[u2 : u2+u4])
	af.AttributeLength = l

	//af.Bytes = bytes[u2+u4: u2+u4+l]
	// just ignore left bytes
	return int(l)
}

func (af *Attribute) ObjLen() int {
	return u2 + u4
}

//u2             attributes_count;
//attribute_info attributes[attributes_count];
type AttributeCount struct {
	Bytes
	Count int32
}

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
