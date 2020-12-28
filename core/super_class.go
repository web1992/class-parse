package core

// SuperClass u2  super_class
type SuperClass struct {
	Bytes
	ClassIndex
	String string
}

func SuperClassNew() *SuperClass {
	return &SuperClass{}
}

func (tc *SuperClass) ReadObj(bytes []byte) int {
	bs := bytes[0:u2]
	tc.Bytes = bs
	tc.ClassIndex = ClassIndex(Byte2U2(bs))
	return u2
}
