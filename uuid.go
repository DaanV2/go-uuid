package uuid

type UUID [16]byte

// Empty returns an empty UUID (00000000-0000-0000-0000-000000000000).
func Empty() UUID {
	return UUID{}
}
