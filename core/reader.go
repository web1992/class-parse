package core

type Reader interface {
	ReadObj(bytes []byte) int
}
