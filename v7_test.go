package uuid_test

import (
	"testing"
	"time"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_V7_Version(t *testing.T) {
	require.Equal(t, uuid.Version7, uuid.V7.Version())
}

func Test_V7_Variant(t *testing.T) {
	require.Equal(t, uuid.Variant1, uuid.V7.Variant())
}

func Test_V7_New(t *testing.T) {
	u := uuid.V7.New()

	Expect_Valid_UUID(t, u, uuid.V7.Version(), uuid.V7.Variant())
}

func Test_V7_NewHex(t *testing.T) {
	u := uuid.V7.NewHex()

	require.Len(t, u, 32)
	require.NotRegexp(t, "[^a-zA-Z0-9]", u)
}

func Test_V7_NewString(t *testing.T) {
	u := uuid.V7.NewString()

	require.Len(t, u, 36)
	require.NotRegexp(t, "[^a-zA-Z0-9-]", u)
	err := uuid.ValidateStringUUID(u)
	require.NoError(t, err)
}

func Test_V7_NewBatch(t *testing.T) {
	sizes := []int{1, 2, 3, 10, 100}

	for _, size := range sizes {
		t.Run("NewBatch", func(t *testing.T) {
			batch := uuid.V7.NewBatch(size)
			require.Len(t, batch, size)

			uni := make(map[string]bool)

			for _, u := range batch {
				Expect_Valid_UUID(t, u, uuid.V7.Version(), uuid.V7.Variant())

				if has, ok := uni[u.String()]; ok {
					require.False(t, has, "UUIDs should be unique")
				} else {
					uni[u.String()] = true
				}
			}
		})
	}
}

func Test_V7_Time_Ordering(t *testing.T) {
	// Create UUIDs with known timestamps
	now := time.Now().UTC()
	u1 := uuid.V7.From(now)
	u2 := uuid.V7.From(now.Add(time.Second))

	// The string representation should be ordered
	require.Less(t, u1.String(), u2.String(), "Later timestamp should produce greater UUID")
}

func Test_V7_Multiple_UUIDs_Are_Ordered(t *testing.T) {
	// Generate multiple UUIDs with small delays
	uuids := make([]uuid.UUID, 10)
	for i := 0; i < 10; i++ {
		uuids[i] = uuid.V7.New()
		time.Sleep(time.Millisecond) // Ensure different timestamps
	}

	// Verify UUIDs are ordered (later UUIDs should be greater than or equal to earlier ones)
	for i := 1; i < len(uuids); i++ {
		current := uuids[i].String()
		previous := uuids[i-1].String()
		require.GreaterOrEqual(t, current, previous, "UUIDs should be time-ordered")
	}
}
