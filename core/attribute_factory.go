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
<p>
Deprecated
</p>

end
*/
func CreateAttribute(name string, cpInfos CpInfos, pointer int) interface{} {

	switch name {

	case StackMapTable:
		return &StackMapTableAttribute{
			Attribute: Attribute{
				CpInfos: cpInfos,
				Name:    name,
			},
		}
	case SourceFile:
		return &SourceFileAttribute{
			AttributeName: SourceFile,
			Attribute: Attribute{
				CpInfos: cpInfos,
				Name:    name,
			},
		}
	case BootstrapMethods:
		return &BootstrapMethodsAttribute{
			Name: BootstrapMethods,
			Attribute: Attribute{
				CpInfos: cpInfos,
				Name:    name,
			}}
	case InnerClasses:
		return &InnerClassesAttribute{
			Name: InnerClasses,
			Attribute: Attribute{
				CpInfos: cpInfos,
				Name:    name,
			}}
	case Code:
		return &CodeAttribute{
			Name: Code,
			Attribute: Attribute{
				CpInfos: cpInfos,
				Offset:  pointer,
				Name:    name,
			}}
	case Exceptions:
		return &ExceptionsAttribute{
			Name: Exceptions,
			Attribute: Attribute{
				CpInfos: cpInfos,
				Name:    name,
			}}
	case LineNumberTable:
		return &LineNumberTableAttribute{
			Name: LineNumberTable,
			Attribute: Attribute{
				CpInfos: cpInfos,
				Name:    name,
			},
		}
	case RuntimeVisibleAnnotations:
		return &RuntimeVisibleAnnotationsAttr{
			Attribute: Attribute{
				CpInfos: cpInfos,
				Name:    name,
			},
		}
	case Signature:
		return &SignatureAttribute{
			Attribute: Attribute{
				CpInfos: cpInfos,
				Name:    name,
			},
		}
	default:
		return AttributeNew(name, cpInfos)
	}
}

func CreateAttributeByIndex(nameIndex int, cpInfos CpInfos, pointer int) interface{} {
	_ = cpInfos[nameIndex]
	name := GetCp(cpInfos, nameIndex)
	return CreateAttribute(name, cpInfos, pointer)
}
