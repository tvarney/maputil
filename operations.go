package maputil

// Keys returns the keys of a map.
func Keys(m map[string]interface{}) []string {
	if len(m) == 0 {
		return nil
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Copy makes a deep copy of the given map.
//
// Only data structures which are JSON-like are handled by this function; a
// struct pointer in a map will not be deep-copied.
func Copy(m map[string]interface{}) map[string]interface{} {
	if len(m) == 0 {
		if m == nil {
			return nil
		}
		return map[string]interface{}{}
	}
	r := make(map[string]interface{}, len(m))
	for k, v := range m {
		switch d := v.(type) {
		case map[string]interface{}:
			r[k] = Copy(d)
		case []interface{}:
			r[k] = CopyArray(d)
		default:
			r[k] = d
		}
	}
	return r
}

// CopyArray makes a deep copy of the given array.
//
// Only data structures which are JSON-like are handled by this function; a
// struct pointer in the array will not be deep-copied.
func CopyArray(a []interface{}) []interface{} {
	if len(a) == 0 {
		return nil
	}
	r := make([]interface{}, 0, len(a))
	for _, v := range a {
		switch d := v.(type) {
		case map[string]interface{}:
			r = append(r, Copy(d))
		case []interface{}:
			r = append(r, CopyArray(d))
		default:
			r = append(r, d)
		}
	}
	return r
}
