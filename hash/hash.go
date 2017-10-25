package hash

import (
	"crypto/sha1"
	"fmt"

	"github.com/zddhub/dit/utils"
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
	buffer, err := utils.ReadFile(path)
	if err != nil {
		return "", err
	}
	return MemHashToString(buffer), nil
}
