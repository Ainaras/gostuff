package dataloader

import (
	"testing"
)

func TestLoadData(t *testing.T) {
	data := LoadData("f1/races_with_tracks.txt")
	if data == "" {
		t.Error("Expected data to be loaded")
	}
}
