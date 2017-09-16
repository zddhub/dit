package dit

type commit struct {
	object
	tree   *tree
	parent *commit
}

func (repo *repository) Commit(message string) {
	_ = repo.NewTree()
	// commit := commit{Type: "commit", tree: tree}
}
