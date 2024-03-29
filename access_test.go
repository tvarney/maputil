package maputil_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/maputil"
)

const (
	keyMissing = "missing"
	keyGood    = "good"
	keyBad     = "bad"
	keyBadVal  = "bad-value"

	testInteger = int64(10)
	testNumber  = float64(14.9)
	testString  = "value"
)

var (
	testArray  = []interface{}{1, 2, 3}
	testObject = map[string]interface{}{"test": "value"}
)

func TestGetArray(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testArray,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		a, ok, err := maputil.GetArray(d, keyMissing)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.Nil(t, a)       // No value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		a, ok, err := maputil.GetArray(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeArray},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.Nil(t, a)      // We returned no value
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		a, ok, err := maputil.GetArray(d, keyGood)
		require.NoError(t, err)        // No error due to type
		require.True(t, ok)            // The key was found
		require.Equal(t, testArray, a) // The returned value is what we expect
		require.Equal(t, m, d)         // The source map wasn't changed
	})
}

func TestGetBoolean(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: true,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		b, ok, err := maputil.GetBoolean(d, keyMissing)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.False(t, b)     // No value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		b, ok, err := maputil.GetBoolean(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeBoolean},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.False(t, b)    // We returned no value
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		b, ok, err := maputil.GetBoolean(d, keyGood)
		require.NoError(t, err) // No error due to type
		require.True(t, ok)     // The key was found
		require.True(t, b)      // The returned value is what we expect
		require.Equal(t, m, d)  // The source map wasn't changed
	})
}

func TestGetInteger(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testInteger,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		i, ok, err := maputil.GetInteger(d, keyMissing)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.Zero(t, i)      // No value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		i, ok, err := maputil.GetInteger(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeInteger},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.Zero(t, i)     // We returned no value
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		i, ok, err := maputil.GetInteger(d, keyGood)
		require.NoError(t, err)          // No error due to type
		require.True(t, ok)              // The key was found
		require.Equal(t, testInteger, i) // The returned value is what we expect
		require.Equal(t, m, d)           // The source map wasn't changed
	})
}

func TestGetNull(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: nil,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ok, err := maputil.GetNull(d, keyMissing)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ok, err := maputil.GetNull(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Actual:   maputil.TypeString,
			Expected: []string{maputil.TypeNull},
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ok, err := maputil.GetNull(d, keyGood)
		require.NoError(t, err) // No error due to type
		require.True(t, ok)     // The key was found
		require.Equal(t, m, d)  // The source map wasn't changed
	})
}

func TestGetNumber(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testNumber,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		n, ok, err := maputil.GetNumber(d, keyMissing)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.Zero(t, n)      // No value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		n, ok, err := maputil.GetNumber(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeNumber},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.Zero(t, n)     // We returned no value
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		n, ok, err := maputil.GetNumber(d, keyGood)
		require.NoError(t, err)         // No error due to type
		require.True(t, ok)             // The key was found
		require.Equal(t, testNumber, n) // The returned value is what we expect
		require.Equal(t, m, d)          // The source map wasn't changed
	})
}

func TestGetObject(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testObject,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		o, ok, err := maputil.GetObject(d, keyMissing)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.Nil(t, o)       // No value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		o, ok, err := maputil.GetObject(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeObject},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.Nil(t, o)      // We returned no value
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		o, ok, err := maputil.GetObject(d, keyGood)
		require.NoError(t, err)         // No error due to type
		require.True(t, ok)             // The key was found
		require.Equal(t, testObject, o) // The returned value is what we expect
		require.Equal(t, m, d)          // The source map wasn't changed
	})
}

func TestGetString(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testString,
		keyBad:  testInteger,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, ok, err := maputil.GetString(d, keyMissing)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.Zero(t, s)      // No value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, ok, err := maputil.GetString(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeString},
			Actual:   maputil.TypeInteger,
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.Zero(t, s)     // We returned no value
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, ok, err := maputil.GetString(d, keyGood)
		require.NoError(t, err)         // No error due to type
		require.True(t, ok)             // The key was found
		require.Equal(t, testString, s) // The returned value is what we expect
		require.Equal(t, m, d)          // The source map wasn't changed
	})
}

func TestGetStringEnum(t *testing.T) {
	t.Parallel()
	const badVal = "four"
	const goodVal = "two"
	enum := []string{"one", "two", "three"}
	m := map[string]interface{}{
		keyGood:   goodVal,
		keyBad:    testInteger,
		keyBadVal: badVal,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, ok, err := maputil.GetStringEnum(d, keyMissing, enum)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.Zero(t, s)      // No value returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, ok, err := maputil.GetStringEnum(d, keyBad, enum)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Actual:   maputil.TypeInteger,
			Expected: []string{maputil.TypeString},
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.Zero(t, s)     // No value returned
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("BadValue", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, ok, err := maputil.GetStringEnum(d, keyBadVal, enum)
		require.EqualError(t, err, maputil.EnumStringError{
			Value: badVal,
			Enum:  enum,
		}.Error()) // Error due to value
		require.True(t, ok)         // The key was found
		require.Equal(t, badVal, s) // Value is returned
		require.Equal(t, m, d)      // The source map wasn't changed
	})
	t.Run("GoodValue", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, ok, err := maputil.GetStringEnum(d, keyGood, enum)
		require.NoError(t, err)      // No error due to type
		require.True(t, ok)          // The key wasn't found
		require.Equal(t, goodVal, s) // No value returned
		require.Equal(t, m, d)       // The source map wasn't changed
	})
}

func TestOptionalArray(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testArray,
		keyBad:  testString,
	}
	dv := []interface{}{1, "two", true}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		a, err := maputil.OptionalArray(d, keyMissing, dv)
		require.NoError(t, err) // No error due to type
		require.Equal(t, dv, a) // We returned the default value
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		a, err := maputil.OptionalArray(d, keyBad, dv)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeArray},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.Equal(t, dv, a) // We returned the default value
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		a, err := maputil.OptionalArray(d, keyGood, dv)
		require.NoError(t, err)        // No error due to type
		require.Equal(t, testArray, a) // The returned value is what we expect
		require.Equal(t, m, d)         // The source map wasn't changed
	})
}

func TestOptionalBoolean(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: true,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		b, err := maputil.OptionalBoolean(d, keyMissing, true)
		require.NoError(t, err) // No error due to type
		require.True(t, b)      // The default value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		b, err := maputil.OptionalBoolean(d, keyBad, true)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeBoolean},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.True(t, b)     // The default value was returned
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		b, err := maputil.OptionalBoolean(d, keyGood, false)
		require.NoError(t, err) // No error due to type
		require.True(t, b)      // The returned value is what we expect
		require.Equal(t, m, d)  // The source map wasn't changed
	})
}

func TestOptionalInteger(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testInteger,
		keyBad:  testString,
	}
	const dv int64 = 65536
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		i, err := maputil.OptionalInteger(d, keyMissing, dv)
		require.NoError(t, err) // No error due to type
		require.Equal(t, dv, i) // The default value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		i, err := maputil.OptionalInteger(d, keyBad, dv)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeInteger},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.Equal(t, dv, i) // The default value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		i, err := maputil.OptionalInteger(d, keyGood, dv)
		require.NoError(t, err)          // No error due to type
		require.Equal(t, testInteger, i) // The returned value is what we expect
		require.Equal(t, m, d)           // The source map wasn't changed
	})
}

func TestOptionalNull(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: nil,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		err := maputil.OptionalNull(d, keyMissing)
		require.NoError(t, err) // No error due to type
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		err := maputil.OptionalNull(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Actual:   maputil.TypeString,
			Expected: []string{maputil.TypeNull},
		}.Error()) // Error due to type
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		err := maputil.OptionalNull(d, keyGood)
		require.NoError(t, err) // No error due to type
		require.Equal(t, m, d)  // The source map wasn't changed
	})
}

func TestOptionalNumber(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testNumber,
		keyBad:  testString,
	}
	const dv float64 = 101.3
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		n, err := maputil.OptionalNumber(d, keyMissing, dv)
		require.NoError(t, err) // No error due to type
		require.Equal(t, dv, n) // The default value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		n, err := maputil.OptionalNumber(d, keyBad, dv)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeNumber},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.Equal(t, dv, n) // The default value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		n, err := maputil.OptionalNumber(d, keyGood, dv)
		require.NoError(t, err)         // No error due to type
		require.Equal(t, testNumber, n) // The returned value is what we expect
		require.Equal(t, m, d)          // The source map wasn't changed
	})
}

func TestOptionalObject(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testObject,
		keyBad:  testString,
	}
	dv := map[string]interface{}{
		"a": "test",
		"b": "value",
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		o, err := maputil.OptionalObject(d, keyMissing, dv)
		require.NoError(t, err) // No error due to type
		require.Equal(t, dv, o) // The default value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		o, err := maputil.OptionalObject(d, keyBad, dv)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeObject},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.Equal(t, dv, o) // The default value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		o, err := maputil.OptionalObject(d, keyGood, dv)
		require.NoError(t, err)         // No error due to type
		require.Equal(t, testObject, o) // The returned value is what we expect
		require.Equal(t, m, d)          // The source map wasn't changed
	})
}

func TestOptionalString(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testString,
		keyBad:  testInteger,
	}
	const dv string = "default-value"
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, err := maputil.OptionalString(d, keyMissing, dv)
		require.NoError(t, err) // No error due to type
		require.Equal(t, dv, s) // The default value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, err := maputil.OptionalString(d, keyBad, dv)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeString},
			Actual:   maputil.TypeInteger,
		}.Error()) // Error due to type
		require.Equal(t, dv, s) // We returned no value
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, err := maputil.OptionalString(d, keyGood, dv)
		require.NoError(t, err)         // No error due to type
		require.Equal(t, testString, s) // The returned value is what we expect
		require.Equal(t, m, d)          // The source map wasn't changed
	})
}

func TestOptionalStringEnum(t *testing.T) {
	t.Parallel()
	const badVal = "four"
	const goodVal = "two"
	enum := []string{"one", "two", "three"}
	m := map[string]interface{}{
		keyGood:   goodVal,
		keyBad:    testInteger,
		keyBadVal: badVal,
	}
	const dv = "three"
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, err := maputil.OptionalStringEnum(d, keyMissing, enum, dv)
		require.NoError(t, err) // No error due to type
		require.Equal(t, dv, s) // The default value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, err := maputil.OptionalStringEnum(d, keyBad, enum, dv)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Actual:   maputil.TypeInteger,
			Expected: []string{maputil.TypeString},
		}.Error()) // Error due to type
		require.Equal(t, dv, s) // The default value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("BadValue", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, err := maputil.OptionalStringEnum(d, keyBadVal, enum, dv)
		require.EqualError(t, err, maputil.EnumStringError{
			Value: badVal,
			Enum:  enum,
		}.Error()) // Error due to value
		require.Equal(t, dv, s) // The default value is returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("GoodValue", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, err := maputil.OptionalStringEnum(d, keyGood, enum, dv)
		require.NoError(t, err)      // No error due to type
		require.Equal(t, goodVal, s) // No value returned
		require.Equal(t, m, d)       // The source map wasn't changed
	})
}
func TestPopArray(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testArray,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		a, ok, err := maputil.PopArray(d, keyMissing)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.Nil(t, a)       // No value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyBad)
		a, ok, err := maputil.PopArray(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeArray},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.Nil(t, a)      // We returned no value
		require.Equal(t, e, d) // The source map had the key removed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyGood)
		a, ok, err := maputil.PopArray(d, keyGood)
		require.NoError(t, err)        // No error due to type
		require.True(t, ok)            // The key was found
		require.Equal(t, testArray, a) // The returned value is what we expect
		require.Equal(t, e, d)         // The source had the key removed
	})
}

func TestPopBoolean(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: true,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		b, ok, err := maputil.PopBoolean(d, keyMissing)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.False(t, b)     // No value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyBad)
		b, ok, err := maputil.PopBoolean(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeBoolean},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.False(t, b)    // We returned no value
		require.Equal(t, e, d) // The source map had the key removed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyGood)
		b, ok, err := maputil.PopBoolean(d, keyGood)
		require.NoError(t, err) // No error due to type
		require.True(t, ok)     // The key was found
		require.True(t, b)      // The returned value is what we expect
		require.Equal(t, e, d)  // The source map had the key removed
	})
}

func TestPopInteger(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testInteger,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		i, ok, err := maputil.PopInteger(d, keyMissing)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.Zero(t, i)      // No value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyBad)
		i, ok, err := maputil.PopInteger(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeInteger},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.Zero(t, i)     // We returned no value
		require.Equal(t, e, d) // The source map had the key removed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyGood)
		i, ok, err := maputil.PopInteger(d, keyGood)
		require.NoError(t, err)          // No error due to type
		require.True(t, ok)              // The key was found
		require.Equal(t, testInteger, i) // The returned value is what we expect
		require.Equal(t, e, d)           // The source map had the key removed
	})
}

func TestPopNull(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: nil,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ok, err := maputil.PopNull(d, keyMissing)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyBad)
		ok, err := maputil.PopNull(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Actual:   maputil.TypeString,
			Expected: []string{maputil.TypeNull},
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.Equal(t, e, d) // The source map had the key removed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyGood)
		ok, err := maputil.PopNull(d, keyGood)
		require.NoError(t, err) // No error due to type
		require.True(t, ok)     // The key was found
		require.Equal(t, e, d)  // The source map had the key removed
	})
}

func TestPopNumber(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testNumber,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		n, ok, err := maputil.PopNumber(d, keyMissing)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.Zero(t, n)      // No value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyBad)
		n, ok, err := maputil.PopNumber(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeNumber},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.Zero(t, n)     // We returned no value
		require.Equal(t, e, d) // The source map had the key removed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyGood)
		n, ok, err := maputil.PopNumber(d, keyGood)
		require.NoError(t, err)         // No error due to type
		require.True(t, ok)             // The key was found
		require.Equal(t, testNumber, n) // The returned value is what we expect
		require.Equal(t, e, d)          // The source map had the key removed
	})
}

func TestPopObject(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testObject,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		o, ok, err := maputil.PopObject(d, keyMissing)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.Nil(t, o)       // No value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyBad)
		o, ok, err := maputil.PopObject(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeObject},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.Nil(t, o)      // We returned no value
		require.Equal(t, e, d) // The source map had the key removed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyGood)
		o, ok, err := maputil.PopObject(d, keyGood)
		require.NoError(t, err)         // No error due to type
		require.True(t, ok)             // The key was found
		require.Equal(t, testObject, o) // The returned value is what we expect
		require.Equal(t, e, d)          // The source map had the key removed
	})
}

func TestPopString(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testString,
		keyBad:  testInteger,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, ok, err := maputil.PopString(d, keyMissing)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.Zero(t, s)      // No value was returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyBad)
		s, ok, err := maputil.PopString(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeString},
			Actual:   maputil.TypeInteger,
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.Zero(t, s)     // We returned no value
		require.Equal(t, e, d) // The source map had the key removed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyGood)
		s, ok, err := maputil.PopString(d, keyGood)
		require.NoError(t, err)         // No error due to type
		require.True(t, ok)             // The key was found
		require.Equal(t, testString, s) // The returned value is what we expect
		require.Equal(t, e, d)          // The source map had the key removed
	})
}

func TestPopStringEnum(t *testing.T) {
	t.Parallel()
	const badVal = "four"
	const goodVal = "two"
	enum := []string{"one", "two", "three"}
	m := map[string]interface{}{
		keyGood:   goodVal,
		keyBad:    testInteger,
		keyBadVal: badVal,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, ok, err := maputil.PopStringEnum(d, keyMissing, enum)
		require.NoError(t, err) // No error due to type
		require.False(t, ok)    // The key wasn't found
		require.Zero(t, s)      // No value returned
		require.Equal(t, m, d)  // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyBad)
		s, ok, err := maputil.PopStringEnum(d, keyBad, enum)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Actual:   maputil.TypeInteger,
			Expected: []string{maputil.TypeString},
		}.Error()) // Error due to type
		require.True(t, ok)    // The key was found
		require.Zero(t, s)     // No value returned
		require.Equal(t, e, d) // The source map wasn't changed
	})
	t.Run("BadValue", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyBadVal)
		s, ok, err := maputil.PopStringEnum(d, keyBadVal, enum)
		require.EqualError(t, err, maputil.EnumStringError{
			Value: badVal,
			Enum:  enum,
		}.Error()) // Error due to value
		require.True(t, ok)         // The key was found
		require.Equal(t, badVal, s) // No value returned
		require.Equal(t, e, d)      // The source map wasn't changed
	})
	t.Run("GoodValue", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		e := maputil.Copy(d)
		delete(e, keyGood)
		s, ok, err := maputil.PopStringEnum(d, keyGood, enum)
		require.NoError(t, err)      // No error due to type
		require.True(t, ok)          // The key was found
		require.Equal(t, goodVal, s) // No value returned
		require.Equal(t, e, d)       // The source map wasn't changed
	})
}

func TestRequireArray(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testArray,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		a, err := maputil.RequireArray(d, keyMissing)
		require.EqualError(t, err, maputil.MissingRequiredValueError{
			Key: keyMissing,
		}.Error()) // Error due to key being required
		require.Nil(t, a)      // No value was returned
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		a, err := maputil.RequireArray(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeArray},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.Nil(t, a)      // We returned no value
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		a, err := maputil.RequireArray(d, keyGood)
		require.NoError(t, err)        // No error due to type
		require.Equal(t, testArray, a) // The returned value is what we expect
		require.Equal(t, m, d)         // The source map wasn't changed
	})
}

func TestRequireBoolean(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: true,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		b, err := maputil.RequireBoolean(d, keyMissing)
		require.EqualError(t, err, maputil.MissingRequiredValueError{
			Key: keyMissing,
		}.Error()) // Error due to key being required
		require.False(t, b)    // No value was returned
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		b, err := maputil.RequireBoolean(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeBoolean},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.False(t, b)    // We returned no value
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		b, err := maputil.RequireBoolean(d, keyGood)
		require.NoError(t, err) // No error due to type
		require.True(t, b)      // The returned value is what we expect
		require.Equal(t, m, d)  // The source map wasn't changed
	})
}

func TestRequireInteger(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testInteger,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		i, err := maputil.RequireInteger(d, keyMissing)
		require.EqualError(t, err, maputil.MissingRequiredValueError{
			Key: keyMissing,
		}.Error()) // Error due to key being required
		require.Zero(t, i)     // No value was returned
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		i, err := maputil.RequireInteger(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeInteger},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.Zero(t, i)     // We returned no value
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		i, err := maputil.RequireInteger(d, keyGood)
		require.NoError(t, err)          // No error due to type
		require.Equal(t, testInteger, i) // The returned value is what we expect
		require.Equal(t, m, d)           // The source map wasn't changed
	})
}

func TestRequireNull(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: nil,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		err := maputil.RequireNull(d, keyMissing)
		require.EqualError(t, err, maputil.MissingRequiredValueError{
			Key: keyMissing,
		}.Error()) // Error due to key being required
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		err := maputil.RequireNull(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Actual:   maputil.TypeString,
			Expected: []string{maputil.TypeNull},
		}.Error()) // Error due to type
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		err := maputil.RequireNull(d, keyGood)
		require.NoError(t, err) // No error due to type
		require.Equal(t, m, d)  // The source map wasn't changed
	})
}

func TestRequireNumber(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testNumber,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		n, err := maputil.RequireNumber(d, keyMissing)
		require.EqualError(t, err, maputil.MissingRequiredValueError{
			Key: keyMissing,
		}.Error()) // Error due to key being required
		require.Zero(t, n)     // No value was returned
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		n, err := maputil.RequireNumber(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeNumber},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.Zero(t, n)     // We returned no value
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		n, err := maputil.RequireNumber(d, keyGood)
		require.NoError(t, err)         // No error due to type
		require.Equal(t, testNumber, n) // The returned value is what we expect
		require.Equal(t, m, d)          // The source map wasn't changed
	})
}

func TestRequireObject(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testObject,
		keyBad:  testString,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		o, err := maputil.RequireObject(d, keyMissing)
		require.EqualError(t, err, maputil.MissingRequiredValueError{
			Key: keyMissing,
		}.Error()) // No error due to type
		require.Nil(t, o)      // No value was returned
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		o, err := maputil.RequireObject(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeObject},
			Actual:   maputil.TypeString,
		}.Error()) // Error due to type
		require.Nil(t, o)      // We returned no value
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		o, err := maputil.RequireObject(d, keyGood)
		require.NoError(t, err)         // No error due to type
		require.Equal(t, testObject, o) // The returned value is what we expect
		require.Equal(t, m, d)          // The source map wasn't changed
	})
}

func TestRequireString(t *testing.T) {
	t.Parallel()
	m := map[string]interface{}{
		keyGood: testString,
		keyBad:  testInteger,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, err := maputil.RequireString(d, keyMissing)
		require.EqualError(t, err, maputil.MissingRequiredValueError{
			Key: keyMissing,
		}.Error()) // No error due to type
		require.Zero(t, s)     // No value was returned
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, err := maputil.RequireString(d, keyBad)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Expected: []string{maputil.TypeString},
			Actual:   maputil.TypeInteger,
		}.Error()) // Error due to type
		require.Zero(t, s)     // We returned no value
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("Present", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, err := maputil.RequireString(d, keyGood)
		require.NoError(t, err)         // No error due to type
		require.Equal(t, testString, s) // The returned value is what we expect
		require.Equal(t, m, d)          // The source map wasn't changed
	})
}

func TestRequireStringEnum(t *testing.T) {
	t.Parallel()
	const badVal = "four"
	const goodVal = "two"
	enum := []string{"one", "two", "three"}
	m := map[string]interface{}{
		keyGood:   goodVal,
		keyBad:    testInteger,
		keyBadVal: badVal,
	}
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, err := maputil.RequireStringEnum(d, keyMissing, enum)
		require.EqualError(t, err, maputil.MissingRequiredValueError{
			Key: keyMissing,
		}.Error()) // No error due to type
		require.Zero(t, s)     // No value returned
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("InvalidType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, err := maputil.RequireStringEnum(d, keyBad, enum)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Actual:   maputil.TypeInteger,
			Expected: []string{maputil.TypeString},
		}.Error()) // Error due to type
		require.Zero(t, s)     // No value returned
		require.Equal(t, m, d) // The source map wasn't changed
	})
	t.Run("BadValue", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, err := maputil.RequireStringEnum(d, keyBadVal, enum)
		require.EqualError(t, err, maputil.EnumStringError{
			Value: badVal,
			Enum:  enum,
		}.Error()) // Error due to value
		require.Equal(t, badVal, s) // Value is returned
		require.Equal(t, m, d)      // The source map wasn't changed
	})
	t.Run("GoodValue", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		s, err := maputil.RequireStringEnum(d, keyGood, enum)
		require.NoError(t, err)      // No error due to type
		require.Equal(t, goodVal, s) // No value returned
		require.Equal(t, m, d)       // The source map wasn't changed
	})
}
