package parse

import (
	"fmt"
	"github.com/web1992/goclass/core"
	"log"
	"os"
	"strings"
	"testing"
)

const file = "../testfiles/Main.class"

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
	expect := core.Hex("CAFEBABE")

	if m.Hex != expect {
		t.Fatalf("magic num is  %s  except is %s", m.Hex, expect)
	}
}

func Test_get_Minor_Version(t *testing.T) {

	var cp ClassParse
	_ = cp.parseFile(file)
	mv := cp.ClassFile().MinorVersion
	v := mv.Version
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
	expect := 58

	if (expect) != mv.Version {
		t.Fatalf("major version is %d,except is %d", mv.Version, expect)
	}

}

func Test_get_cp(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)
	constPoolCount := cp.ClassFile().ConstantPoolCount
	v := constPoolCount.Count

	except := 162
	if v != except {
		t.Fatalf("constPoolCount is %d  except is %d", v, except)
	}

}

func Test_get_cp_info_view(t *testing.T) {

	var cp ClassParse
	_ = cp.parseFile(file)
	classFile := cp.ClassFile()
	cpInfos := classFile.CpInfos

	//sv := cpInfos.View()
	//fmt.Println(cpInfos)
	expect := "#1 = Methodref          #2.#3         // AbstractMain.<init>:()V"

	s := cpInfos.String()
	if !strings.Contains(s, expect) {
		t.Fatalf(" expect %s \n sv is \n%s", expect, s)
	}

}
func Test_get_cp_info(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)
	classFile := cp.ClassFile()
	cpInfos := classFile.CpInfos
	v := cpInfos.String()
	fmt.Println(v)
	f, _ := os.Create("../testfiles/cp_test.txt")
	defer f.Close()
	f.Write([]byte(v))
}

func Test_get_access_flags(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)

	af := cp.ClassFile().AccessFlag
	view := af.String()
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
	expect := "Main"

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

	if ifc.Count != 1 {
		t.Fatalf("ifc is %v expect is %d", ifc, 1)
	}
	ifcc := classFile.Interfaces[0]

	s := ifcc.NameString
	expect := "InterfaceMain"

	if s != expect {
		t.Fatalf("ifcc is %v expect is %s", ifcc, expect)
	}
}

func Test_get_attributes(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)

	classFile := cp.ClassFile()

	core.Info.Println("tes output -------------------")
	fmt.Println(classFile.String())

	file, err := os.Create("../testfiles/Main.txt")
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	defer file.Close()
	file.WriteString("\n")
	file.WriteString("\n")
	file.WriteString("\n")

	file.WriteString(classFile.String())

}

func Test_log(t *testing.T) {
	core.Info.Println("INFO")
	core.Trace.Println("Trace")
	core.Warning.Println("Warning")
	core.Error.Println("Error")
}
