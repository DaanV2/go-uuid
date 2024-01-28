package uuid_test

import (
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	u := uuid.New()

	Expect_Valid_UUID(t, u, uuid.Version4, uuid.Variant1)
}

func Test_NewHex(t *testing.T) {
	u := uuid.NewHex()

	require.Len(t, u, 32)
	require.NotRegexp(t, "[^a-zA-Z0-9]", u)
}

func Test_NewString(t *testing.T) {
	u := uuid.NewString()

	err := uuid.ValidateStringUUID(u)
	require.NoError(t, err)
}
