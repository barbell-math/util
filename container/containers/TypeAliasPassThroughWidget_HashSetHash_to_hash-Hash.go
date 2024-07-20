package containers

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/algo/hash"
import "github.com/barbell-math/util/algo/widgets"

// A pass through widget to represent the aliased type HashSetHash
// This is meant to be used with the containers from the [containers] package.
// Returns true if both HashSetHash's are equal. Uses the Eq operator provided by the widgets.BuiltinHash widget internally.
func (_ *HashSetHash)Eq(l *HashSetHash, r *HashSetHash) bool {
	var tmp widgets.BuiltinHash
	return tmp.Eq((*hash.Hash)(l), (*hash.Hash)(r))
}

// Returns true if a is less than r. Uses the Lt operator provided by the widgets.BuiltinHash widget internally.
func (_ *HashSetHash)Lt(l *HashSetHash, r *HashSetHash) bool {
	var tmp widgets.BuiltinHash
	return tmp.Lt((*hash.Hash)(l), (*hash.Hash)(r))
}

// Provides a hash function for the value that it is wrapping. The value that is returned will be supplied by the widgets.BuiltinHash widget internally.
func (_ *HashSetHash)Hash(other *HashSetHash) hash.Hash {
	var tmp widgets.BuiltinHash
	return tmp.Hash((*hash.Hash)(other))
}

// Zeros the supplied value. The operation that is performed will be determined by the widgets.BuiltinHash widget internally.
func (_ *HashSetHash)Zero(other *HashSetHash) {
	var tmp widgets.BuiltinHash
	tmp.Zero((*hash.Hash)(other))
}
