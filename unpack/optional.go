package unpack

import (
	"github.com/tvarney/maputil"
	"github.com/tvarney/maputil/errctx"
	"github.com/tvarney/maputil/mpath"
)

// OptionalArray fetches a value from the map and converts it to an array,
// sending any errors to the given context.
func OptionalArray(ctx *errctx.Context, m map[string]interface{}, key string, dv []interface{}) []interface{} {
	a, err := maputil.OptionalArray(m, key, dv)
	ctx.ErrorWithKey(err, key)
	return a
}

// OptionalBoolean fetches a value from the map and converts it to a boolean,
// sending any errors to the given context.
func OptionalBoolean(ctx *errctx.Context, m map[string]interface{}, key string, dv bool) bool {
	b, err := maputil.OptionalBoolean(m, key, dv)
	ctx.ErrorWithKey(err, key)
	return b
}

// OptionalInteger fetches a value from the map and converts it to an integer,
// sending any errors to the given context.
func OptionalInteger(ctx *errctx.Context, m map[string]interface{}, key string, dv int64) int64 {
	i, err := maputil.OptionalInteger(m, key, dv)
	ctx.ErrorWithKey(err, key)
	return i
}

// OptionalNull fetches a value from the map and ensures it is nil, sending any
// errors to the given context.
func OptionalNull(ctx *errctx.Context, m map[string]interface{}, key string) {
	ctx.ErrorWithKey(maputil.OptionalNull(m, key), key)
}

// OptionalNumber fetches a value from the map and converts it to a number,
// sending any errors to the given context.
func OptionalNumber(ctx *errctx.Context, m map[string]interface{}, key string, dv float64) float64 {
	n, err := maputil.OptionalNumber(m, key, dv)
	ctx.ErrorWithKey(err, key)
	return n
}

// OptionalObject fetches a value from the map and converts it to an object,
// sending any errors to the given context.
func OptionalObject(
	ctx *errctx.Context,
	m map[string]interface{},
	key string,
	dv map[string]interface{},
) map[string]interface{} {
	o, err := maputil.OptionalObject(m, key, dv)
	ctx.ErrorWithKey(err, key)
	return o
}

// OptionalString fetches a value from the map and converts it to a string,
// sending any errors to the given context.
func OptionalString(ctx *errctx.Context, m map[string]interface{}, key, dv string) string {
	s, err := maputil.OptionalString(m, key, dv)
	ctx.ErrorWithKey(err, key)
	return s
}

// OptionalStringEnum fetches a value from the map and converts it to a string
// and ensures it is one of the allowed values, sending any errors to the given
// context.
func OptionalStringEnum(ctx *errctx.Context, m map[string]interface{}, key string, allowed []string, dv string) string {
	s, err := maputil.OptionalStringEnum(m, key, allowed, dv)
	ctx.ErrorWithKey(err, key)
	return s
}

// OptionalBooleanArray fetches an array from the map and attempts to convert
// all elements to booleans, sending any errors to the given context.
//
// This function will discard any elements which can not be converted to
// booleans, possibly resulting in an array with fewer items than the array in
// the map.
func OptionalBooleanArray(ctx *errctx.Context, m map[string]interface{}, key string) []bool {
	a, err := maputil.OptionalArray(m, key, nil)
	if err != nil {
		ctx.ErrorWithKey(err, key)
		return nil
	}

	if len(a) == 0 {
		return nil
	}

	ctx.Path.Add(mpath.Key(key))
	ba := make([]bool, 0, len(a))
	for i, iv := range a {
		bv, err := maputil.AsBoolean(iv)
		if err != nil {
			ctx.ErrorWithIndex(err, i)
			continue
		}

		ba = append(ba, bv)
	}
	ctx.Path.Pop()
	return ba
}

// OptionalIntegerArray fetches an array from the map and attempts to convert
// all elements to integers, sending any errors to the given context.
//
// This function will discard any elements which can not be converted to
// integers, possibly resulting in an array with fewer items than the array in
// the map.
func OptionalIntegerArray(ctx *errctx.Context, m map[string]interface{}, key string) []int64 {
	a, err := maputil.OptionalArray(m, key, nil)
	if err != nil {
		ctx.ErrorWithKey(err, key)
		return nil
	}
	if len(a) == 0 {
		return nil
	}

	ctx.Path.Add(mpath.Key(key))
	ia := make([]int64, 0, len(a))
	for i, iv := range a {
		v, err := maputil.AsInteger(iv)
		if err != nil {
			ctx.ErrorWithIndex(err, i)
			continue
		}

		ia = append(ia, v)
	}
	ctx.Path.Pop()
	return ia
}

// OptionalNumberArray fetches an array from the map and attempts to convert
// all elements to numbers, sending any errors to the given context.
//
// This function will discard any elements which can not be converted to
// numbers, possibly resulting in an array with fewer items than the array in
// the map.
func OptionalNumberArray(ctx *errctx.Context, m map[string]interface{}, key string) []float64 {
	a, err := maputil.OptionalArray(m, key, nil)
	if err != nil {
		ctx.ErrorWithKey(err, key)
		return nil
	}
	if len(a) == 0 {
		return nil
	}

	ctx.Path.Add(mpath.Key(key))
	na := make([]float64, 0, len(a))
	for i, iv := range a {
		v, err := maputil.AsNumber(iv)
		if err != nil {
			ctx.ErrorWithIndex(err, i)
			continue
		}

		na = append(na, v)
	}
	ctx.Path.Pop()
	return na
}

// OptionalObjectArray fetches an array from the map and attempts to convert
// all elements to objects, sending any errors to the given context.
//
// This function will discard any elements which can not be converted to
// objects, possibly resulting in an array with fewer items than the array in
// the map.
func OptionalObjectArray(ctx *errctx.Context, m map[string]interface{}, key string) []map[string]interface{} {
	a, err := maputil.OptionalArray(m, key, nil)
	if err != nil {
		ctx.ErrorWithKey(err, key)
		return nil
	}
	if len(a) == 0 {
		return nil
	}

	ctx.Path.Add(mpath.Key(key))
	oa := make([]map[string]interface{}, 0, len(a))
	for i, iv := range a {
		v, err := maputil.AsObject(iv)
		if err != nil {
			ctx.ErrorWithIndex(err, i)
			continue
		}

		oa = append(oa, v)
	}
	ctx.Path.Pop()
	return oa
}

// OptionalStringArray fetches an array from the map and attempts to convert
// all elements to strings, sending any errors to the given context.
//
// This function will discard any elements which can not be converted to
// strings, possibly resulting in an array with fewer items than the array in
// the map.
func OptionalStringArray(ctx *errctx.Context, m map[string]interface{}, key string) []string {
	a, err := maputil.OptionalArray(m, key, nil)
	if err != nil {
		ctx.ErrorWithKey(err, key)
		return nil
	}
	if len(a) == 0 {
		return nil
	}

	ctx.Path.Add(mpath.Key(key))
	sa := make([]string, 0, len(a))
	for i, iv := range a {
		v, err := maputil.AsString(iv)
		if err != nil {
			ctx.ErrorWithIndex(err, i)
			continue
		}

		sa = append(sa, v)
	}
	ctx.Path.Pop()
	return sa
}

// OptionalStringEnumArray fetches an array from the map, attempts to
// convert all elmenents to strings and check that they match the allowed enum
// values, sending any errors to the given context.
//
// This function will discard any elements which can not be converted to
// strings, as well as any strings which do not match the allowed enum values.
// This may result in an array with fewer items than the array in the map.
func OptionalStringEnumArray(ctx *errctx.Context, m map[string]interface{}, key string, allowed []string) []string {
	a, err := maputil.OptionalArray(m, key, nil)
	if err != nil {
		ctx.ErrorWithKey(err, key)
		return nil
	}
	if len(a) == 0 {
		return nil
	}

	ctx.Path.Add(mpath.Key(key))
	sa := make([]string, 0, len(a))
	for i, iv := range a {
		v, err := maputil.AsString(iv)
		if err != nil {
			ctx.ErrorWithIndex(err, i)
			continue
		}
		if err := maputil.CheckEnum(v, allowed); err != nil {
			ctx.ErrorWithIndex(err, i)
			continue
		}

		sa = append(sa, v)
	}
	ctx.Path.Pop()
	return sa
}
