package parse

import (
	"fmt"
	"goclass/core"
)

func parseOpCodes(codeLength int, bs []byte) core.OpCodes {

	var ops core.OpCodes
	hadReadLen := 0
	for hadReadLen < codeLength {

		op := core.Byte2U1(bs[hadReadLen : hadReadLen+core.U1_L])
		_bs := bs[hadReadLen:]
		desc := core.GetOpDesc(int(op))
		opObj := core.CreateOpCode(op)
		if o, ok := opObj.(core.Reader); ok {
			i := o.ReadObj(_bs)
			//fmt.Printf("read len %d %d: %s \n", i, hadReadLen, desc)
			fmt.Printf("%d: %s \n", hadReadLen, desc)
			hadReadLen = i + hadReadLen
			ops = append(ops, o)
		}
	}

	return ops
}
