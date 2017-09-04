package utils

import (
	"fmt"
	"os"
)

func ReadFile(path string) (buffer []byte, err error) {
	fileInfo, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		LogE.Println(err, fileInfo.IsDir())
		return nil, err
	}

	if fileInfo.IsDir() {
		return nil, fmt.Errorf("file: %s is directory", path)
	}

	file, err := os.Open(path)
	if err != nil {
		LogE.Println(err)
		return nil, err
	}
	defer file.Close()

	buffer = make([]byte, fileInfo.Size())
	if n, err := file.Read(buffer); err != nil || int64(n) != fileInfo.Size() {
		LogE.Fatalln(err)
		return nil, err
	}
	return buffer, nil
}
