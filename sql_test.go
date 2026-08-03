package uuid_test

import (
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_Value(t *testing.T) {
	u := uuid.V4.New()
	v, err := u.Value()
	require.NoError(t, err)
	require.Equal(t, u.String(), v)
}

func Test_Scan(t *testing.T) {
	u := uuid.V4.New()

	t.Run("string", func(t *testing.T) {
		var got uuid.UUID
		require.NoError(t, got.Scan(u.String()))
		require.Equal(t, u, got)
	})

	t.Run("text bytes", func(t *testing.T) {
		var got uuid.UUID
		require.NoError(t, got.Scan([]byte(u.String())))
		require.Equal(t, u, got)
	})

	t.Run("raw 16 bytes", func(t *testing.T) {
		var got uuid.UUID
		raw := u.Bytes()
		require.NoError(t, got.Scan(raw[:]))
		require.Equal(t, u, got)
	})

	t.Run("nil", func(t *testing.T) {
		var got uuid.UUID
		require.NoError(t, got.Scan(nil))
		require.Equal(t, uuid.Nil, got)
	})

	t.Run("unsupported", func(t *testing.T) {
		var got uuid.UUID
		require.Error(t, got.Scan(42))
	})
}