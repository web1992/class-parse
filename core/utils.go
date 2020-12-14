package core

import (
	"fmt"
	"strings"
)

func HexByte(b []byte) Hex {
	s := fmt.Sprintf("%x", b)
	return Hex(strings.ToUpper(s))
}
