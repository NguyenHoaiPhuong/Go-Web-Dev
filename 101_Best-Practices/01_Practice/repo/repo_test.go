package repo

import (
	"testing"
)

func TestGetDatabase(t *testing.T) {
	_, err := GetDatabase()
	if err != nil {
		t.Error("Database initalization error")
	}
}
