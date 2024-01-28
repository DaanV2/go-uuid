package util_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/DaanV2/go-uuid/util"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	value uint64
	count uint
}

func Test_PopCount(t *testing.T) {
	cases := []testCase{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{0b1111_1111, 8},
		{0b0110_1001, 4},
	}

	t.Run("uin64", PopCount_Test[uint64](cases))
	t.Run("uint32", PopCount_Test[uint32](cases))
	t.Run("uint16", PopCount_Test[uint16](cases))
	t.Run("uint8", PopCount_Test[uint8](cases))

	t.Run("int64", PopCount_Test[int64](cases))
	t.Run("int32", PopCount_Test[int32](cases))
	t.Run("int16", PopCount_Test[int16](cases))
	t.Run("int8", PopCount_Test[int8](cases))
}

func PopCount_Test[T util.Number](values []testCase) func(t *testing.T) {
	return func(t *testing.T) {
		for _, v := range values {
			t.Run(fmt.Sprintf("PopCount(%v) => %v", v.value, v.count), func(t *testing.T) {
				c := util.PopCount(v.value)

				require.Equal(t, v.count, c)
			})
		}
	}
}

func Fuzz_PopCount(f *testing.F) {
	rnd := rand.New(rand.NewSource(0))

	for i := 0; i < 50; i++ {
		f.Add(rnd.Uint64())
	}

	f.Fuzz(func(t *testing.T, value uint64) {
		c := util.PopCount(value)

		require.Equal(t, popCount(value), c)
	})
}

func popCount(value uint64) uint {
	v := value
	c := uint(0)

	for ; v > 0; v >>= 1 {
		d := v & 1
		c += uint(d)
	}

	return c
}
