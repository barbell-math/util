package dynamicContainers

import "github.com/barbell-math/util/container/containerTypes"

// An iterface that only allows read operations on a set.
type ReadSet[K uint64, V any] interface {
	containerTypes.RWSyncable
	containerTypes.Length
	//containerTypes.Capacity
	containerTypes.ReadOps[K,V]
}
// An iterface that only allows write operations on a set.
type WriteSet[K uint64, V any] interface {
	containerTypes.RWSyncable
	containerTypes.Length
	//containerTypes.Capacity
	containerTypes.WriteOps[K,V]
}
// An interface that represents a set with no rectrictions on reading or
// writing.
type Set[K uint64, V any] interface {
	ReadSet[K,V]
	WriteSet[K,V]
}
