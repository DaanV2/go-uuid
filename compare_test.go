package uuid_test

import (
	"slices"
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_Equal(t *testing.T) {
	u := uuid.V4.New()
	require.True(t, u.Equal(u))
	require.False(t, u.Equal(uuid.V4.New()))
}

func Test_Compare(t *testing.T) {
	low := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	high := uuid.MustParse("00000000-0000-0000-0000-000000000002")

	require.Equal(t, 0, low.Compare(low))
	require.Equal(t, -1, low.Compare(high))
	require.Equal(t, 1, high.Compare(low))
}

func Test_Compare_Sort(t *testing.T) {
	a := uuid.MustParse("00000000-0000-0000-0000-000000000003")
	b := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	c := uuid.MustParse("00000000-0000-0000-0000-000000000002")

	list := []uuid.UUID{a, b, c}
	slices.SortFunc(list, uuid.UUID.Compare)

	require.Equal(t, []uuid.UUID{b, c, a}, list)
}

func Test_URN(t *testing.T) {
	u := uuid.MustParse("01234567-89ab-cdef-0123-456789abcdef")
	require.Equal(t, "urn:uuid:01234567-89ab-cdef-0123-456789abcdef", u.URN())
}

func Test_MustFromBytes(t *testing.T) {
	require.Panics(t, func() {
		uuid.MustFromBytes([]byte{0x00})
	})
	require.NotPanics(t, func() {
		uuid.MustFromBytes(make([]byte, 16))
	})
}