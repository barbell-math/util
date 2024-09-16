package containers

// Code generated by ../../bin/passThroughWidget - DO NOT EDIT.

import ()

// Returns true if l equals r. Uses the Eq operator provided by the *HashSetHash widget internally.
func (_ *edgeHash) Eq(l *edgeHash, r *edgeHash) bool {
	var tmp *HashSetHash
	return tmp.Eq((*HashSetHash)(l), (*HashSetHash)(r))
}

// Returns a hash to represent other. The hash that is returned will be supplied by the *HashSetHash widget internally.
func (_ *edgeHash) Hash(other *edgeHash) hash.Hash {
	var tmp *HashSetHash
	return tmp.Hash((*HashSetHash)(other))
}

// Zeros the supplied value. The operation that is performed will be determined by the *HashSetHash widget internally.
func (_ *edgeHash) Zero(other *edgeHash) {
	var tmp *HashSetHash
	tmp.Zero((*HashSetHash)(other))
}
