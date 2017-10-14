package dit

import (
	"time"

	utils "github.com/zddhub/dit/golang/utils"
)

func (repo *repository) Commit(message string) {
	if !repo.validCommit() {
		utils.LogT.Println("nothing added to commit")
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
