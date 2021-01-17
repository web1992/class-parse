package utils

import (
	"io"
	"os"
)

func GetFile(filePath string) (*os.File, error) {
	f, er := os.Open(filePath)

	if er != nil {
		return nil, er
	}

	return f, nil
}

func ReadFile(filePath string) ([]byte, error) {
	f, er := GetFile(filePath)
	if er != nil {
		return nil, er
	}
	return ReadClassFile(f)
}

func ReadClassFile(f *os.File) ([]byte, error) {

	var bs []byte

	buf := make([]byte, 1024)

	for {

		n, er := f.Read(buf)
		bs = append(bs, buf[:n]...)

		if er == io.EOF {
			break
		}
	}

	return bs, nil

}
