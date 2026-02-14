package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"time"
)

var V7 v7

type v7 struct{}

// Version returns the current version of this generator
func (v7) Version() Version {
	return Version7
}

// Variant returns the current variant of this generator
func (v7) Variant() Variant {
	return Variant1
}

// New returns a new UUID version 7 based on Unix timestamp with random data
func (v7) New() UUID {
	return V7.From(time.Now().UTC())
}

// From returns a new UUID version 7 from the given time
func (v7) From(t time.Time) UUID {
	// UUIDv7 Field and Bit Layout:
	// unix_ts_ms (48 bits) - Unix timestamp in milliseconds
	// ver (4 bits) - version field
	// rand_a (12 bits) - random data
	// var (2 bits) - variant field
	// rand_b (62 bits) - random data

	var data [16]byte

	// Get Unix timestamp in milliseconds (48 bits)
	unixMs := uint64(t.UnixMilli())

	// Place timestamp in first 48 bits (6 bytes)
	binary.BigEndian.PutUint64(data[0:], unixMs<<16)

	// Fill remaining bytes with random data
	_, _ = rand.Read(data[6:])

	return createUUID(data, V7.Version(), V7.Variant())
}

// NewHex returns a new UUID version 7 as a hex string
func (v7) NewHex() string {
	return V7.New().StringHex()
}

// NewString returns a new UUID version 7 as a string
func (v7) NewString() string {
	return V7.New().String()
}

// NewBatch returns a batch of UUID version 7
func (v7) NewBatch(n int) []UUID {
	uuids := make([]UUID, n)
	for i := range n {
		uuids[i] = V7.New()
	}
	return uuids
}
