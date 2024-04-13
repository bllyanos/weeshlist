package util

import (
	"testing"
)

func TestPatchMap(t *testing.T) {
	t.Logf("TestPatchMap")

	m := map[string]interface{}{"a": 1, "b": 2, "d": map[string]interface{}{"e": 5, "f": 6}}
	patch := map[string]interface{}{"a": 3, "c": 4, "d": map[string]interface{}{"e": 6}}
	PatchMap(m, patch)

	if m["a"] != 3 {
		t.Errorf("Expected m[a] to be 3, got %v", m["a"])
	}
	if m["b"] != 2 {
		t.Errorf("Expected m[b] to be 2, got %v", m["b"])
	}
	if m["c"] != 4 {
		t.Errorf("Expected m[c] to be 4, got %v", m["c"])
	}
	if m["d"].(map[string]interface{})["e"] != 6 {
		t.Errorf("Expected m[d][e] to be 6, got %v", m["d"].(map[string]interface{})["e"])
	}
}
