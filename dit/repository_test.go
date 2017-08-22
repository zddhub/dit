package dit

import (
	"testing"
)

func TestRepositoryInit(t *testing.T) {
	repo := NewRepository()
	repo.Init()
}
