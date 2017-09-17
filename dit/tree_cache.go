package dit

type cachedTree struct {
	Signature string
	Entries   []*cachedTreeEntry
}

type cachedTreeEntry struct {
	Dirname       string
	Sha1          string
	ObjectsCount  int
	SubTreesCount int
}

func (c commit) newTreeCache() *cachedTree {
	cachedTree := &cachedTree{"TREE", nil}
	cachedTreeEntry := &cachedTreeEntry{c.tree.Path, c.tree.Sha1, len(c.tree.objects), 0}
	cachedTree.Entries = append(cachedTree.Entries, cachedTreeEntry)
	return cachedTree
}

func (repo *repository) NewInvalidTreeCache() *cachedTree {
	return &cachedTree{"TREE", nil}
}

func (repo *repository) SetInvalidTreeCache() {
	repo.cache.Extensions = repo.NewInvalidTreeCache()
}
