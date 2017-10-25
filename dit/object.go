package dit

import (
	"fmt"

	"github.com/zddhub/dit/compress"
	"github.com/zddhub/dit/hash"
	"github.com/zddhub/dit/utils"
)

type object struct {
	Sha1 string
	Type string // object type: object, blob, tree, commit, tag
	Size int
	Mode string
	Path string
}

func (obj object) String() string {
	return fmt.Sprintf("%s %s %s\t%s", obj.Mode, obj.Type, obj.Sha1, obj.Path)
}

func (obj *object) Write(p []byte) (n int, err error) {
	header := []byte(fmt.Sprintf("%s %d\x00", obj.Type, len(p)))
	data := append(header, p...)

	obj.Sha1 = utils.BytesToSha1(hash.MemHashToBytes(data))
	obj.Size = len(p)

	filePath := DIT["objects"] + "/" + obj.Sha1[0:2] + "/" + obj.Sha1[2:]

	if utils.IsExist(filePath) {
		return 0, nil
	}
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
