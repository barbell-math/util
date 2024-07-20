package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/algo/hash"

// A widget to represent the built-in type uint8
// This is meant to be used with the containers from the [containers] package.
type BuiltinUint8 struct{}

// Returns true if both uint8's are equal. Uses the standard == operator internally.
func (a BuiltinUint8) Eq(l *uint8, r *uint8) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (a BuiltinUint8) Lt(l *uint8, r *uint8) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (a BuiltinUint8) Hash(v *uint8) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (a BuiltinUint8) Zero(v *uint8) {
	*v = uint8(0)
}

func (a BuiltinUint8) ZeroVal() uint8 {
	return uint8(0)
}

func (a BuiltinUint8) UnitVal() uint8 {
	return uint8(1)
}

func (a BuiltinUint8) Neg(v *uint8) {
	*v = -(*v)
}

func (a BuiltinUint8) Add(res *uint8, l *uint8, r *uint8) {
	*res = *l + *r
}

func (a BuiltinUint8) Sub(res *uint8, l *uint8, r *uint8) {
	*res = *l - *r
}

func (a BuiltinUint8) Mul(res *uint8, l *uint8, r *uint8) {
	*res = *l * *r
}

func (a BuiltinUint8) Div(res *uint8, l *uint8, r *uint8) {
	*res = *l / *r
}
