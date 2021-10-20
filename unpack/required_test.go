package unpack_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/maputil"
	"github.com/tvarney/maputil/errctx"
	"github.com/tvarney/maputil/unpack"
)

func TestRequireArray(t *testing.T) {
	m := map[string]interface{}{
		testKeyGood: testArray,
		testKeyBad:  testNumber,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testArray, unpack.RequireArray(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireArray(ctx, d, testKeyMissing), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireArray(ctx, d, testKeyBad), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestRequireBoolean(t *testing.T) {
	m := map[string]interface{}{
		testKeyGood: true,
		testKeyBad:  testNumber,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.True(t, unpack.RequireBoolean(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.False(t, unpack.RequireBoolean(ctx, d, testKeyMissing))
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.False(t, unpack.RequireBoolean(ctx, d, testKeyBad))
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestRequireInteger(t *testing.T) {
	m := map[string]interface{}{
		testKeyGood: testInt,
		testKeyBad:  testString,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testInt, unpack.RequireInteger(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Zero(t, unpack.RequireInteger(ctx, d, testKeyMissing))
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Zero(t, unpack.RequireInteger(ctx, d, testKeyBad))
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestRequireNull(t *testing.T) {
	m := map[string]interface{}{
		testKeyGood: nil,
		testKeyBad:  testNumber,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		unpack.RequireNull(ctx, d, testKeyGood)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		unpack.RequireNull(ctx, d, testKeyMissing)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		unpack.RequireNull(ctx, d, testKeyBad)
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestRequireNumber(t *testing.T) {
	m := map[string]interface{}{
		testKeyGood: testNumber,
		testKeyBad:  true,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testNumber, unpack.RequireNumber(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Zero(t, unpack.RequireNumber(ctx, d, testKeyMissing))
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Zero(t, unpack.RequireNumber(ctx, d, testKeyBad))
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestRequireObject(t *testing.T) {
	m := map[string]interface{}{
		testKeyGood: testObject,
		testKeyBad:  testNumber,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testObject, unpack.RequireObject(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Nil(t, unpack.RequireObject(ctx, d, testKeyMissing))
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Nil(t, unpack.RequireObject(ctx, d, testKeyBad))
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestRequireString(t *testing.T) {
	m := map[string]interface{}{
		testKeyGood: testString,
		testKeyBad:  testNumber,
	}
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, testString, unpack.RequireString(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Zero(t, unpack.RequireString(ctx, d, testKeyMissing))
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Zero(t, unpack.RequireString(ctx, d, testKeyBad))
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestRequireStringEnum(t *testing.T) {
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
		require.Equal(t, testString, unpack.RequireStringEnum(ctx, d, testKeyGood, allowed))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Zero(t, unpack.RequireStringEnum(ctx, d, testKeyMissing, allowed))
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadValue", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, badValue, unpack.RequireStringEnum(ctx, d, testKeyBadElement, allowed))
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Zero(t, unpack.RequireStringEnum(ctx, d, testKeyBad, allowed))
		require.Equal(t, 1, ctx.ErrorCount())
	})
}

func TestRequireBooleanArray(t *testing.T) {
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
		require.Equal(t, barray, unpack.RequireBooleanArray(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireBooleanArray(ctx, d, testKeyMissing), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireBooleanArray(ctx, d, testKeyBad), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireBooleanArray(ctx, d, testKeyEmpty), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadElement", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.RequireBooleanArray(ctx, d, testKeyBadElement))
		require.Equal(t, 2, ctx.ErrorCount())
	})
}

func TestRequireIntegerArray(t *testing.T) {
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
		require.Equal(t, barray, unpack.RequireIntegerArray(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireIntegerArray(ctx, d, testKeyMissing), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireIntegerArray(ctx, d, testKeyBad), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireIntegerArray(ctx, d, testKeyEmpty), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadElement", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.RequireIntegerArray(ctx, d, testKeyBadElement))
		require.Equal(t, 2, ctx.ErrorCount())
	})
}

func TestRequireNumberArray(t *testing.T) {
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
		require.Equal(t, barray, unpack.RequireNumberArray(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireNumberArray(ctx, d, testKeyMissing), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireNumberArray(ctx, d, testKeyBad), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireNumberArray(ctx, d, testKeyEmpty), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadElement", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.RequireNumberArray(ctx, d, testKeyBadElement))
		require.Equal(t, 2, ctx.ErrorCount())
	})
}

func TestRequireObjectArray(t *testing.T) {
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
		require.Equal(t, barray, unpack.RequireObjectArray(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireObjectArray(ctx, d, testKeyMissing), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireObjectArray(ctx, d, testKeyBad), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireObjectArray(ctx, d, testKeyEmpty), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadElement", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.RequireObjectArray(ctx, d, testKeyBadElement))
		require.Equal(t, 2, ctx.ErrorCount())
	})
}

func TestRequireStringArray(t *testing.T) {
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
		require.Equal(t, barray, unpack.RequireStringArray(ctx, d, testKeyGood))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireStringArray(ctx, d, testKeyMissing), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireStringArray(ctx, d, testKeyBad), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireStringArray(ctx, d, testKeyEmpty), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadElement", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.RequireStringArray(ctx, d, testKeyBadElement))
		require.Equal(t, 2, ctx.ErrorCount())
	})
}

func TestRequireStringEnumArray(t *testing.T) {
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
		require.Equal(t, barray, unpack.RequireStringEnumArray(ctx, d, testKeyGood, allowed))
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("Missing", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireStringEnumArray(ctx, d, testKeyMissing, allowed), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireStringEnumArray(ctx, d, testKeyBad, allowed), 0)
		require.Equal(t, 1, ctx.ErrorCount())
	})
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Len(t, unpack.RequireStringEnumArray(ctx, d, testKeyEmpty, allowed), 0)
		require.Zero(t, ctx.ErrorCount())
	})
	t.Run("BadElement", func(t *testing.T) {
		t.Parallel()
		d := maputil.Copy(m)
		ctx := errctx.New(errctx.ErrorDiscarder{})
		require.Equal(t, barray, unpack.RequireStringEnumArray(ctx, d, testKeyBadElement, allowed))
		require.Equal(t, 2, ctx.ErrorCount())
	})
}
