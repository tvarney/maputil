package consterr_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/maputil/consterr"
)

func TestError(t *testing.T) {
	t.Parallel()
	t.Run("Error", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, "hello", consterr.Error("hello").Error())
	})
}
