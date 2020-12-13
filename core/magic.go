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
	return U4
}

func (m *Magic) View() interface{} {
	magicStr := fmt.Sprintf("%x", m.Bytes)
	mx := strings.ToUpper(magicStr)
	return mx
}
