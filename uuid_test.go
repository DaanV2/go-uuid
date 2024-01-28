package uuid_test

import (
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_UUID_Empty(t *testing.T) {
	u := uuid.Empty()

	require.True(t, uuid.IsValidStringUUID(u.String()))
}

func Test_UUID_IsEmpty(t *testing.T) {
	u := uuid.Empty()

	require.True(t, u.IsEmpty())
}

func Test_UUID_Bytes(t *testing.T) {
	u := uuid.Empty()

	require.Equal(t, [16]byte{}, u.Bytes())
}
