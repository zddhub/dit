package utils

import (
	"io/ioutil"
	"os"
)

func ReadFile(filename string) (buffer []byte, err error) {
	return ioutil.ReadFile(filename)
}

func WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}
