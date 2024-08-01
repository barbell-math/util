package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/math/basic"
import "github.com/barbell-math/util/hash"

// A widget to represent the built-in type complex64
// This is meant to be used with the containers from the [containers] package.
type BuiltinComplex64 struct{}

// Returns true if both complex64's are equal. Uses the standard == operator internally.
func (_ BuiltinComplex64) Eq(l *complex64, r *complex64) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (_ BuiltinComplex64) Lt(l *complex64, r *complex64) bool {
	panic("Complex values cannot be compared relative to each other!")
}

// Provides a hash function for the value that it is wrapping.
func (_ BuiltinComplex64) Hash(v *complex64) hash.Hash {
	return hash.Hash(basic.LossyConv[float32, int32](basic.RealPart[complex64, float32](*v))).
		Combine(hash.Hash(basic.LossyConv[float32, int32](basic.ImaginaryPart[complex64, float32](*v))))
}

// Zeros the supplied value.
func (_ BuiltinComplex64) Zero(v *complex64) {
	*v = complex64(0)
}

func (_ BuiltinComplex64) ZeroVal() complex64 {
	return complex64(0)
}

func (_ BuiltinComplex64) UnitVal() complex64 {
	return complex64(1)
}

func (_ BuiltinComplex64) Neg(v *complex64) {
	*v = -(*v)
}

func (_ BuiltinComplex64) Add(res *complex64, l *complex64, r *complex64) {
	*res = *l + *r
}

func (_ BuiltinComplex64) Sub(res *complex64, l *complex64, r *complex64) {
	*res = *l - *r
}

func (_ BuiltinComplex64) Mul(res *complex64, l *complex64, r *complex64) {
	*res = *l * *r
}

func (_ BuiltinComplex64) Div(res *complex64, l *complex64, r *complex64) {
	*res = *l / *r
}
