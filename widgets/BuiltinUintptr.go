package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/hash"

// A widget to represent the built-in type uintptr
// This is meant to be used with the containers from the [containers] package.
type BuiltinUintptr struct{}

// Returns true if both uintptr's are equal. Uses the standard == operator internally.
func (_ BuiltinUintptr) Eq(l *uintptr, r *uintptr) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (_ BuiltinUintptr) Lt(l *uintptr, r *uintptr) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (_ BuiltinUintptr) Hash(v *uintptr) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (_ BuiltinUintptr) Zero(v *uintptr) {
	*v = uintptr(0)
}

func (_ BuiltinUintptr) ZeroVal() uintptr {
	return uintptr(0)
}

func (_ BuiltinUintptr) UnitVal() uintptr {
	return uintptr(1)
}

func (_ BuiltinUintptr) Neg(v *uintptr) {
	*v = -(*v)
}

func (_ BuiltinUintptr) Add(res *uintptr, l *uintptr, r *uintptr) {
	*res = *l + *r
}

func (_ BuiltinUintptr) Sub(res *uintptr, l *uintptr, r *uintptr) {
	*res = *l - *r
}

func (_ BuiltinUintptr) Mul(res *uintptr, l *uintptr, r *uintptr) {
	*res = *l * *r
}

func (_ BuiltinUintptr) Div(res *uintptr, l *uintptr, r *uintptr) {
	*res = *l / *r
}
