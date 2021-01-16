package core

import (
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
	if f&JVM_ACC_PUBLIC != 0 {
		fs = append(fs, "public")
	}

	if f&JVM_ACC_FINAL != 0 {
		fs = append(fs, "final")
	}

	//if f&JVM_ACC_SUPER != 0 {
	//	fs = append(fs, "ACC_SUPER")
	//}

	if f&JVM_ACC_INTERFACE != 0 {
		fs = append(fs, "interface")
	}
	if f&JVM_ACC_ABSTRACT != 0 {
		fs = append(fs, "abstract")
	}

	//if f&JVM_ACC_SYNTHETIC != 0 {
	//	fs = append(fs, "ACC_SYNTHETIC")
	//}

	if f&JVM_ACC_ANNOTATION != 0 {
		fs = append(fs, "annotation")
	}
	if f&JVM_ACC_ENUM != 0 {
		fs = append(fs, "enum")
	}
	return strings.Join(fs, " ")
}

func getFlag(f int) string {

	var fs []string
	if f&JVM_ACC_PUBLIC != 0 {
		fs = append(fs, "ACC_PUBLIC")
	}

	if f&JVM_ACC_FINAL != 0 {
		fs = append(fs, "ACC_FINAL")
	}

	if f&JVM_ACC_SUPER != 0 {
		fs = append(fs, "ACC_SUPER")
	}

	if f&JVM_ACC_INTERFACE != 0 {
		fs = append(fs, "ACC_INTERFACE")
	}
	if f&JVM_ACC_ABSTRACT != 0 {
		fs = append(fs, "ACC_ABSTRACT")
	}

	if f&JVM_ACC_SYNTHETIC != 0 {
		fs = append(fs, "ACC_SYNTHETIC")

	}
	if f&JVM_ACC_ANNOTATION != 0 {
		fs = append(fs, "ACC_ANNOTATION")

	}
	if f&JVM_ACC_ENUM != 0 {
		fs = append(fs, "ACC_ENUM")
	}

	//s := fmt.Sprintf("flag is unknown %d", f)
	//e := errors.New(s)
	//panic(e)

	return strings.Join(fs, ",")
}
