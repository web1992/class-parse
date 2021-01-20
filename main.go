package main

import (
	"fmt"

	"github.com/web1992/goclass/parse"
)

func main() {

	const file = "./testfiles/Main.class"
	var cp parse.ClassParse
	_ = cp.Parse(file)
	cf := cp.ClassFile()
	desc := cf.ClassDesc()
	fmt.Println(cp.CpDesc(cf.ThisClass.String))
	fmt.Println(desc)

}
