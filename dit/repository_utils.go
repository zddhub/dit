package dit

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/zddhub/dit/utils"
)

func branch() string {
	head, _ := utils.ReadFile(DIT["HEAD"])
	branch := strings.Join(strings.Split(string(head), "/")[2:], "/")
	return strings.Trim(branch, "\n")
}

func head() string {
	head, _ := utils.ReadFile(DIT["HEAD"])
	branch := strings.Split(string(head), " ")[1]
	sha1, _ := utils.ReadFile(DIT["dir"] + "/" + branch)
	return strings.Trim(string(sha1), "\n")
}

func checkRepositoryExist() bool {
	return utils.IsExist(DIT["HEAD"])
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
