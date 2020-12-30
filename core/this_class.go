package core

// ThisClass u2 this_class
type ThisClass struct {
	ClassIndex
	String string
}

func ThisClassNew() *ThisClass {
	return &ThisClass{}
}

func (tc *ThisClass) ReadObj(bytes []byte) int {
	bs := bytes[0:u2]
	tc.ClassIndex = ClassIndex(Byte2U2(bs))
	return u2
}
