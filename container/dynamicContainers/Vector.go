package dynamicContainers

import "github.com/barbell-math/util/container/containerTypes"

// An interface that only allows read operations on a vector.
type ReadVector[V any] interface {
	containerTypes.RWSyncable
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.ReadOps[V]
	containerTypes.ReadKeyedOps[int,V]
	containerTypes.Comparisons[
		containerTypes.ComparisonsOtherConstraint[V], 
		int, 
		V,
	]
	containerTypes.KeyedComparisons[
		containerTypes.KeyedComparisonsOtherConstraint[int,V],
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
	containerTypes.WriteOps[int,V]
	containerTypes.WriteKeyedOps[int,V]
	containerTypes.DeleteOps[int,V]
	containerTypes.DeleteKeyedOps[int,V]
}
// An interface that represents a vector with no restrictions on reading or
// writing.
type Vector[V any] interface {
	ReadVector[V]
	WriteVector[V]
}
