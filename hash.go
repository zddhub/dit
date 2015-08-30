package dit

import (
	"crypto/sha1"
	"fmt"
	"os"
)

// sha1
func MemHash(data []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(data))
}

func FileHash(path string) (string, error) {
	fileInfo, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		LogE.Println(err)
		return "", err
	}

	file, err := os.Open(path)
	if err != nil {
		LogE.Println(err)
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, fileInfo.Size())
	if n, err := file.Read(buffer); err != nil || int64(n) != fileInfo.Size() {
		LogE.Fatalln(err)
		return "", err
	}
	return MemHash(buffer), nil
}
