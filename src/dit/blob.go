package dit

import (
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"os"
	"strconv"
)

type Blob struct {
	object
}

func writeToObject(data []byte, sha1 string) {
	fileDir := ".dit/objects/" + sha1[0:2]
	os.Mkdir(fileDir, os.ModePerm)

	filePath := fileDir + "/" + sha1[2:]
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("open file error", filePath, err)
		return
	}
	defer file.Close()

	w := zlib.NewWriter(file)
	defer w.Close()

	w.Write(data)
}

func (b *Blob) Hash(filePath string) {
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Printf("fatal: pathspec %x did not match any files\n", filePath)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file", filePath, "error: ", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, fileInfo.Size())
	length, err := file.Read(buffer)
	if err != nil || int64(length) != fileInfo.Size() {
		fmt.Println("read file", filePath, "error:", err)
		return
	}

	header := "blob " + strconv.FormatInt(fileInfo.Size(), 10) + "\x00"
	sha1Buffer := append([]byte(header), buffer...)

	b.sha1 = fmt.Sprintf("%x", sha1.Sum(sha1Buffer))
	b.size = fileInfo.Size()
	b.type_ = "blob"
	b.parsed = true
	b.used = true

	// zlib and write
	writeToObject(sha1Buffer, b.sha1)

	fmt.Printf("sha1: %s\n", b.sha1)
}
