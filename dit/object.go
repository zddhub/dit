package dit

import (
	"fmt"
	"github.com/zddhub/dit/compress"
	"github.com/zddhub/dit/hash"
)

type object struct {
	sha1 [20]byte
	flag string // object type: object, blob, tree, commit, tag
	size int
}

func (o *object) Type() string {
	return o.flag
}

func (o *object) Size() int {
	return o.size
}

func (o *object) Sha1String() string {
	return BytesToSha1(o.sha1)
}

func (o *object) Write(p []byte) (n int, err error) {
	header := []byte(fmt.Sprintf("%s %d\x00", o.flag, len(p)))
	data := append(header, p...)

	o.sha1 = hash.MemHashToBytes(data)

	sha1String := o.Sha1String()

	filePath := DIT["objects"] + "/" + sha1String[0:2] + "/" + sha1String[2:]

	return compress.Compress(filePath, data)
}

func (o *object) ReadAll() ([]byte, error) {
	sha1String := o.Sha1String()

	filePath := DIT["objects"] + "/" + sha1String[0:2] + "/" + sha1String[2:]

	buf, err := compress.Decompress(filePath)
	if err != nil {
		return nil, err
	}

	fmt.Sscanf(fmt.Sprintf("%s", buf), "%s %d", &o.flag, &o.size)
	nf := len(fmt.Sprintf("%s %d\x00", o.flag, o.size))

	return buf[nf:], err
}

func BytesToSha1(bytes [20]byte) (sha1 string) {
	return fmt.Sprintf("%x", bytes)
}

func Sha1ToBytes(sha1 string) (bytes [20]byte) {
	for i := range bytes {
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
