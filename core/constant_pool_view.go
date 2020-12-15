package core

import (
	"fmt"
	"strings"
)

// View print constant pool
func (cpInfos CpInfos) View() interface{} {

	var views []string

	for i, v := range cpInfos {
		if i == 0 || v == nil {
			continue
		}
		views = append(views, getCp(cpInfos, i))
	}
	return strings.Join(views, "\n")
}

func getCp(cpInfos CpInfos, index int) string {

	_ = cpInfos[index]
	cp := cpInfos[index]

	if m, ok := isCpMethodref(cp); ok {
		return fmt.Sprintf("#%d = Methodref %s ", index, cpInfos.viewCpMethodRef(m, true))
	}

	if c, ok := isCpClass(cp); ok {
		return fmt.Sprintf("#%d = Class %s ", index, cpInfos.viewCpClass(c, true))
	}

	if u, ok := isCpUTF8(cp); ok {
		return fmt.Sprintf("#%d = Utf8 %s", index, string(u.String))
	}

	if c, ok := isCpNameAndType(cp); ok {
		return fmt.Sprintf("#%d = NameAndType %s ", index, cpInfos.viewCpNameAndType(c, true))
	}

	if f, ok := isCpFieldRef(cp); ok {
		return fmt.Sprintf("#%d = Fieldref %s ", index, cpInfos.viewCpFieldref(f))
	}

	if s, ok := isCpString(cp); ok {
		return fmt.Sprintf("#%d = String %s ", index, cpInfos.viewCpString(s))
	}

	if id, ok := isCpInvokeDynamic(cp); ok {
		return fmt.Sprintf("#%d = InvokeDynamic %s ", index, cpInfos.viewCpInvokeDynamic(id))
	}

	if im, ok := isCpInterfaceMethodRef(cp); ok {
		return fmt.Sprintf("#%d = InterfaceMethodref %s ", index, cpInfos.viewCpInterfaceMethodRef(im))
	}

	if ok := isNum(cp); ok {
		return cpInfos.viewNum(cp, index)
	}

	if mh, ok := isCpMethodHandle(cp); ok {
		return fmt.Sprintf("#%d = MethodHandle %s ", index, cpInfos.viewMethodHandle(mh))
	}

	if mh, ok := isCpMethodType(cp); ok {
		return fmt.Sprintf("#%d = MethodType %s ", index, cpInfos.viewMethodType(mh))
	}

	return fmt.Sprintf("#%d = ", index)

}

func isCpMethodType(e interface{}) (*CpMethodType, bool) {
	mt, ok := e.(*CpMethodType)
	return mt, ok
}

func isCpMethodHandle(e interface{}) (*CpMethodHandle, bool) {
	mh, ok := e.(*CpMethodHandle)
	return mh, ok
}

//#129 = MethodType         #6            //  ()V
func (cpInfos CpInfos) viewMethodType(mt interface{}) string {

	if m, ok := isCpMethodType(mt); ok {
		di := m.DescriptorIndex
		cp := cpInfos[di]
		if uu, ok := isCpUTF8(cp); ok {
			return fmt.Sprintf("#%d %s", di, string(uu.String))
		}
	}

	return ""
}

//#130 = MethodHandle       6:#131        // REF_invokeStatic Main.lambda$main$0:()V
func (cpInfos CpInfos) viewMethodHandle(mh interface{}) string {

	if m, ok := isCpMethodHandle(mh); ok {
		rk := m.ReferenceKind
		ri := m.ReferenceIndex
		cp := cpInfos[ri]
		rks := getReferenceKind(int32(rk))
		methodRef := cpInfos.viewCpMethodRef(cp, false)

		return fmt.Sprintf("%d:%d // %s %s", rk, ri, rks, methodRef)
	}

	return ""
}
func (cpInfos CpInfos) viewCpInterfaceMethodRef(im interface{}) string {
	if im, ok := isCpInterfaceMethodRef(im); ok {

		ci := im.ClassIndex
		nati := im.NameAndTypeIndex

		cpClass := cpInfos[ci]
		cpNAti := cpInfos[nati]

		s := fmt.Sprintf("#%d.#%d %s.%s",
			int32(ci),
			int32(nati),
			cpInfos.viewCpClass(cpClass, false),
			cpInfos.viewCpNameAndType(cpNAti, false))

		return s
	}
	return ""
}
func (cpInfos CpInfos) viewCpInvokeDynamic(cpInvokeDynamic interface{}) string {
	if id, ok := isCpInvokeDynamic(cpInvokeDynamic); ok {
		_ = id.BootstrapMethodAttrIndex
		nati := id.NameAndTypeIndex

		cpNAti := cpInfos[nati]

		s := fmt.Sprintf("#%d:#%d %s:%s",
			0,
			int32(nati),
			"0",
			cpInfos.viewCpNameAndType(cpNAti, false))

		return s
	}

	return ""
}
func (cpInfos CpInfos) viewCpMethodRef(cpMethodRef interface{}, showIndex bool) string {

	if m, ok := isCpMethodref(cpMethodRef); ok {
		ci := m.ClassIndex
		ti := m.NameAndTypeIndex

		cpClass := cpInfos[ci]
		cpNAti := cpInfos[ti]

		if showIndex {
			return fmt.Sprintf("#%d.#%d %s.%s",
				int32(ci),
				int32(ti),
				cpInfos.viewCpClass(cpClass, false),
				cpInfos.viewCpNameAndType(cpNAti, false))
		}
		return fmt.Sprintf("%s.%s",
			cpInfos.viewCpClass(cpClass, false),
			cpInfos.viewCpNameAndType(cpNAti, false))
	}

	return ""
}

func (cpInfos CpInfos) viewCpFieldref(cpFieldref interface{}) string {

	if f, ok := isCpFieldRef(cpFieldref); ok {
		ci := f.ClassIndex
		nati := f.NameAndTypeIndex

		cpClass := cpInfos[ci]
		cpNAti := cpInfos[nati]

		s := fmt.Sprintf("#%d.#%d %s.%s",
			int32(ci),
			int32(nati),
			cpInfos.viewCpClass(cpClass, false),
			cpInfos.viewCpNameAndType(cpNAti, false))
		return s
	}
	return ""
}

func (cpInfos CpInfos) viewCpClass(cpClass interface{}, showIndex bool) string {
	if c, ok := isCpClass(cpClass); ok {
		ni := c.NameIndex
		u := cpInfos[ni]
		if uu, ok := isCpUTF8(u); ok {
			if showIndex {
				return fmt.Sprintf("#%d %s", ni, string(uu.String))
			}
			return string(uu.String)
		}
	}
	return ""
}

func (cpInfos CpInfos) viewCpNameAndType(cpNameAndType interface{}, showIndex bool) string {
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
		if showIndex {
			return fmt.Sprintf("#%d:#%d %v:%v", ni, di, ds, us)
		}
		s := fmt.Sprintf("%v:%v", ds, us)
		return s
	}
	return ""
}

func (cpInfos CpInfos) viewCpString(cpString interface{}) string {
	if s, ok := isCpString(cpString); ok {
		si := s.StringIndex
		ss := cpInfos[si]

		if uu, ok := isCpUTF8(ss); ok {
			return fmt.Sprintf("%s", string(uu.String))
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
func (cpInfos CpInfos) viewCpInteger(cpString interface{}) int32 {
	if i, ok := cpString.(*CpInteger); ok {
		return int32(i.Integer)
	}

	return -1
}

// Long
// #60 = Long               -9223372036854775808l
func isCpLong(e interface{}) (*CpLong, bool) {
	i, ok := e.(*CpLong)
	return i, ok
}

func (cpInfos CpInfos) viewCpLong(e interface{}) int64 {
	if i, ok := e.(*CpLong); ok {
		return int64(i.Long)
	}
	return -1
}

// Double
// #67 = Double             1.7976931348623157E308d

func isCpDouble(e interface{}) (*CpDouble, bool) {
	i, ok := e.(*CpDouble)
	return i, ok
}

func (cpInfos CpInfos) viewCpDouble(e interface{}) float64 {
	if i, ok := e.(*CpDouble); ok {
		return float64(i.Double)
	}
	return -1
}

// Float
// #91 = Float              1.4E-45f
func isCpFloat(e interface{}) (*CpFloat, bool) {
	i, ok := e.(*CpFloat)
	return i, ok
}

func (cpInfos CpInfos) viewCpFloat(e interface{}) float32 {
	if i, ok := e.(*CpFloat); ok {
		return float32(i.Float)
	}
	return -1
}

func (cpInfos CpInfos) viewNum(e interface{}, index int) string {
	if i, ok := isCpInteger(e); ok {
		// #34 = Integer
		return fmt.Sprintf("#%d = Integer %v", index, i.Integer)
	}

	if l, ok := isCpLong(e); ok {
		return fmt.Sprintf("#%d = Long %v", index, l.Long)
	}

	if d, ok := isCpDouble(e); ok {
		return fmt.Sprintf("#%d = Double %v", index, d.Double)
	}

	if f, ok := isCpFloat(e); ok {
		return fmt.Sprintf("#%d = Float %v", index, f.Float)
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
