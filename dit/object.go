package dit

import (
	"fmt"

	"github.com/zddhub/dit/compress"
	"github.com/zddhub/dit/hash"
)

type object struct {
	Sha1 string
	Type string // object type: object, blob, tree, commit, tag
	Size int
	Mode string
	Path string
}

func (obj *object) Sha1String() string {
	return obj.Sha1
}

func (obj *object) String() string {
	return fmt.Sprintf("%s %s %s\t%s", obj.Mode, obj.Type, obj.Sha1, obj.Path)
}

func (obj *object) Write(p []byte) (n int, err error) {
	header := []byte(fmt.Sprintf("%s %d\x00", obj.Type, len(p)))
	data := append(header, p...)

	obj.Sha1 = BytesToSha1(hash.MemHashToBytes(data))
	obj.Size = len(p)

	filePath := DIT["objects"] + "/" + obj.Sha1[0:2] + "/" + obj.Sha1[2:]

	return compress.Compress(filePath, data)
}

func (obj *object) ReadAll() ([]byte, error) {
	filePath := DIT["objects"] + "/" + obj.Sha1[0:2] + "/" + obj.Sha1[2:]

	buf, err := compress.Decompress(filePath)
	if err != nil {
		return nil, err
	}

	fmt.Sscanf(fmt.Sprintf("%s", buf), "%s %d", &obj.Type, &obj.Size)
	nf := len(fmt.Sprintf("%s %d\x00", obj.Type, obj.Size))

	return buf[nf:], err
}

func BytesToSha1(bytes [20]byte) (sha1 string) {
	return fmt.Sprintf("%x", bytes)
}

func Sha1ToBytes(sha1 string) (bytes [20]byte) {
	for i := range bytes {
		if 2*i+1 >= len(sha1) {
			break
		}
		bytes[i] = htob(sha1[2*i])*16 + htob(sha1[2*i+1])
	}
	return bytes
}

func htob(hex byte) (b byte) {
	if hex <= '9' {
		b = hex - '0'
	} else {
		b = hex - 'a' + 10
	}
	return b
}
