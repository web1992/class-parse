package parse

import (
	"fmt"
	"goclass/core"
	"log"
	"strings"
	"testing"
)

const file = "/Users/zl/Documents/DEV/github/class-parse/testfiles/Main.class"

func Test_class_parse(t *testing.T) {

	var cp ClassParse
	e := cp.parseFile(file)
	if e != nil {
		t.Fatal(e)
	}
	fmt.Println(cp)

}

func Test_get_class_file(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)

	classFile := cp.ClassFile()

	fmt.Println("classFile is", classFile)
}

func Test_get_magic_num(t *testing.T) {

	var cp ClassParse
	_ = cp.parseFile(file)

	m := cp.ClassFile().Magic
	view := m.View()
	expect := core.Hex("CAFEBABE")

	if view != expect {
		t.Fatalf("magic num is  %s  except is %s", view, expect)
	}
}

func Test_get_Minor_Version(t *testing.T) {

	var cp ClassParse
	_ = cp.parseFile(file)
	mv := cp.ClassFile().MinorVersion
	v := mv.View()
	log.Println("minor_Version is ", v)
	expect := 0
	if v != expect {
		t.Fatalf("minor version is %d,except is %d", v, expect)
	}
}

func Test_get_major_version(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)
	mv := cp.ClassFile().MajorVersion
	v := mv.View()
	expect := 58

	log.Println("view is", v)

	if expect != v {
		t.Fatalf("major version is %d,except is %d", v, expect)
	}

}

func Test_get_cp(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)
	constPoolCount := cp.ClassFile().ConstantPoolCount
	v := constPoolCount.View()

	except := 139
	if v != except {
		t.Fatalf("constPoolCount is %d  except is %d", v, except)
	}

}

func Test_get_cp_info_view(t *testing.T) {

	var cp ClassParse
	_ = cp.parseFile(file)
	classFile := cp.ClassFile()
	cpInfos := classFile.CpInfos

	sv := cpInfos.View()
	//fmt.Println(sv)
	expect := "#1 = Methodref #2.#3 AbstractMain.<init>:()V"

	s := sv.(string)
	if !strings.Contains(s, expect) {
		t.Fatalf(" expect %s \n sv is \n%s", expect, s)
	}

}
func Test_get_cp_info(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)
	classFile := cp.ClassFile()
	cpInfos := classFile.CpInfos

	fmt.Println(cpInfos.View())
}

func Test_get_access_flags(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)

	af := cp.ClassFile().AccessFlag
	view := af.View()
	except := "ACC_PUBLIC,ACC_SUPER"

	if view != except {
		t.Fatalf("access flags is %s  except is %s", view, except)
	}

}

func Test_get_this_class(t *testing.T) {

	var cp ClassParse
	_ = cp.parseFile(file)

	classFile := cp.ClassFile()

	tc := classFile.ThisClass

	s := tc.String
	fmt.Println(s)
	expect := "#40 = Class #42 Main "

	if s != expect {
		t.Fatalf("\n expect is %v \n s is %v", expect, s)
	}
}

func Test_get_interface(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)

	classFile := cp.ClassFile()

	fmt.Println(classFile)

	ifc := classFile.InterfacesCount

	if ifc.Count != int32(1) {
		t.Fatalf("ifc is %v expect is %d", ifc, 1)
	}
	ifcc := classFile.Interfaces[0]

	s := ifcc.NameString
	expect := "#107 = Class #108 InterfaceMain "

	if s != expect {
		t.Fatalf("ifcc is %v expect is %s", ifcc, expect)
	}
}

func Test_get_attributes(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)

	classFile := cp.ClassFile()

	fmt.Println(classFile)
}
