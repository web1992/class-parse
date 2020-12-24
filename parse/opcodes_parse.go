package parse

import (
	"fmt"
	"goclass/core"
	"strings"
)

func parseOpCodes(pointer int, codeLength int, bs []byte) core.OpCodes {

	var ops core.OpCodes
	hadReadLen := 0
	for hadReadLen < codeLength {

		op := core.Byte2U1(bs[hadReadLen : hadReadLen+core.U1_L])
		_bs := bs[hadReadLen:]
		desc := core.GetOpDesc(int(op))
		opObj := core.CreateOpCode(op)
		if o, ok := opObj.(*core.OpCodeTableSwitch); ok {
			o.Offset = pointer - codeLength + hadReadLen
			o.Base = int32(hadReadLen)
			i := o.ReadObj(_bs)
			fmt.Printf("%d: %s \n", hadReadLen, getTableSwitchDesc(*o, desc))
			hadReadLen = i + hadReadLen
			ops = append(ops, o)
			continue
		}
		if o, ok := opObj.(*core.OpCodeLookupSwitch); ok {
			o.Offset = pointer - codeLength + hadReadLen
			o.Base = int32(hadReadLen)
			i := o.ReadObj(_bs)
			fmt.Printf("%d: %s \n", hadReadLen, getLookupSwitchDesc(*o, desc))
			hadReadLen = i + hadReadLen
			ops = append(ops, o)
			continue
		}
		if o, ok := opObj.(core.Reader); ok {
			i := o.ReadObj(_bs)
			//fmt.Printf("read len %d %d: %s \n", i, hadReadLen, desc)
			fmt.Printf("%d: %s \n", hadReadLen, desc)
			hadReadLen = i + hadReadLen
			ops = append(ops, o)
			continue
		} else {
			panic(fmt.Sprintf("Error opObj %v", opObj))
		}
	}

	return ops
}

func getTableSwitchDesc(ts core.OpCodeTableSwitch, desc string) string {

	var s []string
	s = append(s, fmt.Sprintf("%s { // %d-%d", desc, ts.Low, ts.High))

	for _, v := range ts.Pairs {
		if !v.Default {
			s = append(s, fmt.Sprintf("%16v:%v", v.Case, v.LineNo))
		}
	}

	for _, v := range ts.Pairs {
		if v.Default {
			s = append(s, fmt.Sprintf("%16v:%v", "default", v.LineNo))
		}
	}
	s = append(s, "}")

	return strings.Join(s, "\n")
}
func getLookupSwitchDesc(lsw core.OpCodeLookupSwitch, desc string) string {

	var s []string
	s = append(s, fmt.Sprintf("%s { // %d", desc, len(lsw.Pairs)-1))

	for _, v := range lsw.Pairs {
		if !v.Default {
			s = append(s, fmt.Sprintf("%16v:%v", v.Case, v.LineNo))
		}
	}

	for _, v := range lsw.Pairs {
		if v.Default {
			s = append(s, fmt.Sprintf("%16v:%v", "default", v.LineNo))
		}
	}
	s = append(s, "}")

	return strings.Join(s, "\n")
}
