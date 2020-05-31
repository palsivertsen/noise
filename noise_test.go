package noise_test

import (
	"testing"

	noise "github.com/palsivertsen/noise"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSmooth_Next(t *testing.T) {
	assert.PanicsWithValue(t, "Size (0) needs to be larger than 0", func() {
		var unit noise.Smooth
		unit.Next()
	}, "zero size not allowed")

	{
		unit := noise.Smooth{
			Size: 1000,
		}

		var a []float64
		for i := 0; i < unit.Size; i++ {
			a = append(a, unit.Next())
		}

		var b []float64
		for i := 0; i < unit.Size; i++ {
			b = append(b, unit.Next())
		}

		assert.Equal(t, len(a), len(b), "Size should not change between calls to Next()")
		assert.NotEqual(t, a[len(a)-1], b[0], "Do not re-use values when creating new pools")
		assert.NotEqual(t, a, b, "New pool should be different from previous pool")
	}

	{
		unit := noise.Smooth{
			Size: 1000,
		}

		zeroCounter := 0

		for i := 0; i < unit.Size; i++ {
			v := unit.Next()
			if v == 0 {
				zeroCounter++
				require.Less(t, zeroCounter, 2, "Too many zero values in generated pool")
			}
			require.Less(t, v, 1.0, "Generated value must be less than 1")
			require.GreaterOrEqual(t, v, 0.0, "Generated value must be greater or equal to 0")
		}
	}
}
