package uuid

// Nil is the nil UUID (00000000-0000-0000-0000-000000000000).
var Nil = Empty()

// UUID is a 128 bit (16 byte) Universal Unique IDentifier as defined in RFC-4122.
type UUID [16]byte

// Empty returns an empty UUID (00000000-0000-0000-0000-000000000000).
func Empty() UUID {
	return UUID{}
}

// IsEmpty returns true if the UUID is empty (00000000-0000-0000-0000-000000000000).
func (u UUID) IsEmpty() bool {
	return u == Empty()
}

// Bytes returns the UUID as a copied byte slice.
func (u UUID) Bytes() [16]byte {
	var b [16]byte
	copy(b[:], u[:])
	return b
}

// Version returns the version of the UUID.
func (u UUID) Version() Version {
	return getVersion(u)
}

// Variant returns the variant of the UUID.
func (u UUID) Variant() Variant {
	return getVariant(u)
}