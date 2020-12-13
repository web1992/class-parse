package parse

type Reader interface {
	ReadObj(bytes []byte)
	ObjLen() int
}
