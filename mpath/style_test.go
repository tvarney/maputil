package mpath_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/maputil/mpath"
)

func TestDotNotation(t *testing.T) {
	t.Parallel()
	dn := mpath.DotNotation{}
	t.Run("Strict", func(t *testing.T) {
		t.Parallel()
		require.True(t, dn.Strict())
	})
	t.Run("Format", func(t *testing.T) {
		t.Parallel()
		t.Run("Empty", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, "", dn.Format(nil))
		})
		t.Run("ArrayStart", func(t *testing.T) {
			t.Parallel()
			require.Equal(
				t, "[1][2].three.four[5]", dn.Format([]mpath.Element{
					mpath.Index(1), mpath.Index(2), mpath.Key("three"),
					mpath.Key("four"), mpath.Index(5),
				}),
			)
		})
		t.Run("KeyStart", func(t *testing.T) {
			t.Parallel()
			require.Equal(
				t, "one.two.three[4].five", dn.Format([]mpath.Element{
					mpath.Key("one"), mpath.Key("two"), mpath.Key("three"),
					mpath.Index(4), mpath.Key("five"),
				}),
			)
		})
	})
	t.Run("Parse", func(t *testing.T) {
		t.Parallel()
		t.Run("Empty", func(t *testing.T) {
			t.Parallel()
			p, err := dn.Parse("")
			require.NoError(t, err)
			require.Nil(t, p)
		})
		t.Run("Dot", func(t *testing.T) {
			t.Parallel()
			p, err := dn.Parse(".")
			require.NoError(t, err)
			require.Equal(t, []mpath.Element{mpath.Key("")}, p)
		})
		t.Run("Good", func(t *testing.T) {
			t.Parallel()
			p, err := dn.Parse("one.two[3][1:2][-]")
			require.NoError(t, err)
			require.Equal(
				t, []mpath.Element{
					mpath.Key("one"), mpath.Key("two"), mpath.Index(3),
					mpath.RangeFull(1, 2), mpath.ArrayEnd{},
				}, p,
			)
		})
		t.Run("Escapes", func(t *testing.T) {
			t.Parallel()
			p, err := dn.Parse(`one\.\[12\].two\\`)
			require.NoError(t, err)
			require.Equal(t, []mpath.Element{mpath.Key(`one.[12]`), mpath.Key(`two\`)}, p)
		})
		t.Run("BadEscape", func(t *testing.T) {
			t.Parallel()
			p, err := dn.Parse(`12\`)
			require.EqualError(t, err, mpath.ErrInvalidEscape.Error())
			require.Nil(t, p)
		})
		t.Run("UnclosedIndex", func(t *testing.T) {
			t.Parallel()
			p, err := dn.Parse("one[123")
			require.EqualError(t, err, mpath.ErrUnmatchedOpenBracket.Error())
			require.Nil(t, p)
		})
		t.Run("UnmatchedClose", func(t *testing.T) {
			t.Parallel()
			p, err := dn.Parse("one.]1[")
			require.EqualError(t, err, mpath.ErrUnmatchedCloseBracket.Error())
			require.Nil(t, p)
		})
		t.Run("MissingSep", func(t *testing.T) {
			t.Parallel()
			p, err := dn.Parse("one[12]two")
			require.EqualError(t, err, mpath.ErrMissingSep.Error())
			require.Nil(t, p)
		})
	})
}
