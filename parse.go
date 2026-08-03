package uuid

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidLength is returned when the input has an unexpected length.
	ErrInvalidLength = errors.New("invalid UUID length")
	// ErrInvalidFormat is returned when the input is not a recognized UUID format.
	ErrInvalidFormat = errors.New("invalid UUID format")
)

// dashPositions are the indexes of the '-' separators in the canonical form
// 00000000-0000-0000-0000-000000000000.
var dashPositions = [4]int{8, 13, 18, 23}

// Parse decodes a string into a UUID. It accepts the most common textual
// representations:
//
//	00000000-0000-0000-0000-000000000000        (canonical, 36 chars)
//	{00000000-0000-0000-0000-000000000000}      (braced, 38 chars)
//	urn:uuid:00000000-0000-0000-0000-000000000000 (URN, 45 chars)
//	00000000000000000000000000000000            (hex, no dashes, 32 chars)
//
// Hexadecimal digits are case-insensitive.
func Parse(s string) (UUID, error) {
	switch len(s) {
	case STRING_LENGTH: // canonical, e.g. 00000000-0000-0000-0000-000000000000
		return parseCanonical(s)
	case STRING_LENGTH + 9: // urn:uuid: prefix
		if !equalFoldASCII(s[:9], "urn:uuid:") {
			return Nil, fmt.Errorf("%w: expected urn:uuid: prefix", ErrInvalidFormat)
		}
		return parseCanonical(s[9:])
	case STRING_LENGTH + 2: // {...}
		if s[0] != '{' || s[len(s)-1] != '}' {
			return Nil, fmt.Errorf("%w: expected surrounding braces", ErrInvalidFormat)
		}
		return parseCanonical(s[1 : len(s)-1])
	case TOTAL_BYTES * 2: // hex, no dashes
		return parseHex(s)
	default:
		return Nil, fmt.Errorf("%w: %d", ErrInvalidLength, len(s))
	}
}

// MustParse is like Parse but panics if the string cannot be parsed.
// It simplifies safe initialisation of package-level UUID variables.
func MustParse(s string) UUID {
	u, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}

// ParseBytes is like Parse but accepts a byte slice holding the textual
// representation instead of a string.
func ParseBytes(b []byte) (UUID, error) {
	return Parse(string(b))
}

// parseCanonical decodes the 36 character canonical form.
func parseCanonical(s string) (UUID, error) {
	for _, p := range dashPositions {
		if s[p] != '-' {
			return Nil, fmt.Errorf("%w: expected '-' at position %d", ErrInvalidFormat, p)
		}
	}

	var u UUID
	// src indexes into s, skipping the four separators.
	for i, src := 0, 0; i < TOTAL_BYTES; i++ {
		if src == 8 || src == 13 || src == 18 || src == 23 {
			src++
		}

		hi, ok := fromHexChar(s[src])
		if !ok {
			return Nil, fmt.Errorf("%w: invalid character '%c' at position %d", ErrInvalidFormat, s[src], src)
		}
		lo, ok := fromHexChar(s[src+1])
		if !ok {
			return Nil, fmt.Errorf("%w: invalid character '%c' at position %d", ErrInvalidFormat, s[src+1], src+1)
		}

		u[i] = hi<<4 | lo
		src += 2
	}

	return u, nil
}

// parseHex decodes the 32 character dash-less form.
func parseHex(s string) (UUID, error) {
	var u UUID
	for i := 0; i < TOTAL_BYTES; i++ {
		hi, ok := fromHexChar(s[i*2])
		if !ok {
			return Nil, fmt.Errorf("%w: invalid character '%c' at position %d", ErrInvalidFormat, s[i*2], i*2)
		}
		lo, ok := fromHexChar(s[i*2+1])
		if !ok {
			return Nil, fmt.Errorf("%w: invalid character '%c' at position %d", ErrInvalidFormat, s[i*2+1], i*2+1)
		}
		u[i] = hi<<4 | lo
	}
	return u, nil
}

// fromHexChar converts a single hex character to its value.
// The second return value reports whether the character was valid.
func fromHexChar(c byte) (byte, bool) {
	switch {
	case c >= '0' && c <= '9':
		return c - '0', true
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10, true
	case c >= 'A' && c <= 'F':
		return c - 'A' + 10, true
	default:
		return 0, false
	}
}

// equalFoldASCII reports whether s and lower are equal, case-insensitively.
// lower is assumed to already be lowercase ASCII.
func equalFoldASCII(s, lower string) bool {
	if len(s) != len(lower) {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		if c != lower[i] {
			return false
		}
	}
	return true
}