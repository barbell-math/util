package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/algo/hash"

// A widget to represent the built-in type float64
// This is meant to be used with the containers from the [containers] package.
type BuiltinFloat64 struct{}

// Returns true if both float64's are equal. Uses the standard == operator internally.
func (a BuiltinFloat64) Eq(l *float64, r *float64) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (a BuiltinFloat64) Lt(l *float64, r *float64) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (a BuiltinFloat64) Hash(v *float64) hash.Hash {
	panic("Floats are not hashable!")
}

// Zeros the supplied value.
func (a BuiltinFloat64) Zero(v *float64) {
	*v = float64(0)
}

func (a BuiltinFloat64) ZeroVal() float64 {
	return float64(0)
}

func (a BuiltinFloat64) UnitVal() float64 {
	return float64(1)
}

func (a BuiltinFloat64) Neg(v *float64) {
	*v = -(*v)
}

func (a BuiltinFloat64) Add(res *float64, l *float64, r *float64) {
	*res = *l + *r
}

func (a BuiltinFloat64) Sub(res *float64, l *float64, r *float64) {
	*res = *l - *r
}

func (a BuiltinFloat64) Mul(res *float64, l *float64, r *float64) {
	*res = *l * *r
}

func (a BuiltinFloat64) Div(res *float64, l *float64, r *float64) {
	*res = *l / *r
}
