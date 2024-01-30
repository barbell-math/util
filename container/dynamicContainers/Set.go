package dynamicContainers

import "github.com/barbell-math/util/container/containerTypes"

// An iterface that only allows read operations on a set.
type ReadSet[V any] interface {
	containerTypes.RWSyncable
	containerTypes.Length
	//containerTypes.Capacity
	containerTypes.ReadOps[uint64,V]
}
// An iterface that only allows write operations on a set.
type WriteSet[V any] interface {
	containerTypes.RWSyncable
	containerTypes.Clear
	containerTypes.Length
	//containerTypes.Capacity
	// containerTypes.WriteOps[K,V]
	containerTypes.WriteUniqueOps[uint64,V]
	containerTypes.DeleteOps[uint64,V]
}
// An interface that represents a set with no rectrictions on reading or
// writing.
type Set[V any] interface {
	ReadSet[V]
	WriteSet[V]
}
