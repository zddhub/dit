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
	repo := &repository{isInitialized: false}
	return repo
}

// dit repository
type repository struct {
	isInitialized bool
}

func (r *repository) Init() {
	absRepoDir, err := filepath.Abs(DIT["dir"])

	if err == nil && isExist(DIT["HEAD"]) {
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
	if err := index.storeCache(DIT["index"]); err != nil {
		LogE.Println(err)
	}
}

// true if file exists
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		return false
	}
	return true
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
