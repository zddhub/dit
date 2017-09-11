package dit

import (
	. "github.com/zddhub/dit/utils"
)

const EmptyFileMessage = `Nothing specified, nothing added.
Maybe you wanted to say 'git add .'?`

func (repo *repository) AddFiles(files []string) {
	if len(files) == 0 {
		LogT.Println(EmptyFileMessage)
		return
	}

	for _, file := range files {
		obj, _ := addFileToObjects(file)
		repo.AddCacheEntry(obj)
	}
}

func addFileToObjects(filename string) (obj *object, err error) {
	buffer, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}

	mode, _ := FileMode(filename)
	object := &object{Type: "blob", Mode: mode, Path: filename}
	object.Write(buffer)
	return object, err
}
