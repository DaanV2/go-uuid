package uuid

import (
	"database/sql/driver"
	"fmt"
)

// Compile-time checks that UUID implements the database interfaces.
var (
	_ driver.Valuer                = UUID{}
	_ interface{ Scan(any) error } = (*UUID)(nil)
)

// Value implements driver.Valuer, allowing a UUID to be used directly as a SQL
// query argument. It is stored as its canonical string representation.
func (u UUID) Value() (driver.Value, error) {
	return u.String(), nil
}

// Scan implements sql.Scanner, allowing a UUID to be read from a database. It
// accepts a string, a []byte holding either the textual representation or the
// raw 16 bytes, and nil (decoded as the Nil UUID).
func (u *UUID) Scan(src any) error {
	switch v := src.(type) {
	case nil:
		*u = Nil
		return nil
	case string:
		parsed, err := Parse(v)
		if err != nil {
			return err
		}
		*u = parsed
		return nil
	case []byte:
		// A 16 byte slice is treated as the raw UUID; anything else is
		// interpreted as text (e.g. a driver returning the string form as bytes).
		if len(v) == TOTAL_BYTES {
			copy(u[:], v)
			return nil
		}
		parsed, err := ParseBytes(v)
		if err != nil {
			return err
		}
		*u = parsed
		return nil
	default:
		return fmt.Errorf("%w: cannot scan %T into UUID", ErrInvalidFormat, src)
	}
}