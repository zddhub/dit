package dit

import (
	"github.com/zddhub/dit/golang/utils"
)

// NewRepository tries to init dit repo
func NewRepository() *repository {
	return &repository{isInitialized: false}
}

// LoadRepository loads an exist dit repo.
func LoadRepository() *repository {
	if !checkRepositoryExist() {
		utils.LogT.Fatalln("fatal: Not a dit repository (or any of the parent directories):", DitRepo)
		return nil
	}
	var cache cache
	cache.loadCache(DIT["index"])

	repo := &repository{true, cache, branch(), head()}
	utils.LogI.Println(*repo)
	return repo
}

// Repository host repo
type repository struct {
	isInitialized bool
	cache         cache
	branch        string
	head          string
}

func (r *repository) StoreCache() {
	if err := r.cache.storeCache(DIT["index"]); err != nil {
		utils.LogE.Println(err)
	}
}

func (r *repository) AddCacheEntry(obj *object) {
	r.cache.addEntry(obj)
}

func (r repository) WriteHead(sha1 string) {
	if err := utils.WriteFile(DIT["refs/heads"]+"/"+r.branch, []byte(sha1+"\n"), 0644); err != nil {
		utils.LogE.Println(err)
	}
}
