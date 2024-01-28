package uuid_test

import (
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_V5_Version(t *testing.T) {
	require.Equal(t, uuid.Version5, uuid.V5.Version())
}

func Test_V5_Variant(t *testing.T) {
	require.Equal(t, uuid.Variant1, uuid.V5.Variant())
}

func Test_V5_New(t *testing.T) {
	u := uuid.V5.New([]byte("abcdefghijklmnopqrstuvwxyz"))

	Expect_Valid_UUID(t, u, uuid.V5.Version(), uuid.V5.Variant())
}

func Fuzz_V5_New(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		u := uuid.V5.New([]byte(s))
		Expect_Valid_UUID(t, u, uuid.V5.Version(), uuid.V5.Variant())
	})
}

func Test_V5_NewHex(t *testing.T) {
	u := uuid.V5.NewHex([]byte("abcdefghijklmnopqrstuvwxyz"))

	require.Len(t, u, 32)
	require.NotRegexp(t, "[^a-zA-Z0-9]", u)
}

func Test_V5_NewString(t *testing.T) {
	u := uuid.V5.NewString([]byte("abcdefghijklmnopqrstuvwxyz"))

	require.Len(t, u, 36)
	require.NotRegexp(t, "[^a-zA-Z0-9-]", u)
	err := uuid.ValidateStringUUID(u)
	require.NoError(t, err)
}

func Test_V5_NewBatch(t *testing.T) {
	sizes := []int{
		1,
		2,
		1 * uuid.TOTAL_BYTES,
		2 * uuid.TOTAL_BYTES,
		3 * uuid.TOTAL_BYTES,
		16 * uuid.TOTAL_BYTES,
		16*uuid.TOTAL_BYTES - 1,
	}

	for _, size := range sizes {
		t.Run(fmt.Sprintf("NewBatch(%v)", size), func(t *testing.T) {
			brand := make([]byte, size)
			rand.Read(brand)

			batch := uuid.V5.NewBatch(brand)

			for _, u := range batch {
				Expect_Valid_UUID(t, u, uuid.V5.Version(), uuid.V5.Variant())
			}
		})
	}
}

func Fuzz_V5_NewBatch(f *testing.F) {
	f.Fuzz(func(t *testing.T, len int) {
		s := make([]byte, len)
		rand.Read(s)
		u := uuid.V5.NewBatch(s)
		require.Len(t, u, len/uuid.TOTAL_BYTES)
		for _, u := range u {
			Expect_Valid_UUID(t, u, uuid.V5.Version(), uuid.V5.Variant())
		}
	})
}

func Test_V5_NewBatchWithSize(t *testing.T) {
	sizes := []int{
		1,
		2,
		1 * uuid.TOTAL_BYTES,
		2 * uuid.TOTAL_BYTES,
		3 * uuid.TOTAL_BYTES,
		16 * uuid.TOTAL_BYTES,
		16*uuid.TOTAL_BYTES - 1,
	}

	for _, size := range sizes {
		t.Run(fmt.Sprintf("NewBatchWithSize(%v)", size), func(t *testing.T) {
			brand := make([]byte, size)
			rand.Read(brand)

			batch := uuid.V5.NewBatchWithSize(make([]byte, size), size/20)

			for _, u := range batch {
				Expect_Valid_UUID(t, u, uuid.V5.Version(), uuid.V5.Variant())
			}
		})
	}
}
