package uuid

import (
	"encoding/binary"
	"net"
	"time"
)

var V6 v6

type v6 struct{}

// Version returns the current version of this generator
func (v6) Version() Version {
	return Version6
}

// Variant returns the current variant of this generator
func (v6) Variant() Variant {
	return Variant1
}

// New returns a new UUID version 6 based on the current time and mac address
func (v6) New() (UUID, error) {
	macAddress, err := getMacAddr()
	if err != nil {
		return Nil, err
	}

	timestamp, nanoSeconds := V6.TimeData()
	return V6.From(macAddress, timestamp, nanoSeconds), nil
}

// From returns a new UUID version 6 based on the given mac address and time
// UUID v6 is a reordered version of UUID v1 for better database indexing
func (v6) From(macAddress net.HardwareAddr, timestamp int64, nanoSeconds uint16) UUID {
	// UUIDv6 Field and Bit Layout:
	// time_high (32 bits) - bits 0-31
	// time_mid (16 bits) - bits 32-47
	// time_low_and_version (16 bits) - bits 48-63
	// clock_seq_hi_and_reserved (8 bits) - bits 64-71
	// clock_seq_low (8 bits) - bits 72-79
	// node (48 bits) - bits 80-127

	var data [16]byte

	// Time fields are reordered from v1 for better sorting:
	// time_high (most significant 32 bits of 60-bit timestamp)
	binary.BigEndian.PutUint32(data[0:], uint32(timestamp>>28))
	// time_mid (middle 16 bits of timestamp)
	binary.BigEndian.PutUint16(data[4:], uint16(timestamp>>12))
	// time_low (least significant 12 bits of timestamp, shifted to make room for version)
	binary.BigEndian.PutUint16(data[6:], uint16(timestamp&0x0FFF))
	// clock sequence
	binary.BigEndian.PutUint16(data[8:], uint16(nanoSeconds))
	// mac address
	copy(data[10:], macAddress)

	return createUUID(data, V6.Version(), V6.Variant())
}

// TimeData returns the current time and nano seconds
func (v6) TimeData() (timestamp int64, nanoSeconds uint16) {
	return V6.TimeDataFrom(time.Now().UTC())
}

// TimeDataFrom returns the time and nano seconds from the given time
func (v6) TimeDataFrom(t time.Time) (timestamp int64, nanoSeconds uint16) {
	timestamp, nanoSeconds = t.UnixNano(), uint16(t.Nanosecond())
	return
}
