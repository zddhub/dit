package utils

import (
	"testing"
)

var (
	// treated as constants
	Sha1Bytes  = [20]byte{143, 44, 150, 173, 103, 109, 116, 35, 210, 195, 25, 255, 251, 120, 207, 184, 124, 120, 195, 226}
	Sha1String = "8f2c96ad676d7423d2c319fffb78cfb87c78c3e2"
)

func TestBytesToSha1(t *testing.T) {
	if BytesToSha1(Sha1Bytes) != Sha1String {
		t.Error("BytesToSha1 error")
	}
}

func TestSha1ToBytes(t *testing.T) {
	if Sha1ToBytes(Sha1String) != Sha1Bytes {
		t.Error("Sha1ToBytes error")
	}
}
