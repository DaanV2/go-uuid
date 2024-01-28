package uuid

import "crypto/rand"

// V4 is a UUID version 4 generator
var V4 v4

type v4 struct{}

func (v4) fill(buf []byte) {
	_, _ = rand.Read(buf)
}

// New returns a new UUID version 4
func (c v4) New() UUID {
	var buf [16]byte
	c.fill(buf[:])

	return createUUID(buf, Version4, Variant1)
}

// NewHex returns a new UUID version 4 as a hex string
func (v4) NewHex() string {
	return V4.New().StringHex()
}

// NewString returns a new UUID version 4 as a string
func (v4) NewString() string {
	return V4.New().String()
}

// NewBatch returns a batch of UUID version 4
func (c v4) NewBatch(n int) []UUID {
	buf := make([]byte, 16*n)
	c.fill(buf)

	return createBatchUUID(buf, Version4, Variant1)
}