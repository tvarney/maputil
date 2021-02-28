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

func TestCopy(t *testing.T) {
	t.Parallel()
	t.Run("Nil", func(t *testing.T) {
		t.Parallel()
		require.Nil(t, maputil.Copy(nil))
	})
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		d := map[string]interface{}{}
		c := maputil.Copy(d)
		require.NotNil(t, c)
		require.Equal(t, d, c)
		// Ensure they are different instances
		c["hello"] = "world"
		require.NotEqual(t, d, c)
	})
	t.Run("NonEmpty", func(t *testing.T) {
		t.Parallel()
		d := map[string]interface{}{
			"one": "two",
			"two": []interface{}{1, 2, 3},
			"three": map[string]interface{}{
				"hello": "world",
			},
		}
		c := maputil.Copy(d)
		require.Equal(t, d, c)
		// Check that the returned map is a different instance.
		c["four"] = "test"
		require.NotEqual(t, d, c)
		// Check that the child array is a different instance
		c["two"].([]interface{})[1] = 0
		require.NotEqual(t, d["two"], c["two"])
		// Check that the child map is a different instance
		c["three"].(map[string]interface{})["hello"] = "test"
		require.NotEqual(t, d["three"], c["three"])
	})
}

func TestCopyArray(t *testing.T) {
	t.Parallel()
	t.Run("Nil", func(t *testing.T) {
		t.Parallel()
		require.Nil(t, maputil.CopyArray(nil))
	})
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		d := []interface{}{}
		require.Nil(t, maputil.CopyArray(d))
	})
	t.Run("NotEmpty", func(t *testing.T) {
		t.Parallel()
		a := []interface{}{1, []interface{}{1, 2, 3}, map[string]interface{}{
			"hello": "world",
		}}
		c := maputil.CopyArray(a)
		require.Equal(t, a, c)
		// Check that the returned array is a different instance.
		c[0] = "test"
		require.NotEqual(t, a, c)
		// Check that the child array is a different instance.
		(c[1].([]interface{}))[0] = 0
		require.NotEqual(t, a[1], c[1])
		// Check that the child map is a different instance.
		(c[2].(map[string]interface{}))["hello"] = "test"
		require.NotEqual(t, a[2], c[2])
	})
}
