package uuid

import (
	"errors"

	"github.com/DaanV2/go-uuid/util"
)

// Variant represents the variant of the UUID.
type Variant uint

const (
	// Variant0 is reserved
	Variant0 Variant = 0b0000_0000
	// Variant1 is the RFC 4122 variant
	Variant1 Variant = 0b1000_0000
	// Variant2 is reserved
	Variant2 Variant = 0b1100_0000
	// Variant3 is future reserved
	Variant3 Variant = 0b1110_0000
)

var (
	// ErrInvalidVariant is returned when the variant is invalid.
	ErrInvalidVariant = errors.New("invalid UUID variant")
)

// String returns the string representation of the variant.
func (v Variant) String() string {
	switch v {
	case Variant0:
		return "V0"
	case Variant1:
		return "V1"
	case Variant2:
		return "V2"
	case Variant3:
		return "V3"
	default:
		return "unknown"
	}
}

// Value returns the value of the variant.
func (v Variant) Value() int {
	switch v {
	case Variant0:
		return 0
	case Variant1:
		return 1
	case Variant2:
		return 2
	case Variant3:
		return 3
	}

	return 0
}

// VariantFromValue returns the variant of the UUID from the value,
// returns ErrInvalidVariant if the value is invalid.
func VariantFromValue(value int) (Variant, error) {
	switch value {
	case 0:
		return Variant0, nil
	case 1:
		return Variant1, nil
	case 2:
		return Variant2, nil
	case 3:
		return Variant3, nil
	default:
		return Variant0, ErrInvalidVariant
	}
}

// getVariant returns the variant of the UUID.
func getVariant(u UUID) Variant {
	data := u[VERSION_BYTE_INDEX]

	if data & getVariantMask(Variant0) == byte(Variant0) {
		return Variant0
	}
	if data & getVariantMask(Variant1) == byte(Variant1) {
		return Variant1
	}
	if data & getVariantMask(Variant2) == byte(Variant2) {
		return Variant2
	}
	if data & getVariantMask(Variant3) == byte(Variant3) {
		return Variant3
	}

	return Variant0
}

// getVariantMask returns the variant mask of the UUID.
func getVariantMask(v Variant) byte {
	// 0 x x. 0 bits => 0b1000_0000
	// 1 0 x. 1 bits => 0b1100_0000
	// 1 1 0. 2 bits => 0b1110_0000
	// 1 1 1. 3 bits => 0b1110_0000

	c := util.PopCount(uint(v))
	r := c + 1
	if r > 3 {
		r = 3
	}

	baseMask := 0b1111_0000_0000
	outputMask := 0b1110_0000

	return byte((baseMask >> r) & outputMask)
}
