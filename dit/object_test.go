package dit

import (
	"os"
	"testing"
)

var (
	// treated as constants
	Sha1_Bytes  = [20]byte{143, 44, 150, 173, 103, 109, 116, 35, 210, 195, 25, 255, 251, 120, 207, 184, 124, 120, 195, 226}
	Sha1_String = "8f2c96ad676d7423d2c319fffb78cfb87c78c3e2"
)

func TestWriteAndReadObject(t *testing.T) {
	data := []byte("dit\n")
	object := object{flag: "blob"}
	object.Write(data)

	if object.Sha1String() != Sha1_String {
		t.Error("sum sha1 error when write object")
	}

	filePath := DIT_REPO_DIR + "/objects/" + string(object.Sha1String()[:2]) + "/" + object.Sha1String()[2:]
	if _, err := os.Stat(filePath); err != nil {
		t.Error("write object to repo error")
	}

	// Read
	object.flag = ""
	object.size = 0

	buffer, err := object.ReadAll()

	if err != nil || object.flag != "blob" || object.size != 4 {
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

	os.RemoveAll(DIT_REPO_DIR)
}

func TestBytesToSha1(t *testing.T) {
	if BytesToSha1(Sha1_Bytes) != Sha1_String {
		t.Error("BytesToSha1 error")
	}
}

func TestSha1ToBytes(t *testing.T) {
	if Sha1ToBytes(Sha1_String) != Sha1_Bytes {
		t.Error("Sha1ToBytes error")
	}
}
