package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// FileSystem to manage local database
type FileSystem struct {
	ditDir        string
	isInitialized bool
}

func (f *FileSystem) createDir(dirName string) error {
	return os.Mkdir(dirName, 0755)
}

func (f *FileSystem) createFile(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
}

func (f *FileSystem) initFileSystem() {
	if _, err := os.Stat(f.ditDir + "/" + ".dit/HEAD"); os.IsNotExist(err) {
		f.isInitialized = false
	} else {
		f.isInitialized = true
	}

	f.createDir(f.ditDir + "/" + ".dit")
	f.ditDir, _ = filepath.Abs(f.ditDir + "/.dit")
	f.ditDir += "/"

	file, err := f.createFile(f.ditDir + "HEAD")
	if err != nil {
		fmt.Println("fatal: create HEAD error: ", err)
	}
	defer file.Close()

	f.createDir(f.ditDir + "objects")
	f.createDir(f.ditDir + "refs")
	f.createDir(f.ditDir + "refs/heads")

	if f.isInitialized == true {
		fmt.Println("Reinitialized empty Dit repository in", f.ditDir)
	} else {
		fmt.Println("Initialized empty Dit repository in", f.ditDir)
	}
}

func main() {
	fs := FileSystem{".", true}
	fs.initFileSystem()
}
