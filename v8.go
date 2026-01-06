package uuid

import "crypto/rand"

var V8 v8

type v8 struct{}

// Version returns the current version of this generator
func (v8) Version() Version {
	return Version8
}

// Variant returns the current variant of this generator
func (v8) Variant() Variant {
	return Variant1
}

// New returns a new UUID version 8 with random data
// UUID v8 provides an RFC-compatible format for experimental or vendor-specific use cases
func (v8) New() UUID {
	var data [16]byte
	_, _ = rand.Read(data[:])

	return createUUID(data, V8.Version(), V8.Variant())
}

// From creates a UUID version 8 from custom data
// The data should be exactly 16 bytes. If it's shorter, it will be padded with zeros.
// If it's longer, only the first 16 bytes will be used.
func (v8) From(data []byte) UUID {
	var buf [16]byte
	copy(buf[:], data)

	return createUUID(buf, V8.Version(), V8.Variant())
}

// NewHex returns a new UUID version 8 as a hex string
func (v8) NewHex() string {
	return V8.New().StringHex()
}

// NewString returns a new UUID version 8 as a string
func (v8) NewString() string {
	return V8.New().String()
}

// NewBatch returns a batch of UUID version 8
func (v8) NewBatch(n int) []UUID {
	buf := make([]byte, 16*n)
	_, _ = rand.Read(buf)

	return createBatchUUID(buf, V8.Version(), V8.Variant())
}
