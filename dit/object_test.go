package dit

import (
	"os"
	"testing"
)

func TestWriteAndReadObject(t *testing.T) {
	data := []byte("dit\n")
	object := object{flag: "blob"}
	object.Write(data)

	if object.Sha1String() != "8f2c96ad676d7423d2c319fffb78cfb87c78c3e2" {
		t.Error(object.Sha1String())
		t.Error("sum sha1 error when write object")
	}

	object.flag = ""

	buffer := make([]byte, len(data))
	object.Read(buffer)

	if object.flag != "blob" {
		t.Error("read object type error")
	}

	filePath := DIT_REPO_DIR + "/objects/" + string(object.Sha1String()[:2]) + "/" + object.Sha1String()[2:]
	if _, err := os.Stat(filePath); err != nil {
		t.Error("write object to repo error")
	}

	for i := range buffer {
		if buffer[i] != data[i] {
			t.Error("read object content error")
		}
	}

	os.RemoveAll(DIT_REPO_DIR)
}
