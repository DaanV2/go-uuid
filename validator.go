package uuid

import (
	"fmt"
)

// IsValidStringUUID returns true if the string is a valid UUID.
func IsValidStringUUID(s string) bool {
	return ValidateStringUUID(s) == nil
}

// ValidateStringUUID returns an error if the string is not a valid UUID.
func ValidateStringUUID(s string) error {
	if len(s) != STRING_LENGTH {
		return fmt.Errorf("invalid length %d, expected %d", len(s), STRING_LENGTH)
	}

	for i, c := range s {
		switch {
		case c >= '0' && c <= '9':
		case c >= 'a' && c <= 'f':
		case c >= 'A' && c <= 'F':
		case c == '-':
		default:
			return fmt.Errorf("invalid character '%c' at position %d", c, i)
		}
	}

	return nil
}
