package compress

import (
	"compress/zlib"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/zddhub/dit/golang/utils"
)

// Use Go default zlib implements, different with Git

// compress data to path file
func Compress(path string, data []byte) (int, error) {
	dir := filepath.Dir(path)
	os.MkdirAll(dir, os.ModePerm)

	file, err := os.Create(path)
	if err != nil {
		utils.LogE.Println(err)
		return 0, err
	}
	defer file.Close()

	w := zlib.NewWriter(file)
	defer w.Close()

	return w.Write(data)
}

// Decompress file
func Decompress(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		utils.LogE.Println(err)
		return nil, err
	}
	defer file.Close()

	r, err := zlib.NewReader(file)
	if err != nil {
		utils.LogE.Println(err)
		return nil, err
	}
	defer r.Close()

	return ioutil.ReadAll(r)
}
