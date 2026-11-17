package aliasdeps

type UnderlyingType struct {
	a uint64
	b uint64
}

// NewTypeName has no generated proofs of its own: its proofs are
// UnderlyingType's.
type NewTypeName UnderlyingType

type Outer struct {
	x NewTypeName
}
