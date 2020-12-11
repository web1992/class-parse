package parse

import (
	"class-parse/utils"
	"errors"
	"strings"
)

const fileSuffix = ".class"

var filePathInvalid = errors.New("file patch is invalid")
var fileSuffixInvalid = errors.New("file must has suffix " + fileSuffix)

type ClassParse struct {
	filePath string
	fileByes []byte
}

func (cp *ClassParse) Name() string {
	return cp.filePath
}

func (cp *ClassParse) Bytes() []byte {
	return cp.fileByes
}

func (cp *ClassParse) String() string {
	return cp.filePath
}
func (cp *ClassParse) parseFile(filePath string) error {

	cp.filePath = filePath

	if len(filePath) == 0 {
		return filePathInvalid
	}

	checkBol := strings.HasSuffix(filePath, fileSuffix)

	if !checkBol {
		return fileSuffixInvalid
	}

	bytes, e := utils.ReadClassFile(filePath)

	if e != nil {
		return e
	}

	cp.fileByes = bytes
	return nil
}
