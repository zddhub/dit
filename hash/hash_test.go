package hash

import (
	"os/exec"
	"testing"
)

func TestHashToBytesFromMemory(t *testing.T) {
	data := []byte("blob 4\x00dit\n")
	sha1_code := [20]byte{143, 44, 150, 173, 103, 109, 116, 35, 210, 195, 25, 255, 251, 120, 207, 184, 124, 120, 195, 226}
	sha1 := MemHashToBytes(data)

	if sha1 != sha1_code {
		t.Error("sha1 to bytes from memory error")
	}
}

func TestHashToStringFromMemory(t *testing.T) {
	data := []byte("blob 4\x00dit\n")
	sha1 := MemHashToString(data)
	if sha1 != "8f2c96ad676d7423d2c319fffb78cfb87c78c3e2" {
		t.Error("sha1 to string from memory error")
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
