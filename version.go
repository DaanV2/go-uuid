package uuid

import "errors"

// Version represents the version of the UUID.
type Version uint

const (
	// VERSION_BYTE_MASK is the mask for the version byte.
	VERSION_BYTE_MASK uint8 = 0b11110000

	// V1 is the version that is made from the timestamp and MACAddress
	Version1 Version = 0b0001_0000
	// V2 Not Implemented
	Version2 Version = 0b0010_0000
	// V3 is the version that hashes (MD5) given data into an UUID
	Version3 Version = 0b0011_0000
	// V4 is the version that generates a random UUID
	Version4 Version = 0b0100_0000
	// V5 is the version that hashes (SHA1) given data into an UUID
	Version5 Version = 0b0101_0000
	// V6 is the version that is made from the reordered timestamp and MACAddress (RFC 9562)
	Version6 Version = 0b0110_0000
	// V7 is the version that is made from Unix timestamp with random data (RFC 9562)
	Version7 Version = 0b0111_0000
	// V8 is the version for custom/vendor-specific UUID format (RFC 9562)
	Version8 Version = 0b1000_0000
)

var (
	// ErrInvalidVersion is returned when the version is invalid.
	ErrInvalidVersion = errors.New("invalid UUID version")
)

// String returns the string representation of the version.
func (v Version) String() string {
	switch v {
	case Version1:
		return "v1"
	case Version2:
		return "v2"
	case Version3:
		return "v3"
	case Version4:
		return "v4"
	case Version5:
		return "v5"
	case Version6:
		return "v6"
	case Version7:
		return "v7"
	case Version8:
		return "v8"
	default:
		return "unknown"
	}
}

// Value returns the value of the version.
func (v Version) Value() int {
	return int(v >> 4)
}

// VersionFromValue returns the version of the UUID from the value,
// returns ErrInvalidVersion if the value is invalid.
func VersionFromValue(value int) (Version, error) {
	switch value {
	case 1:
		return Version1, nil
	case 2:
		return Version2, nil
	case 3:
		return Version3, nil
	case 4:
		return Version4, nil
	case 5:
		return Version5, nil
	case 6:
		return Version6, nil
	case 7:
		return Version7, nil
	case 8:
		return Version8, nil
	}

	return Version1, ErrInvalidVersion
}

// getVersion returns the version of the UUID.
func getVersion(u UUID) Version {
	data := u[VERSION_BYTE_INDEX] & VERSION_BYTE_MASK

	return Version(data)
}

func getVersionMask() byte {
	return VERSION_BYTE_MASK
}