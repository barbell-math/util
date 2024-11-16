package widgets

// Code generated by ../../bin/widgetInterfaceImpl - DO NOT EDIT.

import (
	"github.com/barbell-math/util/src/hash"
)

var ()

// A widget to represent the built-in int32 type
type (
	BuiltinInt32 struct{}
)

// Returns true if both int32's are equal. Uses the standard == operator internally.
func (_ BuiltinInt32) Eq(l *int32, r *int32) bool {
	return *l == *r
}

// Provides a hash function for the value that it is wrapping.
func (_ BuiltinInt32) Hash(v *int32) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (_ BuiltinInt32) Zero(other *int32) {
	*other = (int32)(0)
}

// Returns true if l<r. Uses the standard < operator internally.
func (_ BuiltinInt32) Lt(l *int32, r *int32) bool {
	return *l < *r
}

// Returns the zero value for the int32 type.
func (_ BuiltinInt32) ZeroVal() int32 {
	return int32(0)
}

// Returns the unit value for the int32 type.
func (_ BuiltinInt32) UnitVal() int32 {
	return int32(1)
}

// Negates v, updating the value that v points to.
func (_ BuiltinInt32) Neg(v *int32) {
	*v = -(*v)
}

// Adds l to r, placing the result in the value that res points to.
func (_ BuiltinInt32) Add(res *int32, l *int32, r *int32) {
	*res = *l + *r
}

// Subtracts l to r, placing the result in the value that res points to.
func (_ BuiltinInt32) Sub(res *int32, l *int32, r *int32) {
	*res = *l - *r
}

// Multiplies l to r, placing the result in the value that res points to.
func (_ BuiltinInt32) Mul(res *int32, l *int32, r *int32) {
	*res = *l * *r
}

// Divides l to r, placing the result in the value that res points to.
func (_ BuiltinInt32) Div(res *int32, l *int32, r *int32) {
	*res = *l / *r
}