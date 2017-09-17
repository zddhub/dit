package dit

import (
	"os"
)

const (
	DefaultDitPath = "."
	DefaultDitRepo = ".dit"
)

var (
	DitPath = os.Getenv("DIT_PATH")
	DitRepo = os.Getenv("DIT_REPO")

	DIT map[string]string
)

func init() {
	if DitPath == "" {
		DitPath = DefaultDitPath
	}
	if DitRepo == "" {
		DitRepo = DefaultDitRepo
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
	return DitPath + "/" + DitRepo
}

func GetSubPath(name string) string {
	return GetRepoPath() + "/" + name
}
