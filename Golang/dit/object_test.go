package dit

import (
	"os"
	"testing"
)

func TestWriteAndReadObject(t *testing.T) {
	data := []byte("dit\n")
	object := object{Type: "blob"}
	object.Write(data)

	if object.Sha1 != "8f2c96ad676d7423d2c319fffb78cfb87c78c3e2" {
		t.Error("sum sha1 error when write object")
	}

	filePath := DIT["objects"] + "/" + object.Sha1[:2] + "/" + object.Sha1[2:]
	if _, err := os.Stat(filePath); err != nil {
		t.Error("write object to repo error")
	}

	// Read
	object.Type = ""
	object.Size = 0

	buffer, err := object.ReadAll()

	if err != nil || object.Type != "blob" || object.Size != 4 {
		t.Error("read object info error: ", err)
	}

	if len(buffer) != len(data) {
		t.Error("read object buffer error")
	}

	for i := range buffer {
		if buffer[i] != data[i] {
			t.Error("read object content error")
		}
	}

	os.RemoveAll(DIT["dir"])
}
