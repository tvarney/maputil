package mpath_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/maputil/mpath"
)

func TestPath(t *testing.T) {
	t.Parallel()
	dn := mpath.DotNotation{}
	t.Run("New", func(t *testing.T) {
		t.Parallel()
		p := mpath.New(dn, mpath.Key("one"), mpath.Index(2))
		require.Equal(
			t, &mpath.Path{
				Elements: []mpath.Element{mpath.Key("one"), mpath.Index(2)},
				Style:    dn,
			}, p,
		)
	})
	t.Run("Parse", func(t *testing.T) {
		t.Parallel()
		t.Run("Good", func(t *testing.T) {
			p, err := mpath.Parse(dn, "one.two[3]")
			require.NoError(t, err)
			require.Equal(t, &mpath.Path{
				Elements: []mpath.Element{mpath.Key("one"), mpath.Key("two"), mpath.Index(3)},
				Style:    dn,
			}, p)
		})
		t.Run("Error", func(t *testing.T) {
			p, err := mpath.Parse(dn, "one[2]three")
			require.EqualError(t, err, mpath.ErrMissingSep.Error())
			require.Nil(t, p)
		})
	})
	t.Run("Add", func(t *testing.T) {
		t.Parallel()
		p := mpath.New(dn, mpath.Key("one"))
		require.Equal(t, p, p.Add(mpath.Index(2)))
		require.Equal(t, []mpath.Element{mpath.Key("one"), mpath.Index(2)}, p.Elements)
	})
	t.Run("Pop", func(t *testing.T) {
		t.Parallel()
		t.Run("Empty", func(t *testing.T) {
			t.Parallel()
			p := mpath.New(dn)
			require.Equal(t, p, p.Pop())
			require.Len(t, p.Elements, 0)
		})
		t.Run("SingleElement", func(t *testing.T) {
			t.Parallel()
			p := mpath.New(dn, mpath.Key("one"))
			require.Equal(t, p, p.Pop())
			require.Len(t, p.Elements, 0)
		})
		t.Run("MultiElement", func(t *testing.T) {
			t.Parallel()
			p := mpath.New(dn, mpath.Key("one"), mpath.Key("two"), mpath.Index(3))
			require.Equal(t, p, p.Pop())
			require.Equal(t, []mpath.Element{mpath.Key("one"), mpath.Key("two")}, p.Elements)
		})
	})
	t.Run("PopN", func(t *testing.T) {
		t.Parallel()
		t.Run("MoreThanSize", func(t *testing.T) {
			t.Parallel()
			p := mpath.New(dn, mpath.Key("one"), mpath.Key("two"), mpath.Key("three"))
			require.Equal(t, p, p.PopN(5))
			require.Len(t, p.Elements, 0)
		})
		t.Run("SameAsSize", func(t *testing.T) {
			t.Parallel()
			p := mpath.New(dn, mpath.Key("one"), mpath.Key("two"), mpath.Key("three"))
			require.Equal(t, p, p.PopN(3))
			require.Len(t, p.Elements, 0)
		})
		t.Run("LessThanSize", func(t *testing.T) {
			t.Parallel()
			p := mpath.New(dn, mpath.Key("one"), mpath.Key("two"), mpath.Key("three"))
			require.Equal(t, p, p.PopN(2))
			require.Equal(t, []mpath.Element{mpath.Key("one")}, p.Elements)
		})
	})
	t.Run("Clear", func(t *testing.T) {
		t.Parallel()
		t.Run("Empty", func(t *testing.T) {
			t.Parallel()
			p := mpath.New(dn)
			require.Equal(t, p, p.Clear())
			require.Len(t, p.Elements, 0)
		})
		t.Run("NonEmpty", func(t *testing.T) {
			t.Parallel()
			p := mpath.New(dn, mpath.Key("one"), mpath.Index(2), mpath.Index(3))
			require.Equal(t, p, p.Clear())
			require.Len(t, p.Elements, 0)
		})
	})
	t.Run("Copy", func(t *testing.T) {
		t.Parallel()
		t.Run("Empty", func(t *testing.T) {
			t.Parallel()
			p := mpath.New(dn)
			c := p.Copy()
			require.Equal(t, c, p)
			c.Add(mpath.Key("one"))
			require.NotEqual(t, c, p)
		})
		t.Run("NonEmpty", func(t *testing.T) {
			t.Parallel()
			p := mpath.New(dn, mpath.Key("one"), mpath.Index(2), mpath.RangeEmpty())
			c := p.Copy()
			require.Equal(t, p, c)
			c.PopN(2).Add(mpath.Key("three")).Add(mpath.Index(-1))
			require.NotEqual(t, c, p)
		})
	})
	t.Run("String", func(t *testing.T) {
		t.Parallel()
		t.Run("NoFile", func(t *testing.T) {
			t.Parallel()
			p := mpath.New(dn, mpath.Key("one"), mpath.Index(2))
			require.Equal(t, "one[2]", p.String())
		})
		t.Run("WithFile", func(t *testing.T) {
			t.Parallel()
			p := mpath.New(dn, mpath.Key("one"), mpath.Index(2))
			p.Filename = "file.json"
			require.Equal(t, "file.json: one[2]", p.String())
		})
	})
}
