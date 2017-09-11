package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	Regular   uint32 = 0100000
	Directory uint32 = 0040000
	Symbolic  uint32 = 0120000
)

func ReadFile(filename string) (buffer []byte, err error) {
	return ioutil.ReadFile(filename)
}

func WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}

func FileMode(filename string) (string, error) {
	if fileInfo, err := os.Lstat(filename); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("%o", uint32(fileInfo.Mode().Perm())|Regular), nil
	}
}
