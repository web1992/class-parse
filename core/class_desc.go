package core

import (
	"fmt"
	"strings"
)

/**
public class Main extends AbstractMain<java.lang.String> implements InterfaceMain

*/
func (cf ClassFile) ClassDesc() string {

	// public class Main extends AbstractMain implements InterfaceMain
	var str []string
	str = append(str, cf.ClassNameDesc())
	str = append(str, NewLine)
	str = append(str, cf.VersionDesc())
	str = append(str, NewLine)
	str = append(str, cf.FlagDesc())
	str = append(str, NewLine)
	str = append(str, cf.ConstantPool())
	str = append(str, NewLine)
	str = append(str, "{")
	str = append(str, NewLine)
	str = append(str, cf.FieldDesc())
	str = append(str, NewLine)
	str = append(str, cf.MethodDesc())
	str = append(str, NewLine)
	str = append(str, "}")
	str = append(str, NewLine)
	str = append(str, cf.AttrDesc())
	return strings.Join(str, "")
}
func (cf ClassFile) AttrDesc() string {
	var str []string

	ac := int(cf.AttributeCount.Count)
	if ac > 0 {
		for _, a := range cf.Attributes {

			if attr, ok := a.(fmt.Stringer); ok {
				str = append(str, attr.String())
			}
		}
	}
	return strings.Join(str, NewLine)
}

/**

  public Main();
    descriptor: ()V
    flags: ACC_PUBLIC
    Code:
      stack=1, locals=1, args_size=1
         0: aload_0
         1: invokespecial #1                  // Method AbstractMain."<init>":()V
         4: return
      LineNumberTable:
        line 9: 0
*/
func (cf ClassFile) MethodDesc() string {
	var str []string

	mc := int(cf.MethodCount.Count)
	if mc > 0 {
		for _, m := range cf.Methods {
			md := getMDesc(m)
			str = append(str, md)
		}
	}
	return strings.Join(str, NewLine)
}

func getMDesc(m Method) string {
	var str []string

	s1 := GetFlagDesc(m.AccessFlag)
	str = append(str, fmt.Sprintf("%s%s %s();", GetSpace(2), s1, m.NameString))

	s2 := fmt.Sprintf("%sdescriptor: %s", GetSpace(6), m.DescriptorString)
	str = append(str, s2)
	s3 := fmt.Sprintf("%sflags: %s", GetSpace(6), m.AccessFlag.String())
	str = append(str, s3)
	ac := int(m.AttributeCount.Count)
	if ac > 0 {
		str = append(str, fmt.Sprintf(GetSpace(6)+"Code:"))
		for _, v := range m.Attributes {
			if s, ok := v.(fmt.Stringer); ok {
				str = append(str, fmt.Sprintf("%s%s", GetSpace(8), s.String()))
			}
		}
	}

	return strings.Join(str, NewLine) + NewLine
}

/**
  public static final java.lang.Integer INT_MAX;
    descriptor: Ljava/lang/Integer;
    flags: ACC_PUBLIC, ACC_STATIC, ACC_FINAL
    RuntimeVisibleAnnotations:
      0: #65()
*/
func (cf ClassFile) FieldDesc() string {
	var str []string
	fc := int(cf.FieldsCount.Count)
	if fc > 0 {
		for _, f := range cf.Fields {
			fd := getFDesc(f)
			str = append(str, fd)
		}
	}
	return strings.Join(str, NewLine)
}

func GetSpace(num int) string {
	var str []string

	for i := 0; i < num; i++ {
		str = append(str, OneSpace)
	}
	return strings.Join(str, "")
}

func getFDesc(f Field) string {
	var str []string

	s1 := GetFlagDesc(f.AccessFlag)
	s2 := f.DescriptorString
	s3 := f.NameString
	s4 := fmt.Sprintf("%s%s %s %s;", GetSpace(2), s1, s2, s3)
	str = append(str, s4)

	s5 := fmt.Sprintf("%sdescriptor: %s", GetSpace(4), s2)
	str = append(str, s5)

	s6 := fmt.Sprintf("%sflags: %s", GetSpace(4), f.AccessFlag.String())
	str = append(str, s6)

	ac := int(f.AttributeCount.Count)

	if ac > 0 {
		for _, v := range f.Attributes {
			if s, ok := v.(fmt.Stringer); ok {
				str = append(str, fmt.Sprintf("%s%s", GetSpace(4), s.String()))
			}
		}
	}

	return strings.Join(str, NewLine) + NewLine

}

func (cf ClassFile) ConstantPool() string {
	var str []string
	str = append(str, fmt.Sprintf("%s", "Constant pool:"+NewLine))
	str = append(str, cf.CpInfos.String())
	return strings.Join(str, "")
}
func (cf ClassFile) FlagDesc() string {
	return fmt.Sprintf("%s %s", "  flags:", cf.AccessFlag.FlagString)
}

func (cf ClassFile) VersionDesc() string {
	var str []string
	v1 := cf.MinorVersion.Version
	v2 := cf.MajorVersion.Version

	vs1 := fmt.Sprintf("%15s: %d", _minor_version, v1)
	vs2 := fmt.Sprintf("%15s: %d", _major_version, v2)
	str = append(str, vs1)
	str = append(str, vs2)
	return strings.Join(str, NewLine)

}

func (cf ClassFile) ClassNameDesc() string {
	var str []string
	str = append(str, cf.ThisDesc())
	str = append(str, cf.SuperDesc())
	str = append(str, cf.InterfaceDesc())
	return strings.Join(str, " ")
}

func (cf ClassFile) ThisDesc() string {
	var str []string
	afDesc := GetFlagDesc(cf.AccessFlag)
	str = append(str, afDesc)

	str = append(str, _class)
	str = append(str, cf.ThisClass.String)
	return strings.Join(str, OneSpace)
}

func (cf ClassFile) SuperDesc() string {
	var str []string
	if cf.AccessFlag.HasSuper() {
		if !cf.SuperClass.isJavaLangObject() {
			str = append(str, _extends)
			str = append(str, cf.SuperClass.String)
		}
	}
	return strings.Join(str, OneSpace)
}

func (cf ClassFile) InterfaceDesc() string {
	var str []string
	ifc := cf.InterfacesCount.Count
	var interfaceDesc string
	if ifc > 0 {
		str = append(str, _implements)
		var s []string
		for _, v := range cf.Interfaces {
			s = append(s, v.NameString)
		}
		interfaceDesc = strings.Join(s, P)
		str = append(str, interfaceDesc)
	}

	return strings.Join(str, OneSpace)
}
