package uuid_test

import (
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_V3_Version(t *testing.T) {
	require.Equal(t, uuid.Version3, uuid.V3.Version())
}

func Test_V3_Variant(t *testing.T) {
	require.Equal(t, uuid.Variant1, uuid.V3.Variant())
}

func Test_V3_New(t *testing.T) {
	u := uuid.V3.New([]byte("abcdefghijklmnopqrstuvwxyz"))

	Expect_Valid_UUID(t, u, uuid.V3.Version(), uuid.V3.Variant())
}

func Fuzz_V3_New(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		u := uuid.V3.New([]byte(s))
		Expect_Valid_UUID(t, u, uuid.V3.Version(), uuid.V3.Variant())
	})
}

func Test_V3_NewHex(t *testing.T) {
	u := uuid.V3.NewHex([]byte("abcdefghijklmnopqrstuvwxyz"))

	require.Len(t, u, 32)
	require.NotRegexp(t, "[^a-zA-Z0-9]", u)
}

func Test_V3_NewString(t *testing.T) {
	u := uuid.V3.NewString([]byte("abcdefghijklmnopqrstuvwxyz"))

	require.Len(t, u, 36)
	require.NotRegexp(t, "[^a-zA-Z0-9-]", u)
	err := uuid.ValidateStringUUID(u)
	require.NoError(t, err)
}

func Test_V3_NewBatch(t *testing.T) {
	sizes := []int{
		1,
		2,
		1 * uuid.TOTAL_BYTES,
		2 * uuid.TOTAL_BYTES,
		3 * uuid.TOTAL_BYTES,
		16 * uuid.TOTAL_BYTES,
		16 * uuid.TOTAL_BYTES - 1,
	}

	for _, size := range sizes {
		t.Run(fmt.Sprintf("NewBatch(%v)", size), func(t *testing.T) {
			brand := make([]byte, size)
			rand.Read(brand)

			batch := uuid.V3.NewBatch(make([]byte, size))
			require.Len(t, batch, size/uuid.TOTAL_BYTES)

			for _, u := range batch {
				Expect_Valid_UUID(t, u, uuid.V3.Version(), uuid.V3.Variant())
			}
		})
	}
}

func Fuzz_V3_NewBatch(f *testing.F) {
	f.Fuzz(func(t *testing.T, len int) {
		s := make([]byte, len)
		rand.Read(s)
		u := uuid.V3.NewBatch(s)
		require.Len(t, u, len/uuid.TOTAL_BYTES)
		for _, u := range u {
			Expect_Valid_UUID(t, u, uuid.V3.Version(), uuid.V3.Variant())
		}
	})
}

func Test_V3_NewBatchWithSize(t *testing.T) {
	sizes := []int{
		1,
		2,
		1 * uuid.TOTAL_BYTES,
		2 * uuid.TOTAL_BYTES,
		3 * uuid.TOTAL_BYTES,
		16 * uuid.TOTAL_BYTES,
		16 * uuid.TOTAL_BYTES - 1,
	}

	for _, size := range sizes {
		t.Run(fmt.Sprintf("NewBatchWithSize(%v)", size), func(t *testing.T) {
			brand := make([]byte, size)
			rand.Read(brand)

			batch := uuid.V3.NewBatchWithSize(brand, size/uuid.TOTAL_BYTES)
			require.Len(t, batch, size/uuid.TOTAL_BYTES)

			for _, u := range batch {
				Expect_Valid_UUID(t, u, uuid.V3.Version(), uuid.V3.Variant())
			}
		})
	}
}