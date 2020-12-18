package parse

import "class-parse/core"

func (cp *ClassParse) attributeCount() core.AttributeCount {

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
func (cp *ClassParse) attributes(cpInfos core.CpInfos, attributeCount core.AttributeCount) core.Attributes {

	var attrs core.Attributes
	c := int(attributeCount.Count)
	for i := 0; i < c; i++ {
		attr := core.AttributeNew()
		cp.Read(attr)
		attr.Name = core.GetCp(cpInfos, int(attr.AttributeNameIndex))
		attrs = append(attrs, attr)

	}
	return attrs
}

/**

Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}

*/
