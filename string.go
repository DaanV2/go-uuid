package uuid

// String returns a string representation of the UUID in the following format:
// 00000000-0000-0000-0000-000000000000
func (u UUID) String() string {
	return bytesToString(u)
}

// StringHex returns a hex string representation of the UUID.
func (u UUID) StringHex() string {
	data := bytesToHexData(u)
	return string(data[:])
}

// bytesToString converts a byte array of uuid data to a string.
func bytesToString(b [16]byte) string {
	hexData := bytesToHexData(b)

	// Now copy the data into the proper format
	var result [STRING_LENGTH]byte

	for i, j := 0, 0; i < len(result); i++ {
		switch i {
		case 8, 13, 18, 23:
			result[i] = '-'
		default:
			result[i] = hexData[j]
			j++
		}
	}

	return string(result[:])
}

const (
	zeroByte = byte('0')
	nineByte = byte('9')
	aByte    = byte('a')
	offsetNineA = aByte - nineByte - 1
)

// bytesToHexData converts a byte array of data to the hex representation of that data.
func bytesToHexData(b [16]byte) [32]byte {
	var upperBits [16]byte
	var lowerBits [16]byte

	copy(upperBits[:], b[:])
	copy(lowerBits[:], b[:])

	for i, v := range upperBits {
		upperBits[i] = v >> 4
	}
	for i, v := range lowerBits {
		lowerBits[i] = v & 0x0F
	}

	// Raise both by '0'
	for i, v := range upperBits {
		upperBits[i] = v + zeroByte
	}
	for i, v := range lowerBits {
		lowerBits[i] = v + zeroByte
	}

	// If larger than '9', raise by 'a' - '9' - 1, so that the character is in the range 'a' - 'f'
	for i, v := range upperBits {
		if v > nineByte {
			upperBits[i] = v + offsetNineA
		} else {
			upperBits[i] = v
		}
	}
	for i, v := range lowerBits {
		if v > nineByte {
			lowerBits[i] = v + offsetNineA
		} else {
			lowerBits[i] = v
		}
	}

	// Now copy the data into the result
	var result [32]byte

	for i := range upperBits {
		j := i * 2
		result[j] = upperBits[i]
		result[j + 1] = lowerBits[i]
	}

	return result
}
