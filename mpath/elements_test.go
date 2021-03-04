package mpath_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/maputil/mpath"
)

func TestKey(t *testing.T) {
	t.Parallel()
	t.Run("Type", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, mpath.KeyType, mpath.Key("a").Type())
	})
	t.Run("String", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, "test", mpath.Key("test").String())
	})
	t.Run("Copy", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, mpath.Key("test"), mpath.Key("test").Copy())
	})
}

func TestIndex(t *testing.T) {
	t.Parallel()
	t.Run("Type", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, mpath.IndexType, mpath.Index(1).Type())
	})
	t.Run("String", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, "-2", mpath.Index(-2).String())
	})
	t.Run("Copy", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, mpath.Index(13), mpath.Index(13).Copy())
	})
}

func TestRange(t *testing.T) {
	t.Parallel()
	t.Run("RangeFull", func(t *testing.T) {
		t.Parallel()
		expected := mpath.Range{Start: 13, End: -2, Tag: mpath.RangeTagFull}
		require.Equal(t, expected, mpath.RangeFull(13, -2))
	})
	t.Run("RangeStart", func(t *testing.T) {
		t.Parallel()
		expected := mpath.Range{Start: 6, End: 0, Tag: mpath.RangeTagNoEnd}
		require.Equal(t, expected, mpath.RangeStart(6))
	})
	t.Run("RangeEnd", func(t *testing.T) {
		t.Parallel()
		expected := mpath.Range{Start: 0, End: -2, Tag: mpath.RangeTagNoStart}
		require.Equal(t, expected, mpath.RangeEnd(-2))
	})
	t.Run("RangeEmpty", func(t *testing.T) {
		t.Parallel()
		expected := mpath.Range{Start: 0, End: 0, Tag: mpath.RangeTagEmpty}
		require.Equal(t, expected, mpath.RangeEmpty())
	})
	t.Run("Type", func(t *testing.T) {
		t.Parallel()
		t.Run("Empty", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, mpath.IndexType, mpath.RangeEmpty().Type())
		})
		t.Run("Full", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, mpath.IndexType, mpath.RangeFull(1, 2).Type())
		})
		t.Run("Start", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, mpath.IndexType, mpath.RangeStart(1).Type())
		})
		t.Run("End", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, mpath.IndexType, mpath.RangeEnd(2).Type())
		})
	})
	t.Run("String", func(t *testing.T) {
		t.Parallel()
		t.Run("Full", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, "1:2", mpath.RangeFull(1, 2).String())
		})
		t.Run("Empty", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, ":", mpath.RangeEmpty().String())
		})
		t.Run("NoStart", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, ":-1", mpath.RangeEnd(-1).String())
		})
		t.Run("NoEnd", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, "-5:", mpath.RangeStart(-5).String())
		})
	})
	t.Run("Copy", func(t *testing.T) {
		t.Parallel()
		t.Run("Full", func(t *testing.T) {
			t.Parallel()
			expected := mpath.RangeFull(1, 2)
			require.Equal(t, expected, expected.Copy())
		})
		t.Run("Empty", func(t *testing.T) {
			t.Parallel()
			expected := mpath.RangeEmpty()
			require.Equal(t, expected, expected.Copy())
		})
		t.Run("NoStart", func(t *testing.T) {
			t.Parallel()
			expected := mpath.RangeEnd(2)
			require.Equal(t, expected, expected.Copy())
		})
		t.Run("NoEnd", func(t *testing.T) {
			t.Parallel()
			expected := mpath.RangeStart(1)
			require.Equal(t, expected, expected.Copy())
		})
	})
}

func TestArrayEnd(t *testing.T) {
	t.Parallel()
	t.Run("Type", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, mpath.IndexType, mpath.ArrayEnd{}.Type())
	})
	t.Run("String", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, "-", mpath.ArrayEnd{}.String())
	})
	t.Run("Copy", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, mpath.ArrayEnd{}, mpath.ArrayEnd{}.Copy())
	})
}
