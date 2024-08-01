package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/hash"

// A widget to represent the built-in type byte
// This is meant to be used with the containers from the [containers] package.
type BuiltinByte struct{}

// Returns true if both byte's are equal. Uses the standard == operator internally.
func (_ BuiltinByte) Eq(l *byte, r *byte) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (_ BuiltinByte) Lt(l *byte, r *byte) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (_ BuiltinByte) Hash(v *byte) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (_ BuiltinByte) Zero(v *byte) {
	*v = byte(0)
}

func (_ BuiltinByte) ZeroVal() byte {
	return byte(0)
}

func (_ BuiltinByte) UnitVal() byte {
	return byte(1)
}

func (_ BuiltinByte) Neg(v *byte) {
	*v = -(*v)
}

func (_ BuiltinByte) Add(res *byte, l *byte, r *byte) {
	*res = *l + *r
}

func (_ BuiltinByte) Sub(res *byte, l *byte, r *byte) {
	*res = *l - *r
}

func (_ BuiltinByte) Mul(res *byte, l *byte, r *byte) {
	*res = *l * *r
}

func (_ BuiltinByte) Div(res *byte, l *byte, r *byte) {
	*res = *l / *r
}
