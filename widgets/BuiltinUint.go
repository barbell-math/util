package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/hash"

// A widget to represent the built-in type uint
// This is meant to be used with the containers from the [containers] package.
type BuiltinUint struct{}

// Returns true if both uint's are equal. Uses the standard == operator internally.
func (_ BuiltinUint) Eq(l *uint, r *uint) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (_ BuiltinUint) Lt(l *uint, r *uint) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (_ BuiltinUint) Hash(v *uint) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (_ BuiltinUint) Zero(v *uint) {
	*v = uint(0)
}

func (_ BuiltinUint) ZeroVal() uint {
	return uint(0)
}

func (_ BuiltinUint) UnitVal() uint {
	return uint(1)
}

func (_ BuiltinUint) Neg(v *uint) {
	*v = -(*v)
}

func (_ BuiltinUint) Add(res *uint, l *uint, r *uint) {
	*res = *l + *r
}

func (_ BuiltinUint) Sub(res *uint, l *uint, r *uint) {
	*res = *l - *r
}

func (_ BuiltinUint) Mul(res *uint, l *uint, r *uint) {
	*res = *l * *r
}

func (_ BuiltinUint) Div(res *uint, l *uint, r *uint) {
	*res = *l / *r
}
