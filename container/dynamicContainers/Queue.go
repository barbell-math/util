package dynamicContainers

import "github.com/barbell-math/util/container/containerTypes"

// An interface that only allows read operations on a vector.
type ReadQueue[K ~int, V any] interface {
	containerTypes.RWSyncable
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.FirstElemRead[V]
}
// An interface that only allows write operations on a vector.
type WriteQueue[K ~int, V any] interface {
	containerTypes.RWSyncable
	containerTypes.Clear
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.LastElemWrite[V]
	containerTypes.FirstElemDelete[V]
}
// An interface that represents a vector with no restrictions on reading or
// writing.
type Queue[K ~int, V any] interface {
	ReadQueue[K,V]
	WriteQueue[K,V]
}
