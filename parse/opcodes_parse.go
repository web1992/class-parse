package parse

import (
	"fmt"
	"goclass/core"
)

func parseOpCodes(pointer int, codeLength int, bs []byte) core.OpCodes {

	var ops core.OpCodes
	hadReadLen := 0
	for hadReadLen < codeLength {

		op := core.Byte2U1(bs[hadReadLen : hadReadLen+core.U1_L])
		_bs := bs[hadReadLen:]
		//desc := core.GetOpDesc(int(op))
		opObj := core.CreateOpCode(op)
		if o, ok := opObj.(*core.OpCodeTableSwitch); ok {
			o.Offset = pointer - codeLength + hadReadLen
			o.Base = int32(hadReadLen)
			o.LineNo = hadReadLen
			readLen := o.ReadObj(_bs)
			//fmt.Printf("%d: %s \n", hadReadLen, core.GetTableSwitchDesc(*o, desc))
			hadReadLen = readLen + hadReadLen
			ops = append(ops, o)
			continue
		}
		if o, ok := opObj.(*core.OpCodeLookupSwitch); ok {
			o.Offset = pointer - codeLength + hadReadLen
			o.Base = int32(hadReadLen)
			o.LineNo = hadReadLen
			readLen := o.ReadObj(_bs)
			//fmt.Printf("%d: %s \n", hadReadLen, core.GetLookupSwitchDesc(*o, desc))
			hadReadLen = readLen + hadReadLen
			ops = append(ops, o)
			continue
		}
		if o, ok := opObj.(core.Reader); ok {
			readLen := o.ReadObj(_bs)
			//fmt.Printf("%d: %s \n", hadReadLen, desc)

			if opc, ok := opObj.(core.OpCoder); ok {
				opc.SetLineNo(hadReadLen)
			}

			hadReadLen = readLen + hadReadLen
			ops = append(ops, o)
			continue
		} else {
			panic(fmt.Sprintf("Error opObj %v", opObj))
		}
	}

	return ops
}
