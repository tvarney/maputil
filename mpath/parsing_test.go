package mpath_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/maputil/mpath"
)

func TestParseRange(t *testing.T) {
	t.Parallel()
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		const value = ":"
		const sep = 0
		expected := mpath.RangeEmpty()
		r, err := mpath.ParseRange([]rune(value), sep)
		require.NoError(t, err)
		require.Equal(t, expected, r)
	})
	t.Run("NoStart", func(t *testing.T) {
		t.Parallel()
		const value = ":15"
		const sep = 0
		expected := mpath.RangeEnd(15)
		r, err := mpath.ParseRange([]rune(value), sep)
		require.NoError(t, err)
		require.Equal(t, expected, r)
	})
	t.Run("NoEnd", func(t *testing.T) {
		t.Parallel()
		const value = "14:"
		const sep = 2
		expected := mpath.RangeStart(14)
		r, err := mpath.ParseRange([]rune(value), sep)
		require.NoError(t, err)
		require.Equal(t, expected, r)
	})
	t.Run("BadStart", func(t *testing.T) {
		t.Parallel()
		const value = "as:10"
		const sep = 2
		r, err := mpath.ParseRange([]rune(value), sep)
		require.EqualError(t, err, mpath.BadRangeStartError{Value: "as"}.Error())
		require.Nil(t, r)
	})
	t.Run("BadEnd", func(t *testing.T) {
		t.Parallel()
		const value = "13:oh"
		const sep = 2
		r, err := mpath.ParseRange([]rune(value), sep)
		require.EqualError(t, err, mpath.BadRangeEndError{Value: "oh"}.Error())
		require.Nil(t, r)
	})
	t.Run("Full", func(t *testing.T) {
		t.Parallel()
		const value = "4:23"
		const sep = 1
		expected := mpath.RangeFull(4, 23)
		r, err := mpath.ParseRange([]rune(value), sep)
		require.NoError(t, err)
		require.Equal(t, expected, r)
	})
}

func TestParseIndexExt(t *testing.T) {
	t.Parallel()
	t.Run("Range", func(t *testing.T) {
		t.Run("Full", func(t *testing.T) {
			t.Parallel()
			const value = "-5:-2"
			const sep = 2
			expected := mpath.RangeFull(-5, -2)
			r, err := mpath.ParseIndexExt([]rune(value), sep)
			require.NoError(t, err)
			require.Equal(t, expected, r)
		})
		t.Run("Empty", func(t *testing.T) {
			t.Parallel()
			const value = ":"
			const sep = 0
			expected := mpath.RangeEmpty()
			r, err := mpath.ParseIndexExt([]rune(value), sep)
			require.NoError(t, err)
			require.Equal(t, expected, r)
		})
		t.Run("NoEnd", func(t *testing.T) {
			t.Parallel()
			const value = "1:"
			const sep = 1
			expected := mpath.RangeStart(1)
			r, err := mpath.ParseIndexExt([]rune(value), sep)
			require.NoError(t, err)
			require.Equal(t, expected, r)
		})
		t.Run("NoStart", func(t *testing.T) {
			t.Parallel()
			const value = ":192"
			const sep = 0
			expected := mpath.RangeEnd(192)
			r, err := mpath.ParseIndexExt([]rune(value), sep)
			require.NoError(t, err)
			require.Equal(t, expected, r)
		})
		t.Run("Error", func(t *testing.T) {
			t.Parallel()
			const value = "a:b"
			const sep = 1
			r, err := mpath.ParseIndexExt([]rune(value), sep)
			require.Error(t, err)
			require.Nil(t, r)
		})
	})
	t.Run("ArrayEnd", func(t *testing.T) {
		t.Parallel()
		const value = "-"
		const sep = -1
		expected := mpath.ArrayEnd{}
		r, err := mpath.ParseIndexExt([]rune(value), sep)
		require.NoError(t, err)
		require.Equal(t, expected, r)
	})
	t.Run("Index", func(t *testing.T) {
		t.Parallel()
		t.Run("Good", func(t *testing.T) {
			t.Parallel()
			const value = "43"
			const sep = -1
			expected := mpath.Index(43)
			r, err := mpath.ParseIndexExt([]rune(value), sep)
			require.NoError(t, err)
			require.Equal(t, expected, r)
		})
		t.Run("Bad", func(t *testing.T) {
			t.Parallel()
			const value = "blah"
			const sep = -1
			r, err := mpath.ParseIndexExt([]rune(value), sep)
			require.EqualError(t, err, mpath.ErrBadIndex.Error())
			require.Nil(t, r)
		})
	})
}

func TestParseIndex(t *testing.T) {
	t.Parallel()
	t.Run("Index", func(t *testing.T) {
		t.Parallel()
		const value = "1234"
		expected := mpath.Index(1234)
		r, err := mpath.ParseIndex([]rune(value))
		require.NoError(t, err)
		require.Equal(t, expected, r)
	})
	t.Run("Range", func(t *testing.T) {
		t.Parallel()
		t.Run("Full", func(t *testing.T) {
			t.Parallel()
			const value = "123:45"
			expected := mpath.RangeFull(123, 45)
			r, err := mpath.ParseIndex([]rune(value))
			require.NoError(t, err)
			require.Equal(t, expected, r)
		})
		t.Run("NoEnd", func(t *testing.T) {
			t.Parallel()
			const value = "-13:"
			expected := mpath.RangeStart(-13)
			r, err := mpath.ParseIndex([]rune(value))
			require.NoError(t, err)
			require.Equal(t, expected, r)
		})
		t.Run("NoStart", func(t *testing.T) {
			t.Parallel()
			const value = ":777"
			expected := mpath.RangeEnd(777)
			r, err := mpath.ParseIndex([]rune(value))
			require.NoError(t, err)
			require.Equal(t, expected, r)
		})
		t.Run("Empty", func(t *testing.T) {
			t.Parallel()
			const value = ":"
			expected := mpath.RangeEmpty()
			r, err := mpath.ParseIndex([]rune(value))
			require.NoError(t, err)
			require.Equal(t, expected, r)
		})
		t.Run("MultipleSeparators", func(t *testing.T) {
			t.Parallel()
			const value = "1:2:3:4"
			r, err := mpath.ParseIndex([]rune(value))
			require.EqualError(t, err, mpath.ErrBadRange.Error())
			require.Nil(t, r)
		})
	})
	t.Run("ArrayEnd", func(t *testing.T) {
		t.Parallel()
		const value = "-"
		expected := mpath.ArrayEnd{}
		r, err := mpath.ParseIndex([]rune(value))
		require.NoError(t, err)
		require.Equal(t, expected, r)
	})
}
