package uuid

// createUUID creates a UUID from the given bytes, version and variant.
func createUUID(bytes [TOTAL_BYTES]byte, version Version, variant Variant) UUID {
	uuid := UUID(bytes)

	return StampVersionVariant(uuid, version, variant)
}

func createBatchUUID(bytes []byte, version Version, variant Variant) []UUID {
	c := len(bytes) / TOTAL_BYTES
	uuids := make([]UUID, c)

	for i := range c {
		var data [TOTAL_BYTES]byte
		copy(data[:], bytes[i*TOTAL_BYTES:])

		uuids[i] = UUID(data)
	}

	return batchStampVersionVariant(uuids, version, variant)

}
