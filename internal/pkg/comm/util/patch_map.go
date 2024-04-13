package util

func PatchMap(m map[string]interface{}, patch map[string]interface{}) {
	// recursively patch map
	for k, v := range patch {
		if _, ok := m[k]; ok {
			if _, ok := v.(map[string]interface{}); ok {
				PatchMap(m[k].(map[string]interface{}), v.(map[string]interface{}))
			} else {
				m[k] = v
			}
		} else {
			m[k] = v
		}
	}
}
