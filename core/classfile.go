package core

import (
	"fmt"
	"strings"
)

const (
	U1 = 1
	U2 = 2
	U4 = 4
)

/**
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type ClassFile interface {
}

type ClassPart interface {
	ByteLen() int
}

type PartName interface {
	Name() string
}

type PartShow interface {
	View() interface{}
}

type ByteAble struct {
	Bytes []byte
}

type Magic struct {
	ByteAble
}

type MinorVersion struct {
	ByteAble
}

type MajorVersion struct {
	ByteAble
}

type ConstantPoolCount struct {
	ByteAble
}

func (m *Magic) ByteLen() int {
	return U4
}

func (m *Magic) View() interface{} {
	magicStr := fmt.Sprintf("%x", m.Bytes)
	mx := strings.ToUpper(magicStr)
	return mx
}

func (mv *MinorVersion) ByteLen() int {
	return U2
}

func (m *MinorVersion) View() interface{} {
	// 0xFF = 1111 1111
	b := m.Bytes
	n := ((b[0] & 0xFF) << 8) + ((b[1]) << 0)
	return int(n)
}

func (mv *MajorVersion) ByteLen() int {
	return U2
}

func (m *MajorVersion) View() interface{} {
	// 0xFF = 1111 1111
	b := m.Bytes
	n := ((b[0] & 0xFF) << 8) + ((b[1]) << 0)
	return int(n)
}

func (cp *ConstantPoolCount) ByteLen() int {
	return U2
}

// always stored in big-endian order, where the high bytes come first
func (cp *ConstantPoolCount) View() interface{} {
	// 0xFF = 1111 1111
	b := cp.Bytes
	n := ((b[0] & 0xFF) << 8) + ((b[1]) << 0)
	return int(n)
}
