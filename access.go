package maputil

func CheckEnum(s string, allowed []string) error {
	for _, v := range allowed {
		if v == s {
			return nil
		}
	}
	return EnumStringError{
		Value: s,
		Enum:  allowed,
	}
}

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

	return s, true, CheckEnum(s, values)
}

// OptionalArray fetches a value from the map and converts it to an array.
func OptionalArray(m map[string]interface{}, key string, dv []interface{}) ([]interface{}, error) {
	v, ok := m[key]
	if !ok {
		return dv, nil
	}
	a, err := AsArray(v)
	if err != nil {
		return dv, err
	}
	return a, nil
}

// OptionalBoolean fetches a value from the map and converts it to a boolean.
func OptionalBoolean(m map[string]interface{}, key string, dv bool) (bool, error) {
	v, ok := m[key]
	if !ok {
		return dv, nil
	}
	b, err := AsBoolean(v)
	if err != nil {
		return dv, err
	}
	return b, nil
}

// OptionalInteger fetches a value from the map and converts it to an integer.
func OptionalInteger(m map[string]interface{}, key string, dv int64) (int64, error) {
	v, ok := m[key]
	if !ok {
		return dv, nil
	}
	i, err := AsInteger(v)
	if err != nil {
		return dv, err
	}
	return i, nil
}

// OptionalNull fetches a value from the map and ensures it was null.
func OptionalNull(m map[string]interface{}, key string) error {
	v, ok := m[key]
	if !ok {
		return nil
	}
	if v != nil {
		return InvalidTypeError{
			Expected: []string{TypeNull},
			Actual:   TypeName(v),
		}
	}
	return nil
}

// OptionalNumber fetches a value from the map and converts it to a number.
func OptionalNumber(m map[string]interface{}, key string, dv float64) (float64, error) {
	v, ok := m[key]
	if !ok {
		return dv, nil
	}
	n, err := AsNumber(v)
	if err != nil {
		return dv, err
	}
	return n, nil
}

// OptionalObject fetches a value from the map and converts it to an object.
func OptionalObject(m map[string]interface{}, key string, dv map[string]interface{}) (map[string]interface{}, error) {
	v, ok := m[key]
	if !ok {
		return dv, nil
	}
	o, err := AsObject(v)
	if err != nil {
		return dv, err
	}
	return o, nil
}

// OptionalString fetches a value from the map and converts it to a string.
func OptionalString(m map[string]interface{}, key, dv string) (string, error) {
	v, ok := m[key]
	if !ok {
		return dv, nil
	}
	s, err := AsString(v)
	if err != nil {
		return dv, err
	}
	return s, nil
}

// GetStringEnum fetches a value from the map, converts it to a string, and
// ensures it is one of the given values.
func OptionalStringEnum(m map[string]interface{}, key string, values []string, dv string) (string, error) {
	v, ok := m[key]
	if !ok {
		return dv, nil
	}
	s, err := AsString(v)
	if err != nil {
		return dv, err
	}
	if err := CheckEnum(s, values); err != nil {
		return dv, err
	}
	return s, nil
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
	return s, true, CheckEnum(s, values)
}

// RequireArray fetches a value from the map and converts it to an array.
func RequireArray(m map[string]interface{}, key string) ([]interface{}, error) {
	v, ok := m[key]
	if !ok {
		return nil, MissingRequiredValueError{Key: key}
	}
	return AsArray(v)
}

// RequireBoolean fetches a value from the map and converts it to a boolean.
func RequireBoolean(m map[string]interface{}, key string) (bool, error) {
	v, ok := m[key]
	if !ok {
		return false, MissingRequiredValueError{Key: key}
	}
	return AsBoolean(v)
}

// RequireInteger fetches a value from the map and converts it to an integer.
func RequireInteger(m map[string]interface{}, key string) (int64, error) {
	v, ok := m[key]
	if !ok {
		return 0, MissingRequiredValueError{Key: key}
	}
	return AsInteger(v)
}

// RequireNull fetches a value from the map and ensures it was null.
func RequireNull(m map[string]interface{}, key string) error {
	v, ok := m[key]
	if !ok {
		return MissingRequiredValueError{Key: key}
	}
	if v != nil {
		return InvalidTypeError{
			Expected: []string{TypeNull},
			Actual:   TypeName(v),
		}
	}
	return nil
}

// RequireNumber fetches a value from the map and converts it to a number.
func RequireNumber(m map[string]interface{}, key string) (float64, error) {
	v, ok := m[key]
	if !ok {
		return 0, MissingRequiredValueError{Key: key}
	}
	return AsNumber(v)
}

// RequireObject fetches a value from the map and converts it to an object.
func RequireObject(m map[string]interface{}, key string) (map[string]interface{}, error) {
	v, ok := m[key]
	if !ok {
		return nil, MissingRequiredValueError{Key: key}
	}
	return AsObject(v)
}

// RequireString fetches a value from the map and converts it to a string.
func RequireString(m map[string]interface{}, key string) (string, error) {
	v, ok := m[key]
	if !ok {
		return "", MissingRequiredValueError{Key: key}
	}
	return AsString(v)
}

// RequireStringEnum fetches a value from the map, converts it to a string, and
// ensures it is one of the given values.
func RequireStringEnum(m map[string]interface{}, key string, values []string) (string, error) {
	v, ok := m[key]
	if !ok {
		return "", MissingRequiredValueError{Key: key}
	}
	s, err := AsString(v)
	if err != nil {
		return "", err
	}

	return s, CheckEnum(s, values)
}
