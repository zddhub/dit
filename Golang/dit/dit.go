package dit

import (
	"os"
)

const (
	defaultDitPath = "."
	defaultDitRepo = ".dit"
)

var (
	// DitPath get DIT_PATH from env, default value is "."
	DitPath = os.Getenv("DIT_PATH")

	// DitRepo get DIT_REPO from env, default value is ".dit"
	DitRepo = os.Getenv("DIT_REPO")

	// DIT holds dit repo file system path
	DIT map[string]string
)

func init() {
	if DitPath == "" {
		DitPath = defaultDitPath
	}
	if DitRepo == "" {
		DitRepo = defaultDitRepo
	}

	setDitEnv()
}

func setDitEnv() {
	DIT = map[string]string{
		"dir":        repoPath(),
		"objects":    subPath("objects"),
		"refs":       subPath("refs"),
		"refs/heads": subPath("refs/heads"),
		"HEAD":       subPath("HEAD"),
		"index":      subPath("index"),
	}
}

func repoPath() string {
	return DitPath + "/" + DitRepo
}

func subPath(name string) string {
	return repoPath() + "/" + name
}
