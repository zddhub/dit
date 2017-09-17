package dit

import (
	"bytes"

	. "github.com/zddhub/dit/utils"
)

type tree struct {
	object
	objects []*object
}

func (repo repository) NewTree() *tree {
	tree := &tree{object{Type: "tree"}, nil}

	tree.objects = make([]*object, len(repo.cache.Entries))
	copy(tree.objects, repo.cache.Entries)

	return tree
}

func (t tree) Content() []byte {
	var b bytes.Buffer
	for _, obj := range t.objects {
		LogI.Println(obj.String())
		b.Write([]byte(obj.String() + "\n"))
	}
	return b.Bytes()
}
