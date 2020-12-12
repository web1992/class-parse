package core

type CpReader interface {
	ReadMethod() CpMethodRef
	ReadClass() CpClassInfo
}

/**
tag =10
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/

type CpMethodRef struct {
}

/**
tag =7
CONSTANT_Class_info {
u1 tag;
u2 name_index;
}
*/

type CpClassInfo struct {
}
