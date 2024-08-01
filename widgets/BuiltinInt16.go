package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/hash"

// A widget to represent the built-in type int16
// This is meant to be used with the containers from the [containers] package.
type BuiltinInt16 struct{}

// Returns true if both int16's are equal. Uses the standard == operator internally.
func (_ BuiltinInt16) Eq(l *int16, r *int16) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (_ BuiltinInt16) Lt(l *int16, r *int16) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (_ BuiltinInt16) Hash(v *int16) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (_ BuiltinInt16) Zero(v *int16) {
	*v = int16(0)
}

func (_ BuiltinInt16) ZeroVal() int16 {
	return int16(0)
}

func (_ BuiltinInt16) UnitVal() int16 {
	return int16(1)
}

func (_ BuiltinInt16) Neg(v *int16) {
	*v = -(*v)
}

func (_ BuiltinInt16) Add(res *int16, l *int16, r *int16) {
	*res = *l + *r
}

func (_ BuiltinInt16) Sub(res *int16, l *int16, r *int16) {
	*res = *l - *r
}

func (_ BuiltinInt16) Mul(res *int16, l *int16, r *int16) {
	*res = *l * *r
}

func (_ BuiltinInt16) Div(res *int16, l *int16, r *int16) {
	*res = *l / *r
}
