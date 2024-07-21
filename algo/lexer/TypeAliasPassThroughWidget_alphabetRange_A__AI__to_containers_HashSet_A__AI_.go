package lexer

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.

import "github.com/barbell-math/util/algo/hash"
import "github.com/barbell-math/util/container/containers"

// A pass through widget to represent the aliased type alphabetRange[A, AI]
// This is meant to be used with the containers from the [containers] package.
// Returns true if both alphabetRange[A, AI]'s are equal. Uses the Eq operator provided by the containers.HashSet[A, AI] widget internally.
func (_ *alphabetRange[A, AI]) Eq(l *alphabetRange[A, AI], r *alphabetRange[A, AI]) bool {
	var tmp containers.HashSet[A, AI]
	return tmp.Eq((*containers.HashSet[A, AI])(l), (*containers.HashSet[A, AI])(r))
}

// Returns true if a is less than r. Uses the Lt operator provided by the containers.HashSet[A, AI] widget internally.
func (_ *alphabetRange[A, AI]) Lt(l *alphabetRange[A, AI], r *alphabetRange[A, AI]) bool {
	var tmp containers.HashSet[A, AI]
	return tmp.Lt((*containers.HashSet[A, AI])(l), (*containers.HashSet[A, AI])(r))
}

// Provides a hash function for the value that it is wrapping. The value that is returned will be supplied by the containers.HashSet[A, AI] widget internally.
func (_ *alphabetRange[A, AI]) Hash(other *alphabetRange[A, AI]) hash.Hash {
	var tmp containers.HashSet[A, AI]
	return tmp.Hash((*containers.HashSet[A, AI])(other))
}

// Zeros the supplied value. The operation that is performed will be determined by the containers.HashSet[A, AI] widget internally.
func (_ *alphabetRange[A, AI]) Zero(other *alphabetRange[A, AI]) {
	var tmp containers.HashSet[A, AI]
	tmp.Zero((*containers.HashSet[A, AI])(other))
}
