package uuid

import "fmt"

// Compile-time checks that UUID implements the standard encoding interfaces.
var (
	_ interface{ MarshalText() ([]byte, error) }   = UUID{}
	_ interface{ UnmarshalText([]byte) error }     = (*UUID)(nil)
	_ interface{ MarshalBinary() ([]byte, error) } = UUID{}
	_ interface{ UnmarshalBinary([]byte) error }   = (*UUID)(nil)
	_ interface{ MarshalJSON() ([]byte, error) }   = UUID{}
	_ interface{ UnmarshalJSON([]byte) error }     = (*UUID)(nil)
)

// MarshalText implements encoding.TextMarshaler. It encodes the UUID in its
// canonical string form, which makes UUID work with encoding/xml, and any
// other package that relies on the interface.
func (u UUID) MarshalText() ([]byte, error) {
	return []byte(u.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler. It parses any of the
// textual representations accepted by Parse.
func (u *UUID) UnmarshalText(data []byte) error {
	parsed, err := ParseBytes(data)
	if err != nil {
		return err
	}
	*u = parsed
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler. It returns the raw 16
// bytes of the UUID.
func (u UUID) MarshalBinary() ([]byte, error) {
	b := make([]byte, TOTAL_BYTES)
	copy(b, u[:])
	return b, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler. It expects exactly 16
// bytes.
func (u *UUID) UnmarshalBinary(data []byte) error {
	if len(data) != TOTAL_BYTES {
		return fmt.Errorf("%w: %d", ErrInvalidLength, len(data))
	}
	copy(u[:], data)
	return nil
}

// MarshalJSON implements json.Marshaler. The UUID is encoded as a quoted
// canonical string.
func (u UUID) MarshalJSON() ([]byte, error) {
	// STRING_LENGTH digits plus two surrounding quotes.
	b := make([]byte, 0, STRING_LENGTH+2)
	b = append(b, '"')
	b = append(b, u.String()...)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements json.Unmarshaler. It accepts a quoted UUID string
// as well as a JSON null (decoded as the Nil UUID).
func (u *UUID) UnmarshalJSON(data []byte) error {
	// Treat null as the Nil UUID, matching common Go conventions.
	if string(data) == "null" {
		*u = Nil
		return nil
	}
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("%w: expected a JSON string", ErrInvalidFormat)
	}
	return u.UnmarshalText(data[1 : len(data)-1])
}