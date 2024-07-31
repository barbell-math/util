package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/hash"

// A widget to represent the built-in type int8
// This is meant to be used with the containers from the [containers] package.
type BuiltinInt8 struct{}

// Returns true if both int8's are equal. Uses the standard == operator internally.
func (a BuiltinInt8) Eq(l *int8, r *int8) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (a BuiltinInt8) Lt(l *int8, r *int8) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (a BuiltinInt8) Hash(v *int8) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (a BuiltinInt8) Zero(v *int8) {
	*v = int8(0)
}

func (a BuiltinInt8) ZeroVal() int8 {
	return int8(0)
}

func (a BuiltinInt8) UnitVal() int8 {
	return int8(1)
}

func (a BuiltinInt8) Neg(v *int8) {
	*v = -(*v)
}

func (a BuiltinInt8) Add(res *int8, l *int8, r *int8) {
	*res = *l + *r
}

func (a BuiltinInt8) Sub(res *int8, l *int8, r *int8) {
	*res = *l - *r
}

func (a BuiltinInt8) Mul(res *int8, l *int8, r *int8) {
	*res = *l * *r
}

func (a BuiltinInt8) Div(res *int8, l *int8, r *int8) {
	*res = *l / *r
}
