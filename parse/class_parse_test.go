package parse

import (
	"fmt"
	"log"
	"testing"
)

const file = "/Users/zl/Documents/DEV/github/class-parse/Main.class"

func Test_class_parse(t *testing.T) {

	var cp ClassParse
	e := cp.parseFile(file)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(cp)

}

func Test_get_magic_num(t *testing.T) {

	var cp ClassParse
	_ = cp.parseFile(file)

	m := cp.Magic()

	log.Println(m)
	log.Println(m.View())
}

func Test_get_Minor_Version(t *testing.T) {

	var cp ClassParse
	_ = cp.parseFile(file)
	mv := cp.MinorVersion()
	v := mv.View()
	log.Println(v)
	log.Println("view is", v)
}

func Test_get_major_version(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)
	mv := cp.MajorVersion()
	v := mv.View()
	log.Println("view is", v)

}

func Test_get_cp(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)
	cpool := cp.ConstantPoolCount()
	v := cpool.View()
	log.Println("view is", v)

	if v != 35 {
		t.Fatalf("view is %d  except is %d", v, 35)
	}

}

func Test_get_cp_info(t *testing.T) {
	var cp ClassParse
	_ = cp.parseFile(file)

	cpInfos := cp.CpInfos()

	fmt.Println(cpInfos)
}
