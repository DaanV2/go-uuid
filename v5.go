package uuid

import (
	"crypto/sha1"
)

// V5 is a UUID version 4 generator, using SHA-1 hashing
var V5 v5

type v5 struct{}

// Version returns the current version of this generator
func (v5) Version() Version {
	return Version5
}

// Variant returns the current variant of this generator
func (v5) Variant() Variant {
	return Variant1
}

func (v5) New(toHash []byte) UUID {
	data := sha1.Sum(toHash)
	// 20 bytes to 16 bytes
	var data16 [16]byte
	copy(data16[:], data[:])

	return createUUID(data16, V5.Version(), V5.Variant())
}

// NewHex returns a new UUID version 4 as a hex string
func (v5) NewHex(toHash []byte) string {
	return V5.New(toHash).StringHex()
}

// NewString returns a new UUID version 4 as a string
func (v5) NewString(toHash []byte) string {
	return V5.New(toHash).String()
}

// NewBatch returns a batch of UUID version 4
func (v5) NewBatch(toHash []byte) []UUID {
	n := len(toHash) / sha1.Size

	return V5.NewBatchWithSize(toHash, n)
}

// NewBatchWithSize returns a batch of UUID version 4 with a specific size
func (v5) NewBatchWithSize(toHash []byte, n int) []UUID {
	max := len(toHash) / sha1.Size
	if n > max {
		n = max
	}

	data := make([]byte, 0, n*sha1.Size)
	for i := 0; i < n; i++ {
		s := i * sha1.Size
		e := s + sha1.Size
		d := sha1.Sum(toHash[s:e])
		data = append(data, d[:]...)
	}

	return createBatchUUID(data, V5.Version(), V5.Variant())
}
