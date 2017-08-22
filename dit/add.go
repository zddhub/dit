package dit

import (
	. "github.com/zddhub/dit/utils"
)

func AddFileToObjects(filePath string) (obj *object, err error) {
	buffer, err := ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	object := &object{flag: "blob"}
	object.Write(buffer)
	return object, err
}
