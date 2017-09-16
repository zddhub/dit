package dit

import (
	"bytes"

	. "github.com/zddhub/dit/utils"
)

type tree struct {
	object
	objects []*object
}

func (repo *repository) NewTree() *tree {
	var b bytes.Buffer
	tree := &tree{object{Type: "tree"}, nil}

	tree.objects = make([]*object, len(repo.index.Entries))
	copy(tree.objects, repo.index.Entries)

	for _, obj := range tree.objects {
		LogI.Println(obj.String())
		b.Write([]byte(obj.String() + "\n"))
	}

	tree.Write(b.Bytes())
	return tree
}
