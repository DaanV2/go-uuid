package uuid

func FromBytes(b [16]byte) UUID {
	var result UUID
	_ = copy(result[:], b[:])
	return result
}