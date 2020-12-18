package core

import (
	"fmt"
	"strings"
)

const F = "%6s = %-18s%-14s// %s"

// View print constant pool
func (cpInfos CpInfos) View() interface{} {

	var views []string

	for i, v := range cpInfos {
		if i == 0 || v == nil {
			continue
		}
		views = append(views, GetCpView(cpInfos, i))
	}
	return strings.Join(views, "\n")
}

func GetCp(cpInfos CpInfos, index int) string {

	_ = cpInfos[index]
	cp := cpInfos[index]

	if m, ok := isCpMethodref(cp); ok {
		return cpInfos.getCpMethodRef(m)
	}

	if c, ok := isCpClass(cp); ok {
		return cpInfos.getCpClass(c)
	}

	if u, ok := isCpUTF8(cp); ok {
		return string(u.String)
	}

	if c, ok := isCpNameAndType(cp); ok {
		return cpInfos.getCpNameAndType(c)
	}

	if f, ok := isCpFieldRef(cp); ok {
		return cpInfos.getCpFieldRef(f)
	}

	if s, ok := isCpString(cp); ok {
		return cpInfos.getCpString(s)
	}

	if id, ok := isCpInvokeDynamic(cp); ok {
		return cpInfos.getCpInvokeDynamic(id)
	}

	if im, ok := isCpInterfaceMethodRef(cp); ok {
		return cpInfos.getCpInterfaceMethodRef(im)
	}

	if ok := isNum(cp); ok {
		return cpInfos.getNum(cp)
	}

	if mh, ok := isCpMethodHandle(cp); ok {
		return cpInfos.getMethodHandle(mh)
	}

	if mh, ok := isCpMethodType(cp); ok {
		return cpInfos.getMethodType(mh)
	}

	return fmt.Sprintf("#%d = ", index)
}

func FmtM(a interface{}, b interface{}) string {
	return fmt.Sprintf("#%d.#%d", a, b)
}
func Fmt2(a interface{}, b interface{}) string {
	return fmt.Sprintf("#%d:#%d", a, b)
}
func Fmt22(a interface{}, b interface{}) string {
	return fmt.Sprintf("%d.#%d", a, b)
}

func Fmt1(ci interface{}) string {
	return fmt.Sprintf("#%d", ci)
}

func Fmt11(ci interface{}) string {
	return fmt.Sprintf("#0:%d", ci)
}

func FmtIndex(i interface{}) string {
	return fmt.Sprintf("#%d", i)
}
func GetCpView(cpInfos CpInfos, index int) string {

	_ = cpInfos[index]
	cp := cpInfos[index]

	if m, ok := isCpMethodref(cp); ok {
		return fmt.Sprintf(F, FmtIndex(index), "Methodref", FmtM(m.ClassIndex, m.NameAndTypeIndex), cpInfos.getCpMethodRef(m))
	}

	if c, ok := isCpClass(cp); ok {
		return fmt.Sprintf(F, FmtIndex(index), "Class", Fmt1(c.NameIndex), cpInfos.getCpClass(c))
	}

	if u, ok := isCpUTF8(cp); ok {
		return fmt.Sprintf("%6s = %-18s %-s", FmtIndex(index), "Utf8 ", string(u.String))
	}

	if c, ok := isCpNameAndType(cp); ok {
		return fmt.Sprintf(F, FmtIndex(index), "NameAndType", Fmt2(c.NameIndex, c.DescriptorIndex), cpInfos.getCpNameAndType(c))
	}

	if f, ok := isCpFieldRef(cp); ok {
		return fmt.Sprintf(F, FmtIndex(index), "Fieldref", FmtM(f.ClassIndex, f.NameAndTypeIndex), cpInfos.getCpFieldRef(f))
	}

	if s, ok := isCpString(cp); ok {
		return fmt.Sprintf(F, FmtIndex(index), "String", Fmt1(s.StringIndex), cpInfos.getCpString(s))
	}

	if id, ok := isCpInvokeDynamic(cp); ok {
		return fmt.Sprintf(F, FmtIndex(index), "InvokeDynamic", Fmt11(id.NameAndTypeIndex), cpInfos.getCpInvokeDynamic(id))
	}

	if im, ok := isCpInterfaceMethodRef(cp); ok {
		return fmt.Sprintf(F, FmtIndex(index), "InterfaceMethodref", FmtM(im.ClassIndex, im.NameAndTypeIndex), cpInfos.getCpInterfaceMethodRef(im))
	}

	if ok := isNum(cp); ok {
		return cpInfos.viewNum(cp, index)
	}

	if mh, ok := isCpMethodHandle(cp); ok {
		return fmt.Sprintf(F, FmtIndex(index), "MethodHandle", Fmt22(mh.ReferenceKind, mh.ReferenceIndex), cpInfos.getMethodHandle(mh))
	}

	if mh, ok := isCpMethodType(cp); ok {
		return fmt.Sprintf(F, FmtIndex(index), "MethodType", Fmt1(mh.DescriptorIndex), cpInfos.getMethodType(mh))
	}

	return fmt.Sprintf("#%d = ", FmtIndex(index))

}

func isCpMethodType(e interface{}) (*CpMethodType, bool) {
	mt, ok := e.(*CpMethodType)
	return mt, ok
}

func isCpMethodHandle(e interface{}) (*CpMethodHandle, bool) {
	mh, ok := e.(*CpMethodHandle)
	return mh, ok
}

func (cpInfos CpInfos) getMethodType(mt interface{}) string {

	if m, ok := isCpMethodType(mt); ok {
		di := m.DescriptorIndex
		cp := cpInfos[di]
		if uu, ok := isCpUTF8(cp); ok {
			return string(uu.String)
		}
	}

	return ""
}

//#129 = MethodType         #6            //  ()V
func (cpInfos CpInfos) getMethodHandle(mh interface{}) string {

	if m, ok := isCpMethodHandle(mh); ok {
		rk := m.ReferenceKind
		ri := m.ReferenceIndex
		cp := cpInfos[ri]
		rks := getReferenceKind(int32(rk))
		methodRef := cpInfos.getCpMethodRef(cp)

		return fmt.Sprintf("%s %s", rks, methodRef)
	}

	return ""
}

func (cpInfos CpInfos) getCpInterfaceMethodRef(im interface{}) string {
	if im, ok := isCpInterfaceMethodRef(im); ok {

		ci := im.ClassIndex
		nati := im.NameAndTypeIndex

		cpClass := cpInfos[ci]
		cpNAti := cpInfos[nati]

		s := fmt.Sprintf("%s.%s",
			cpInfos.getCpClass(cpClass),
			cpInfos.getCpNameAndType(cpNAti))

		return s
	}
	return ""
}

func (cpInfos CpInfos) getCpInvokeDynamic(cpInvokeDynamic interface{}) string {
	if id, ok := isCpInvokeDynamic(cpInvokeDynamic); ok {
		_ = id.BootstrapMethodAttrIndex
		nati := id.NameAndTypeIndex
		cpNAti := cpInfos[nati]
		s := fmt.Sprintf("%s:%s",
			"0",
			cpInfos.getCpNameAndType(cpNAti))
		return s
	}

	return ""
}

func (cpInfos CpInfos) getCpMethodRef(cpMethodRef interface{}) string {

	if m, ok := isCpMethodref(cpMethodRef); ok {
		ci := m.ClassIndex
		ti := m.NameAndTypeIndex
		cpClass := cpInfos[ci]
		cpNAti := cpInfos[ti]

		return fmt.Sprintf("%s.%s",
			cpInfos.getCpClass(cpClass),
			cpInfos.getCpNameAndType(cpNAti))
	}

	return ""
}

func (cpInfos CpInfos) getCpFieldRef(cpFieldref interface{}) string {

	if f, ok := isCpFieldRef(cpFieldref); ok {
		ci := f.ClassIndex
		nati := f.NameAndTypeIndex

		cpClass := cpInfos[ci]
		cpNAti := cpInfos[nati]

		s := fmt.Sprintf("%s.%s",
			cpInfos.getCpClass(cpClass),
			cpInfos.getCpNameAndType(cpNAti))
		return s
	}
	return ""
}

func (cpInfos CpInfos) getCpClass(cpClass interface{}) string {
	if c, ok := isCpClass(cpClass); ok {
		ni := c.NameIndex
		u := cpInfos[ni]
		if uu, ok := isCpUTF8(u); ok {
			return string(uu.String)
		}
	}
	return ""
}

func (cpInfos CpInfos) getCpNameAndType(cpNameAndType interface{}) string {
	if c, ok := isCpNameAndType(cpNameAndType); ok {
		ni := c.NameIndex
		di := c.DescriptorIndex
		u1 := cpInfos[ni]

		ds := ""
		us := ""
		if uu, ok := isCpUTF8(u1); ok {
			ds = string(uu.String)
		}

		u2 := cpInfos[di]

		if uu, ok := isCpUTF8(u2); ok {
			us = string(uu.String)
		}

		s := fmt.Sprintf("%v:%v", ds, us)
		return s
	}
	return ""
}

func (cpInfos CpInfos) getCpString(cpString interface{}) string {
	if s, ok := isCpString(cpString); ok {
		si := s.StringIndex
		ss := cpInfos[si]

		if uu, ok := isCpUTF8(ss); ok {
			return string(uu.String)
		}

	}

	return ""
}

func isCpNameAndType(e interface{}) (*CpNameAndType, bool) {
	c, ok := e.(*CpNameAndType)
	return c, ok
}

func isCpClass(e interface{}) (*CpClass, bool) {
	c, ok := e.(*CpClass)
	return c, ok
}

func isCpUTF8(e interface{}) (*CpUTF8, bool) {
	u, ok := e.(*CpUTF8)
	return u, ok
}

func isCpMethodref(e interface{}) (*CpMethodRef, bool) {
	m, ok := e.(*CpMethodRef)
	return m, ok
}

func isCpFieldRef(e interface{}) (*CpFieldRef, bool) {
	f, ok := e.(*CpFieldRef)
	return f, ok
}

func isCpString(e interface{}) (*CpString, bool) {
	s, ok := e.(*CpString)
	return s, ok
}

// CpInvokeDynamic
func isCpInvokeDynamic(e interface{}) (*CpInvokeDynamic, bool) {
	id, ok := e.(*CpInvokeDynamic)
	return id, ok
}

// CpInterfaceMethodRef
func isCpInterfaceMethodRef(e interface{}) (*CpInterfaceMethodRef, bool) {
	im, ok := e.(*CpInterfaceMethodRef)
	return im, ok
}

// CpInteger
// #34 = Integer            2147483647
func isCpInteger(e interface{}) (*CpInteger, bool) {
	i, ok := e.(*CpInteger)
	return i, ok
}

// Long
// #60 = Long               -9223372036854775808l
func isCpLong(e interface{}) (*CpLong, bool) {
	i, ok := e.(*CpLong)
	return i, ok
}

// Double
// #67 = Double             1.7976931348623157E308d

func isCpDouble(e interface{}) (*CpDouble, bool) {
	i, ok := e.(*CpDouble)
	return i, ok
}

// Float
// #91 = Float              1.4E-45f
func isCpFloat(e interface{}) (*CpFloat, bool) {
	i, ok := e.(*CpFloat)
	return i, ok
}

func (cpInfos CpInfos) getNum(e interface{}) string {
	if i, ok := isCpInteger(e); ok {
		// #34 = Integer
		return fmt.Sprintf("%v", i.Integer)
	}

	if l, ok := isCpLong(e); ok {
		return fmt.Sprintf("%v", l.Long)
	}

	if d, ok := isCpDouble(e); ok {
		return fmt.Sprintf("%v", d.Double)
	}

	if f, ok := isCpFloat(e); ok {
		return fmt.Sprintf("%v", f.Float)
	}
	return ""
}

func (cpInfos CpInfos) viewNum(e interface{}, index int) string {
	if i, ok := isCpInteger(e); ok {
		// #34 = Integer
		return fmt.Sprintf("%6s = %-18s %-18v", FmtIndex(index), "Integer", i.Integer)
	}

	if l, ok := isCpLong(e); ok {
		return fmt.Sprintf("%6s = %-18s %-18v", FmtIndex(index), "Long", l.Long)
	}

	if d, ok := isCpDouble(e); ok {
		return fmt.Sprintf("%6s = %-18s %-18v", FmtIndex(index), "Double", d.Double)
	}

	if f, ok := isCpFloat(e); ok {
		return fmt.Sprintf("%6s = %-18s %-18v", FmtIndex(index), "Float", f.Float)
	}
	return ""
}

func isNum(e interface{}) bool {

	if _, ok := isCpInteger(e); ok {
		return ok
	}

	if _, ok := isCpLong(e); ok {
		return ok
	}

	if _, ok := isCpDouble(e); ok {
		return ok
	}

	if _, ok := isCpFloat(e); ok {
		return ok
	}
	return false
}
