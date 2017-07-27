package hash

import (
	"os/exec"
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
	cmd := exec.Command("bash", "-c", "mkdir -p testdata; echo -e 'blob 4\\0dit' > testdata/dit")
	cmd.Run()

	sha1, _ := FileHash("testdata/dit")

	if sha1 != "8f2c96ad676d7423d2c319fffb78cfb87c78c3e2" {
		t.Error("sha1 from file error")
	}
}

func TestHashFromFileErr(t *testing.T) {
	sha1, err := FileHash("testdata/notexistfile")

	if sha1 != "" || err == nil {
		t.Error("")
	}
}
