package widgets

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/hash"

// A widget to represent the built-in type bool
// This is meant to be used with the containers from the [containers] package.
type BuiltinBool struct{}

// Returns true if both bool's are equal. Uses the standard == operator internally.
func (_ BuiltinBool) Eq(l *bool, r *bool) bool {
	return *l == *r
}

// Returns true if a is less than r. Uses the standard < operator internally.
func (_ BuiltinBool) Lt(l *bool, r *bool) bool {
	return (!*l && *r)
}

// Provides a hash function for the value that it is wrapping.
func (_ BuiltinBool) Hash(v *bool) hash.Hash {
	if *v {
		return hash.Hash(1)
	}
	return hash.Hash(0)
}

// Zeros the supplied value.
func (_ BuiltinBool) Zero(v *bool) {
	*v = false
}

// A bool is not an arithmetic aware widget. Bools are only base widgets.
