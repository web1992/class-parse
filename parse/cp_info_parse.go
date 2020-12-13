package parse

import (
	"class-parse/core"
	"class-parse/utils"
	"fmt"
)

const (
	u1 = core.U1_L
	u2 = core.U2_L
)

// read constant_pool
// The constant_pool table is indexed from 1 to constant_pool_count - 1.
func (cp *ClassParse) CpInfos() core.CpInfos {

	var cpInfos core.CpInfos
	cpc := cp.ConstantPoolCount()
	count := (cpc.View()).(int)
	fmt.Printf("constant pool count is %d \n", count)

	for i := 1; i <= count-1; i++ {
		bytes := cp.Bytes()

		p := cp.pointer
		tagByte := bytes[p : p+core.U1_L]
		tag := utils.U1(tagByte)
		//fmt.Println(tag)

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
		//case core.TAG_CONSTANT_InterfaceMethodref:
		//	cpInfos = append(cpInfos, parseCpInterfaceMethodRef())
		//case core.TAG_CONSTANT_String:
		//	cpInfos = append(cpInfos, parseCpString())
		//case core.TAG_CONSTANT_Integer:
		//	cpInfos = append(cpInfos, parseCpInter())
		//case core.TAG_CONSTANT_Float:
		//	cpInfos = append(cpInfos, parseCpFloat())
		//case core.TAG_CONSTANT_Long:
		//	cpInfos = append(cpInfos, parseCpLong())
		//case core.TAG_CONSTANT_Double:
		//	cpInfos = append(cpInfos, parseCpDouble())
		case core.TAG_CONSTANT_NameAndType:
			var cnat = core.CpNameAndTypeNew()
			cp.Read(cnat)
			cpInfos = append(cpInfos, cnat)
		//case core.TAG_CONSTANT_Utf8:
		//	cpInfos = append(cpInfos, parseCpUTF8())
		//case core.TAG_CONSTANT_MethodHandle:
		//	cpInfos = append(cpInfos, parseCpMethodHandle())
		//case core.TAG_CONSTANT_MethodType:
		//	cpInfos = append(cpInfos, parseCpMethodType())
		//case core.TAG_CONSTANT_InvokeDynamic:
		//	cpInfos = append(cpInfos, parseCpInvokeDynamic())

		default:
			//fmt.Sprintf("tag %d is undefiend", tag)

		}
	}
	return cpInfos
}

/**
/*
tag =12
CONSTANT_NameAndType_info {
u1 tag;
u2 name_index;
u2 descriptor_index;
}*/

func parseCpNameAndType(cp *ClassParse) core.CpNameAndType {

	var cpnt core.CpNameAndType

	bytes := cp.Bytes()
	p := cp.pointer
	b := bytes[p : p+u1+u2+u2]

	cpnt.Tag = core.Tag(utils.U1(b[:u1]))
	cpnt.NameIndex = core.NameIndex(utils.U2(b[u1 : u1+u2]))
	cpnt.DescriptorIndex = core.DescriptorIndex(utils.U2(b[u1+u2 : u1+u2+u2]))

	cp.pointer += u1 + u2 + u2
	return cpnt
}

/**
parseCpMethodRef

tag =10
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
func parseCpMethodRef(cp *ClassParse) core.CpMethodRef {
	var cpM core.CpMethodRef

	bytes := cp.Bytes()
	p := cp.pointer
	b := bytes[p : p+u1+u2+u2]

	tag := utils.U1(b[:u1])
	cpM.Tag = core.Tag(tag)

	ci := utils.U2(b[u1 : u1+u2])
	cpM.ClassIndex = core.ClassIndex(ci)

	ni := utils.U2(b[u1+u2 : u1+u2+u2])
	cpM.NameAndTypeIndex = core.NameAndTypeIndex(ni)

	cp.pointer += u1 + u2 + u2
	return cpM
}

/**
tag =7
CONSTANT_Class_info {
u1 tag;
u2 name_index;
}
*/
func parseCpClass(cp *ClassParse) core.CpClass {
	var cpClass core.CpClass

	bytes := cp.Bytes()
	p := cp.pointer
	b := bytes[p : p+u1+u2]

	tag := utils.U1(b[:u1])
	cpClass.Tag = core.Tag(tag)

	ni := utils.U2(b[u1 : u1+u2])
	cpClass.NameIndex = core.NameIndex(ni)

	cp.pointer += u1 + u2
	return cpClass
}
