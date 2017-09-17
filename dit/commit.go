package dit

import (
	"bytes"
	"fmt"
	"time"

	. "github.com/zddhub/dit/utils"
)

type commit struct {
	object
	tree      *tree
	parent    *commit
	message   string
	author    *userInfo
	committer *userInfo
}

func (repo *repository) Commit(message string) {
	if !repo.validCommit() {
		LogT.Println("nothing added to commit")
		return
	}
	commit := repo.NewCommit()
	commit.message = message
	commit.Write(commit.Content())

	repo.WriteHead(commit.Sha1)

	repo.cache.Extensions = commit.newTreeCache()
	repo.StoreCache()
}

func (repo repository) NewCommit() *commit {
	tree := repo.NewTree()
	tree.Write(tree.Content())

	parent := &commit{object{Type: "commit", Sha1: repo.head}, nil, nil, "", nil, nil}

	user.timestamp = time.Now()
	return &commit{object{Type: "commit"}, tree, parent, "", user, user}
}

func (repo repository) validCommit() bool {
	extensions := repo.cache.Extensions.(map[string]interface{})
	if extensions["Signature"] == "TREE" && extensions["Entries"] != nil {
		return false
	}
	return true
}

func (c commit) Content() []byte {
	var b bytes.Buffer
	var content string
	content += fmt.Sprintf("tree %s\n", c.tree.Sha1)
	if c.parent.Sha1 != "" {
		content += fmt.Sprintf("parent %s\n", c.parent.Sha1)
	}
	content += fmt.Sprintf("author %s\n", c.author)
	content += fmt.Sprintf("committer %s\n", c.committer)
	content += fmt.Sprintf("\n%s\n", c.message)
	LogI.Println(content)

	b.Write([]byte(content))
	return b.Bytes()
}
