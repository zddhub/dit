package dit

import (
	"fmt"
	"github.com/zddhub/dit/compress"
	"github.com/zddhub/dit/hash"
	"os"
)

type object struct {
	sha1 [20]byte
	flag string // object category: object, blob, tree, commit, tag
}

func (o *object) Sha1String() string {
	return BytesToSha1(o.sha1)
}

func (o *object) Write(p []byte) (n int, err error) {
	header := []byte(fmt.Sprintf("%s %d\x00", o.flag, len(p)))
	data := append(header, p...)

	o.sha1 = hash.MemHashToBytes(data)

	sha1String := o.Sha1String()

	fileDir := DIT_REPO_DIR + "/objects/" + sha1String[0:2]
	filePath := fileDir + "/" + sha1String[2:]

	return compress.Compress(filePath, data)
}

func (o *object) Read(b []byte) (n int, err error) {
	sha1String := o.Sha1String()

	fileDir := DIT_REPO_DIR + "/objects/" + sha1String[0:2]
	filePath := fileDir + "/" + sha1String[2:]

	var fileSize int64
	if fileInfo, err := os.Stat(filePath); err == nil {
		fileSize = fileInfo.Size()
	} else {
		return -1, err
	}

	// TODO: Actually fileSize is larger then buf size
	buf := make([]byte, fileSize)
	n, err = compress.Uncompress(buf, filePath)

	flag, size := "", 0
	fmt.Sscanf(fmt.Sprintf("%s", buf), "%s %d", &flag, &size)
	nf := len(fmt.Sprintf("%s %d\x00", flag, size))

	o.flag = flag
	copy(b, buf[nf:n])

	return n, err
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
