package compress

import (
	"testing"
)

// Test compress and decompress algorithm
func TestCompress(t *testing.T) {
	data := []byte("blob 4\x00dit\n")
	Compress("../testdata/compress_test", data)

	buffer, err := Decompress("../testdata/compress_test")

	if err != nil || string(buffer) != string(data) {
		t.Error(buffer, "compress and decompress error")
	}
}
