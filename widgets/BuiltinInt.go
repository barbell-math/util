package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/hash"

// A widget to represent the built-in type int
// This is meant to be used with the containers from the [containers] package.
type BuiltinInt struct{}

// Returns true if both int's are equal. Uses the standard == operator internally.
func (_ BuiltinInt) Eq(l *int, r *int) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (_ BuiltinInt) Lt(l *int, r *int) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (_ BuiltinInt) Hash(v *int) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (_ BuiltinInt) Zero(v *int) {
	*v = int(0)
}

func (_ BuiltinInt) ZeroVal() int {
	return int(0)
}

func (_ BuiltinInt) UnitVal() int {
	return int(1)
}

func (_ BuiltinInt) Neg(v *int) {
	*v = -(*v)
}

func (_ BuiltinInt) Add(res *int, l *int, r *int) {
	*res = *l + *r
}

func (_ BuiltinInt) Sub(res *int, l *int, r *int) {
	*res = *l - *r
}

func (_ BuiltinInt) Mul(res *int, l *int, r *int) {
	*res = *l * *r
}

func (_ BuiltinInt) Div(res *int, l *int, r *int) {
	*res = *l / *r
}
