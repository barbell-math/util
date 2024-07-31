package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/hash"

// A widget to represent the built-in type uint16
// This is meant to be used with the containers from the [containers] package.
type BuiltinUint16 struct{}

// Returns true if both uint16's are equal. Uses the standard == operator internally.
func (a BuiltinUint16) Eq(l *uint16, r *uint16) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (a BuiltinUint16) Lt(l *uint16, r *uint16) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (a BuiltinUint16) Hash(v *uint16) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (a BuiltinUint16) Zero(v *uint16) {
	*v = uint16(0)
}

func (a BuiltinUint16) ZeroVal() uint16 {
	return uint16(0)
}

func (a BuiltinUint16) UnitVal() uint16 {
	return uint16(1)
}

func (a BuiltinUint16) Neg(v *uint16) {
	*v = -(*v)
}

func (a BuiltinUint16) Add(res *uint16, l *uint16, r *uint16) {
	*res = *l + *r
}

func (a BuiltinUint16) Sub(res *uint16, l *uint16, r *uint16) {
	*res = *l - *r
}

func (a BuiltinUint16) Mul(res *uint16, l *uint16, r *uint16) {
	*res = *l * *r
}

func (a BuiltinUint16) Div(res *uint16, l *uint16, r *uint16) {
	*res = *l / *r
}
