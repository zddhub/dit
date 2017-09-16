package dit

import (
	. "github.com/zddhub/dit/utils"
	"os"
	"path/filepath"
)

type Repository interface {
	Init()
	AddFiles()
}

func NewRepository() *repository {
	return &repository{isInitialized: false}
}

func LoadRepository() *repository {
	if !checkRepositoryExist() {
		LogT.Fatalln("fatal: Not a dit repository (or any of the parent directories):", DIT_REPO)
		return nil
	}
	var index cache
	index.loadCache(DIT["index"])
	return &repository{isInitialized: true, index: index}
}

// dit repository
type repository struct {
	isInitialized bool
	index         cache
}

func (r *repository) Init() {
	absRepoDir, err := filepath.Abs(DIT["dir"])

	if err == nil && checkRepositoryExist() {
		r.isInitialized = true
		LogT.Println("Reinitialized empty Dit repository in", absRepoDir)
		return
	}

	createDir(DIT["objects"])
	createDir(DIT["refs"] + "/heads")
	file, err := safeCreateFile(DIT["HEAD"])
	if err != nil {
		LogE.Fatalln(err)
	}
	defer file.Close()
	file.Write([]byte("ref: refs/heads/master"))
	LogT.Println("Initialized empty Dit repository in", absRepoDir)
}

func (r *repository) StoreCache() {
	if err := r.index.storeCache(DIT["index"]); err != nil {
		LogE.Println(err)
	}
}

func (r *repository) AddCacheEntry(obj *object) {
	if yes, i := r.includes(obj); yes {
		r.index.Entries[i] = obj
	} else {
		r.index.Entries = append(r.index.Entries, obj)
	}
}

func (r *repository) includes(obj *object) (bool, int) {
	for i, entry := range r.index.Entries {
		if obj.Sha1 == entry.Sha1 {
			return true, i
		}
	}
	return false, 0
}

func checkRepositoryExist() bool {
	return IsExist(DIT["HEAD"])
}

// create file, if no directory, create directory
func safeCreateFile(path string) (*os.File, error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// create dir
func createDir(path string) {
	os.MkdirAll(path, 0755)
}
