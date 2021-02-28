package maputil

import (
	"fmt"
	"math"
)

// Type name constants encode the JSON-like names of the types handled by the
// maputil package.
const (
	TypeArray   = "array"
	TypeBoolean = "boolean"
	TypeInteger = "integer"
	TypeNull    = "null"
	TypeNumber  = "number"
	TypeObject  = "object"
	TypeString  = "string"
)

// GenericNumber is a number type which allows for querying for floating point
// or integer values.
//
// The expectation of the Int64 and Float64 methods are that they return an
// error if the underlying value is not parsable as that type - for Float64,
// this is assumed to not happen, but the error case is used from Int64 to
// distinguish between the two types.
//
// This type allows for using the json.Number type, as well as any other
// similar types.
type GenericNumber interface {
	Int64() (int64, error)
	Float64() (float64, error)
	String() string
}

// TypeName converts a value to a type name.
//
// If the values type matches one of the JSON-like types, that name is
// returned. Otherwise, the name will be in the form `golang<%T>`.
func TypeName(v interface{}) string {
	if v == nil {
		return TypeNull
	}

	switch d := v.(type) {
	case []interface{}:
		return TypeArray
	case bool:
		return TypeBoolean
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return TypeInteger
	case float32:
		if math.Trunc(float64(d)) == float64(d) {
			return TypeInteger
		}
		return TypeNumber
	case float64:
		if math.Trunc(d) == d {
			return TypeInteger
		}
		return TypeNumber
	case GenericNumber:
		_, err := d.Int64()
		if err == nil {
			return TypeInteger
		}
		return TypeNumber
	case map[string]interface{}:
		return TypeObject
	case string:
		return TypeString
	}
	return fmt.Sprintf("golang<%T>", v)
}

// Is returns an error if the given value is not one of the given types.
func Is(v interface{}, types ...string) error {
	typename := TypeName(v)
	for _, t := range types {
		if typename == t {
			return nil
		}
		// If we accept a Number and the type given is an integer, accept that
		// as an Integer is a Number.
		if t == TypeNumber && typename == TypeInteger {
			return nil
		}
	}

	return InvalidTypeError{
		Expected: types,
		Actual:   typename,
	}
}

// AsArray attempts to coerce the value into an array.
func AsArray(v interface{}) ([]interface{}, error) {
	a, ok := v.([]interface{})
	if !ok {
		return nil, InvalidTypeError{
			Expected: []string{TypeArray},
			Actual:   TypeName(v),
		}
	}
	return a, nil
}

// AsBoolean attempts to coerce the value into a boolean.
func AsBoolean(v interface{}) (bool, error) {
	b, ok := v.(bool)
	if !ok {
		return false, InvalidTypeError{
			Expected: []string{TypeBoolean},
			Actual:   TypeName(v),
		}
	}
	return b, nil
}

func truncFloat(v float64) (int64, error) {
	if math.Trunc(v) == v {
		return int64(v), nil
	}
	return 0, InvalidTypeError{
		Expected: []string{TypeInteger},
		Actual:   TypeNumber,
	}
}

// AsInteger attempts to coerce the value into an integer.
func AsInteger(v interface{}) (int64, error) {
	switch d := v.(type) {
	case int:
		return int64(d), nil
	case int8:
		return int64(d), nil
	case int16:
		return int64(d), nil
	case int32:
		return int64(d), nil
	case int64:
		return d, nil
	case uint:
		return int64(d), nil
	case uint8:
		return int64(d), nil
	case uint16:
		return int64(d), nil
	case uint32:
		return int64(d), nil
	case uint64:
		return int64(d), nil
	case float32:
		return truncFloat(float64(d))
	case float64:
		return truncFloat(d)
	case GenericNumber:
		if i, err := d.Int64(); err == nil {
			return i, nil
		}
		f, _ := d.Float64()
		return truncFloat(f)
	}
	return 0, InvalidTypeError{
		Expected: []string{TypeInteger},
		Actual:   TypeName(v),
	}
}

// AsNumber attempts to coerce the value into a number.
func AsNumber(v interface{}) (float64, error) {
	switch d := v.(type) {
	case int:
		return float64(d), nil
	case int8:
		return float64(d), nil
	case int16:
		return float64(d), nil
	case int32:
		return float64(d), nil
	case int64:
		return float64(d), nil
	case uint:
		return float64(d), nil
	case uint8:
		return float64(d), nil
	case uint16:
		return float64(d), nil
	case uint32:
		return float64(d), nil
	case uint64:
		return float64(d), nil
	case float32:
		return float64(d), nil
	case float64:
		return d, nil
	case GenericNumber:
		f, err := d.Float64()
		if err != nil {
			return 0, InvalidTypeError{
				Expected: []string{TypeNumber},
				Actual:   TypeName(v),
			}
		}
		return f, nil
	}
	return 0, InvalidTypeError{
		Expected: []string{TypeNumber},
		Actual:   TypeName(v),
	}
}

// AsObject attempts to coerce the value to an object.
func AsObject(v interface{}) (map[string]interface{}, error) {
	m, ok := v.(map[string]interface{})
	if !ok {
		return nil, InvalidTypeError{
			Expected: []string{TypeObject},
			Actual:   TypeName(v),
		}
	}
	return m, nil
}

// AsString attempts to coerce the value to a string.
func AsString(v interface{}) (string, error) {
	s, ok := v.(string)
	if !ok {
		return "", InvalidTypeError{
			Expected: []string{TypeString},
			Actual:   TypeName(v),
		}
	}
	return s, nil
}
