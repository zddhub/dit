package dit

import (
	"testing"
)

func TestRepositoryInit(t *testing.T) {
	repo := CreateRepository()
	repo.Init()
}
