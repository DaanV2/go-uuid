package uuid

// New returns a new UUID version 4
func New() UUID {
	return V4.New()
}

// NewHex returns a new UUID version 4 as a hex string
func NewHex() string {
	return V4.NewHex()
}

// NewString returns a new UUID version 4 as a string
func NewString() string {
	return V4.NewString()
}
