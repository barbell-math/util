package dynamicContainers

import "github.com/barbell-math/util/container/containerTypes"

// An interface that only allows read operations on a set.
type ReadSet[V any] interface {
	containerTypes.RWSyncable
	containerTypes.Addressable
	containerTypes.Length
	containerTypes.ReadOps[V]
	containerTypes.ReadUniqueOps[V]
	containerTypes.Comparisons[
		containerTypes.ComparisonsOtherConstraint[V],
		V,
	]
}

// An interface that only allows write operations on a set.
type WriteSet[V any] interface {
	containerTypes.RWSyncable
	containerTypes.Clear
	containerTypes.Length
	containerTypes.WriteUniqueOps[uint64, V]
	containerTypes.DeleteOps[uint64, V]
	containerTypes.SetOperations[
		containerTypes.ComparisonsOtherConstraint[V],
		V,
	]
}

// An interface that represents a set with no rectrictions on reading or
// writing.
type Set[V any] interface {
	ReadSet[V]
	WriteSet[V]
}
