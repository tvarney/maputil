package unpack_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/maputil"
	"github.com/tvarney/maputil/errctx"
	"github.com/tvarney/maputil/unpack"
)

const (
	testKeyMissing    = "missing"
	testKeyGood       = "good"
	testKeyBad        = "bad"
	testKeyBadElement = "bad-element"
	testKeyEmpty      = "empty"

	testDefaultInt    = int64(314)
	testDefaultNumber = float64(1.15)
	testDefaultString = "default-value"

	testInt    = int64(-389)
	testNumber = float64(42.42)
	testString = "test-value"
)

var (
	testDefaultArray  = []interface{}{1, 2, 3}
	testDefaultObject = map[string]interface{}{"a": "A", "b": "B"}
	testArray         = []interface{}{"one", 2, "drei"}
	testObject        = map[string]interface{}{"A": 1, "B": 2}
)

func TestOptionalArray(t *testing.T) {
	m := map[string]interface{}{
		testKeyGood: testArray,
		testKeyBad:  testNumber,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testArray, unpack.OptionalArray(ctx, d, testKeyGood, testDefaultArray))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testDefaultArray, unpack.OptionalArray(ctx, d, testKeyMissing, testDefaultArray))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testDefaultArray, unpack.OptionalArray(ctx, d, testKeyBad, testDefaultArray))
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestOptionalBoolean(t *testing.T) {
	m := map[string]interface{}{
		testKeyGood: false,
		testKeyBad:  testNumber,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.False(t, unpack.OptionalBoolean(ctx, d, testKeyGood, true))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.True(t, unpack.OptionalBoolean(ctx, d, testKeyMissing, true))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.True(t, unpack.OptionalBoolean(ctx, d, testKeyBad, true))
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestOptionalInteger(t *testing.T) {
	m := map[string]interface{}{
		testKeyGood: testInt,
		testKeyBad:  testString,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testInt, unpack.OptionalInteger(ctx, d, testKeyGood, testDefaultInt))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testDefaultInt, unpack.OptionalInteger(ctx, d, testKeyMissing, testDefaultInt))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testDefaultInt, unpack.OptionalInteger(ctx, d, testKeyBad, testDefaultInt))
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestOptionalNull(t *testing.T) {
	m := map[string]interface{}{
		testKeyGood: nil,
		testKeyBad:  testNumber,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		unpack.OptionalNull(ctx, d, testKeyGood)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		unpack.OptionalNull(ctx, d, testKeyMissing)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		unpack.OptionalNull(ctx, d, testKeyBad)
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestOptionalNumber(t *testing.T) {
	m := map[string]interface{}{
		testKeyGood: testNumber,
		testKeyBad:  true,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testNumber, unpack.OptionalNumber(ctx, d, testKeyGood, testDefaultNumber))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testDefaultNumber, unpack.OptionalNumber(ctx, d, testKeyMissing, testDefaultNumber))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testDefaultNumber, unpack.OptionalNumber(ctx, d, testKeyBad, testDefaultNumber))
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestOptionalObject(t *testing.T) {
	m := map[string]interface{}{
		testKeyGood: testObject,
		testKeyBad:  testNumber,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testObject, unpack.OptionalObject(ctx, d, testKeyGood, testDefaultObject))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testDefaultObject, unpack.OptionalObject(ctx, d, testKeyMissing, testDefaultObject))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testDefaultObject, unpack.OptionalObject(ctx, d, testKeyBad, testDefaultObject))
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestOptionalString(t *testing.T) {
	m := map[string]interface{}{
		testKeyGood: testString,
		testKeyBad:  testNumber,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testString, unpack.OptionalString(ctx, d, testKeyGood, testDefaultString))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testDefaultString, unpack.OptionalString(ctx, d, testKeyMissing, testDefaultString))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testDefaultString, unpack.OptionalString(ctx, d, testKeyBad, testDefaultString))
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestOptionalStringEnum(t *testing.T) {
	const badValue = "bad-value"
	allowed := []string{testString, testDefaultString}
	m := map[string]interface{}{
		testKeyGood:       testString,
		testKeyBad:        testNumber,
		testKeyBadElement: badValue,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testString, unpack.OptionalStringEnum(ctx, d, testKeyGood, allowed, testDefaultString))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testDefaultString, unpack.OptionalStringEnum(ctx, d, testKeyMissing, allowed, testDefaultString))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadValue", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testDefaultString, unpack.OptionalStringEnum(ctx, d, testKeyBadElement, allowed, testDefaultString))
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testDefaultString, unpack.OptionalStringEnum(ctx, d, testKeyBad, allowed, testDefaultString))
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestOptionalBooleanArray(t *testing.T) {
	barray := []bool{true, false, true}
	m := map[string]interface{}{
		testKeyGood:       []interface{}{true, false, true},
		testKeyEmpty:      []interface{}{},
		testKeyBadElement: []interface{}{true, false, "hello", true, 1.4},
		testKeyBad:        testNumber,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.OptionalBooleanArray(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalBooleanArray(ctx, d, testKeyMissing), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalBooleanArray(ctx, d, testKeyBad), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalBooleanArray(ctx, d, testKeyEmpty), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadElement", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.OptionalBooleanArray(ctx, d, testKeyBadElement))
		require.Equal(t, 2, ctx.ErrorCount())
	})
}

func TestOptionalIntegerArray(t *testing.T) {
	barray := []int64{1, 2, 3}
	m := map[string]interface{}{
		testKeyGood:       []interface{}{1, 2, 3},
		testKeyEmpty:      []interface{}{},
		testKeyBadElement: []interface{}{1, 2, "hello", 3, false},
		testKeyBad:        testString,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.OptionalIntegerArray(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalIntegerArray(ctx, d, testKeyMissing), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalIntegerArray(ctx, d, testKeyBad), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalIntegerArray(ctx, d, testKeyEmpty), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadElement", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.OptionalIntegerArray(ctx, d, testKeyBadElement))
		require.Equal(t, 2, ctx.ErrorCount())
	})
}

func TestOptionalNumberArray(t *testing.T) {
	barray := []float64{1.4, 2.8, -100.0}
	m := map[string]interface{}{
		testKeyGood:       []interface{}{1.4, 2.8, -100.0},
		testKeyEmpty:      []interface{}{},
		testKeyBadElement: []interface{}{1.4, 2.8, "hello", -100.0, true},
		testKeyBad:        testString,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.OptionalNumberArray(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalNumberArray(ctx, d, testKeyMissing), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalNumberArray(ctx, d, testKeyBad), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalNumberArray(ctx, d, testKeyEmpty), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadElement", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.OptionalNumberArray(ctx, d, testKeyBadElement))
		require.Equal(t, 2, ctx.ErrorCount())
	})
}

func TestOptionalObjectArray(t *testing.T) {
	m1 := map[string]interface{}{"a": "A"}
	m2 := map[string]interface{}{"b": "B"}
	m3 := map[string]interface{}{"c": "C"}
	barray := []map[string]interface{}{m1, m2, m3}
	m := map[string]interface{}{
		testKeyGood:       []interface{}{m1, m2, m3},
		testKeyEmpty:      []interface{}{},
		testKeyBadElement: []interface{}{m1, m2, "hello", m3, 1.4},
		testKeyBad:        testNumber,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.OptionalObjectArray(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalObjectArray(ctx, d, testKeyMissing), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalObjectArray(ctx, d, testKeyBad), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalObjectArray(ctx, d, testKeyEmpty), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadElement", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.OptionalObjectArray(ctx, d, testKeyBadElement))
		require.Equal(t, 2, ctx.ErrorCount())
	})
}

func TestOptionalStringArray(t *testing.T) {
	barray := []string{"a", "b", "c"}
	m := map[string]interface{}{
		testKeyGood:       []interface{}{"a", "b", "c"},
		testKeyEmpty:      []interface{}{},
		testKeyBadElement: []interface{}{"a", "b", false, "c", 1.4},
		testKeyBad:        testNumber,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.OptionalStringArray(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalStringArray(ctx, d, testKeyMissing), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalStringArray(ctx, d, testKeyBad), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalStringArray(ctx, d, testKeyEmpty), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadElement", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.OptionalStringArray(ctx, d, testKeyBadElement))
		require.Equal(t, 2, ctx.ErrorCount())
	})
}

func TestOptionalStringEnumArray(t *testing.T) {
	allowed := []string{"a", "b", "c"}
	barray := []string{"c", "b", "a"}
	m := map[string]interface{}{
		testKeyGood:       []interface{}{"c", "b", "a"},
		testKeyEmpty:      []interface{}{},
		testKeyBadElement: []interface{}{"c", "b", "hello", "a", 1.4},
		testKeyBad:        testNumber,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.OptionalStringEnumArray(ctx, d, testKeyGood, allowed))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalStringEnumArray(ctx, d, testKeyMissing, allowed), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalStringEnumArray(ctx, d, testKeyBad, allowed), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.OptionalStringEnumArray(ctx, d, testKeyEmpty, allowed), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadElement", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.OptionalStringEnumArray(ctx, d, testKeyBadElement, allowed))
		require.Equal(t, 2, ctx.ErrorCount())
	})
}
