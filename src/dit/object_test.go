package dit

import (
	"crypto/sha1"
	"os"
	"testing"
)

func TestWriteAndReadObject(t *testing.T) {
	data := []byte("dit\n")
	object := object{flag: "blob", sha1: sha1.Sum(data)}
	object.Write(data)

	object.flag = ""

	buffer := make([]byte, len(data))
	object.Read(buffer)

	if object.flag != "blob" {
		t.Error("read object type error")
	}

	filePath := DIT_REPO_DIR + "/objects/" + string(object.Sha1String()[:2]) + "/" + object.Sha1String()[2:]
	if _, err := os.Stat(filePath); err != nil {
		t.Error("write and read object error")
	}

	os.RemoveAll(DIT_REPO_DIR)
}
