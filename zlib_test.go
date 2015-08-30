package dit

import (
	"testing"
)

// Test compress and uncompress algorithm
func TestZlib(t *testing.T) {
	data := []byte("blob 4\x00dit\n")
	Compress(data, "tests/zlib_test")

	buffer := make([]byte, len(data))
	Uncompress("tests/zlib_test", buffer)

	if string(buffer) != string(data) {
		t.Error("compress and uncompress error")
	}
}
