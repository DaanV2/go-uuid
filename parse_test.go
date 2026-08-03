package uuid_test

import (
	"math/rand"
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_Parse(t *testing.T) {
	expected := uuid.MustFromBytes([]byte{
		0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
	})

	cases := []string{
		"01234567-89ab-cdef-0123-456789abcdef",          // canonical
		"01234567-89AB-CDEF-0123-456789ABCDEF",          // upper case
		"{01234567-89ab-cdef-0123-456789abcdef}",        // braced
		"urn:uuid:01234567-89ab-cdef-0123-456789abcdef", // urn
		"URN:UUID:01234567-89ab-cdef-0123-456789abcdef", // urn upper prefix
		"0123456789abcdef0123456789abcdef",              // no dashes
	}

	for _, c := range cases {
		t.Run(c, func(t *testing.T) {
			u, err := uuid.Parse(c)
			require.NoError(t, err)
			require.Equal(t, expected, u)
		})
	}
}

func Test_Parse_Invalid(t *testing.T) {
	cases := []string{
		"",
		"not-a-uuid",
		"01234567-89ab-cdef-0123-456789abcde",           // too short
		"01234567-89ab-cdef-0123-456789abcdeff",         // too long
		"01234567x89ab-cdef-0123-456789abcdef",          // wrong separator
		"0123456789abcdefg123456789abcdef",              // invalid hex (32 len)
		"{01234567-89ab-cdef-0123-456789abcdef]",        // mismatched braces
		"xxx:uuid:01234567-89ab-cdef-0123-456789abcdef", // bad urn prefix
	}

	for _, c := range cases {
		t.Run(c, func(t *testing.T) {
			_, err := uuid.Parse(c)
			require.Error(t, err)
		})
	}
}

func Test_MustParse(t *testing.T) {
	require.NotPanics(t, func() {
		uuid.MustParse("01234567-89ab-cdef-0123-456789abcdef")
	})
	require.Panics(t, func() {
		uuid.MustParse("nope")
	})
}

func Test_ParseBytes(t *testing.T) {
	u, err := uuid.ParseBytes([]byte("01234567-89ab-cdef-0123-456789abcdef"))
	require.NoError(t, err)
	require.Equal(t, "01234567-89ab-cdef-0123-456789abcdef", u.String())
}

// Fuzz_Parse_Roundtrip verifies that every generated UUID survives a
// String -> Parse round trip, in all accepted formats.
func Fuzz_Parse_Roundtrip(f *testing.F) {
	rnd := rand.New(rand.NewSource(0))

	for range 50 {
		b := make([]byte, 16)
		_, _ = rnd.Read(b)
		u := uuid.MustFromBytes(b)
		f.Add(u.String())
	}

	f.Fuzz(func(t *testing.T, s string) {
		u, err := uuid.Parse(s)
		if err != nil {
			return
		}

		// Every accepted form must round trip back to the same UUID.
		for _, form := range []string{u.String(), u.StringHex(), u.URN()} {
			got, err := uuid.Parse(form)
			require.NoError(t, err, "form %q", form)
			require.Equal(t, u, got, "form %q", form)
		}
	})
}