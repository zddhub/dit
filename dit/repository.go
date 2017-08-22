package dit

import (
	. "github.com/zddhub/dit/utils"
	"os"
	"path/filepath"
)

type Repository interface {
	Init()
}

func NewRepository() *repository {
	repo := &repository{dir: DEFAULT_REPO_PATH, isInitialized: false}
	return repo
}

// dit repository
type repository struct {
	dir           string
	isInitialized bool
}

func (r *repository) Init() {
	absRepoDir, err := filepath.Abs(r.dir + "/" + DIT_REPO_DIR)

	if err == nil && isExist(absRepoDir+"/HEAD") {
		r.isInitialized = true
		LogT.Println("Reinitialized empty Dit repository in", absRepoDir)
		return
	}

	createDir(absRepoDir + "/objects")
	createDir(absRepoDir + "/refs/heads")
	file, err := safeCreateFile(absRepoDir + "/HEAD")
	if err != nil {
		LogE.Fatalln(err)
	}
	defer file.Close()
	file.Write([]byte("ref: refs/heads/master"))
	LogT.Println("Initialized empty Dit repository in", absRepoDir)
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
