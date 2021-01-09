package core

import (
	"fmt"
	"strings"
)

const (

	// Reserved Opcodes
	impdep1    = 254 // 0xfe
	impdep2    = 255 // 0xff
	breakpoint = 202 // 0xca

	JVM_OPC_nop             = 0
	JVM_OPC_aconst_null     = 1
	JVM_OPC_iconst_m1       = 2
	JVM_OPC_iconst_0        = 3
	JVM_OPC_iconst_1        = 4
	JVM_OPC_iconst_2        = 5
	JVM_OPC_iconst_3        = 6
	JVM_OPC_iconst_4        = 7
	JVM_OPC_iconst_5        = 8
	JVM_OPC_lconst_0        = 9
	JVM_OPC_lconst_1        = 10
	JVM_OPC_fconst_0        = 11
	JVM_OPC_fconst_1        = 12
	JVM_OPC_fconst_2        = 13
	JVM_OPC_dconst_0        = 14
	JVM_OPC_dconst_1        = 15
	JVM_OPC_bipush          = 16
	JVM_OPC_sipush          = 17
	JVM_OPC_ldc             = 18
	JVM_OPC_ldc_w           = 19
	JVM_OPC_ldc2_w          = 20
	JVM_OPC_iload           = 21
	JVM_OPC_lload           = 22
	JVM_OPC_fload           = 23
	JVM_OPC_dload           = 24
	JVM_OPC_aload           = 25
	JVM_OPC_iload_0         = 26
	JVM_OPC_iload_1         = 27
	JVM_OPC_iload_2         = 28
	JVM_OPC_iload_3         = 29
	JVM_OPC_lload_0         = 30
	JVM_OPC_lload_1         = 31
	JVM_OPC_lload_2         = 32
	JVM_OPC_lload_3         = 33
	JVM_OPC_fload_0         = 34
	JVM_OPC_fload_1         = 35
	JVM_OPC_fload_2         = 36
	JVM_OPC_fload_3         = 37
	JVM_OPC_dload_0         = 38
	JVM_OPC_dload_1         = 39
	JVM_OPC_dload_2         = 40
	JVM_OPC_dload_3         = 41
	JVM_OPC_aload_0         = 42
	JVM_OPC_aload_1         = 43
	JVM_OPC_aload_2         = 44
	JVM_OPC_aload_3         = 45
	JVM_OPC_iaload          = 46
	JVM_OPC_laload          = 47
	JVM_OPC_faload          = 48
	JVM_OPC_daload          = 49
	JVM_OPC_aaload          = 50
	JVM_OPC_baload          = 51
	JVM_OPC_caload          = 52
	JVM_OPC_saload          = 53
	JVM_OPC_istore          = 54
	JVM_OPC_lstore          = 55
	JVM_OPC_fstore          = 56
	JVM_OPC_dstore          = 57
	JVM_OPC_astore          = 58
	JVM_OPC_istore_0        = 59
	JVM_OPC_istore_1        = 60
	JVM_OPC_istore_2        = 61
	JVM_OPC_istore_3        = 62
	JVM_OPC_lstore_0        = 63
	JVM_OPC_lstore_1        = 64
	JVM_OPC_lstore_2        = 65
	JVM_OPC_lstore_3        = 66
	JVM_OPC_fstore_0        = 67
	JVM_OPC_fstore_1        = 68
	JVM_OPC_fstore_2        = 69
	JVM_OPC_fstore_3        = 70
	JVM_OPC_dstore_0        = 71
	JVM_OPC_dstore_1        = 72
	JVM_OPC_dstore_2        = 73
	JVM_OPC_dstore_3        = 74
	JVM_OPC_astore_0        = 75
	JVM_OPC_astore_1        = 76
	JVM_OPC_astore_2        = 77
	JVM_OPC_astore_3        = 78
	JVM_OPC_iastore         = 79
	JVM_OPC_lastore         = 80
	JVM_OPC_fastore         = 81
	JVM_OPC_dastore         = 82
	JVM_OPC_aastore         = 83
	JVM_OPC_bastore         = 84
	JVM_OPC_castore         = 85
	JVM_OPC_sastore         = 86
	JVM_OPC_pop             = 87
	JVM_OPC_pop2            = 88
	JVM_OPC_dup             = 89
	JVM_OPC_dup_x1          = 90
	JVM_OPC_dup_x2          = 91
	JVM_OPC_dup2            = 92
	JVM_OPC_dup2_x1         = 93
	JVM_OPC_dup2_x2         = 94
	JVM_OPC_swap            = 95
	JVM_OPC_iadd            = 96
	JVM_OPC_ladd            = 97
	JVM_OPC_fadd            = 98
	JVM_OPC_dadd            = 99
	JVM_OPC_isub            = 100
	JVM_OPC_lsub            = 101
	JVM_OPC_fsub            = 102
	JVM_OPC_dsub            = 103
	JVM_OPC_imul            = 104
	JVM_OPC_lmul            = 105
	JVM_OPC_fmul            = 106
	JVM_OPC_dmul            = 107
	JVM_OPC_idiv            = 108
	JVM_OPC_ldiv            = 109
	JVM_OPC_fdiv            = 110
	JVM_OPC_ddiv            = 111
	JVM_OPC_irem            = 112
	JVM_OPC_lrem            = 113
	JVM_OPC_frem            = 114
	JVM_OPC_drem            = 115
	JVM_OPC_ineg            = 116
	JVM_OPC_lneg            = 117
	JVM_OPC_fneg            = 118
	JVM_OPC_dneg            = 119
	JVM_OPC_ishl            = 120
	JVM_OPC_lshl            = 121
	JVM_OPC_ishr            = 122
	JVM_OPC_lshr            = 123
	JVM_OPC_iushr           = 124
	JVM_OPC_lushr           = 125
	JVM_OPC_iand            = 126
	JVM_OPC_land            = 127
	JVM_OPC_ior             = 128
	JVM_OPC_lor             = 129
	JVM_OPC_ixor            = 130
	JVM_OPC_lxor            = 131
	JVM_OPC_iinc            = 132
	JVM_OPC_i2l             = 133
	JVM_OPC_i2f             = 134
	JVM_OPC_i2d             = 135
	JVM_OPC_l2i             = 136
	JVM_OPC_l2f             = 137
	JVM_OPC_l2d             = 138
	JVM_OPC_f2i             = 139
	JVM_OPC_f2l             = 140
	JVM_OPC_f2d             = 141
	JVM_OPC_d2i             = 142
	JVM_OPC_d2l             = 143
	JVM_OPC_d2f             = 144
	JVM_OPC_i2b             = 145
	JVM_OPC_i2c             = 146
	JVM_OPC_i2s             = 147
	JVM_OPC_lcmp            = 148
	JVM_OPC_fcmpl           = 149
	JVM_OPC_fcmpg           = 150
	JVM_OPC_dcmpl           = 151
	JVM_OPC_dcmpg           = 152
	JVM_OPC_ifeq            = 153
	JVM_OPC_ifne            = 154
	JVM_OPC_iflt            = 155
	JVM_OPC_ifge            = 156
	JVM_OPC_ifgt            = 157
	JVM_OPC_ifle            = 158
	JVM_OPC_if_icmpeq       = 159
	JVM_OPC_if_icmpne       = 160
	JVM_OPC_if_icmplt       = 161
	JVM_OPC_if_icmpge       = 162
	JVM_OPC_if_icmpgt       = 163
	JVM_OPC_if_icmple       = 164
	JVM_OPC_if_acmpeq       = 165
	JVM_OPC_if_acmpne       = 166
	JVM_OPC_goto            = 167
	JVM_OPC_jsr             = 168
	JVM_OPC_ret             = 169
	JVM_OPC_tableswitch     = 170
	JVM_OPC_lookupswitch    = 171
	JVM_OPC_ireturn         = 172
	JVM_OPC_lreturn         = 173
	JVM_OPC_freturn         = 174
	JVM_OPC_dreturn         = 175
	JVM_OPC_areturn         = 176
	JVM_OPC_return          = 177
	JVM_OPC_getstatic       = 178
	JVM_OPC_putstatic       = 179
	JVM_OPC_getfield        = 180
	JVM_OPC_putfield        = 181
	JVM_OPC_invokevirtual   = 182
	JVM_OPC_invokespecial   = 183
	JVM_OPC_invokestatic    = 184
	JVM_OPC_invokeinterface = 185
	JVM_OPC_invokedynamic   = 186
	JVM_OPC_new             = 187
	JVM_OPC_newarray        = 188
	JVM_OPC_anewarray       = 189
	JVM_OPC_arraylength     = 190
	JVM_OPC_athrow          = 191
	JVM_OPC_checkcast       = 192
	JVM_OPC_instanceof      = 193
	JVM_OPC_monitorenter    = 194
	JVM_OPC_monitorexit     = 195
	JVM_OPC_wide            = 196
	JVM_OPC_multianewarray  = 197
	JVM_OPC_ifnull          = 198
	JVM_OPC_ifnonnull       = 199
	JVM_OPC_goto_w          = 200
	JVM_OPC_jsr_w           = 201
)

type OpCodes []interface{}

type OpCoder interface {
	SetLineNo(lineNo int)
}

// OpCode 1 byte
type OpCode struct {
	LineNo int
	Desc   string
	Opc    int32
	Args   []int32
}

func (op *OpCode) String() string {
	desc := GetOpDesc(int(op.Opc))
	return fmt.Sprintf("%d: %s \n", op.LineNo, desc)
}
func (oc *OpCode) SetLineNo(lineNo int) {
	oc.LineNo = lineNo
}

// OpCode2 2 bytes
type OpCode2 struct {
	OpCode
}

// OpCode3 3 bytes
type OpCode3 struct {
	OpCode
}

// 4 bytes
type OpCode4 struct {
	OpCode
}

// 5 bytes
type OpCode5 struct {
	OpCode
}

type OpCodeJsrW struct {
	OpCode
}

type OpCodeLookupSwitch struct {
	OpCode
	Base  int32
	Pairs []Pair
}

type Pair struct {
	Default bool
	Case    int32
	Offset  int32
	LineNo  int32
}

type OpCodeTableSwitch struct {
	OpCode
	Offset int
	Base   int32
	Low    int32
	High   int32
	Pairs  []Pair
}

func (op *OpCodeTableSwitch) String() string {
	desc := GetOpDesc(int(op.Opc))
	return fmt.Sprintf("%d: %s \n", op.LineNo, GetTableSwitchDesc(*op, desc))

}

func (op *OpCodeTableSwitch) ReadObj(bytes []byte) int {
	readLen := 0
	op.Opc = Byte2U1(bytes[0:u1])
	readLen += u1
	op.Desc = GetOpDesc(int(op.Opc))
	var pad = int(op.Base-1) % 4
	readLen += pad
	defaultOffset := Byte2U4(bytes[readLen : readLen+u4])
	readLen += u4

	var p Pair
	p.Default = true
	p.Offset = defaultOffset
	p.LineNo = p.Offset + op.Base
	op.Pairs = append(op.Pairs, p)

	low := Byte2U4(bytes[readLen : readLen+u4])
	readLen += u4
	high := Byte2U4(bytes[readLen : readLen+u4])
	readLen += u4
	op.Low = low
	op.High = high
	targetsLength := int((high - low) + 1)

	for i := 0; i < targetsLength; i++ {
		var p Pair
		p.Case = low
		p.Offset = Byte2U4(bytes[readLen : readLen+u4])
		readLen += u4
		p.LineNo = p.Offset + op.Base
		op.Pairs = append(op.Pairs, p)
		low = low + 1
	}

	return readLen
}

/**
26: tableswitch   { // 1 to 3
                 1: 52
                 2: 55
                 3: 58
           default: 61
   }
*/
func GetTableSwitchDesc(ts OpCodeTableSwitch, desc string) string {

	var s []string
	s = append(s, fmt.Sprintf("%s { // %d-%d", desc, ts.Low, ts.High))

	for _, v := range ts.Pairs {
		if !v.Default {
			s = append(s, fmt.Sprintf("%16v:%v", v.Case, v.LineNo))
		}
	}

	for _, v := range ts.Pairs {
		if v.Default {
			s = append(s, fmt.Sprintf("%16v:%v", "default", v.LineNo))
		}
	}
	s = append(s, "}")

	return strings.Join(s, "\n")
}

func (op *OpCodeLookupSwitch) ReadObj(bytes []byte) int {
	readLen := 0
	op.Opc = Byte2U1(bytes[readLen:u1])
	readLen += u1
	op.Desc = GetOpDesc(int(op.Opc))
	var pad = int(op.Base-1) % 4
	readLen += pad
	defaultOffset := Byte2U4(bytes[readLen : readLen+u4])
	readLen += u4

	var p Pair
	p.Default = true
	p.LineNo = defaultOffset + op.Base
	p.Offset = defaultOffset
	op.Pairs = append(op.Pairs, p)
	npairsLen := int(Byte2U4(bytes[readLen : readLen+u4]))
	readLen += u4

	for i := 0; i < npairsLen; i++ {
		var p Pair
		p.Default = false
		p.Case = Byte2U4(bytes[readLen : readLen+u4])
		readLen += u4
		p.Offset = Byte2U4(bytes[readLen : readLen+u4])
		readLen += u4
		p.LineNo = op.Base + p.Offset
		op.Pairs = append(op.Pairs, p)
	}
	return readLen
}

func (op *OpCodeLookupSwitch) String() string {
	desc := GetOpDesc(int(op.Opc))
	return fmt.Sprintf("%d: %s \n", op.LineNo, GetLookupSwitchDesc(*op, desc))
}

func GetLookupSwitchDesc(lsw OpCodeLookupSwitch, desc string) string {

	var s []string
	s = append(s, fmt.Sprintf("%s { // %d", desc, len(lsw.Pairs)-1))

	for _, v := range lsw.Pairs {
		if !v.Default {
			s = append(s, fmt.Sprintf("%16v:%v", v.Case, v.LineNo))
		}
	}

	for _, v := range lsw.Pairs {
		if v.Default {
			s = append(s, fmt.Sprintf("%16v:%v", "default", v.LineNo))
		}
	}
	s = append(s, "}")

	return strings.Join(s, "\n")
}

func (op *OpCode) ReadObj(bytes []byte) int {
	op.Opc = Byte2U1(bytes[0:u1])
	op.Desc = GetOpDesc(int(op.Opc))
	return u1
}

func (op *OpCode2) ReadObj(bytes []byte) int {
	op.Opc = Byte2U1(bytes[0:u1])
	op.Desc = GetOpDesc(int(op.Opc))
	op.Args = append(op.Args, Byte2U1(bytes[u1:u1+u1]))
	return u1 + u1
}

func (op *OpCode3) ReadObj(bytes []byte) int {
	op.Opc = Byte2U1(bytes[0:u1])
	op.Desc = GetOpDesc(int(op.Opc))
	op.Args = append(op.Args, Byte2U2(bytes[u1:u1+u2]))
	return u1 + u2
}

func (op *OpCode4) ReadObj(bytes []byte) int {
	op.Opc = Byte2U1(bytes[0:u1])
	op.Desc = GetOpDesc(int(op.Opc))
	op.Args = append(op.Args, Byte2U2(bytes[u1:u1+u2]))
	op.Args = append(op.Args, Byte2U1(bytes[u1+u2:u1+u2+u1]))
	return u1 + u2 + u1
}

func (op *OpCode5) ReadObj(bytes []byte) int {
	op.Opc = Byte2U1(bytes[0:u1])
	op.Desc = GetOpDesc(int(op.Opc))
	op.Args = append(op.Args, Byte2U2(bytes[u1:u1+u2]))
	op.Args = append(op.Args, Byte2U1(bytes[u1+u2:u1+u2+u1]))
	op.Args = append(op.Args, Byte2U1(bytes[u1+u2+u1:u1+u2+u1+u1]))
	return u1 * 5
}

func (op *OpCodeJsrW) ReadObj(bytes []byte) int {
	op.Opc = Byte2U1(bytes[0:u1])
	op.Desc = GetOpDesc(int(op.Opc))
	op.Args = append(op.Args, Byte2U4(bytes[u1:u1*5]))
	return u1 * 5
}

func GetOpDesc(opcode int) string {

	switch opcode {

	case JVM_OPC_nop:
		return "nop"
	case JVM_OPC_aconst_null:
		return "aconst_null"
	case JVM_OPC_iconst_m1:
		return "iconst_m1"
	case JVM_OPC_iconst_0:
		return "iconst_0"
	case JVM_OPC_iconst_1:
		return "iconst_1"
	case JVM_OPC_iconst_2:
		return "iconst_2"
	case JVM_OPC_iconst_3:
		return "iconst_3"
	case JVM_OPC_iconst_4:
		return "iconst_4"
	case JVM_OPC_iconst_5:
		return "iconst_5"
	case JVM_OPC_lconst_0:
		return "lconst_0"
	case JVM_OPC_lconst_1:
		return "lconst_1"
	case JVM_OPC_fconst_0:
		return "fconst_0"
	case JVM_OPC_fconst_1:
		return "fconst_1"
	case JVM_OPC_fconst_2:
		return "fconst_2"
	case JVM_OPC_dconst_0:
		return "dconst_0"
	case JVM_OPC_dconst_1:
		return "dconst_1"
	case JVM_OPC_bipush:
		return "bipush"
	case JVM_OPC_sipush:
		return "sipush"
	case JVM_OPC_ldc:
		return "ldc"
	case JVM_OPC_ldc_w:
		return "ldc_w"
	case JVM_OPC_ldc2_w:
		return "ldc2_w"
	case JVM_OPC_iload:
		return "iload"
	case JVM_OPC_lload:
		return "lload"
	case JVM_OPC_fload:
		return "fload"
	case JVM_OPC_dload:
		return "dload"
	case JVM_OPC_aload:
		return "aload"
	case JVM_OPC_iload_0:
		return "iload_0"
	case JVM_OPC_iload_1:
		return "iload_1"
	case JVM_OPC_iload_2:
		return "iload_2"
	case JVM_OPC_iload_3:
		return "iload_3"
	case JVM_OPC_lload_0:
		return "lload_0"
	case JVM_OPC_lload_1:
		return "lload_1"
	case JVM_OPC_lload_2:
		return "lload_2"
	case JVM_OPC_lload_3:
		return "lload_3"
	case JVM_OPC_fload_0:
		return "fload_0"
	case JVM_OPC_fload_1:
		return "fload_1"
	case JVM_OPC_fload_2:
		return "fload_2"
	case JVM_OPC_fload_3:
		return "fload_3"
	case JVM_OPC_dload_0:
		return "dload_0"
	case JVM_OPC_dload_1:
		return "dload_1"
	case JVM_OPC_dload_2:
		return "dload_2"
	case JVM_OPC_dload_3:
		return "dload_3"
	case JVM_OPC_aload_0:
		return "aload_0"
	case JVM_OPC_aload_1:
		return "aload_1"
	case JVM_OPC_aload_2:
		return "aload_2"
	case JVM_OPC_aload_3:
		return "aload_3"
	case JVM_OPC_iaload:
		return "iaload"
	case JVM_OPC_laload:
		return "laload"
	case JVM_OPC_faload:
		return "faload"
	case JVM_OPC_daload:
		return "daload"
	case JVM_OPC_aaload:
		return "aaload"
	case JVM_OPC_baload:
		return "baload"
	case JVM_OPC_caload:
		return "caload"
	case JVM_OPC_saload:
		return "saload"
	case JVM_OPC_istore:
		return "istore"
	case JVM_OPC_lstore:
		return "lstore"
	case JVM_OPC_fstore:
		return "fstore"
	case JVM_OPC_dstore:
		return "dstore"
	case JVM_OPC_astore:
		return "astore"
	case JVM_OPC_istore_0:
		return "istore_0"
	case JVM_OPC_istore_1:
		return "istore_1"
	case JVM_OPC_istore_2:
		return "istore_2"
	case JVM_OPC_istore_3:
		return "istore_3"
	case JVM_OPC_lstore_0:
		return "lstore_0"
	case JVM_OPC_lstore_1:
		return "lstore_1"
	case JVM_OPC_lstore_2:
		return "lstore_2"
	case JVM_OPC_lstore_3:
		return "lstore_3"
	case JVM_OPC_fstore_0:
		return "fstore_0"
	case JVM_OPC_fstore_1:
		return "fstore_1"
	case JVM_OPC_fstore_2:
		return "fstore_2"
	case JVM_OPC_fstore_3:
		return "fstore_3"
	case JVM_OPC_dstore_0:
		return "dstore_0"
	case JVM_OPC_dstore_1:
		return "dstore_1"
	case JVM_OPC_dstore_2:
		return "dstore_2"
	case JVM_OPC_dstore_3:
		return "dstore_3"
	case JVM_OPC_astore_0:
		return "astore_0"
	case JVM_OPC_astore_1:
		return "astore_1"
	case JVM_OPC_astore_2:
		return "astore_2"
	case JVM_OPC_astore_3:
		return "astore_3"
	case JVM_OPC_iastore:
		return "iastore"
	case JVM_OPC_lastore:
		return "lastore"
	case JVM_OPC_fastore:
		return "fastore"
	case JVM_OPC_dastore:
		return "dastore"
	case JVM_OPC_aastore:
		return "aastore"
	case JVM_OPC_bastore:
		return "bastore"
	case JVM_OPC_castore:
		return "castore"
	case JVM_OPC_sastore:
		return "sastore"
	case JVM_OPC_pop:
		return "pop"
	case JVM_OPC_pop2:
		return "pop2"
	case JVM_OPC_dup:
		return "dup"
	case JVM_OPC_dup_x1:
		return "dup_x1"
	case JVM_OPC_dup_x2:
		return "dup_x2"
	case JVM_OPC_dup2:
		return "dup2"
	case JVM_OPC_dup2_x1:
		return "dup2_x1"
	case JVM_OPC_dup2_x2:
		return "dup2_x2"
	case JVM_OPC_swap:
		return "swap"
	case JVM_OPC_iadd:
		return "iadd"
	case JVM_OPC_ladd:
		return "ladd"
	case JVM_OPC_fadd:
		return "fadd"
	case JVM_OPC_dadd:
		return "dadd"
	case JVM_OPC_isub:
		return "isub"
	case JVM_OPC_lsub:
		return "lsub"
	case JVM_OPC_fsub:
		return "fsub"
	case JVM_OPC_dsub:
		return "dsub"
	case JVM_OPC_imul:
		return "imul"
	case JVM_OPC_lmul:
		return "lmul"
	case JVM_OPC_fmul:
		return "fmul"
	case JVM_OPC_dmul:
		return "dmul"
	case JVM_OPC_idiv:
		return "idiv"
	case JVM_OPC_ldiv:
		return "ldiv"
	case JVM_OPC_fdiv:
		return "fdiv"
	case JVM_OPC_ddiv:
		return "ddiv"
	case JVM_OPC_irem:
		return "irem"
	case JVM_OPC_lrem:
		return "lrem"
	case JVM_OPC_frem:
		return "frem"
	case JVM_OPC_drem:
		return "drem"
	case JVM_OPC_ineg:
		return "ineg"
	case JVM_OPC_lneg:
		return "lneg"
	case JVM_OPC_fneg:
		return "fneg"
	case JVM_OPC_dneg:
		return "dneg"
	case JVM_OPC_ishl:
		return "ishl"
	case JVM_OPC_lshl:
		return "lshl"
	case JVM_OPC_ishr:
		return "ishr"
	case JVM_OPC_lshr:
		return "lshr"
	case JVM_OPC_iushr:
		return "iushr"
	case JVM_OPC_lushr:
		return "lushr"
	case JVM_OPC_iand:
		return "iand"
	case JVM_OPC_land:
		return "land"
	case JVM_OPC_ior:
		return "ior"
	case JVM_OPC_lor:
		return "lor"
	case JVM_OPC_ixor:
		return "ixor"
	case JVM_OPC_lxor:
		return "lxor"
	case JVM_OPC_iinc:
		return "iinc"
	case JVM_OPC_i2l:
		return "i2l"
	case JVM_OPC_i2f:
		return "i2f"
	case JVM_OPC_i2d:
		return "i2d"
	case JVM_OPC_l2i:
		return "l2i"
	case JVM_OPC_l2f:
		return "l2f"
	case JVM_OPC_l2d:
		return "l2d"
	case JVM_OPC_f2i:
		return "f2i"
	case JVM_OPC_f2l:
		return "f2l"
	case JVM_OPC_f2d:
		return "f2d"
	case JVM_OPC_d2i:
		return "d2i"
	case JVM_OPC_d2l:
		return "d2l"
	case JVM_OPC_d2f:
		return "d2f"
	case JVM_OPC_i2b:
		return "i2b"
	case JVM_OPC_i2c:
		return "i2c"
	case JVM_OPC_i2s:
		return "i2s"
	case JVM_OPC_lcmp:
		return "lcmp"
	case JVM_OPC_fcmpl:
		return "fcmpl"
	case JVM_OPC_fcmpg:
		return "fcmpg"
	case JVM_OPC_dcmpl:
		return "dcmpl"
	case JVM_OPC_dcmpg:
		return "dcmpg"
	case JVM_OPC_ifeq:
		return "ifeq"
	case JVM_OPC_ifne:
		return "ifne"
	case JVM_OPC_iflt:
		return "iflt"
	case JVM_OPC_ifge:
		return "ifge"
	case JVM_OPC_ifgt:
		return "ifgt"
	case JVM_OPC_ifle:
		return "ifle"
	case JVM_OPC_if_icmpeq:
		return "if_icmpeq"
	case JVM_OPC_if_icmpne:
		return "if_icmpne"
	case JVM_OPC_if_icmplt:
		return "if_icmplt"
	case JVM_OPC_if_icmpge:
		return "if_icmpge"
	case JVM_OPC_if_icmpgt:
		return "if_icmpgt"
	case JVM_OPC_if_icmple:
		return "if_icmple"
	case JVM_OPC_if_acmpeq:
		return "if_acmpeq"
	case JVM_OPC_if_acmpne:
		return "if_acmpne"
	case JVM_OPC_goto:
		return "goto"
	case JVM_OPC_jsr:
		return "jsr"
	case JVM_OPC_ret:
		return "ret"
	case JVM_OPC_tableswitch:
		return "tableswitch"
	case JVM_OPC_lookupswitch:
		return "lookupswitch"
	case JVM_OPC_ireturn:
		return "ireturn"
	case JVM_OPC_lreturn:
		return "lreturn"
	case JVM_OPC_freturn:
		return "freturn"
	case JVM_OPC_dreturn:
		return "dreturn"
	case JVM_OPC_areturn:
		return "areturn"
	case JVM_OPC_return:
		return "return"
	case JVM_OPC_getstatic:
		return "getstatic"
	case JVM_OPC_putstatic:
		return "putstatic"
	case JVM_OPC_getfield:
		return "getfield"
	case JVM_OPC_putfield:
		return "putfield"
	case JVM_OPC_invokevirtual:
		return "invokevirtual"
	case JVM_OPC_invokespecial:
		return "invokespecial"
	case JVM_OPC_invokestatic:
		return "invokestatic"
	case JVM_OPC_invokeinterface:
		return "invokeinterface"
	case JVM_OPC_invokedynamic:
		return "invokedynamic"
	case JVM_OPC_new:
		return "new"
	case JVM_OPC_newarray:
		return "newarray"
	case JVM_OPC_anewarray:
		return "anewarray"
	case JVM_OPC_arraylength:
		return "arraylength"
	case JVM_OPC_athrow:
		return "athrow"
	case JVM_OPC_checkcast:
		return "checkcast"
	case JVM_OPC_instanceof:
		return "instanceof"
	case JVM_OPC_monitorenter:
		return "monitorenter"
	case JVM_OPC_monitorexit:
		return "monitorexit"
	case JVM_OPC_wide:
		return "wide"
	case JVM_OPC_multianewarray:
		return "multianewarray"
	case JVM_OPC_ifnull:
		return "ifnull"
	case JVM_OPC_ifnonnull:
		return "ifnonnull"
	case JVM_OPC_goto_w:
		return "goto_w"
	case JVM_OPC_jsr_w:
		return "jsr_w"
	}

	return "unknown"
}
