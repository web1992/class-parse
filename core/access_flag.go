package core

import (
	"sort"
	"strings"
)

// https://github.com/openjdk/jdk/blob/master/src/java.base/share/native/include/classfile_constants.h.template

const (
	JVM_ACC_PUBLIC       = 0x0001
	JVM_ACC_PRIVATE      = 0x0002
	JVM_ACC_PROTECTED    = 0x0004
	JVM_ACC_STATIC       = 0x0008
	JVM_ACC_FINAL        = 0x0010
	JVM_ACC_SYNCHRONIZED = 0x0020
	JVM_ACC_SUPER        = 0x0020
	JVM_ACC_VOLATILE     = 0x0040
	JVM_ACC_BRIDGE       = 0x0040
	JVM_ACC_TRANSIENT    = 0x0080
	JVM_ACC_VARARGS      = 0x0080
	JVM_ACC_NATIVE       = 0x0100
	JVM_ACC_INTERFACE    = 0x0200
	JVM_ACC_ABSTRACT     = 0x0400
	JVM_ACC_STRICT       = 0x0800
	JVM_ACC_SYNTHETIC    = 0x1000
	JVM_ACC_ANNOTATION   = 0x2000
	JVM_ACC_ENUM         = 0x4000
	JVM_ACC_MODULE       = 0x8000
)

var accMap map[int]string

func init() {

	accMap = make(map[int]string)

	accMap[JVM_ACC_PUBLIC] = "ACC_PUBLIC"
	accMap[JVM_ACC_PRIVATE] = "ACC_PRIVATE"
	accMap[JVM_ACC_PROTECTED] = "ACC_PROTECTED"
	accMap[JVM_ACC_STATIC] = "ACC_STATIC"
	accMap[JVM_ACC_FINAL] = "ACC_FINAL"
	accMap[JVM_ACC_SYNCHRONIZED] = "ACC_SYNCHRONIZED"
	accMap[JVM_ACC_SUPER] = "ACC_SUPER"
	accMap[JVM_ACC_VOLATILE] = "ACC_VOLATILE"
	accMap[JVM_ACC_BRIDGE] = "ACC_BRIDGE"
	accMap[JVM_ACC_TRANSIENT] = "ACC_TRANSIENT"
	accMap[JVM_ACC_VARARGS] = "ACC_VARARGS"
	accMap[JVM_ACC_NATIVE] = "ACC_NATIVE"
	accMap[JVM_ACC_INTERFACE] = "ACC_INTERFACE"
	accMap[JVM_ACC_ABSTRACT] = "ACC_ABSTRACT"
	accMap[JVM_ACC_STRICT] = "ACC_STRICT"
	accMap[JVM_ACC_SYNTHETIC] = "ACC_SYNTHETIC"
	accMap[JVM_ACC_ANNOTATION] = "ACC_ANNOTATION"
	accMap[JVM_ACC_ENUM] = "ACC_ENUM"
	accMap[JVM_ACC_MODULE] = "ACC_MODULE"
}

/*
Value	Interpretation
ACC_PUBLIC	0x0001	Declared public; may be accessed from outside its package.
ACC_FINAL	0x0010	Declared final; no subclasses allowed.
ACC_SUPER	0x0020	Treat superclass methods specially when invoked by the invokespecial instruction.
ACC_INTERFACE	0x0200	Is an interface, not a class.
ACC_ABSTRACT	0x0400	Declared abstract; must not be instantiated.
ACC_SYNTHETIC	0x1000	Declared synthetic; not present in the source code.
ACC_ANNOTATION	0x2000	Declared as an annotation type.
ACC_ENUM	0x4000	Declared as an enum type.
*/

// access_flags u2
type AccessFlag struct {
	Flag       int
	FlagString string
	FlagDesc   string
}

func AccessFlagNew() *AccessFlag {
	return &AccessFlag{}
}

func (af *AccessFlag) ReadObj(bytes []byte) int {
	af.Flag = int(Byte2U2(bytes[0:u2]))
	return u2
}

/**
ACC_SUPER
*/
func (af *AccessFlag) HasSuper() bool {
	return af.Flag&JVM_ACC_SUPER != 0
}

func (af *AccessFlag) HasPublic() bool {
	return af.Flag&JVM_ACC_PUBLIC != 0
}

func (af *AccessFlag) HasAbstract() bool {
	return af.Flag&JVM_ACC_ABSTRACT != 0
}
func (af *AccessFlag) String() string {
	return getFlag(af.Flag)
}

// GetFlag get int vale to convert to string
func GetFlag(f AccessFlag) string {
	return getFlag(f.Flag)
}

func GetFlagDesc(f AccessFlag) string {
	return getFlagDesc(f.Flag)
}

func getFlagDesc(f int) string {

	var fs []string
	var keys []int
	for k := range accMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		v := accMap[k]
		if k == JVM_ACC_SUPER {
			continue
		}
		if k&f != 0 {
			s := strings.Replace(v, "ACC_", "", -1)
			cv := strings.ToLower(s)
			fs = append(fs, cv)
		}
	}
	return strings.Join(fs, " ")
}

func getFlag(f int) string {

	var fs []string

	var keys []int
	for k := range accMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		v := accMap[k]
		if k&f != 0 {
			fs = append(fs, v)
		}
	}
	return strings.Join(fs, ", ")
}
