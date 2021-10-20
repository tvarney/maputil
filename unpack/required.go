package unpack

import (
	"github.com/tvarney/maputil"
	"github.com/tvarney/maputil/errctx"
	"github.com/tvarney/maputil/mpath"
)

// RequireArray fetches a value from the map and converts it to an array,
// sending any errors to the given context.
func RequireArray(ctx *errctx.Context, m map[string]interface{}, key string) []interface{} {
	a, err := maputil.RequireArray(m, key)
	ctx.ErrorWithKey(err, key)
	return a
}

// RequireBoolean fetches a value from the map and converts it to a boolean,
// sending any errors to the given context.
func RequireBoolean(ctx *errctx.Context, m map[string]interface{}, key string) bool {
	b, err := maputil.RequireBoolean(m, key)
	ctx.ErrorWithKey(err, key)
	return b
}

// RequireInteger fetches a value from the map and converts it to an integer,
// sending any errors to the given context.
func RequireInteger(ctx *errctx.Context, m map[string]interface{}, key string) int64 {
	i, err := maputil.RequireInteger(m, key)
	ctx.ErrorWithKey(err, key)
	return i
}

// RequireNull fetches a value from the map and ensures it is nil, sending any
// errors to the given context.
func RequireNull(ctx *errctx.Context, m map[string]interface{}, key string) {
	ctx.ErrorWithKey(maputil.RequireNull(m, key), key)
}

// RequireNumber fetches a value from the map and converts it to a number,
// sending any errors to the given context.
func RequireNumber(ctx *errctx.Context, m map[string]interface{}, key string) float64 {
	n, err := maputil.RequireNumber(m, key)
	ctx.ErrorWithKey(err, key)
	return n
}

// RequireObject fetches a value from the map and converts it to an object,
// sending any errors to the given context.
func RequireObject(ctx *errctx.Context, m map[string]interface{}, key string) map[string]interface{} {
	o, err := maputil.RequireObject(m, key)
	ctx.ErrorWithKey(err, key)
	return o
}

// RequireString fetches a value from the map and converts it to a string,
// sending any errors to the given context.
func RequireString(ctx *errctx.Context, m map[string]interface{}, key string) string {
	s, err := maputil.RequireString(m, key)
	ctx.ErrorWithKey(err, key)
	return s
}

// RequireStringEnum fetches a value from the map and converts it to a string
// and ensures it is one of the allowed values, sending any errors to the given
// context.
func RequireStringEnum(ctx *errctx.Context, m map[string]interface{}, key string, allowed []string) string {
	s, err := maputil.RequireStringEnum(m, key, allowed)
	ctx.ErrorWithKey(err, key)
	return s
}

// RequireBooleanArray fetches an array from the map and attempts to convert
// all elements to booleans, sending any errors to the given context.
//
// This function will discard any elements which can not be converted to
// booleans, possibly resulting in an array with fewer items than the array in
// the map.
func RequireBooleanArray(ctx *errctx.Context, m map[string]interface{}, key string) []bool {
	a, err := maputil.RequireArray(m, key)
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

// RequireIntegerArray fetches an array from the map and attempts to convert
// all elements to integers, sending any errors to the given context.
//
// This function will discard any elements which can not be converted to
// integers, possibly resulting in an array with fewer items than the array in
// the map.
func RequireIntegerArray(ctx *errctx.Context, m map[string]interface{}, key string) []int64 {
	a, err := maputil.RequireArray(m, key)
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

// RequireNumberArray fetches an array from the map and attempts to convert
// all elements to numbers, sending any errors to the given context.
//
// This function will discard any elements which can not be converted to
// numbers, possibly resulting in an array with fewer items than the array in
// the map.
func RequireNumberArray(ctx *errctx.Context, m map[string]interface{}, key string) []float64 {
	a, err := maputil.RequireArray(m, key)
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

// RequireObjectArray fetches an array from the map and attempts to convert
// all elements to objects, sending any errors to the given context.
//
// This function will discard any elements which can not be converted to
// objects, possibly resulting in an array with fewer items than the array in
// the map.
func RequireObjectArray(ctx *errctx.Context, m map[string]interface{}, key string) []map[string]interface{} {
	a, err := maputil.RequireArray(m, key)
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

// RequireStringArray fetches an array from the map and attempts to convert
// all elements to strings, sending any errors to the given context.
//
// This function will discard any elements which can not be converted to
// strings, possibly resulting in an array with fewer items than the array in
// the map.
func RequireStringArray(ctx *errctx.Context, m map[string]interface{}, key string) []string {
	a, err := maputil.RequireArray(m, key)
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

// RequireStringEnumArray fetches an array from the map, attempts to
// convert all elmenents to strings and check that they match the allowed enum
// values, sending any errors to the given context.
//
// This function will discard any elements which can not be converted to
// strings, as well as any strings which do not match the allowed enum values.
// This may result in an array with fewer items than the array in the map.
func RequireStringEnumArray(ctx *errctx.Context, m map[string]interface{}, key string, allowed []string) []string {
	a, err := maputil.RequireArray(m, key)
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
