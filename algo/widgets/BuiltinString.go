package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "hash/maphash"
import "github.com/barbell-math/util/algo/hash"

// The random seed will be different every time the program runs// meaning that between runs the hash values will not be consistent.
// This was done for security purposes.
var RANDOM_SEED_string maphash.Seed = maphash.MakeSeed()

// A widget to represent the built-in type string
// This is meant to be used with the containers from the [containers] package.
type BuiltinString struct{}

// Returns true if both string's are equal. Uses the standard == operator internally.
func (a BuiltinString) Eq(l *string, r *string) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (a BuiltinString) Lt(l *string, r *string) bool {
	return *l < *r
}

// Provides a hash function for the value that it is wrapping.
func (a BuiltinString) Hash(v *string) hash.Hash {
	return hash.Hash(maphash.String(RANDOM_SEED_string, *(v)))
}

// Zeros the supplied value.
func (a BuiltinString) Zero(v *string) {
	*v = ""
}

// A string is not an arithmetic aware widget. Strings are only base widgets.
