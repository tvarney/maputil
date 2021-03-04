package errctx_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/maputil/errctx"
	"github.com/tvarney/maputil/mpath"
)

func TestErrorPrinter(t *testing.T) {
	t.Parallel()
	t.Run("Add", func(t *testing.T) {
		t.Parallel()
		sb := &strings.Builder{}
		ep := &errctx.ErrorPrinter{Stream: sb}
		p := mpath.New(mpath.DotNotation{}, mpath.Key("one"))
		ep.Add(p, errors.New("error one"))
		p.Add(mpath.Key("two"))
		ep.Add(p, errors.New("error two"))
		require.Equal(t, "one: error one\none.two: error two\n", sb.String())
	})
}

func TestMultiHandler(t *testing.T) {
	t.Parallel()
	t.Run("Add", func(t *testing.T) {
		t.Parallel()
		sb1 := &strings.Builder{}
		sb2 := &strings.Builder{}
		h := errctx.NewMultiHandler(
			&errctx.ErrorPrinter{Stream: sb1},
			&errctx.ErrorPrinter{Stream: sb2},
		)
		p := mpath.New(mpath.DotNotation{}, mpath.Key("one"))
		h.Add(p, errors.New("error one"))
		p.Add(mpath.Key("two"))
		h.Add(p, errors.New("error two"))

		const expected = "one: error one\none.two: error two\n"
		require.Equal(t, expected, sb1.String())
		require.Equal(t, expected, sb2.String())
	})
}
