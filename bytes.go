package uuid

import "errors"

// FromBytes returns a UUID from a byte slice.
func FromBytes(data []byte) (UUID, error) {
	if len(data) != 16 {
		return Empty(), errors.New("invalid byte slice length")
	}

	var result UUID
	_ = copy(result[:], data[:])
	return result, nil
}

// MustFromBytes is like FromBytes but panics if the slice is not 16 bytes long.
// It is convenient for tests and package-level variable initialisation.
func MustFromBytes(data []byte) UUID {
	u, err := FromBytes(data)
	if err != nil {
		panic(err)
	}
	return u
}
