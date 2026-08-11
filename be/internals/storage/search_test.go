package storage

import "testing"

func TestSearchRepository_Interface(t *testing.T) {
	var _ SearchRepository = (*Repository)(nil)
}
