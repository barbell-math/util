package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/algo/hash"

// A widget to represent the built-in type uint32
// This is meant to be used with the containers from the [containers] package.
type BuiltinUint32 struct{}

// Returns true if both uint32's are equal. Uses the standard == operator internally.
func (a BuiltinUint32) Eq(l *uint32, r *uint32) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (a BuiltinUint32) Lt(l *uint32, r *uint32) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (a BuiltinUint32) Hash(v *uint32) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (a BuiltinUint32) Zero(v *uint32) {
	*v = uint32(0)
}

func (a BuiltinUint32) ZeroVal() uint32 {
	return uint32(0)
}

func (a BuiltinUint32) UnitVal() uint32 {
	return uint32(1)
}

func (a BuiltinUint32) Neg(v *uint32) {
	*v = -(*v)
}

func (a BuiltinUint32) Add(res *uint32, l *uint32, r *uint32) {
	*res = *l + *r
}

func (a BuiltinUint32) Sub(res *uint32, l *uint32, r *uint32) {
	*res = *l - *r
}

func (a BuiltinUint32) Mul(res *uint32, l *uint32, r *uint32) {
	*res = *l * *r
}

func (a BuiltinUint32) Div(res *uint32, l *uint32, r *uint32) {
	*res = *l / *r
}
