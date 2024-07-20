package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/algo/hash"

// A widget to represent the built-in type uint64
// This is meant to be used with the containers from the [containers] package.
type BuiltinUint64 struct{}

// Returns true if both uint64's are equal. Uses the standard == operator internally.
func (a BuiltinUint64) Eq(l *uint64, r *uint64) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (a BuiltinUint64) Lt(l *uint64, r *uint64) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (a BuiltinUint64) Hash(v *uint64) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (a BuiltinUint64) Zero(v *uint64) {
	*v = uint64(0)
}

func (a BuiltinUint64) ZeroVal() uint64 {
	return uint64(0)
}

func (a BuiltinUint64) UnitVal() uint64 {
	return uint64(1)
}

func (a BuiltinUint64) Neg(v *uint64) {
	*v = -(*v)
}

func (a BuiltinUint64) Add(res *uint64, l *uint64, r *uint64) {
	*res = *l + *r
}

func (a BuiltinUint64) Sub(res *uint64, l *uint64, r *uint64) {
	*res = *l - *r
}

func (a BuiltinUint64) Mul(res *uint64, l *uint64, r *uint64) {
	*res = *l * *r
}

func (a BuiltinUint64) Div(res *uint64, l *uint64, r *uint64) {
	*res = *l / *r
}
