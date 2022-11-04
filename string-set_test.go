package mirage

import "testing"

func TestHas(t *testing.T) {
	x := NewStringSet([]string{"hello", "world", "golf"})
	if !x.Has("world") {
		t.Errorf("Expected member of set was not found")
	}
}
