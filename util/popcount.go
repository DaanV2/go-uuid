package util

type Number interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

// PopCount returns the number of bits set to 1 in the given value.
func PopCount[T Number](value T) uint {
	v := value
	c := uint(0)

	for ; v > 0; v >>= 1 {
		d := v & 1
		c += uint(d)
	}

	return c
}
