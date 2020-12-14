package core

import (
	"fmt"
	"strings"
)

// Magic class file Magic num
type Magic struct {
	Bytes
}

func (m *Magic) ByteLen() int {
	return U4_L
}

func (m *Magic) View() interface{} {
	magicStr := fmt.Sprintf("%x", m.Bytes)
	mx := strings.ToUpper(magicStr)
	return mx
}

func MagicNew() *Magic {
	return &Magic{}
}

func (magic *Magic) ReadObj(bytes []byte) int {
	magic.Bytes = bytes[0:u4]
	return 0
}

func (magic *Magic) ObjLen() int {
	return u4
}
