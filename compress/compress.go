package compress

import (
	"compress/zlib"
	. "github.com/zddhub/dit/utils"
	"os"
	"path/filepath"
)

// Use Go default zlib implements, different with Git

// compress data to path file
func Compress(path string, data []byte) (int, error) {
	dir := filepath.Dir(path)
	os.MkdirAll(dir, os.ModePerm)

	file, err := os.Create(path)
	if err != nil {
		LogE.Fatalln(err)
		return -1, err
	}
	defer file.Close()

	w := zlib.NewWriter(file)
	defer w.Close()

	return w.Write(data)
}

// uncompress file to data
func Uncompress(data []byte, path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		LogE.Fatalln(err)
		return -1, err
	}
	defer file.Close()

	r, err := zlib.NewReader(file)
	if err != nil {
		LogE.Fatalln("Read file error!")
		return -1, err
	}
	r.Close()

	return r.Read(data)
}
