package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	l := len(os.Args)
	if l == 1 {
		fmt.Println("Please enter file name! Exit!")
		return
	}
	fileName := os.Args[1]
	fmt.Println("main", fileName)

	bs, error := ioutil.ReadFile(fileName)

	if error != nil {
		fmt.Println("error ", error)
		return
	}

	magic := bs[0:4]
	minor_version := bs[4:6]

	major_version := bs[6:8]
	c_p_c := bs[8:10]

	// with "%x" format byte array into hex string
	hexStr := fmt.Sprintf("%x", magic)
	minor_version_srt := fmt.Sprintf("%x", minor_version)
	major_version_str := fmt.Sprintf("%x", major_version)
	c_p_c_str := fmt.Sprintf("%x", c_p_c)
	fmt.Println(hexStr)
	fmt.Println(minor_version_srt)
	fmt.Println(major_version_str)
	fmt.Println(c_p_c_str)

}
