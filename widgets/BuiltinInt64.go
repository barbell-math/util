package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/hash"

// A widget to represent the built-in type int64
// This is meant to be used with the containers from the [containers] package.
type BuiltinInt64 struct{}

// Returns true if both int64's are equal. Uses the standard == operator internally.
func (a BuiltinInt64) Eq(l *int64, r *int64) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (a BuiltinInt64) Lt(l *int64, r *int64) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (a BuiltinInt64) Hash(v *int64) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (a BuiltinInt64) Zero(v *int64) {
	*v = int64(0)
}

func (a BuiltinInt64) ZeroVal() int64 {
	return int64(0)
}

func (a BuiltinInt64) UnitVal() int64 {
	return int64(1)
}

func (a BuiltinInt64) Neg(v *int64) {
	*v = -(*v)
}

func (a BuiltinInt64) Add(res *int64, l *int64, r *int64) {
	*res = *l + *r
}

func (a BuiltinInt64) Sub(res *int64, l *int64, r *int64) {
	*res = *l - *r
}

func (a BuiltinInt64) Mul(res *int64, l *int64, r *int64) {
	*res = *l * *r
}

func (a BuiltinInt64) Div(res *int64, l *int64, r *int64) {
	*res = *l / *r
}
