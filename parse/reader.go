package parse

type Reader interface {
	ReadObj(bytes []byte) int
	ObjLen() int
}
