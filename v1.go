package uuid

import (
	"encoding/binary"
	"errors"
	"net"
	"time"
)

var (
	// ErrMacAddrNotFound is returned when the mac address is not found
	ErrMacAddrNotFound = errors.New("mac address not found")
)

var V1 v1

type v1 struct{}

// Version returns the current version of this generator
func (v1) Version() Version {
	return Version1
}

// Variant returns the current variant of this generator
func (v1) Variant() Variant {
	return Variant1
}

// New returns a new UUID based on the current time and mac address
func (v1) New() (UUID, error) {
	macAddress, err := getMacAddr()
	if err != nil {
		return Nil, err
	}

	timestamp, nanoSeconds := V1.TimeData()
	return V1.From(macAddress, timestamp, nanoSeconds), nil
}

// From returns a new UUID based on the given mac address and time
func (v1) From(macAddress net.HardwareAddr, timestamp int64, nanoSeconds uint16) UUID {
	// first 32 bites = time_low
	// next 16 bites = time_mid
	// next 16 bites = time_hi_version = time_high | version
	// next 8 bites = clock_seq_hi_variant
	// next 8 bites = clock_seq_low | node
	// last 48 bites mac address

	var data [16]byte

	// time low
	binary.BigEndian.PutUint32(data[0:], uint32(timestamp))
	// time mid
	binary.BigEndian.PutUint16(data[4:], uint16(timestamp>>32))
	// time hi version
	binary.BigEndian.PutUint16(data[6:], uint16(timestamp>>48))
	// clock seq hi variant
	binary.BigEndian.PutUint16(data[8:], uint16(nanoSeconds))
	// mac address
	copy(data[10:], macAddress)

	return createUUID(data, V1.Version(), V1.Variant())
}

// TimeData returns the current time and nano seconds
func (v1) TimeData() (timestamp int64, nanoSeconds uint16) {
	return V1.TimeDataFrom(time.Now().UTC())
}

// TimeDataFrom returns the time and nano seconds from the given time
func (v1) TimeDataFrom(t time.Time) (timestamp int64, nanoSeconds uint16) {
	timestamp, nanoSeconds = t.UnixNano(), uint16(t.Nanosecond())
	return
}

// getMacAddr returns the mac address of the current machine, if it exists
func getMacAddr() (net.HardwareAddr, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, ifa := range interfaces {
		addr := ifa.HardwareAddr
		if addr != nil && !isEmpty(addr) {
			return addr, nil
		}
	}
	return nil, ErrMacAddrNotFound
}

// isEmpty returns true if the given mac address is empty
func isEmpty(d net.HardwareAddr) bool {
	for _, b := range d {
		if b != 0 {
			return false
		}
	}
	return true
}
