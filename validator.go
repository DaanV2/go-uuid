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
		if i == 8 || i == 13 || i == 18 || i == 23 {
			if c != '-' {
				return fmt.Errorf("invalid character '%c' at position %d, expected '-'", c, i)
			}

			continue
		}

		switch {
		case c >= '0' && c <= '9':
		case c >= 'a' && c <= 'f':
		case c >= 'A' && c <= 'F':
		default:
			return fmt.Errorf("invalid character '%c' at position %d", c, i)
		}
	}

	return nil
}
