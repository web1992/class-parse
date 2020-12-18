package utils

import (
	"goclass/core"
	"testing"
)

func Test_get_double(t *testing.T) {
	bs := []byte{64, 9, 33, 251, 84, 68, 45, 24}
	d := core.Byte2Double(bs)
	except := 3.141592653589793

	if except != d {
		t.Fatalf("d is %f except is  %f", d, except)
	}
}

func Test_get_integer(t *testing.T) {

}
