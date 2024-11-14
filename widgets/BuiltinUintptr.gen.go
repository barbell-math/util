package widgets

// Code generated by ../bin/widgetInterfaceImpl - DO NOT EDIT.

import (
	"github.com/barbell-math/util/hash"
)

var ()

// A widget to represent the built-in uintptr type
type (
	BuiltinUintptr struct{}
)

// Returns true if both uintptr's are equal. Uses the standard == operator internally.
func (_ BuiltinUintptr) Eq(l *uintptr, r *uintptr) bool {
	return *l == *r
}

// Provides a hash function for the value that it is wrapping.
func (_ BuiltinUintptr) Hash(v *uintptr) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (_ BuiltinUintptr) Zero(other *uintptr) {
	*other = (uintptr)(0)
}

// Returns true if l<r. Uses the standard < operator internally.
func (_ BuiltinUintptr) Lt(l *uintptr, r *uintptr) bool {
	return *l < *r
}

// Returns the zero value for the uintptr type.
func (_ BuiltinUintptr) ZeroVal() uintptr {
	return uintptr(0)
}

// Returns the unit value for the uintptr type.
func (_ BuiltinUintptr) UnitVal() uintptr {
	return uintptr(1)
}

// Negates v, updating the value that v points to.
func (_ BuiltinUintptr) Neg(v *uintptr) {
	*v = -(*v)
}

// Adds l to r, placing the result in the value that res points to.
func (_ BuiltinUintptr) Add(res *uintptr, l *uintptr, r *uintptr) {
	*res = *l + *r
}

// Subtracts l to r, placing the result in the value that res points to.
func (_ BuiltinUintptr) Sub(res *uintptr, l *uintptr, r *uintptr) {
	*res = *l - *r
}

// Multiplies l to r, placing the result in the value that res points to.
func (_ BuiltinUintptr) Mul(res *uintptr, l *uintptr, r *uintptr) {
	*res = *l * *r
}

// Divides l to r, placing the result in the value that res points to.
func (_ BuiltinUintptr) Div(res *uintptr, l *uintptr, r *uintptr) {
	*res = *l / *r
}
