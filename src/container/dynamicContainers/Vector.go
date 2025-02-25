package dynamicContainers

import "github.com/barbell-math/util/src/container/containerTypes"

// An interface that only allows read operations on a vector.
type ReadVector[V any] interface {
	containerTypes.RWSyncable
	containerTypes.Addressable
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.ReadOps[V]
	containerTypes.ReadKeyedOps[int, V]
	containerTypes.Comparisons[
		containerTypes.ComparisonsOtherConstraint[V],
		V,
	]
	containerTypes.KeyedComparisons[
		containerTypes.KeyedComparisonsOtherConstraint[int, V],
		int,
		V,
	]
}

// An interface that only allows write operations on a vector.
type WriteVector[V any] interface {
	containerTypes.RWSyncable
	containerTypes.Clear
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.WriteOps[V]
	containerTypes.WriteKeyedOps[int, V]
	containerTypes.WriteKeyedSequentialOps[int, V]
	containerTypes.WriteDynKeyedOps[int, V]
	containerTypes.DeleteOps[int, V]
	containerTypes.DeleteKeyedOps[int, V]
	containerTypes.DeleteSequentialOps[int, V]
	containerTypes.DeleteKeyedSequentialOps[int, V]
	containerTypes.SetOperations[
		containerTypes.ComparisonsOtherConstraint[V],
		V,
	]
}

// An interface that represents a vector with no restrictions on reading or
// writing.
type Vector[V any] interface {
	ReadVector[V]
	WriteVector[V]
}
