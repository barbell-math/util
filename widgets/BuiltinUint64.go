package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/hash"

// A widget to represent the built-in type uint64
// This is meant to be used with the containers from the [containers] package.
type BuiltinUint64 struct{}

// Returns true if both uint64's are equal. Uses the standard == operator internally.
func (_ BuiltinUint64) Eq(l *uint64, r *uint64) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (_ BuiltinUint64) Lt(l *uint64, r *uint64) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (_ BuiltinUint64) Hash(v *uint64) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (_ BuiltinUint64) Zero(v *uint64) {
	*v = uint64(0)
}

func (_ BuiltinUint64) ZeroVal() uint64 {
	return uint64(0)
}

func (_ BuiltinUint64) UnitVal() uint64 {
	return uint64(1)
}

func (_ BuiltinUint64) Neg(v *uint64) {
	*v = -(*v)
}

func (_ BuiltinUint64) Add(res *uint64, l *uint64, r *uint64) {
	*res = *l + *r
}

func (_ BuiltinUint64) Sub(res *uint64, l *uint64, r *uint64) {
	*res = *l - *r
}

func (_ BuiltinUint64) Mul(res *uint64, l *uint64, r *uint64) {
	*res = *l * *r
}

func (_ BuiltinUint64) Div(res *uint64, l *uint64, r *uint64) {
	*res = *l / *r
}
