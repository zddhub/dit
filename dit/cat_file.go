package dit

import (
	"fmt"
)

func (repo *repository) CatFile(sha1 string) {
	object := &object{sha1: Sha1ToBytes(sha1)}
	buffer := make([]byte, 512)
	object.Read(buffer)
	fmt.Printf("%s", buffer)
}
