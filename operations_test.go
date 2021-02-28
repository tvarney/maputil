package maputil_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/maputil"
)

func TestKeys(t *testing.T) {
	t.Parallel()
	t.Run("Nil", func(t *testing.T) {
		t.Parallel()
		require.Nil(t, maputil.Keys(nil))
	})
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		require.Nil(t, maputil.Keys(map[string]interface{}{}))
	})
	t.Run("NonEmpty", func(t *testing.T) {
		t.Parallel()
		keys := maputil.Keys(map[string]interface{}{
			"one":   1,
			"two":   2,
			"three": 3,
		})
		require.ElementsMatch(t, keys, []string{"one", "two", "three"})
	})
}
