package parse

import (
	"class-parse/core"
	"fmt"
	"log"
	"testing"
)

const file = "/Users/zl/Documents/DEV/github/class-parse/java/Main.class"

func Test_class_parse(t *testing.T) {

	var cp ClassParse
	e := cp.parseFile(file)
	if e != nil {
		t.Fatal(e)
	}
	fmt.Println(cp)

}

func Test_get_magic_num(t *testing.T) {

	var cp ClassParse
	_ = cp.parseFile(file)

	m := cp.Magic()
	view := m.View()
	expect := core.Hex("CAFEBABE")

	if view != expect {
		t.Fatalf("magic num is  %s  except is %s", view, expect)
	}
}

func Test_get_Minor_Version(t *testing.T) {

	var cp ClassParse
	_ = cp.parseFile(file)
	mv := cp.MinorVersion()
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
	mv := cp.MajorVersion()
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
	constPoolCount := cp.ConstantPoolCount()
	v := constPoolCount.View()

	except := 139
	if v != except {
		t.Fatalf("constPoolCount is %d  except is %d", v, except)
	}

}

func Test_get_cp_info_view(t *testing.T) {

	var cp ClassParse
	_ = cp.parseFile(file)

	cpInfos := cp.CpInfos()

	s := cpInfos.View()
	fmt.Println(s)

	// 	}
	// for i, v := range cpInfos {
	// 	view, ok := v.(core.View)
	// 	if ok {
	// 		fmt.Printf("#%d = %v \n", i, view.View())
	// 	}
	// }

}
func Test_get_cp_info(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)
	cpInfos := cp.CpInfos()

	fmt.Println(cpInfos)
	integerMax := cpInfos[34]
	fmt.Println("integerMax is", integerMax)

	integerMin := cpInfos[45]
	fmt.Println("integerMin is", integerMin)

	longMax := cpInfos[51]
	fmt.Println("longMax is", longMax)

	longMin := cpInfos[60]
	cpLong := longMin.(*core.CpLong)
	fmt.Println("longMin is", longMin)

	expect := core.Long(-9223372036854775808)
	if cpLong.Long != expect {
		t.Fatalf("Long Min is %d except %d", cpLong.Long, expect)
	}
}

func Test_get_access_flags(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)

	af := cp.AccessFlag()
	view := af.View()
	except := "ACC_PUBLIC,ACC_SUPER"

	if view != except {
		t.Fatalf("access flags is %s  except is %s", view, except)
	}

}
