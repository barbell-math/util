package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/algo/hash"

// A widget to represent the built-in type int32
// This is meant to be used with the containers from the [containers] package.
type BuiltinInt32 struct{}

// Returns true if both int32's are equal. Uses the standard == operator internally.
func (a BuiltinInt32) Eq(l *int32, r *int32) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (a BuiltinInt32) Lt(l *int32, r *int32) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (a BuiltinInt32) Hash(v *int32) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (a BuiltinInt32) Zero(v *int32) {
	*v = int32(0)
}

func (a BuiltinInt32) ZeroVal() int32 {
	return int32(0)
}

func (a BuiltinInt32) UnitVal() int32 {
	return int32(1)
}

func (a BuiltinInt32) Neg(v *int32) {
	*v = -(*v)
}

func (a BuiltinInt32) Add(res *int32, l *int32, r *int32) {
	*res = *l + *r
}

func (a BuiltinInt32) Sub(res *int32, l *int32, r *int32) {
	*res = *l - *r
}

func (a BuiltinInt32) Mul(res *int32, l *int32, r *int32) {
	*res = *l * *r
}

func (a BuiltinInt32) Div(res *int32, l *int32, r *int32) {
	*res = *l / *r
}
