package containers

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/algo/hash"

// A pass through widget to represent the aliased type vertexHash
// This is meant to be used with the containers from the [containers] package.
// Returns true if both vertexHash's are equal. Uses the Eq operator provided by the *HashSetHash widget internally.
func (_ *vertexHash)Eq(l *vertexHash, r *vertexHash) bool {
	var tmp *HashSetHash
	return tmp.Eq((*HashSetHash)(l), (*HashSetHash)(r))
}

// Returns true if a is less than r. Uses the Lt operator provided by the *HashSetHash widget internally.
func (_ *vertexHash)Lt(l *vertexHash, r *vertexHash) bool {
	var tmp *HashSetHash
	return tmp.Lt((*HashSetHash)(l), (*HashSetHash)(r))
}

// Provides a hash function for the value that it is wrapping. The value that is returned will be supplied by the *HashSetHash widget internally.
func (_ *vertexHash)Hash(other *vertexHash) hash.Hash {
	var tmp *HashSetHash
	return tmp.Hash((*HashSetHash)(other))
}

// Zeros the supplied value. The operation that is performed will be determined by the *HashSetHash widget internally.
func (_ *vertexHash)Zero(other *vertexHash) {
	var tmp *HashSetHash
	tmp.Zero((*HashSetHash)(other))
}
