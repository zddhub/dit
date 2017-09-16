package dit

import (
	"os"
)

const (
	DEFAULT_DIT_PATH = "."
	DEFAULT_DIT_REPO = ".dit"
)

var (
	DIT_PATH = os.Getenv("DIT_PATH")
	DIT_REPO = os.Getenv("DIT_REPO")

	DIT map[string]string
)

func init() {
	if DIT_PATH == "" {
		DIT_PATH = DEFAULT_DIT_PATH
	}
	if DIT_REPO == "" {
		DIT_REPO = DEFAULT_DIT_REPO
	}

	setDitEnv()
}

func setDitEnv() {
	DIT = map[string]string{
		"dir":        GetRepoPath(),
		"objects":    GetSubPath("objects"),
		"refs":       GetSubPath("refs"),
		"refs/heads": GetSubPath("refs/heads"),
		"HEAD":       GetSubPath("HEAD"),
		"index":      GetSubPath("index"),
	}
}

func GetRepoPath() string {
	return DIT_PATH + "/" + DIT_REPO
}

func GetSubPath(name string) string {
	return GetRepoPath() + "/" + name
}
