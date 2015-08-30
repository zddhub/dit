package dit

import (
	"testing"
)

func TestHashFromMemory(t *testing.T) {
	data := []byte("blob 4\x00dit\n")
	sha1 := MemHash(data)
	if sha1 != "8f2c96ad676d7423d2c319fffb78cfb87c78c3e2" {
		t.Error("sha1 from memory error")
	}
}

func TestHashFromFile(t *testing.T) {
	sha1, _ := FileHash("tests/dit")

	if sha1 != "8f2c96ad676d7423d2c319fffb78cfb87c78c3e2" {
		t.Error("sha1 from file error")
	}
}

func TestHashFromFileErr(t *testing.T) {
	sha1, err := FileHash("tests/notexistfile")

	if sha1 != "" || err == nil {
		t.Error("")
	}
}
