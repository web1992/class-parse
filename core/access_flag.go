package core

import (
	"strings"
)

const (
	ACC_PUBLIC     = 0x0001
	ACC_FINAL      = 0x0010
	ACC_SUPER      = 0x0020
	ACC_INTERFACE  = 0x0200
	ACC_ABSTRACT   = 0x0400
	ACC_SYNTHETIC  = 0x1000
	ACC_ANNOTATION = 0x2000
	ACC_ENUM       = 0x4000
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
	Bytes
}

func AccessFlagNew() *AccessFlag {
	return &AccessFlag{}
}

func (af *AccessFlag) ReadObj(bytes []byte) int {
	af.Bytes = bytes[0:u2]
	return 0
}

func (af *AccessFlag) ObjLen() int {
	return u2
}

// getFlag get int vale to convert to string
func getFlag(f int32) string {

	var fs []string
	if f&ACC_PUBLIC != 0 {
		fs = append(fs, "ACC_PUBLIC")
	}

	if f&ACC_FINAL != 0 {
		fs = append(fs, "ACC_FINAL")
	}

	if f&ACC_SUPER != 0 {
		fs = append(fs, "ACC_SUPER")
	}

	if f&ACC_INTERFACE != 0 {
		fs = append(fs, "ACC_INTERFACE")
	}
	if f&ACC_ABSTRACT != 0 {
		fs = append(fs, "ACC_ABSTRACT")
	}

	if f&ACC_SYNTHETIC != 0 {
		fs = append(fs, "ACC_SYNTHETIC")

	}
	if f&ACC_ANNOTATION != 0 {
		fs = append(fs, "ACC_ANNOTATION")

	}
	if f&ACC_ENUM != 0 {
		fs = append(fs, "ACC_ENUM")
	}

	//s := fmt.Sprintf("flag is unknown %d", f)
	//e := errors.New(s)
	//panic(e)

	return strings.Join(fs, ",")
}
