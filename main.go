package main

import (
	"fmt"
	"goclass/core"
)

func main() {

	fmt.Printf("goclass")
	a := []byte{0, 32, 0, 0}
	fmt.Println(core.Byte2U4(a))
}
