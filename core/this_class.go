package core

// ThisClass u2 this_class
type ThisClass struct {
	Bytes
	ClassIndex
	String string
}

func ThisClassNew() *ThisClass {
	return &ThisClass{}
}

func (tc *ThisClass) ReadObj(bytes []byte) int {
	bs := bytes[0:u2]
	tc.Bytes = bs
	tc.ClassIndex = ClassIndex(Byte2U2(bs))
	return 0
}

func (tc *ThisClass) ObjLen() int {
	return u2
}
