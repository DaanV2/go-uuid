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
