package widgets

// Code generated by ../../bin/widgetInterfaceImpl - DO NOT EDIT.

import (
	"github.com/barbell-math/util/src/hash"

	"unsafe"
)

var ()

// A widget to represent the built-in float32 type
type (
	BuiltinFloat32 struct{}
)

// Returns true if both float32's are equal. Uses the standard == operator internally.
func (_ BuiltinFloat32) Eq(l *float32, r *float32) bool {
	return *l == *r
}

// Provides a hash function for the value that it is wrapping.
func (_ BuiltinFloat32) Hash(v *float32) hash.Hash {
	return hash.Hash(*(*uint32)(unsafe.Pointer(v)))
}

// Zeros the supplied value.
func (_ BuiltinFloat32) Zero(other *float32) {
	*other = (float32)(0)
}

// Returns true if l<r. Uses the standard < operator internally.
func (_ BuiltinFloat32) Lt(l *float32, r *float32) bool {
	return *l < *r
}

// Returns the zero value for the float32 type.
func (_ BuiltinFloat32) ZeroVal() float32 {
	return float32(0)
}

// Returns the unit value for the float32 type.
func (_ BuiltinFloat32) UnitVal() float32 {
	return float32(1)
}

// Negates v, updating the value that v points to.
func (_ BuiltinFloat32) Neg(v *float32) {
	*v = -(*v)
}

// Adds l to r, placing the result in the value that res points to.
func (_ BuiltinFloat32) Add(res *float32, l *float32, r *float32) {
	*res = *l + *r
}

// Subtracts l to r, placing the result in the value that res points to.
func (_ BuiltinFloat32) Sub(res *float32, l *float32, r *float32) {
	*res = *l - *r
}

// Multiplies l to r, placing the result in the value that res points to.
func (_ BuiltinFloat32) Mul(res *float32, l *float32, r *float32) {
	*res = *l * *r
}

// Divides l to r, placing the result in the value that res points to.
func (_ BuiltinFloat32) Div(res *float32, l *float32, r *float32) {
	*res = *l / *r
}
