package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/hash"

// A widget to represent the built-in type hash.Hash
// This is meant to be used with the containers from the [containers] package.
type BuiltinHash struct{}

// Returns true if both hash.Hash's are equal. Uses the standard == operator internally.
func (a BuiltinHash) Eq(l *hash.Hash, r *hash.Hash) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (a BuiltinHash) Lt(l *hash.Hash, r *hash.Hash) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (a BuiltinHash) Hash(v *hash.Hash) hash.Hash {
	return hash.Hash(*v)
}

// Zeros the supplied value.
func (a BuiltinHash) Zero(v *hash.Hash) {
	*v = hash.Hash(0)
}

func (a BuiltinHash) ZeroVal() hash.Hash {
	return hash.Hash(0)
}

func (a BuiltinHash) UnitVal() hash.Hash {
	return hash.Hash(1)
}

func (a BuiltinHash) Neg(v *hash.Hash) {
	*v = -(*v)
}

func (a BuiltinHash) Add(res *hash.Hash, l *hash.Hash, r *hash.Hash) {
	*res = *l + *r
}

func (a BuiltinHash) Sub(res *hash.Hash, l *hash.Hash, r *hash.Hash) {
	*res = *l - *r
}

func (a BuiltinHash) Mul(res *hash.Hash, l *hash.Hash, r *hash.Hash) {
	*res = *l * *r
}

func (a BuiltinHash) Div(res *hash.Hash, l *hash.Hash, r *hash.Hash) {
	*res = *l / *r
}
