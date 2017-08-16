package hashx

import (
	"crypto/sha1"
	"fmt"
	"os"
	. "utils"
)

const Size = 20

// sha1
func MemHashToBytes(data []byte) [Size]byte {
	return sha1.Sum(data)
}

func MemHashToString(data []byte) string {
	return fmt.Sprintf("%x", MemHashToBytes(data))
}

func FileHash(path string) (sha1 string, err error) {
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
	fmt.Println(buffer)
	return MemHashToString(buffer), nil
}
