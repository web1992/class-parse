package parse

import (
	"fmt"
	"log"
	"testing"
)

func Test_class_parse(t *testing.T) {

	var cp ClassParse
	s := "/Users/zl/Documents/DEV/github/class-parse/Main.class"
	e := cp.parseFile(s)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(cp)

}
