package dit

import (
	"os/exec"
	"testing"
)

func TestAddFileToObjects(t *testing.T) {
	cmd := exec.Command("bash", "-c", "mkdir -p ../testdata; echo 'dit' > ../testdata/dit")
	cmd.Run()

	object, _ := addFileToObjects("../testdata/dit")

	if object.Sha1String() != "8f2c96ad676d7423d2c319fffb78cfb87c78c3e2" {
		t.Error("sha1 from file error")
	}
}
