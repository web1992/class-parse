package core

const (
	ConstantValue    = "ConstantValue"
	Code             = "Code"
	StackMapTable    = "StackMapTable"
	Exceptions       = "Exceptions"
	BootstrapMethods = "BootstrapMethods"

	InnerClasses                         = "InnerClasses"
	EnclosingMethod                      = "EnclosingMethod"
	Synthetic                            = "Synthetic"
	Signature                            = "Signature"
	RuntimeVisibleAnnotations            = "RuntimeVisibleAnnotations"
	RuntimeInvisibleAnnotations          = "RuntimeInvisibleAnnotations"
	RuntimeVisibleParameterAnnotations   = "RuntimeVisibleParameterAnnotations"
	RuntimeInvisibleParameterAnnotations = "RuntimeInvisibleParameterAnnotations"
	RuntimeVisibleTypeAnnotations        = "RuntimeVisibleTypeAnnotations"
	RuntimeInvisibleTypeAnnotations      = "RuntimeInvisibleTypeAnnotations"
	AnnotationDefault                    = "AnnotationDefault"
	MethodParameters                     = "MethodParameters"

	SourceFile             = "SourceFile"
	SourceDebugExtension   = "SourceDebugExtension"
	LineNumberTable        = "LineNumberTable"
	LocalVariableTable     = "LocalVariableTable"
	LocalVariableTypeTable = "LocalVariableTypeTable"
	Deprecated             = "Deprecated"
)

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

end
*/
func CreateAttribute(name string) interface{} {

	if name == SourceFile {
		var sfa SourceFileAttribute

		return &sfa
	}

	if name == BootstrapMethods {
		var bmt BootstrapMethodsAttribute

		return &bmt
	}

	if name == InnerClasses {
		var sfa InnerClassesAttribute

		return &sfa
	}

	attr := AttributeNew()
	return attr
}
