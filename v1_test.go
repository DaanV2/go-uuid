package uuid_test

import (
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_V1_Version(t *testing.T) {
	require.Equal(t, uuid.Version1, uuid.V1.Version())
}

func Test_V1_Variant(t *testing.T) {
	require.Equal(t, uuid.Variant1, uuid.V1.Variant())
}

func Test_V1_New(t *testing.T) {
	u, err := uuid.V1.New()
	require.NoError(t, err)

	Expect_Valid_UUID(t, u, uuid.V1.Version(), uuid.V1.Variant())
}
