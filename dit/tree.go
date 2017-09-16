package dit

import (
	"bytes"

	. "github.com/zddhub/dit/utils"
)

type tree struct {
	object
	objects []*object
}

func (t tree) Content() []byte {
	var b bytes.Buffer
	for _, obj := range t.objects {
		LogI.Println(obj.String())
		b.Write([]byte(obj.String() + "\n"))
	}
	return b.Bytes()
}

func (repo *repository) NewTree() *tree {
	tree := &tree{object{Type: "tree"}, nil}

	tree.objects = make([]*object, len(repo.index.Entries))
	copy(tree.objects, repo.index.Entries)

	return tree
}
