package dit

import (
	. "github.com/zddhub/dit/utils"
)

func NewRepository() *repository {
	return &repository{isInitialized: false}
}

func LoadRepository() *repository {
	if !checkRepositoryExist() {
		LogT.Fatalln("fatal: Not a dit repository (or any of the parent directories):", DitRepo)
		return nil
	}
	var cache cache
	cache.loadCache(DIT["index"])

	repo := &repository{true, cache, Branch(), Head()}
	LogI.Println(*repo)
	return repo
}

// dit repository
type repository struct {
	isInitialized bool
	cache         cache
	branch        string
	head          string
}

func (r *repository) StoreCache() {
	if err := r.cache.storeCache(DIT["index"]); err != nil {
		LogE.Println(err)
	}
}

func (r *repository) AddCacheEntry(obj *object) {
	if yes, i := r.cache.includes(obj); yes {
		r.cache.Entries[i] = obj
	} else {
		r.cache.Entries = append(r.cache.Entries, obj)
	}
}

func (r repository) WriteHead(sha1 string) {
	if err := WriteFile(DIT["refs/heads"]+"/"+r.branch, []byte(sha1+"\n"), 0644); err != nil {
		LogE.Println(err)
	}
}
