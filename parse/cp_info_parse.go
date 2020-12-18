package parse

import (
	"class-parse/core"
	"fmt"
)

// read constant_pool
// The constant_pool table is indexed from 1 to constant_pool_count - 1.
func (cp *ClassParse) cpInfos(cpc core.ConstantPoolCount) core.CpInfos {

	var cpInfos core.CpInfos
	cpInfos = append(cpInfos, nil)
	view := cpc.View()
	count := view.(int)
	fmt.Printf("constant pool count is %d \n", count)

	for i := 1; i <= count-1; i++ {
		bytes := cp.Bytes()

		//fmt.Println("CpInfos c is", i)
		p := cp.pointer
		tagByte := bytes[p : p+core.U1_L]
		tag := core.Byte2U1(tagByte)
		switch tag {

		case core.TAG_CONSTANT_Class:
			c := core.CpClassNew()
			cp.Read(c)
			cpInfos = append(cpInfos, c)
		case core.TAG_CONSTANT_Fieldref:
			f := core.CpFieldRefNew()
			cp.Read(f)
			cpInfos = append(cpInfos, f)
		case core.TAG_CONSTANT_Methodref:
			m := core.CpMethodRefNew()
			cp.Read(m)
			cpInfos = append(cpInfos, m)
		case core.TAG_CONSTANT_InterfaceMethodref:
			imr := core.CpInterfaceMethodRefNew()
			cp.Read(imr)
			cpInfos = append(cpInfos, imr)
		case core.TAG_CONSTANT_String:
			s := core.CpStringNew()
			cp.Read(s)
			cpInfos = append(cpInfos, s)
		case core.TAG_CONSTANT_Integer:
			i := core.CpIntegerNew()
			cp.Read(i)
			cpInfos = append(cpInfos, i)
		case core.TAG_CONSTANT_Float:
			f := core.CpFloatNew()
			cp.Read(f)
			cpInfos = append(cpInfos, f)
		case core.TAG_CONSTANT_Long:
			l := core.CpLongNew()
			cp.Read(l)
			cpInfos = append(cpInfos, l)
			cpInfos = append(cpInfos, nil)
			i++
		case core.TAG_CONSTANT_Double:
			d := core.CpDoubleNew()
			cp.Read(d)
			cpInfos = append(cpInfos, d)
			cpInfos = append(cpInfos, nil)
			i++
		case core.TAG_CONSTANT_NameAndType:
			var cnat = core.CpNameAndTypeNew()
			cp.Read(cnat)
			cpInfos = append(cpInfos, cnat)
		case core.TAG_CONSTANT_Utf8:
			u := core.CpUTF8New()
			cp.Read(u)
			cpInfos = append(cpInfos, u)
		case core.TAG_CONSTANT_MethodHandle:
			m := core.CpMethodHandleNew()
			cp.Read(m)
			cpInfos = append(cpInfos, m)
		case core.TAG_CONSTANT_MethodType:
			mt := core.CpMethodTypeNew()
			cp.Read(mt)
			cpInfos = append(cpInfos, mt)
		case core.TAG_CONSTANT_InvokeDynamic:
			id := core.CpInvokeDynamicNew()
			cp.Read(id)
			cpInfos = append(cpInfos, id)

		default:
			fmt.Sprintf("tag %d is undefiend", tag)

		}
	}
	return cpInfos
}
