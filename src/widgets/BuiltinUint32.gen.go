package widgets

// Code generated by ../../bin/widgetInterfaceImpl - DO NOT EDIT.

import (
	"github.com/barbell-math/util/src/hash"
)

var ()

// A widget to represent the built-in uint32 type
type (
	BuiltinUint32 struct{}
)

// Returns true if both uint32's are equal. Uses the standard == operator internally.
func (_ BuiltinUint32) Eq(l *uint32, r *uint32) bool {
	return *l == *r
}

// Provides a hash function for the value that it is wrapping.
func (_ BuiltinUint32) Hash(v *uint32) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (_ BuiltinUint32) Zero(other *uint32) {
	*other = (uint32)(0)
}

// Returns true if l<r. Uses the standard < operator internally.
func (_ BuiltinUint32) Lt(l *uint32, r *uint32) bool {
	return *l < *r
}

// Returns the zero value for the uint32 type.
func (_ BuiltinUint32) ZeroVal() uint32 {
	return uint32(0)
}

// Returns the unit value for the uint32 type.
func (_ BuiltinUint32) UnitVal() uint32 {
	return uint32(1)
}

// Negates v, updating the value that v points to.
func (_ BuiltinUint32) Neg(v *uint32) {
	*v = -(*v)
}

// Adds l to r, placing the result in the value that res points to.
func (_ BuiltinUint32) Add(res *uint32, l *uint32, r *uint32) {
	*res = *l + *r
}

// Subtracts l to r, placing the result in the value that res points to.
func (_ BuiltinUint32) Sub(res *uint32, l *uint32, r *uint32) {
	*res = *l - *r
}

// Multiplies l to r, placing the result in the value that res points to.
func (_ BuiltinUint32) Mul(res *uint32, l *uint32, r *uint32) {
	*res = *l * *r
}

// Divides l to r, placing the result in the value that res points to.
func (_ BuiltinUint32) Div(res *uint32, l *uint32, r *uint32) {
	*res = *l / *r
}
