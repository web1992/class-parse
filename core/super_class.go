package core

const JavaLangObject = "java/lang/Object"

// SuperClass u2  super_class
type SuperClass struct {
	ClassIndex
	String string
}

func SuperClassNew() *SuperClass {
	return &SuperClass{}
}
func (sc SuperClass) isJavaLangObject() bool {
	return JavaLangObject == sc.String
}
func (sc *SuperClass) ReadObj(bytes []byte) int {
	bs := bytes[0:u2]
	sc.ClassIndex = ClassIndex(Byte2U2(bs))
	return u2
}
