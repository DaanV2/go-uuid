package uuid_test

import (
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_Version_String(t *testing.T) {
	require.Equal(t, "v1", uuid.Version1.String())
	require.Equal(t, "v2", uuid.Version2.String())
	require.Equal(t, "v3", uuid.Version3.String())
	require.Equal(t, "v4", uuid.Version4.String())
	require.Equal(t, "v5", uuid.Version5.String())
	require.Equal(t, "v6", uuid.Version6.String())
	require.Equal(t, "v7", uuid.Version7.String())
	require.Equal(t, "v8", uuid.Version8.String())
	require.Equal(t, "unknown", uuid.Version(0).String())
}

func Test_Version_Value(t *testing.T) {
	require.Equal(t, 1, uuid.Version1.Value())
	require.Equal(t, 2, uuid.Version2.Value())
	require.Equal(t, 3, uuid.Version3.Value())
	require.Equal(t, 4, uuid.Version4.Value())
	require.Equal(t, 5, uuid.Version5.Value())
	require.Equal(t, 6, uuid.Version6.Value())
	require.Equal(t, 7, uuid.Version7.Value())
	require.Equal(t, 8, uuid.Version8.Value())
	require.Equal(t, 0, uuid.Version(0).Value())
}

func Test_VersionFromValue(t *testing.T) {
	test := func(value int, expected uuid.Version) {
		v, err := uuid.VersionFromValue(value)
		require.NoError(t, err)
		require.Equal(t, expected, v)
	}

	test(1, uuid.Version1)
	test(2, uuid.Version2)
	test(3, uuid.Version3)
	test(4, uuid.Version4)
	test(5, uuid.Version5)
	test(6, uuid.Version6)
	test(7, uuid.Version7)
	test(8, uuid.Version8)
}

func Test_VersionFromValue_Error(t *testing.T) {
	test := func(value int) {
		v, err := uuid.VersionFromValue(value)
		require.Error(t, err)
		require.Equal(t, uuid.Version1, v)
	}

	test(0)
	test(9)
}