package uuid_test

import (
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_V4_Version(t *testing.T) {
	require.Equal(t, uuid.Version4, uuid.V4.Version())
}

func Test_V4_Variant(t *testing.T) {
	require.Equal(t, uuid.Variant1, uuid.V4.Variant())
}

func Test_V4_New(t *testing.T) {
	u := uuid.V4.New()

	Expect_Valid_UUID(t, u, uuid.V4.Version(), uuid.V4.Variant())
}

func Test_V4_NewHex(t *testing.T) {
	u := uuid.V4.NewHex()

	require.Len(t, u, 32)
	require.NotRegexp(t, "[^a-zA-Z0-9]", u)
}

func Test_V4_NewString(t *testing.T) {
	u := uuid.V4.NewString()

	require.Len(t, u, 36)
	require.NotRegexp(t, "[^a-zA-Z0-9-]", u)
	err := uuid.ValidateStringUUID(u)
	require.NoError(t, err)
}

func Test_V4_NewBatch(t *testing.T) {
	sizes := []int{
		1 * uuid.TOTAL_BYTES,
		2 * uuid.TOTAL_BYTES,
		3 * uuid.TOTAL_BYTES,
		16 * uuid.TOTAL_BYTES,
		16 * uuid.TOTAL_BYTES - 1,
		1000_000,
	}

	for _, size := range sizes {
		t.Run("NewBatch", func(t *testing.T) {
			batch := uuid.V4.NewBatch(size)
			require.Len(t, batch, size)

			uni := make(map[string]bool)

			for _, u := range batch {
				Expect_Valid_UUID(t, u, uuid.V4.Version(), uuid.V4.Variant())

				if has, ok := uni[u.String()]; ok {
					require.False(t, has)
				} else {
					uni[u.String()] = true
				}
			}
		})
	}
}
