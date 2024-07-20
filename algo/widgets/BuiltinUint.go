package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/algo/hash"

// A widget to represent the built-in type uint
// This is meant to be used with the containers from the [containers] package.
type BuiltinUint struct{}

// Returns true if both uint's are equal. Uses the standard == operator internally.
func (a BuiltinUint) Eq(l *uint, r *uint) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (a BuiltinUint) Lt(l *uint, r *uint) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (a BuiltinUint) Hash(v *uint) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (a BuiltinUint) Zero(v *uint) {
	*v = uint(0)
}

func (a BuiltinUint) ZeroVal() uint {
	return uint(0)
}

func (a BuiltinUint) UnitVal() uint {
	return uint(1)
}

func (a BuiltinUint) Neg(v *uint) {
	*v = -(*v)
}

func (a BuiltinUint) Add(res *uint, l *uint, r *uint) {
	*res = *l + *r
}

func (a BuiltinUint) Sub(res *uint, l *uint, r *uint) {
	*res = *l - *r
}

func (a BuiltinUint) Mul(res *uint, l *uint, r *uint) {
	*res = *l * *r
}

func (a BuiltinUint) Div(res *uint, l *uint, r *uint) {
	*res = *l / *r
}
