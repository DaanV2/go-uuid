package uuid

// StampVersion stamps the version and variant of the UUID.
func StampVersionVariant(uuid UUID, version Version, variant Variant) UUID {
	versionMask := getVersionMask()
	variantMask := getVariantMask(variant)

	uuid[VERSION_BYTE_INDEX] = (uuid[VERSION_BYTE_INDEX] & ^versionMask) | byte(version)
	uuid[VARIANT_BYTE_INDEX] = (uuid[VARIANT_BYTE_INDEX] & ^variantMask) | byte(variant)

	return uuid
}

// batchStampVersionVariant stamps the version and variant of the UUIDs.
func batchStampVersionVariant(uuids []UUID, version Version, variant Variant) []UUID {
	versionMask := getVersionMask()
	variantMask := getVariantMask(variant)

	for i := range uuids {
		uuid := uuids[i]
		uuid[VERSION_BYTE_INDEX] = (uuid[VERSION_BYTE_INDEX] & ^versionMask) | byte(version)
		uuid[VARIANT_BYTE_INDEX] = (uuid[VARIANT_BYTE_INDEX] & ^variantMask) | byte(variant)
		uuids[i] = uuid
	}

	return uuids
}
