package core

// FieldsCount
// u2 fields_count
type FieldsCount struct {
}

// Field
/*
field_info {
u2             access_flags;
u2             name_index;
u2             descriptor_index;
u2             attributes_count;
attribute_info attributes[attributes_count];
}
*/
type Field struct {
}

type Fields []Field
