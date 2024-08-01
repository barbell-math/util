package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/math/basic"
import "github.com/barbell-math/util/hash"

// A widget to represent the built-in type complex128
// This is meant to be used with the containers from the [containers] package.
type BuiltinComplex128 struct{}

// Returns true if both complex128's are equal. Uses the standard == operator internally.
func (_ BuiltinComplex128) Eq(l *complex128, r *complex128) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (_ BuiltinComplex128) Lt(l *complex128, r *complex128) bool {
	panic("Complex values cannot be compared relative to each other!")
}

// Provides a hash function for the value that it is wrapping.
func (_ BuiltinComplex128) Hash(v *complex128) hash.Hash {
	return hash.Hash(basic.LossyConv[float64, int64](basic.RealPart[complex128, float64](*v))).
		Combine(hash.Hash(basic.LossyConv[float64, int64](basic.ImaginaryPart[complex128, float64](*v))))
}

// Zeros the supplied value.
func (_ BuiltinComplex128) Zero(v *complex128) {
	*v = complex128(0)
}

func (_ BuiltinComplex128) ZeroVal() complex128 {
	return complex128(0)
}

func (_ BuiltinComplex128) UnitVal() complex128 {
	return complex128(1)
}

func (_ BuiltinComplex128) Neg(v *complex128) {
	*v = -(*v)
}

func (_ BuiltinComplex128) Add(res *complex128, l *complex128, r *complex128) {
	*res = *l + *r
}

func (_ BuiltinComplex128) Sub(res *complex128, l *complex128, r *complex128) {
	*res = *l - *r
}

func (_ BuiltinComplex128) Mul(res *complex128, l *complex128, r *complex128) {
	*res = *l * *r
}

func (_ BuiltinComplex128) Div(res *complex128, l *complex128, r *complex128) {
	*res = *l / *r
}
