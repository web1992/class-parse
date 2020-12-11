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

		if er == io.EOF {
			break
		}

		bs = append(bs, buf[:n]...)

	}

	return bs, nil

}
