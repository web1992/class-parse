package core

func CreateOpCode(op int32) interface{} {

	switch int(op) {
	case JVM_OPC_nop:
		return &OpCode{}
	case JVM_OPC_aconst_null:
		return &OpCode{}
	case JVM_OPC_iconst_m1,
		JVM_OPC_iconst_0,
		JVM_OPC_iconst_1,
		JVM_OPC_iconst_2,
		JVM_OPC_iconst_3,
		JVM_OPC_iconst_4,
		JVM_OPC_iconst_5:
		return &OpCode{}
	case JVM_OPC_lconst_0,
		JVM_OPC_lconst_1:
		return &OpCode{}
	case JVM_OPC_fconst_0,
		JVM_OPC_fconst_1,
		JVM_OPC_fconst_2:
		return &OpCode{}
	case JVM_OPC_dconst_0,
		JVM_OPC_dconst_1:
		return &OpCode{}
		// todo
	case JVM_OPC_bipush:
		return &OpCode{}
		// todo
	case JVM_OPC_sipush:
		return &OpCode3{}
	case JVM_OPC_ldc:
		return &OpCode2{}
	case JVM_OPC_ldc_w:
		return &OpCode3{}
	case JVM_OPC_ldc2_w:
		return &OpCode3{}
	case JVM_OPC_iload:
		return &OpCode2{}
	case JVM_OPC_lload:
		return &OpCode2{}
	case JVM_OPC_fload:
		return &OpCode2{}
	case JVM_OPC_dload:
		return &OpCode2{}
	case JVM_OPC_aload:
		return &OpCode{}
	case JVM_OPC_iload_0,
		JVM_OPC_iload_1,
		JVM_OPC_iload_2,
		JVM_OPC_iload_3:
		return &OpCode{}
	case JVM_OPC_lload_0,
		JVM_OPC_lload_1,
		JVM_OPC_lload_2,
		JVM_OPC_lload_3:
		return &OpCode{}
	case JVM_OPC_fload_0,
		JVM_OPC_fload_1,
		JVM_OPC_fload_2,
		JVM_OPC_fload_3:
		return &OpCode{}
	case JVM_OPC_dload_0,
		JVM_OPC_dload_1,
		JVM_OPC_dload_2,
		JVM_OPC_dload_3:
		return &OpCode{}
	case JVM_OPC_aload_0,
		JVM_OPC_aload_1,
		JVM_OPC_aload_2,
		JVM_OPC_aload_3:
		return &OpCode{}
	case JVM_OPC_iaload,
		JVM_OPC_laload,
		JVM_OPC_faload,
		JVM_OPC_daload,
		JVM_OPC_aaload,
		JVM_OPC_baload,
		JVM_OPC_caload,
		JVM_OPC_saload:
		return &OpCode{}
	case JVM_OPC_istore,
		JVM_OPC_lstore,
		JVM_OPC_fstore,
		JVM_OPC_dstore,
		JVM_OPC_astore:
		return &OpCode2{}
	case JVM_OPC_istore_0,
		JVM_OPC_istore_1,
		JVM_OPC_istore_2,
		JVM_OPC_istore_3:
		return &OpCode{}
	case JVM_OPC_lstore_0,
		JVM_OPC_lstore_1,
		JVM_OPC_lstore_2,
		JVM_OPC_lstore_3:
		return &OpCode{}
	case JVM_OPC_fstore_0,
		JVM_OPC_fstore_1,
		JVM_OPC_fstore_2,
		JVM_OPC_fstore_3:
		return &OpCode{}
	case JVM_OPC_dstore_0,
		JVM_OPC_dstore_1,
		JVM_OPC_dstore_2,
		JVM_OPC_dstore_3:
		return &OpCode{}
	case JVM_OPC_astore_0,
		JVM_OPC_astore_1,
		JVM_OPC_astore_2,
		JVM_OPC_astore_3:
		return &OpCode{}
	case JVM_OPC_iastore,
		JVM_OPC_lastore,
		JVM_OPC_fastore,
		JVM_OPC_dastore,
		JVM_OPC_aastore,
		JVM_OPC_bastore,
		JVM_OPC_castore,
		JVM_OPC_sastore:
		return &OpCode{}
	case JVM_OPC_pop,
		JVM_OPC_pop2,
		JVM_OPC_dup,
		JVM_OPC_dup_x1,
		JVM_OPC_dup_x2,
		JVM_OPC_dup2,
		JVM_OPC_dup2_x1,
		JVM_OPC_dup2_x2:
		return &OpCode{}
	case JVM_OPC_swap:
		return &OpCode{}
	case JVM_OPC_iadd,
		JVM_OPC_ladd,
		JVM_OPC_fadd,
		JVM_OPC_dadd:
		return &OpCode{}
	case JVM_OPC_isub,
		JVM_OPC_lsub,
		JVM_OPC_fsub,
		JVM_OPC_dsub:
		return &OpCode{}
	case JVM_OPC_imul,
		JVM_OPC_lmul,
		JVM_OPC_fmul,
		JVM_OPC_dmul,
		JVM_OPC_idiv,
		JVM_OPC_ldiv,
		JVM_OPC_fdiv,
		JVM_OPC_ddiv,
		JVM_OPC_irem,
		JVM_OPC_lrem,
		JVM_OPC_frem,
		JVM_OPC_drem,
		JVM_OPC_ineg,
		JVM_OPC_lneg,
		JVM_OPC_fneg,
		JVM_OPC_dneg,
		JVM_OPC_ishl,
		JVM_OPC_lshl,
		JVM_OPC_ishr,
		JVM_OPC_lshr,
		JVM_OPC_iushr,
		JVM_OPC_lushr,
		JVM_OPC_iand,
		JVM_OPC_land,
		JVM_OPC_ior,
		JVM_OPC_lor,
		JVM_OPC_ixor,
		JVM_OPC_lxor:
		return &OpCode{}
		// todo
	case JVM_OPC_iinc:
		return &OpCode{}
	case JVM_OPC_i2l:
		return &OpCode{}
	case JVM_OPC_i2f,
		JVM_OPC_i2d,
		JVM_OPC_l2i,
		JVM_OPC_l2f,
		JVM_OPC_l2d,
		JVM_OPC_f2i,
		JVM_OPC_f2l,
		JVM_OPC_f2d,
		JVM_OPC_d2i,
		JVM_OPC_d2l,
		JVM_OPC_d2f,
		JVM_OPC_i2b,
		JVM_OPC_i2c,
		JVM_OPC_i2s,
		JVM_OPC_lcmp:
		return &OpCode{}
	case JVM_OPC_fcmpl,
		JVM_OPC_fcmpg,
		JVM_OPC_dcmpl,
		JVM_OPC_dcmpg,
		JVM_OPC_ifeq,
		JVM_OPC_ifne,
		JVM_OPC_iflt,
		JVM_OPC_ifge,
		JVM_OPC_ifgt,
		JVM_OPC_ifle:
		return &OpCode3{}
	case JVM_OPC_if_icmpeq,
		JVM_OPC_if_icmpne,
		JVM_OPC_if_icmplt,
		JVM_OPC_if_icmpge,
		JVM_OPC_if_icmpgt,
		JVM_OPC_if_icmple,
		JVM_OPC_if_acmpeq,
		JVM_OPC_if_acmpne:
		return &OpCode3{}
	case JVM_OPC_goto:
		return &OpCode3{}
	case JVM_OPC_jsr:
		return &OpCode3{}
	case JVM_OPC_ret:
		return &OpCode2{}
		// todo
	case JVM_OPC_tableswitch:
		return &OpCode{}
		// todo
	case JVM_OPC_lookupswitch:
		return &OpCode{}
	case JVM_OPC_ireturn,
		JVM_OPC_lreturn,
		JVM_OPC_freturn,
		JVM_OPC_dreturn,
		JVM_OPC_areturn,
		JVM_OPC_return:
		return &OpCode{}
	case JVM_OPC_getstatic,
		JVM_OPC_putstatic,
		JVM_OPC_getfield,
		JVM_OPC_putfield:
		return &OpCode3{}
	case JVM_OPC_invokevirtual:
		return &OpCode3{}
	case JVM_OPC_invokespecial:
		return &OpCode3{}
	case JVM_OPC_invokestatic:
		return &OpCode3{}
	case JVM_OPC_invokeinterface:
		return &OpCode5{}
	case JVM_OPC_invokedynamic:
		return &OpCode5{}
	case JVM_OPC_new:
		return &OpCode3{}
	case JVM_OPC_newarray:
		return &OpCode{}
	case JVM_OPC_anewarray:
		return &OpCode3{}
	case JVM_OPC_arraylength:
		return &OpCode{}
	case JVM_OPC_athrow:
		return &OpCode{}
	case JVM_OPC_checkcast:
		return &OpCode3{}
	case JVM_OPC_instanceof:
		return &OpCode3{}
	case JVM_OPC_monitorenter,
		JVM_OPC_monitorexit:
		return &OpCode{}
		// todo
	case JVM_OPC_wide:
		return &OpCode{}
		// todo
	case JVM_OPC_multianewarray:
		return &OpCode{}
	case JVM_OPC_ifnull,
		JVM_OPC_ifnonnull:
		return &OpCode3{}
	case JVM_OPC_goto_w,
		JVM_OPC_jsr_w:
		return &OpCode5{}
	default:
		return &OpCode{}
	}

	return &OpCode{}
}
