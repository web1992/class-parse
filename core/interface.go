package core

// Interface u2  interfaces[interfaces_count]
type Interface struct {
	ClassIndex
	NameString string
}

// Interface u2  interfaces[interfaces_count]
type Interfaces []Interface

func (tc *Interface) ReadObj(bytes []byte) int {
	bs := bytes[0:u2]
	tc.ClassIndex = ClassIndex(Byte2U2(bs))
	return u2
}

// InterfacesCount u2 interfaces_count
type InterfacesCount struct {
	Count int32
}

func InterfacesCountNew() *InterfacesCount {
	return &InterfacesCount{}
}

func (tc *InterfacesCount) ReadObj(bytes []byte) int {
	bs := bytes[0:u2]
	tc.Count = Byte2U2(bs)
	return u2
}
