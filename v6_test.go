package uuid_test

import (
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_V6_Version(t *testing.T) {
	require.Equal(t, uuid.Version6, uuid.V6.Version())
}

func Test_V6_Variant(t *testing.T) {
	require.Equal(t, uuid.Variant1, uuid.V6.Variant())
}

func Test_V6_New(t *testing.T) {
	u, err := uuid.V6.New()
	require.NoError(t, err)

	Expect_Valid_UUID(t, u, uuid.V6.Version(), uuid.V6.Variant())
}

func Test_V6_Multiple_UUIDs_Are_Ordered(t *testing.T) {
	// Generate multiple UUIDs and verify they are time-ordered
	uuids := make([]uuid.UUID, 10)
	for i := range 10 {
		u, err := uuid.V6.New()
		require.NoError(t, err)
		uuids[i] = u
	}

	// Verify UUIDs are ordered (later UUIDs should be greater than earlier ones)
	for i := 1; i < len(uuids); i++ {
		// Compare the first 8 bytes which contain the timestamp
		current := uuids[i].String()
		previous := uuids[i-1].String()
		require.GreaterOrEqual(t, current, previous, "UUIDs should be time-ordered")
	}
}
