package errctx_test

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/maputil/errctx"
	"github.com/tvarney/maputil/mpath"
)

func TestContext(t *testing.T) {
	t.Parallel()
	t.Run("New", func(t *testing.T) {
		t.Parallel()
		t.Run("NoHandlers", func(t *testing.T) {
			t.Parallel()
			expected := &errctx.Context{
				Path:    mpath.New(mpath.DotNotation{}),
				Handler: nil,
			}
			ctx := errctx.New()
			require.Equal(t, expected, ctx)
		})
		t.Run("OneHandler", func(t *testing.T) {
			t.Parallel()
			ep1 := &errctx.ErrorPrinter{Stream: &strings.Builder{}}
			expected := &errctx.Context{
				Path:    mpath.New(mpath.DotNotation{}),
				Handler: ep1,
			}
			ctx := errctx.New(ep1)
			require.Equal(t, expected, ctx)
		})
		t.Run("MultipleHandlers", func(t *testing.T) {
			t.Parallel()
			ep1 := &errctx.ErrorPrinter{Stream: os.Stderr}
			ep2 := &errctx.ErrorPrinter{Stream: os.Stdout}
			expected := &errctx.Context{
				Path:    mpath.New(mpath.DotNotation{}),
				Handler: errctx.NewMultiHandler(ep1, ep2),
			}
			ctx := errctx.New(ep1, ep2)
			require.Equal(t, expected, ctx)
		})
	})
	t.Run("ErrorCount", func(t *testing.T) {
		t.Parallel()
		ctx := errctx.New()
		require.Zero(t, ctx.ErrorCount())
		ctx.Error(errors.New("test"))
		require.Equal(t, 1, ctx.ErrorCount())
		ctx.Error(errors.New("test"))
		require.Equal(t, 2, ctx.ErrorCount())
	})
	t.Run("LastError", func(t *testing.T) {
		t.Parallel()
		ctx := errctx.New()
		require.Nil(t, ctx.LastError())
		err1 := errors.New("test 1")
		ctx.Error(err1)
		require.Equal(t, err1, ctx.LastError())
		err2 := errors.New("test 2")
		ctx.Error(err2)
		require.Equal(t, err2, ctx.LastError())
	})
	t.Run("Reset", func(t *testing.T) {
		t.Parallel()
		ctx := errctx.New()
		ctx.Error(errors.New("test 1"))
		require.Equal(t, 1, ctx.ErrorCount())
		require.NotNil(t, ctx.LastError())
		ctx.Reset()
		require.Zero(t, ctx.ErrorCount())
		require.Nil(t, ctx.LastError())
	})
	t.Run("Error", func(t *testing.T) {
		t.Parallel()
		t.Run("NotNil", func(t *testing.T) {
			sb := &strings.Builder{}
			ctx := errctx.New(&errctx.ErrorPrinter{Stream: sb})
			err := errors.New("one")
			ctx.Error(err)
			require.Equal(t, 1, ctx.ErrorCount())
			require.Equal(t, ": one\n", sb.String())
		})
		t.Run("Nil", func(t *testing.T) {
			sb := &strings.Builder{}
			ctx := errctx.New(&errctx.ErrorPrinter{Stream: sb})
			ctx.Error(nil)
			require.Zero(t, ctx.ErrorCount())
			require.Zero(t, sb.String())
		})
	})
	t.Run("ErrorWith", func(t *testing.T) {
		t.Parallel()
		t.Run("NotNil", func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			ctx := errctx.New(&errctx.ErrorPrinter{Stream: sb})
			ctx.ErrorWith(errors.New("test"), mpath.Key("one"))
			require.Equal(t, "one: test\n", sb.String())
			require.Len(t, ctx.Path.Elements, 0)
		})
		t.Run("Nil", func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			ctx := errctx.New(&errctx.ErrorPrinter{Stream: sb})
			ctx.ErrorWith(nil, mpath.Key("one"))
			require.Zero(t, sb.String())
			require.Len(t, ctx.Path.Elements, 0)
		})
	})
	t.Run("ErrorWithKey", func(t *testing.T) {
		t.Parallel()
		t.Run("NotNil", func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			ctx := errctx.New(&errctx.ErrorPrinter{Stream: sb})
			ctx.ErrorWithKey(errors.New("test"), "one")
			require.Equal(t, "one: test\n", sb.String())
			require.Len(t, ctx.Path.Elements, 0)
		})
		t.Run("Nil", func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			ctx := errctx.New(&errctx.ErrorPrinter{Stream: sb})
			ctx.ErrorWithKey(nil, "one")
			require.Zero(t, sb.String())
			require.Len(t, ctx.Path.Elements, 0)
		})
	})
	t.Run("ErrorWithIndex", func(t *testing.T) {
		t.Parallel()
		t.Run("NotNil", func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			ctx := errctx.New(&errctx.ErrorPrinter{Stream: sb})
			ctx.Path.Add(mpath.Key("one"))
			ctx.ErrorWithIndex(errors.New("test"), 1)
			require.Equal(t, "one[1]: test\n", sb.String())
			require.Len(t, ctx.Path.Elements, 1)
		})
		t.Run("Nil", func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			ctx := errctx.New(&errctx.ErrorPrinter{Stream: sb})
			ctx.Path.Add(mpath.Key("one"))
			ctx.ErrorWithIndex(nil, 1)
			require.Zero(t, sb.String())
			require.Len(t, ctx.Path.Elements, 1)
		})
	})
}
