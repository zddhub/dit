package dit

import (
	"github.com/zddhub/dit/utils"
)

const emptyFileMessage = `Nothing specified, nothing added.
Maybe you wanted to say 'dit add .'?`

func (repo *repository) Add(args []string) {
	files := args
	if len(files) == 0 {
		utils.LogT.Println(emptyFileMessage)
		return
	}

	for _, file := range files {
		obj, _ := addFileToObjects(file)
		repo.AddCacheEntry(obj)
	}
	repo.SetInvalidTreeCache()
	repo.StoreCache()
}

func addFileToObjects(filename string) (obj *object, err error) {
	buffer, err := utils.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	mode, _ := utils.FileMode(filename)
	object := &object{Type: "blob", Mode: mode, Path: filename}
	object.Write(buffer)
	return object, err
}
