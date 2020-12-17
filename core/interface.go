package core

// Interface u2  interfaces[interfaces_count]
type Interface struct {
	Bytes
	ClassIndex
	String string
}

// Interface u2  interfaces[interfaces_count]
type Interfaces []Interface

func (tc *Interface) ReadObj(bytes []byte) int {
	bs := bytes[0:u2]
	tc.Bytes = bs
	tc.ClassIndex = ClassIndex(U2(bs))
	return 0
}

func (tc *Interface) ObjLen() int {
	return u2
}

// InterfacesCount u2 interfaces_count
type InterfacesCount struct {
	Bytes
	Count int32
}

func InterfacesCountNew() *InterfacesCount {
	return &InterfacesCount{}
}

func (tc *InterfacesCount) ReadObj(bytes []byte) int {
	bs := bytes[0:u2]
	tc.Bytes = bs
	tc.Count = U2(bs)
	return 0
}

func (tc *InterfacesCount) ObjLen() int {
	return u2
}
