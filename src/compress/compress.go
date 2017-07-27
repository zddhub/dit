package compress

import (
	"compress/zlib"
	"os"
	"path/filepath"
	. "utils"
)

// Use Go default zlib implements, different with Git

// compress data to path file
func Compress(path string, data []byte) {
	dir := filepath.Dir(path)
	os.MkdirAll(dir, os.ModePerm)

	file, err := os.Create(path)
	if err != nil {
		LogE.Fatalln(err)
		return
	}
	defer file.Close()

	w := zlib.NewWriter(file)
	defer w.Close()

	w.Write(data)
}

// uncompress file to data
func Uncompress(data []byte, path string) {
	file, err := os.Open(path)
	if err != nil {
		LogE.Fatalln(err)
		return
	}
	defer file.Close()

	r, err := zlib.NewReader(file)
	if err != nil {
		LogE.Fatalln("Read file error!")
		return
	}
	r.Close()

	r.Read(data)
}
