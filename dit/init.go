package dit

import (
	"path/filepath"

	utils "github.com/zddhub/dit/utils"
)

func (r *repository) Init() {
	absRepoDir, err := filepath.Abs(DIT["dir"])

	if err == nil && checkRepositoryExist() {
		r.isInitialized = true
		utils.LogT.Println("Reinitialized empty Dit repository in", absRepoDir)
		return
	}

	createDir(DIT["objects"])
	createDir(DIT["refs"] + "/heads")
	file, err := safeCreateFile(DIT["HEAD"])
	if err != nil {
		utils.LogE.Fatalln(err)
	}
	defer file.Close()
	file.Write([]byte("ref: refs/heads/master\n"))
	utils.LogT.Println("Initialized empty Dit repository in", absRepoDir)
	r.isInitialized = true
}
