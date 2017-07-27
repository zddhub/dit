package dit

import (
	"testing"
)

// Test compress and uncompress algorithm
func TestCompress(t *testing.T) {
	data := []byte("blob 4\x00dit\n")
	Compress("testdata/compress_test", data)

	buffer := make([]byte, len(data))
	Uncompress(buffer, "testdata/compress_test")

	if string(buffer) != string(data) {
		t.Error("compress and uncompress error")
	}
}
