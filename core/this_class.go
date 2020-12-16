package core

// ThisClass u2 this_class
type ThisClass struct {
	Bytes
}

func ThisClassNew() *ThisClass {
	return &ThisClass{}
}

func (tc *ThisClass) ReadObj(bytes []byte) int {
	tc.Bytes = bytes[0:u2]
	return 0
}

func (tc *ThisClass) ObjLen() int {
	return u2
}
