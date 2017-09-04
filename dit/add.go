package dit

import (
	_ "fmt"
	. "github.com/zddhub/dit/utils"
)

const EmptyFileMessage = `Nothing specified, nothing added.
Maybe you wanted to say 'git add .'?`

func (repo *repository) AddFiles(files ...string) {
	if len(files) == 0 {
		LogT.Println(EmptyFileMessage)
		return
	}

	for _, file := range files {
		addFileToObjects(file)
	}
}

func addFileToObjects(filePath string) (obj *object, err error) {
	buffer, err := ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	object := &object{flag: "blob"}
	object.Write(buffer)
	return object, err
}
