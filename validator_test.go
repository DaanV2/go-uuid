package uuid_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_IsValidStringUUID(t *testing.T) {
	test_case := func(s string, expected bool) {
		title := fmt.Sprintf("Test_IsValidStringUUID(%s)", s)
		t.Run(title, func(t *testing.T) {
			require.Equal(t, expected, uuid.IsValidStringUUID(s))
		})
	}

	test_case("00000000-0000-0000-0000-000000000000", true)

	test_case("", false)
	test_case("00000000-0000-0000-0000-00000000000", false)
	test_case("00000000-0000-0000-0000-0000000000000", false)
	test_case("X0000000-0000-0000-0000-000000000000", false)
}

func Fuzz_IsValidStringUUID(f *testing.F) {
	rnd := rand.New(rand.NewSource(0))

	for i := 0; i < 50; i++ {
		b := make([]byte, 16)
		_, _ = rnd.Read(b)
		u, _ := uuid.FromBytes(b)
		f.Add(u.String())
	}

	f.Fuzz(func(t *testing.T, s string) {
		result := uuid.IsValidStringUUID(s)
		require.True(t, result)
	})
}
