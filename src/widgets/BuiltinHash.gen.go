package widgets

// Code generated by ../../bin/widgetInterfaceImpl - DO NOT EDIT.

import (
	"github.com/barbell-math/util/src/hash"
)

var ()

// A widget to represent the built-in hash.Hash type
type (
	BuiltinHash struct{}
)

// Returns true if both hash.Hash's are equal. Uses the standard == operator internally.
func (_ BuiltinHash) Eq(l *hash.Hash, r *hash.Hash) bool {
	return *l == *r
}

// Provides a hash function for the value that it is wrapping.
func (_ BuiltinHash) Hash(v *hash.Hash) hash.Hash {
	return *v
}

// Zeros the supplied value.
func (_ BuiltinHash) Zero(other *hash.Hash) {
	*other = (hash.Hash)(0)
}

// Returns true if l<r. Uses the standard < operator internally.
func (_ BuiltinHash) Lt(l *hash.Hash, r *hash.Hash) bool {
	return *l < *r
}

// Returns the zero value for the hash.Hash type.
func (_ BuiltinHash) ZeroVal() hash.Hash {
	return hash.Hash(0)
}

// Returns the unit value for the hash.Hash type.
func (_ BuiltinHash) UnitVal() hash.Hash {
	return hash.Hash(1)
}

// Negates v, updating the value that v points to.
func (_ BuiltinHash) Neg(v *hash.Hash) {
	*v = -(*v)
}

// Adds l to r, placing the result in the value that res points to.
func (_ BuiltinHash) Add(res *hash.Hash, l *hash.Hash, r *hash.Hash) {
	*res = *l + *r
}

// Subtracts l to r, placing the result in the value that res points to.
func (_ BuiltinHash) Sub(res *hash.Hash, l *hash.Hash, r *hash.Hash) {
	*res = *l - *r
}

// Multiplies l to r, placing the result in the value that res points to.
func (_ BuiltinHash) Mul(res *hash.Hash, l *hash.Hash, r *hash.Hash) {
	*res = *l * *r
}

// Divides l to r, placing the result in the value that res points to.
func (_ BuiltinHash) Div(res *hash.Hash, l *hash.Hash, r *hash.Hash) {
	*res = *l / *r
}