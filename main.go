package main

import (
	"fmt"
	"github.com/web1992/goclass/core"
)

func main() {

	fmt.Printf("goclass \n")
	a := []byte{0, 32, 0, 0}
	fmt.Println(core.Byte2U4(a))
}
