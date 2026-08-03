package uuid

// Equal reports whether two UUIDs are identical.
func (u UUID) Equal(other UUID) bool {
	return u == other
}

// Compare returns an integer comparing two UUIDs lexicographically by their
// bytes. The result is 0 if u == other, -1 if u < other, and +1 if u > other.
// This makes UUIDs easy to sort, for example with slices.SortFunc.
func (u UUID) Compare(other UUID) int {
	for i := 0; i < TOTAL_BYTES; i++ {
		switch {
		case u[i] < other[i]:
			return -1
		case u[i] > other[i]:
			return 1
		}
	}
	return 0
}