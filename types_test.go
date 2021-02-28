package maputil_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/maputil"
)

func TestTypeName(t *testing.T) {
	t.Parallel()
	t.Run("null", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, maputil.TypeNull, maputil.TypeName(nil))
	})
	t.Run("array", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, maputil.TypeArray, maputil.TypeName([]interface{}{1, 2}))
	})
	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, maputil.TypeBoolean, maputil.TypeName(true))
	})
	t.Run("Int", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, maputil.TypeInteger, maputil.TypeName(int(1)))
	})
	t.Run("Int8", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, maputil.TypeInteger, maputil.TypeName(int8(1)))
	})
	t.Run("Int16", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, maputil.TypeInteger, maputil.TypeName(int16(1)))
	})
	t.Run("Int32", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, maputil.TypeInteger, maputil.TypeName(int32(1)))
	})
	t.Run("Int64", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, maputil.TypeInteger, maputil.TypeName(int64(1)))
	})
	t.Run("Uint", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, maputil.TypeInteger, maputil.TypeName(uint(1)))
	})
	t.Run("Uint8", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, maputil.TypeInteger, maputil.TypeName(uint8(1)))
	})
	t.Run("Uint16", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, maputil.TypeInteger, maputil.TypeName(uint16(1)))
	})
	t.Run("Uint32", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, maputil.TypeInteger, maputil.TypeName(uint32(1)))
	})
	t.Run("Uint64", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, maputil.TypeInteger, maputil.TypeName(uint64(1)))
	})
	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		t.Run("NoDecimal", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, maputil.TypeInteger, maputil.TypeName(float32(10)))
		})
		t.Run("Decimal", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, maputil.TypeNumber, maputil.TypeName(float32(10.1)))
		})
	})
	t.Run("float64", func(t *testing.T) {
		t.Parallel()
		t.Run("NoDecimal", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, maputil.TypeInteger, maputil.TypeName(float64(10)))
		})
		t.Run("Decimal", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, maputil.TypeNumber, maputil.TypeName(float64(10.1)))
		})
	})
	t.Run("Generic", func(t *testing.T) {
		t.Parallel()
		t.Run("Int64", func(t *testing.T) {
			t.Parallel()
			n := json.Number("10")
			require.Equal(t, maputil.TypeInteger, maputil.TypeName(n))
		})
		t.Run("Float64", func(t *testing.T) {
			t.Parallel()
			n := json.Number("10.1")
			require.Equal(t, maputil.TypeNumber, maputil.TypeName(n))
		})
	})
	t.Run("Map", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, maputil.TypeObject, maputil.TypeName(map[string]interface{}{}))
	})
	t.Run("String", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, maputil.TypeString, maputil.TypeName(""))
	})
	t.Run("Other", func(t *testing.T) {
		t.Parallel()
		var c chan int
		require.Equal(t, "golang<chan int>", maputil.TypeName(c))
	})
}

func TestIs(t *testing.T) {
	t.Parallel()
	t.Run("Array", func(t *testing.T) {
		t.Parallel()
		require.NoError(t, maputil.Is([]interface{}{}, maputil.TypeArray))
	})
	t.Run("Bool", func(t *testing.T) {
		t.Parallel()
		require.NoError(t, maputil.Is(true, maputil.TypeBoolean))
	})
	t.Run("Integer", func(t *testing.T) {
		t.Parallel()
		t.Run("Float", func(t *testing.T) {
			t.Parallel()
			t.Run("NoDecimal", func(t *testing.T) {
				t.Parallel()
				require.NoError(t, maputil.Is(float64(10), maputil.TypeInteger))
			})
			t.Run("Decimal", func(t *testing.T) {
				t.Parallel()
				require.EqualError(
					t, maputil.Is(float64(4.2), maputil.TypeInteger),
					maputil.InvalidTypeError{
						Actual:   maputil.TypeNumber,
						Expected: []string{maputil.TypeInteger},
					}.Error(),
				)
			})
		})
		t.Run("Int64", func(t *testing.T) {
			t.Parallel()
			require.NoError(t, maputil.Is(int64(10), maputil.TypeInteger))
		})
	})
	t.Run("Number", func(t *testing.T) {
		t.Parallel()
		t.Run("Float", func(t *testing.T) {
			t.Parallel()
			require.NoError(t, maputil.Is(float64(10.1), maputil.TypeNumber))
		})
		t.Run("Integer", func(t *testing.T) {
			t.Parallel()
			require.NoError(t, maputil.Is(int64(10), maputil.TypeNumber))
		})
	})
	t.Run("Null", func(t *testing.T) {
		t.Parallel()
		require.NoError(t, maputil.Is(nil, maputil.TypeNull))
	})
	t.Run("Object", func(t *testing.T) {
		t.Parallel()
		require.NoError(t, maputil.Is(map[string]interface{}{}, maputil.TypeObject))
	})
	t.Run("String", func(t *testing.T) {
		t.Parallel()
		require.NoError(t, maputil.Is("", maputil.TypeString))
	})
}

func TestAsArray(t *testing.T) {
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		var data interface{} = []interface{}{1, 2, 3}
		a, err := maputil.AsArray(data)
		require.NoError(t, err)
		require.Equal(t, data, a)
	})
	t.Run("Bad", func(t *testing.T) {
		t.Parallel()
		var data interface{} = true
		a, err := maputil.AsArray(data)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Actual:   maputil.TypeBoolean,
			Expected: []string{maputil.TypeArray},
		}.Error())
		require.Nil(t, a)
	})
}

func TestAsBoolean(t *testing.T) {
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		var data interface{} = true
		b, err := maputil.AsBoolean(data)
		require.NoError(t, err)
		require.True(t, b)
	})
	t.Run("Bad", func(t *testing.T) {
		t.Parallel()
		var data interface{} = []interface{}{1, 2, 3}
		b, err := maputil.AsBoolean(data)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Actual:   maputil.TypeArray,
			Expected: []string{maputil.TypeBoolean},
		}.Error())
		require.False(t, b)
	})
}

func TestAsInteger(t *testing.T) {
	t.Parallel()
	t.Run("Int", func(t *testing.T) {
		t.Parallel()
		const data = int(10)
		i, err := maputil.AsInteger(data)
		require.NoError(t, err)
		require.Equal(t, int64(data), i)
	})
	t.Run("Int8", func(t *testing.T) {
		t.Parallel()
		const data = int8(10)
		i, err := maputil.AsInteger(data)
		require.NoError(t, err)
		require.Equal(t, int64(data), i)
	})
	t.Run("Int16", func(t *testing.T) {
		t.Parallel()
		const data = int16(10)
		i, err := maputil.AsInteger(data)
		require.NoError(t, err)
		require.Equal(t, int64(data), i)
	})
	t.Run("Int32", func(t *testing.T) {
		t.Parallel()
		const data = int32(10)
		i, err := maputil.AsInteger(data)
		require.NoError(t, err)
		require.Equal(t, int64(data), i)
	})
	t.Run("Int64", func(t *testing.T) {
		t.Parallel()
		const data = int64(10)
		i, err := maputil.AsInteger(data)
		require.NoError(t, err)
		require.Equal(t, data, i)
	})
	t.Run("Uint", func(t *testing.T) {
		t.Parallel()
		const data = uint(10)
		i, err := maputil.AsInteger(data)
		require.NoError(t, err)
		require.Equal(t, int64(data), i)
	})
	t.Run("Uint8", func(t *testing.T) {
		t.Parallel()
		const data = uint8(10)
		i, err := maputil.AsInteger(data)
		require.NoError(t, err)
		require.Equal(t, int64(data), i)
	})
	t.Run("Uint16", func(t *testing.T) {
		t.Parallel()
		const data = uint16(10)
		i, err := maputil.AsInteger(data)
		require.NoError(t, err)
		require.Equal(t, int64(data), i)
	})
	t.Run("Uint32", func(t *testing.T) {
		t.Parallel()
		const data = uint32(10)
		i, err := maputil.AsInteger(data)
		require.NoError(t, err)
		require.Equal(t, int64(data), i)
	})
	t.Run("Uint64", func(t *testing.T) {
		t.Parallel()
		const data = uint64(10)
		i, err := maputil.AsInteger(data)
		require.NoError(t, err)
		require.Equal(t, int64(data), i)
	})
	t.Run("Float32", func(t *testing.T) {
		t.Parallel()
		t.Run("NoDecimal", func(t *testing.T) {
			t.Parallel()
			const data = float32(10)
			i, err := maputil.AsInteger(data)
			require.NoError(t, err)
			require.Equal(t, int64(data), i)
		})
		t.Run("Decimal", func(t *testing.T) {
			t.Parallel()
			const data = float32(10.5)
			i, err := maputil.AsInteger(data)
			require.EqualError(t, err, maputil.InvalidTypeError{
				Actual:   maputil.TypeNumber,
				Expected: []string{maputil.TypeInteger},
			}.Error())
			require.Zero(t, i)
		})
	})
	t.Run("Float64", func(t *testing.T) {
		t.Parallel()
		t.Run("NoDecimal", func(t *testing.T) {
			t.Parallel()
			const data = float64(10)
			i, err := maputil.AsInteger(data)
			require.NoError(t, err)
			require.Equal(t, int64(data), i)
		})
		t.Run("Decimal", func(t *testing.T) {
			t.Parallel()
			const data = float64(10.5)
			i, err := maputil.AsInteger(data)
			require.EqualError(t, err, maputil.InvalidTypeError{
				Actual:   maputil.TypeNumber,
				Expected: []string{maputil.TypeInteger},
			}.Error())
			require.Zero(t, i)
		})
	})
	t.Run("GenericNumber", func(t *testing.T) {
		t.Parallel()
		t.Run("NoDecimal", func(t *testing.T) {
			t.Parallel()
			var data interface{} = json.Number("10")
			i, err := maputil.AsInteger(data)
			require.NoError(t, err)
			require.Equal(t, int64(10), i)
		})
		t.Run("Decimal", func(t *testing.T) {
			t.Parallel()
			var data interface{} = json.Number("10.5")
			i, err := maputil.AsInteger(data)
			require.EqualError(t, err, maputil.InvalidTypeError{
				Actual:   maputil.TypeNumber,
				Expected: []string{maputil.TypeInteger},
			}.Error())
			require.Zero(t, i)
		})
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		var data interface{} = true
		i, err := maputil.AsInteger(data)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Actual:   maputil.TypeBoolean,
			Expected: []string{maputil.TypeInteger},
		}.Error())
		require.Zero(t, i)
	})
}

func TestAsNumber(t *testing.T) {
	t.Parallel()
	t.Run("Int", func(t *testing.T) {
		t.Parallel()
		const data = int(10)
		i, err := maputil.AsNumber(data)
		require.NoError(t, err)
		require.Equal(t, float64(data), i)
	})
	t.Run("Int8", func(t *testing.T) {
		t.Parallel()
		const data = int8(10)
		i, err := maputil.AsNumber(data)
		require.NoError(t, err)
		require.Equal(t, float64(data), i)
	})
	t.Run("Int16", func(t *testing.T) {
		t.Parallel()
		const data = int16(10)
		i, err := maputil.AsNumber(data)
		require.NoError(t, err)
		require.Equal(t, float64(data), i)
	})
	t.Run("Int32", func(t *testing.T) {
		t.Parallel()
		const data = int32(10)
		i, err := maputil.AsNumber(data)
		require.NoError(t, err)
		require.Equal(t, float64(data), i)
	})
	t.Run("Int64", func(t *testing.T) {
		t.Parallel()
		const data = int64(10)
		i, err := maputil.AsNumber(data)
		require.NoError(t, err)
		require.Equal(t, float64(data), i)
	})
	t.Run("Uint", func(t *testing.T) {
		t.Parallel()
		const data = uint(10)
		i, err := maputil.AsNumber(data)
		require.NoError(t, err)
		require.Equal(t, float64(data), i)
	})
	t.Run("Uint8", func(t *testing.T) {
		t.Parallel()
		const data = uint8(10)
		i, err := maputil.AsNumber(data)
		require.NoError(t, err)
		require.Equal(t, float64(data), i)
	})
	t.Run("Uint16", func(t *testing.T) {
		t.Parallel()
		const data = uint16(10)
		i, err := maputil.AsNumber(data)
		require.NoError(t, err)
		require.Equal(t, float64(data), i)
	})
	t.Run("Uint32", func(t *testing.T) {
		t.Parallel()
		const data = uint32(10)
		i, err := maputil.AsNumber(data)
		require.NoError(t, err)
		require.Equal(t, float64(data), i)
	})
	t.Run("Uint64", func(t *testing.T) {
		t.Parallel()
		const data = uint64(10)
		i, err := maputil.AsNumber(data)
		require.NoError(t, err)
		require.Equal(t, float64(data), i)
	})
	t.Run("Float32", func(t *testing.T) {
		t.Run("NoDecimal", func(t *testing.T) {
			t.Parallel()
			const data = float32(10)
			i, err := maputil.AsNumber(data)
			require.NoError(t, err)
			require.Equal(t, float64(data), i)
		})
		t.Run("Decimal", func(t *testing.T) {
			t.Parallel()
			const data = float32(10.5)
			i, err := maputil.AsNumber(data)
			require.NoError(t, err)
			require.Equal(t, float64(data), i)
		})
	})
	t.Run("Float64", func(t *testing.T) {
		t.Run("NoDecimal", func(t *testing.T) {
			t.Parallel()
			const data = float64(10)
			i, err := maputil.AsNumber(data)
			require.NoError(t, err)
			require.Equal(t, data, i)
		})
		t.Run("Decimal", func(t *testing.T) {
			t.Parallel()
			const data = float64(10.5)
			i, err := maputil.AsNumber(data)
			require.NoError(t, err)
			require.Equal(t, data, i)
		})
	})
	t.Run("GenericNumber", func(t *testing.T) {
		t.Run("NoDecimal", func(t *testing.T) {
			t.Parallel()
			var data interface{} = json.Number("10")
			i, err := maputil.AsNumber(data)
			require.NoError(t, err)
			require.Equal(t, float64(10), i)
		})
		t.Run("Decimal", func(t *testing.T) {
			t.Parallel()
			var data interface{} = json.Number("10.5")
			i, err := maputil.AsNumber(data)
			require.NoError(t, err)
			require.Equal(t, float64(10.5), i)
		})
		t.Run("BadValue", func(t *testing.T) {
			t.Parallel()
			var data interface{} = json.Number("invalid")
			i, err := maputil.AsNumber(data)
			require.Error(t, err)
			require.Zero(t, i)
		})
	})
	t.Run("BadType", func(t *testing.T) {
		t.Parallel()
		const data = true
		i, err := maputil.AsNumber(data)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Actual:   maputil.TypeBoolean,
			Expected: []string{maputil.TypeNumber},
		}.Error())
		require.Zero(t, i)
	})
}

func TestAsObject(t *testing.T) {
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		var data interface{} = map[string]interface{}{}
		m, err := maputil.AsObject(data)
		require.NoError(t, err)
		require.Equal(t, data, m)
	})
	t.Run("Bad", func(t *testing.T) {
		t.Parallel()
		var data interface{} = true
		o, err := maputil.AsObject(data)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Actual:   maputil.TypeBoolean,
			Expected: []string{maputil.TypeObject},
		}.Error())
		require.Nil(t, o)
	})
}

func TestAsString(t *testing.T) {
	t.Parallel()
	t.Run("Good", func(t *testing.T) {
		t.Parallel()
		var data interface{} = "hello"
		s, err := maputil.AsString(data)
		require.NoError(t, err)
		require.Equal(t, data, s)
	})
	t.Run("Bad", func(t *testing.T) {
		t.Parallel()
		var data interface{} = true
		s, err := maputil.AsString(data)
		require.EqualError(t, err, maputil.InvalidTypeError{
			Actual:   maputil.TypeBoolean,
			Expected: []string{maputil.TypeString},
		}.Error())
		require.Zero(t, s)
	})
}
