package core

type AttributeCount struct {
}

/*
Attribute

attribute_info {
u2 attribute_name_index;
u4 attribute_length;
u1 info[attribute_length];
}
*/
type Attribute struct {
}

type Attributes []Attribute
