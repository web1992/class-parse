package core

// https://github.com/openjdk/jdk/blob/master/src/java.base/share/native/include/classfile_constants.h.template

const (
	//TAG_CONSTANT_Class              = 7
	//TAG_CONSTANT_Fieldref           = 9
	//TAG_CONSTANT_Methodref          = 10
	//TAG_CONSTANT_InterfaceMethodref = 11
	//TAG_CONSTANT_String             = 8
	//TAG_CONSTANT_Integer            = 3
	//TAG_CONSTANT_Float              = 4
	//TAG_CONSTANT_Long               = 5
	//TAG_CONSTANT_Double             = 6
	//TAG_CONSTANT_NameAndType        = 12
	//TAG_CONSTANT_Utf8               = 1
	//TAG_CONSTANT_MethodHandle       = 15
	//TAG_CONSTANT_MethodType         = 16
	//TAG_CONSTANT_InvokeDynamic      = 18

	JVM_CONSTANT_Utf8               = 1
	JVM_CONSTANT_Unicode            = 2 /* unused */
	JVM_CONSTANT_Integer            = 3
	JVM_CONSTANT_Float              = 4
	JVM_CONSTANT_Long               = 5
	JVM_CONSTANT_Double             = 6
	JVM_CONSTANT_Class              = 7
	JVM_CONSTANT_String             = 8
	JVM_CONSTANT_Fieldref           = 9
	JVM_CONSTANT_Methodref          = 10
	JVM_CONSTANT_InterfaceMethodref = 11
	JVM_CONSTANT_NameAndType        = 12
	JVM_CONSTANT_MethodHandle       = 15 // JSR 292
	JVM_CONSTANT_MethodType         = 16 // JSR 292
	JVM_CONSTANT_Dynamic            = 17
	JVM_CONSTANT_InvokeDynamic      = 18
	JVM_CONSTANT_Module             = 19
	JVM_CONSTANT_Package            = 20
	JVM_CONSTANT_ExternalMax        = 20
)
