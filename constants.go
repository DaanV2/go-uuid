package uuid

import "unsafe"

const (
	TOTAL_BITS    = 128
	TOTAL_BYTES   = TOTAL_BITS / int(unsafe.Sizeof(uint8(0))*8)
	STRING_LENGTH = TOTAL_BYTES*2 + 4

	VERSION_BYTE_INDEX = 6
	VARIANT_BYTE_INDEX = 8
)
