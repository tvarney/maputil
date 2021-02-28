package maputil

// GetArray fetches a value from the map and converts it to an array.
func GetArray(m map[string]interface{}, key string) ([]interface{}, bool, error) {
	v, ok := m[key]
	if !ok {
		return nil, false, nil
	}
	a, err := AsArray(v)
	return a, true, err
}

// GetBoolean fetches a value from the map and converts it to a boolean.
func GetBoolean(m map[string]interface{}, key string) (bool, bool, error) {
	v, ok := m[key]
	if !ok {
		return false, false, nil
	}
	b, err := AsBoolean(v)
	return b, true, err
}

// GetInteger fetches a value from the map and converts it to an integer.
func GetInteger(m map[string]interface{}, key string) (int64, bool, error) {
	v, ok := m[key]
	if !ok {
		return 0, false, nil
	}
	i, err := AsInteger(v)
	return i, true, err
}

// GetNull fetches a value from the map and ensures it was null.
func GetNull(m map[string]interface{}, key string) (bool, error) {
	v, ok := m[key]
	if !ok {
		return false, nil
	}
	if v != nil {
		return true, InvalidTypeError{
			Expected: []string{TypeNull},
			Actual:   TypeName(v),
		}
	}
	return true, nil
}

// GetNumber fetches a value from the map and converts it to a number.
func GetNumber(m map[string]interface{}, key string) (float64, bool, error) {
	v, ok := m[key]
	if !ok {
		return 0, false, nil
	}
	f, err := AsNumber(v)
	return f, true, err
}

// GetObject fetches a value from the map and converts it to an object.
func GetObject(m map[string]interface{}, key string) (map[string]interface{}, bool, error) {
	v, ok := m[key]
	if !ok {
		return nil, false, nil
	}
	m, err := AsObject(v)
	return m, true, err
}

// GetString fetches a value from the map and converts it to a string.
func GetString(m map[string]interface{}, key string) (string, bool, error) {
	v, ok := m[key]
	if !ok {
		return "", false, nil
	}
	s, err := AsString(v)
	return s, true, err
}

// GetStringEnum fetches a value from the map, converts it to a string, and
// ensures it is one of the given values.
func GetStringEnum(m map[string]interface{}, key string, values []string) (string, bool, error) {
	v, ok := m[key]
	if !ok {
		return "", false, nil
	}
	s, err := AsString(v)
	if err != nil {
		return "", true, err
	}

	for _, v := range values {
		if v == s {
			return s, true, nil
		}
	}
	return s, true, EnumStringError{
		Value: s,
		Enum:  values,
	}
}

// PopArray fetches a value from the map and converts it to an array.
func PopArray(m map[string]interface{}, key string) ([]interface{}, bool, error) {
	v, ok := m[key]
	if !ok {
		return nil, false, nil
	}
	delete(m, key)
	a, err := AsArray(v)
	return a, true, err
}

// PopBoolean fetches a value from the map and converts it to a boolean.
func PopBoolean(m map[string]interface{}, key string) (bool, bool, error) {
	v, ok := m[key]
	if !ok {
		return false, false, nil
	}
	delete(m, key)
	b, err := AsBoolean(v)
	return b, true, err
}

// PopInteger fetches a value from the map and converts it to an integer.
func PopInteger(m map[string]interface{}, key string) (int64, bool, error) {
	v, ok := m[key]
	if !ok {
		return 0, false, nil
	}
	delete(m, key)
	i, err := AsInteger(v)
	return i, true, err
}

// PopNull fetches a value from the map and ensures it was null.
func PopNull(m map[string]interface{}, key string) (bool, error) {
	v, ok := m[key]
	if !ok {
		return false, nil
	}
	delete(m, key)
	if v != nil {
		return true, InvalidTypeError{
			Expected: []string{TypeNull},
			Actual:   TypeName(v),
		}
	}
	return true, nil
}

// PopNumber fetches a value from the map and converts it to a number.
func PopNumber(m map[string]interface{}, key string) (float64, bool, error) {
	v, ok := m[key]
	if !ok {
		return 0, false, nil
	}
	delete(m, key)
	n, err := AsNumber(v)
	return n, true, err
}

// PopObject fetches a value from the map and converts it to an object.
func PopObject(m map[string]interface{}, key string) (map[string]interface{}, bool, error) {
	v, ok := m[key]
	if !ok {
		return nil, false, nil
	}
	delete(m, key)
	o, err := AsObject(v)
	return o, true, err
}

// PopString fetches a value from the map and converts it to a string.
func PopString(m map[string]interface{}, key string) (string, bool, error) {
	v, ok := m[key]
	if !ok {
		return "", false, nil
	}
	delete(m, key)
	s, err := AsString(v)
	return s, true, err
}

// PopStringEnum fetches a value from the map, converts it to a string, and
// ensures it is one of the given values.
func PopStringEnum(m map[string]interface{}, key string, values []string) (string, bool, error) {
	v, ok := m[key]
	if !ok {
		return "", false, nil
	}
	delete(m, key)
	s, err := AsString(v)
	if err != nil {
		return "", true, err
	}
	for _, v := range values {
		if v == s {
			return s, true, nil
		}
	}
	return s, false, EnumStringError{
		Value: s,
		Enum:  values,
	}
}
