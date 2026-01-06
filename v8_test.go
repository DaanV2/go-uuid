package uuid_test

import (
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_V8_Version(t *testing.T) {
	require.Equal(t, uuid.Version8, uuid.V8.Version())
}

func Test_V8_Variant(t *testing.T) {
	require.Equal(t, uuid.Variant1, uuid.V8.Variant())
}

func Test_V8_New(t *testing.T) {
	u := uuid.V8.New()

	Expect_Valid_UUID(t, u, uuid.V8.Version(), uuid.V8.Variant())
}

func Test_V8_From(t *testing.T) {
	data := []byte("test-data-12345")
	u := uuid.V8.From(data)

	Expect_Valid_UUID(t, u, uuid.V8.Version(), uuid.V8.Variant())
}

func Test_V8_From_With_Short_Data(t *testing.T) {
	data := []byte("short")
	u := uuid.V8.From(data)

	Expect_Valid_UUID(t, u, uuid.V8.Version(), uuid.V8.Variant())
}

func Test_V8_From_With_Long_Data(t *testing.T) {
	data := []byte("this is a very long data string that exceeds 16 bytes")
	u := uuid.V8.From(data)

	Expect_Valid_UUID(t, u, uuid.V8.Version(), uuid.V8.Variant())
}

func Test_V8_NewHex(t *testing.T) {
	u := uuid.V8.NewHex()

	require.Len(t, u, 32)
	require.NotRegexp(t, "[^a-zA-Z0-9]", u)
}

func Test_V8_NewString(t *testing.T) {
	u := uuid.V8.NewString()

	require.Len(t, u, 36)
	require.NotRegexp(t, "[^a-zA-Z0-9-]", u)
	err := uuid.ValidateStringUUID(u)
	require.NoError(t, err)
}

func Test_V8_NewBatch(t *testing.T) {
	sizes := []int{1, 2, 3, 10, 100}

	for _, size := range sizes {
		t.Run("NewBatch", func(t *testing.T) {
			batch := uuid.V8.NewBatch(size)
			require.Len(t, batch, size)

			uni := make(map[string]bool)

			for _, u := range batch {
				Expect_Valid_UUID(t, u, uuid.V8.Version(), uuid.V8.Variant())

				if has, ok := uni[u.String()]; ok {
					require.False(t, has, "UUIDs should be unique")
				} else {
					uni[u.String()] = true
				}
			}
		})
	}
}
