package maputil_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/maputil"
)

func TestConstError(t *testing.T) {
	t.Parallel()
	t.Run("Error", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, string(maputil.ErrInvalidType), maputil.ErrInvalidType.Error())
	})
}

func TestInvalidTypeError(t *testing.T) {
	t.Parallel()
	t.Run("Error", func(t *testing.T) {
		t.Parallel()
		t.Run("NoTypes", func(t *testing.T) {
			t.Parallel()
			e := maputil.InvalidTypeError{
				Actual:   "string",
				Expected: nil,
			}
			require.Equal(t, string(maputil.ErrInvalidType)+" string", e.Error())
		})
		t.Run("OneType", func(t *testing.T) {
			t.Parallel()
			e := maputil.InvalidTypeError{
				Actual:   "string",
				Expected: []string{"array"},
			}
			require.Equal(
				t, string(maputil.ErrInvalidType)+" string; expected array",
				e.Error(),
			)
		})
		t.Run("TwoTypes", func(t *testing.T) {
			t.Parallel()
			e := maputil.InvalidTypeError{
				Actual:   "string",
				Expected: []string{"array", "object"},
			}
			require.Equal(
				t, string(maputil.ErrInvalidType)+" string; expected array or object",
				e.Error(),
			)
		})
		t.Run("ThreeTypes", func(t *testing.T) {
			t.Parallel()
			e := maputil.InvalidTypeError{
				Actual:   "string",
				Expected: []string{"array", "object", "bool"},
			}
			require.Equal(
				t, string(maputil.ErrInvalidType)+" string; expected array, object, or bool",
				e.Error(),
			)
		})
		t.Run("ManyTypes", func(t *testing.T) {
			t.Parallel()
			e := maputil.InvalidTypeError{
				Actual:   maputil.TypeString,
				Expected: []string{"array", "bool", "integer", "number", "object"},
			}
			require.Equal(
				t, string(maputil.ErrInvalidType)+" string; expected array, bool, integer, number, or object",
				e.Error(),
			)
		})
	})
	t.Run("Unwrap", func(t *testing.T) {
		t.Parallel()
		require.True(t, errors.Is(maputil.InvalidTypeError{}, maputil.ErrInvalidType))
	})
}

func TestEnumStringError(t *testing.T) {
	t.Parallel()
	t.Run("Error", func(t *testing.T) {
		t.Parallel()
		const prefix = string(maputil.ErrInvalidValue)
		t.Run("NoValues", func(t *testing.T) {
			e := maputil.EnumStringError{
				Value: "ten",
				Enum:  []string{},
			}
			require.Equal(t, prefix+` "ten"`, e.Error())
		})
		t.Run("OneValue", func(t *testing.T) {
			t.Parallel()
			e := maputil.EnumStringError{
				Value: "ten",
				Enum:  []string{"one"},
			}
			require.Equal(t, prefix+` "ten"; expected "one"`, e.Error())
		})
		t.Run("TwoValues", func(t *testing.T) {
			t.Parallel()
			e := maputil.EnumStringError{
				Value: "ten",
				Enum:  []string{"one", "two"},
			}
			require.Equal(t, prefix+` "ten"; expected "one" or "two"`, e.Error())
		})
		t.Run("ThreeValues", func(t *testing.T) {
			t.Parallel()
			e := maputil.EnumStringError{
				Value: "ten",
				Enum:  []string{"one", "two", "three"},
			}
			require.Equal(t, prefix+` "ten"; expected one of "one", "two", or "three"`, e.Error())
		})
		t.Run("ManyValues", func(t *testing.T) {
			t.Parallel()
			e := maputil.EnumStringError{
				Value: "ten",
				Enum:  []string{"one", "two", "three", "four", "five", "six"},
			}
			require.Equal(
				t, prefix+` "ten"; expected one of "one", "two", "three", "four", "five", or "six"`,
				e.Error(),
			)
		})
	})
	t.Run("Unwrap", func(t *testing.T) {
		t.Parallel()
		require.True(t, errors.Is(maputil.EnumStringError{}, maputil.ErrInvalidValue))
	})
}

func TestMissingRequiredValueError(t *testing.T) {
	t.Parallel()
	t.Run("Error", func(t *testing.T) {
		t.Parallel()
		e := maputil.MissingRequiredValueError{
			Key: "one",
		}
		require.Equal(
			t, string(maputil.ErrMissingRequiredValue)+` "one"`,
			e.Error(),
		)
	})
	t.Run("Unwrap", func(t *testing.T) {
		t.Parallel()
		require.True(t, errors.Is(maputil.MissingRequiredValueError{}, maputil.ErrMissingRequiredValue))
	})
}
