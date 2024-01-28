package uuid

import "crypto/md5"

// V3 is a UUID version 3 generator, using MD5 hashing
var V3 v3

type v3 struct{}

// Version returns the current version of this generator
func (v3) Version() Version {
	return Version3
}

// Variant returns the current variant of this generator
func (v3) Variant() Variant {
	return Variant1
}

func (v3) New(toHash []byte) UUID {
	data := md5.Sum(toHash)

	return createUUID(data, V3.Version(), V3.Variant())
}

// NewHex returns a new UUID version 4 as a hex string
func (v3) NewHex(toHash []byte) string {
	return V3.New(toHash).StringHex()
}

// NewString returns a new UUID version 4 as a string
func (v3) NewString(toHash []byte) string {
	return V3.New(toHash).String()
}

// NewBatch returns a batch of UUID version 4
func (v3) NewBatch(toHash []byte) []UUID {
	n := len(toHash) / md5.Size

	return V3.NewBatchWithSize(toHash, n)
}

// NewBatchWithSize returns a batch of UUID version 4 with a specific size
func (v3) NewBatchWithSize(toHash []byte, n int) []UUID {
	max := len(toHash) / md5.Size
	if n > max {
		n = max
	}

	data := make([]byte, 0, n*md5.Size)
	for i := 0; i < n; i++ {
		s := i * md5.Size
		e := s + md5.Size
		d := md5.Sum(toHash[s:e])
		data = append(data, d[:]...)
	}

	return createBatchUUID(data, V3.Version(), V3.Variant())
}
