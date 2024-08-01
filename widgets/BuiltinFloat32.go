package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/hash"

// A widget to represent the built-in type float32
// This is meant to be used with the containers from the [containers] package.
type BuiltinFloat32 struct{}

// Returns true if both float32's are equal. Uses the standard == operator internally.
func (_ BuiltinFloat32) Eq(l *float32, r *float32) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (_ BuiltinFloat32) Lt(l *float32, r *float32) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (_ BuiltinFloat32) Hash(v *float32) hash.Hash {
	panic("Floats are not hashable!")
}

// Zeros the supplied value.
func (_ BuiltinFloat32) Zero(v *float32) {
	*v = float32(0)
}

func (_ BuiltinFloat32) ZeroVal() float32 {
	return float32(0)
}

func (_ BuiltinFloat32) UnitVal() float32 {
	return float32(1)
}

func (_ BuiltinFloat32) Neg(v *float32) {
	*v = -(*v)
}

func (_ BuiltinFloat32) Add(res *float32, l *float32, r *float32) {
	*res = *l + *r
}

func (_ BuiltinFloat32) Sub(res *float32, l *float32, r *float32) {
	*res = *l - *r
}

func (_ BuiltinFloat32) Mul(res *float32, l *float32, r *float32) {
	*res = *l * *r
}

func (_ BuiltinFloat32) Div(res *float32, l *float32, r *float32) {
	*res = *l / *r
}
