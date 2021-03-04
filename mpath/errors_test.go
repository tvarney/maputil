package mpath_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/maputil/mpath"
)

func TestBadRangeStartError(t *testing.T) {
	t.Parallel()
	t.Run("Error", func(t *testing.T) {
		t.Parallel()
		require.Equal(
			t, `invalid range; invalid start value "as\tdf"`,
			mpath.BadRangeStartError{Value: "as\tdf"}.Error(),
		)
	})
	t.Run("Unwrap", func(t *testing.T) {
		t.Parallel()
		require.ErrorIs(t, mpath.BadRangeStartError{Value: "asdf"}, mpath.ErrBadRange)
	})
}

func TestBadRangeEndError(t *testing.T) {
	t.Parallel()
	t.Run("Error", func(t *testing.T) {
		t.Parallel()
		require.Equal(
			t, `invalid range; invalid end value "as\tdf"`,
			mpath.BadRangeEndError{Value: "as\tdf"}.Error(),
		)
	})
	t.Run("Unwrap", func(t *testing.T) {
		t.Parallel()
		require.ErrorIs(t, mpath.BadRangeEndError{Value: "asdf"}, mpath.ErrBadRange)
	})
}
