package dit

import (
	"path/filepath"

	. "github.com/zddhub/dit/utils"
)

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
	file.Write([]byte("ref: refs/heads/master\n"))
	LogT.Println("Initialized empty Dit repository in", absRepoDir)
	r.isInitialized = true
}
