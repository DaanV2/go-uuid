package uuid

// createUUID creates a UUID from the given bytes, version and variant.
func createUUID(bytes [16]byte, version Version, variant Variant) UUID {
	uuid := UUID(bytes)

	return StampVersionVariant(uuid, version, variant)
}

func createBatchUUID(bytes []byte, version Version, variant Variant) []UUID {
	c := len(bytes) / 16

	uuids := make([]UUID, c)
	

	for i := 0; i < c; i++ {
		var data [16]byte
		copy(data[:], bytes[i*16:])

		uuids[i] = UUID(data)
	}

	return batchStampVersionVariant(uuids, version, variant)
	
}