package parse

import (
	"github.com/web1992/goclass/core"
)

func (cp *ClassReader) attributeCount() core.AttributeCount {

	ac := core.AttributeCountNew()
	cp.Read(ac)

	return *ac
}

// 23 attributes are predefined by this specification. They are listed three times, for ease of navigation:

/*
5 + 12 +6

// 5
// Five attributes are critical to correct interpretation of the class file by the Java Virtual Machine:
ConstantValue
Code
StackMapTable
Exceptions
BootstrapMethods

// 12
// Twelve attributes are critical to correct interpretation of the class file by the class libraries of the Java SE platform:
InnerClasses
EnclosingMethod
Synthetic
Signature
RuntimeVisibleAnnotations
RuntimeInvisibleAnnotations
RuntimeVisibleParameterAnnotations
RuntimeInvisibleParameterAnnotations
RuntimeVisibleTypeAnnotations
RuntimeInvisibleTypeAnnotations
AnnotationDefault
MethodParameters

// 6
// Six attributes are not critical to correct interpretation of the class file
// by either the Java Virtual Machine or the class libraries of the Java SE platform, but are useful for tools:

SourceFile
SourceDebugExtension
LineNumberTable
LocalVariableTable
LocalVariableTypeTable
Deprecated

*/

/*
attribute_info {
u2 attribute_name_index;
u4 attribute_length;
u1 info[attribute_length];
}*/
func (cp *ClassReader) attributes(cpInfos core.CpInfos, attributeCount core.AttributeCount) core.Attributes {

	var attrs core.Attributes
	c := int(attributeCount.Count)
	for i := 0; i < c; i++ {

		bytes := cp.bytes
		p := cp.pointer
		tagByte := bytes[p : p+core.U2]

		attributeNameIndex := core.Byte2U2(tagByte)
		name := core.GetCp(cpInfos, int(attributeNameIndex))
		ca := core.CreateAttribute(name, cpInfos, cp.pointer)

		attr := cp.parseAttr(ca, cpInfos, name)
		attrs = append(attrs, attr)

	}
	return attrs
}

func (cp *ClassReader) parseAttr(ca interface{}, cpInfos core.CpInfos, name string) interface{} {

	if reader, ok := ca.(core.Reader); ok {
		cp.Read(reader)
		return reader
	} else {
		// default
		attr := core.AttributeNew(name, cpInfos)
		cp.Read(attr)
		attr.Name = name
		core.Warning.Println("skip attribute name: " + name)
		return attr
	}
}
