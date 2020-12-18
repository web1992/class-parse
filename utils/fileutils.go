package utils

import (
	"io"
	"os"
)

func ReadClassFile(filePath string) ([]byte, error) {

	f, er := os.Open(filePath)

	defer f.Close()

	if er != nil {
		return nil, er
	}
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
