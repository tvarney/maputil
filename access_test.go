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
