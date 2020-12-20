package parse

import "goclass/core"

func (cp *ClassParse) parseOpCodes(codeLength int, bs []byte) core.OpCodes {

	var ops core.OpCodes
	for i := 1; i <= codeLength; i++ {
		// 182-186
		op := core.Byte2U1(bs[0:core.U1_L])
		if int(op) >= 182 || int(op) <= 186 {

		}
	}

	return ops
}
