package util

import "testing"

func TestCopyMap(t *testing.T) {
	t.Logf("TestCopyMap")

	m := map[string]interface{}{
		"key1": "value1",
		"key2": map[string]interface{}{
			"key3": "value3",
		},
	}

	cp := CopyMap(m)

	cp["key1"] = "value1_changed"
	cp["key2"].(map[string]interface{})["key3"] = "value3_changed"

	if m["key1"] == cp["key1"] {
		t.Errorf("expected m['key1'] to be different from cp['key1']")
	}

	if m["key2"].(map[string]interface{})["key3"] == cp["key2"].(map[string]interface{})["key3"] {
		t.Errorf("expected m['key2']['key3'] to be different from cp['key2']['key3']")
	}
}
