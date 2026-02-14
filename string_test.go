package uuid_test

import (
	"encoding/hex"
	"math/rand"
	"strings"
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Fuzz_String(f *testing.F) {
	rnd := rand.New(rand.NewSource(0))

	gen_case := func(length int) {
		b := make([]byte, length)
		_, _ = rnd.Read(b)
		f.Add(b)
	}

	gen_case(0)
	gen_case(15)
	gen_case(17)

	for range 25 {
		gen_case(16)
	}

	for range 25 {
		length := (rnd.Int() % 8) + 12
		gen_case(length)
	}

	f.Fuzz(func(t *testing.T, b []byte) {
		u, err := uuid.FromBytes(b[:])

		if len(b) != 16 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)

		actual := u.String()
		require.Equal(t, len(u), uuid.TOTAL_BYTES)
		require.Equal(t, len(actual), uuid.STRING_LENGTH, "Length of the string should be doubled")

		expect := hex.EncodeToString(b)
		actual_hex := strings.ReplaceAll(actual, "-", "")
		require.Equal(t, expect, actual_hex, "expected %s, got %s", expect, actual)
	})
}

func Fuzz_StringHex(f *testing.F) {
	rnd := rand.New(rand.NewSource(0))

	gen_case := func(length int) {
		b := make([]byte, length)
		_, _ = rnd.Read(b)
		f.Add(b)
	}

	gen_case(0)
	gen_case(15)
	gen_case(17)

	for range 25 {
		gen_case(16)
	}

	for range 25 {
		length := (rnd.Int() % 8) + 12
		gen_case(length)
	}

	f.Fuzz(func(t *testing.T, b []byte) {
		u, err := uuid.FromBytes(b[:])

		if len(b) != 16 {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)

		actual := u.StringHex()
		require.Equal(t, len(u), uuid.TOTAL_BYTES)
		require.Equal(t, len(actual), uuid.TOTAL_BYTES*2, "Length of the string should be doubled")

		expect := hex.EncodeToString(b)
		require.Equal(t, expect, actual, "expected %s, got %s", expect, actual)
	})
}

func Benchmark_StringHex(b *testing.B) {
	rnd := rand.New(rand.NewSource(0))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b := make([]byte, 16)
		_, _ = rnd.Read(b)
		u, _ := uuid.FromBytes(b[:])
		_ = u.StringHex()
	}
}
