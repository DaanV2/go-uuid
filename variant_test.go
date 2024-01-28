package uuid_test

import (
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_Variant_String(t *testing.T) {
	require.Equal(t, "v0", uuid.Variant0.String())
	require.Equal(t, "v1", uuid.Variant1.String())
	require.Equal(t, "v2", uuid.Variant2.String())
	require.Equal(t, "v3", uuid.Variant3.String())
	require.Equal(t, "unknown", uuid.Variant(4).String())
}

func Test_Variant_Value(t *testing.T) {
	require.Equal(t, 0, uuid.Variant0.Value())
	require.Equal(t, 1, uuid.Variant1.Value())
	require.Equal(t, 2, uuid.Variant2.Value())
	require.Equal(t, 3, uuid.Variant3.Value())
	require.Equal(t, 0, uuid.Variant(4).Value())
	require.Equal(t, 0, uuid.Variant(0).Value())
}

func Test_VariantFromValue(t *testing.T) {
	test := func(value int, expected uuid.Variant) {
		v, err := uuid.VariantFromValue(value)
		require.NoError(t, err)
		require.Equal(t, expected, v)
	}

	test(0, uuid.Variant0)
	test(1, uuid.Variant1)
	test(2, uuid.Variant2)
	test(3, uuid.Variant3)
}

func Test_VariantFromValue_Error(t *testing.T) {
	test := func(value int) {
		v, err := uuid.VariantFromValue(value)
		require.Error(t, err)
		require.Equal(t, uuid.Variant0, v)
	}

	test(4)
	test(5)
}